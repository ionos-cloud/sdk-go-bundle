/*
 * IONOS Shared Libraries – Failover RoundTripper
 */

package shared

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FailoverRoundTripper is an http.RoundTripper wrapper that retries the request
// against multiple configured servers when the underlying transport returns a
// network-level error or an HTTP status code listed in
// FailoverOptions.FailoverOnStatusCodes.
//
// It is controlled via Configuration.FailoverStrategy.
//
// Notes:
//   - Network errors trigger retries with exponential backoff, cycling to the
//     next server.
//   - Status codes in FailoverOnStatusCodes also trigger retries: the response
//     body is drained and the request is sent to the next server. Response
//     headers (e.g. Retry-After) are not inspected at this layer.
//   - If the request carries a body, the request must have GetBody set (the SDK
//     generates requests in a way that supports this).
//   - For non-idempotent requests (POST, PATCH, etc.), enabling failover on
//     timeouts can produce duplicates. This wrapper currently treats context
//     deadline/canceled as retryable network errors, so use with care.
//
// The request URL is rewritten by swapping scheme/host with each server URL,
// preserving path and query.
type FailoverRoundTripper struct {
	cfg  *Configuration
	base http.RoundTripper
}

// NewFailoverRoundTripper creates a new FailoverRoundTripper.
// If opts is nil, it will fall back to cfg.Failover.
func NewFailoverRoundTripper(cfg *Configuration, base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &FailoverRoundTripper{
		cfg:  cfg,
		base: base,
	}
}

// RoundTrip implements http.RoundTripper with failover logic based on the
// configured strategy and servers.
//
// This method is called by HTTPClient.Do() inside product-level callAPI. Each
// callAPI retry triggers a fresh RoundTrip invocation that cycles through all
// servers from the beginning.
//
// When the strategy is FailoverNone/empty, or there is only one server, the
// call passes through to the base transport with no retry logic.
//
// With an active multi-server strategy (e.g. FailoverRoundRobin):
//   - Network errors: retries with exponential backoff, cycling to the next
//     server.
//   - Status codes in FailoverOnStatusCodes: drains the response body and
//     cycles to the next server. Response headers (e.g. Retry-After) are not
//     inspected.
//   - If all attempts are exhausted, returns an error (not an HTTP response).
//     This means callAPI receives err != nil and returns immediately — its own
//     status-code-based retry logic (for 502/503/504/429) is never reached.
func (t *FailoverRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	if t == nil {
		return nil, errors.New("nil FailoverRoundTripper")
	}
	if t.base == nil {
		// Be resilient if instantiated without constructor.
		t.base = http.DefaultTransport
	}
	if t.cfg == nil {
		// No config => behave like the base transport.
		return t.base.RoundTrip(req)
	}

	fo := t.cfg.Failover
	if fo == nil {
		return t.base.RoundTrip(req)
	}

	bo := fo.ExponentialBackoff.NewExponentialBackoff()

	servers := t.cfg.Servers
	order := serverOrderFor(fo.Strategy, len(servers))
	if order == nil {
		// Unknown or disabled strategy => pass through.
		return t.base.RoundTrip(req)
	}

	// Check if method is allowed for failover retries.
	if !isRetryableMethod(fo, req.Method) {
		return t.base.RoundTrip(req)
	}

	maxRetries := fo.MaxRetries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}
	var lastErr error
	for attempt := range maxRetries {
		serverURL := servers[order(attempt)].URL

		targetURL, err := url.Parse(serverURL)
		if err != nil {
			return nil, fmt.Errorf("invalid server URL at Servers[%d]=%q: %w", order(attempt), serverURL, err)
		}

		attemptReq, err := cloneRequestForRetry(req)
		if err != nil {
			return nil, err
		}

		attemptReq.URL.Scheme = targetURL.Scheme
		attemptReq.URL.Host = targetURL.Host
		attemptReq.Host = targetURL.Host

		if SdkLogLevel.Satisfies(Debug) {
			SdkLogger.Printf("[Failover] attempt=%d method=%s url=%s", attempt+1, attemptReq.Method, attemptReq.URL.String())
		}

		resp, err := t.base.RoundTrip(attemptReq)
		if err != nil {
			lastErr = err
			retryable := isNetworkErrorRT(attemptReq.Context(), err, fo.RetryOnTimeout)
			if !retryable {
				return nil, err
			}
			if SdkLogLevel.Satisfies(Debug) {
				SdkLogger.Printf("[Failover] network error: %v; trying next server", err)
			}

			backoff(attemptReq.Context(), bo.NextBackOff())
			continue
		}

		if !shouldFailoverOnStatus(fo, resp.StatusCode) {
			return resp, nil
		}

		if SdkLogLevel.Satisfies(Debug) {
			SdkLogger.Printf("[Failover] status=%d triggers failover to next server", resp.StatusCode)
		}
		// Drain/close body to allow connection reuse.
		if resp.Body != nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}
		lastErr = fmt.Errorf("failover status: %s", resp.Status)
	}

	return nil, lastErr
}

func cloneRequestForRetry(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.Body != nil {
		if req.GetBody == nil {
			return nil, errors.New("request body is not replayable (GetBody is nil)")
		}
		b, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		clone.Body = b
	}
	return clone, nil
}

func isRetryableMethod(fo *FailoverOptions, method string) bool {
	m := strings.ToUpper(strings.TrimSpace(method))
	if fo == nil {
		return defaultRetryableMethods[m]
	}

	if len(fo.RetryableMethods) == 0 {
		return defaultRetryableMethods[m]
	}

	for _, v := range fo.RetryableMethods {
		if strings.ToUpper(strings.TrimSpace(v)) == m {
			return true
		}
	}
	return false
}

var defaultRetryableMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodPut:     true,
	http.MethodDelete:  true,
	http.MethodOptions: true,
}

func isNetworkErrorRT(ctx context.Context, err error, retryOnTimeout bool) bool {
	if err == nil {
		return false
	}

	// 1. Check for other transport-level errors (connection refused, reset, etc).
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return true
	}

	// 2. Handle url.Error cases not already caught by the typed checks above.
	// The string-based checks are intentionally omitted: "no such host" and
	// "connection refused" are always carried by net.DNSError / net.OpError
	// respectively, so errors.As in block 1 already handles them.
	// The only remaining retryable case here is DeadlineExceeded (subject to
	// the retryOnTimeout flag). All other url.Error variants – TLS certificate
	// failures, redirect-limit-exceeded, protocol mismatches – are non-transient
	// and must not trigger failover.
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if errors.Is(urlErr.Err, context.DeadlineExceeded) {
			return retryOnTimeout
		}
		return false
	}

	// 3. Handle bare timeout errors if enabled.
	if retryOnTimeout && ctx != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return true
		}
	}

	return false
}

func backoff(ctx context.Context, t time.Duration) {
	if ctx == nil {
		time.Sleep(t)
		return
	}
	select {
	case <-time.After(t):
	case <-ctx.Done():
	}
}

func shouldFailoverOnStatus(fo *FailoverOptions, statusCode int) bool {
	if fo == nil || len(fo.FailoverOnStatusCodes) == 0 {
		return false
	}
	for _, sc := range fo.FailoverOnStatusCodes {
		if sc == statusCode {
			return true
		}
	}
	return false
}

// serverOrder maps an attempt index (0, 1, 2, …) to a server index.
// Different strategies produce different orderings.
type serverOrder func(attempt int) int

// serverOrderFor returns a serverOrder for the given strategy, or nil when
// failover should not be applied (unknown/disabled strategy, ≤1 server).
func serverOrderFor(strategy FailoverStrategy, numServers int) serverOrder {
	s := strings.TrimSpace(strings.ToLower(string(strategy)))
	if s == "" || s == string(FailoverNone) || numServers <= 1 {
		return nil
	}
	switch s {
	case strings.ToLower(string(FailoverRoundRobin)):
		// Sequential: 0, 1, 2, …
		return func(attempt int) int { return attempt % numServers }
	default:
		return nil
	}
}
