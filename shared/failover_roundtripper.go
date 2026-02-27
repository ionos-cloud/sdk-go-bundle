/*
 * IONOS Shared Libraries – Failover RoundTripper
 */

package shared

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
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

// NewFailoverRoundTripper wraps base in a FailoverRoundTripper controlled by
// cfg. When failover is inactive (nil config/Failover or strategy
// "none"/empty), RoundTrip passes through to base with no retry overhead.
func NewFailoverRoundTripper(cfg *Configuration, base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &FailoverRoundTripper{
		cfg:  cfg,
		base: base,
	}
}

// RoundTrip implements http.RoundTripper. It validates preconditions and
// dispatches to a strategy-specific method based on the configured
// FailoverStrategy:
//
//   - FailoverRoundRobin (with >1 server): delegates to orderedRoundTrip,
//     which cycles through servers sequentially.
//   - FailoverNone / empty / unknown: passes through to the base
//     transport with no retry logic.
//
// New strategies extend the switch in this method — each case builds its own
// server-selection function (or calls an entirely different method).
//
// This method is called by HTTPClient.Do() inside product-level callAPI. Each
// callAPI retry triggers a fresh RoundTrip invocation that cycles through all
// servers from the beginning.
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

	servers := t.cfg.Servers
	numServers := len(servers)

	switch normalizeStrategy(fo.Strategy) {
	case normalizeStrategy(FailoverRoundRobin):
		order := func(attempt int) int { return attempt % numServers }
		return t.orderedRoundTrip(req, fo, servers, order)
	default: // FailoverNone, "", or unknown
		return t.base.RoundTrip(req)
	}
}

// orderedRoundTrip is the shared retry loop for strategies that pick servers
// deterministically by attempt index. The order function maps each attempt
// (0, 1, 2, …) to a server index.
//
// Behaviour:
//   - Non-retryable HTTP methods (e.g. POST by default): passes through to the
//     base transport without retry.
//   - Network errors (connection refused, reset, etc.): retries with
//     exponential backoff, cycling to the next server via order. DNS errors
//     are never retried. Timeout errors are only retried when
//     FailoverOptions.RetryOnTimeout is set.
//   - Status codes in FailoverOnStatusCodes: drains the response body and
//     retries against the next server. Response headers (e.g. Retry-After) are
//     not inspected at this layer.
//   - All attempts exhausted: returns an error (not an HTTP response). This
//     means callAPI receives err != nil and returns immediately — its own
//     status-code-based retry logic (for 502/503/504/429) is never reached.
func (t *FailoverRoundTripper) orderedRoundTrip(
	req *http.Request,
	fo *FailoverOptions,
	servers ServerConfigurations,
	order serverOrder,
) (*http.Response, error) {
	if !isRetryableMethod(fo, req.Method) {
		return t.base.RoundTrip(req)
	}

	bo := fo.ExponentialBackoff.NewExponentialBackoff()

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

			backOff(attemptReq.Context(), bo.NextBackOff())
			continue
		}

		if !shouldFailoverOnStatus(fo, resp.StatusCode) {
			return resp, nil
		}

		if SdkLogLevel.Satisfies(Debug) {
			SdkLogger.Printf("[Failover] status=%d triggers failover to next server", resp.StatusCode)
		}
		// Drain/close body to allow connection reuse.
		drainBody(resp)
		backOff(attemptReq.Context(), bo.NextBackOff())
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

	// 1. DNS errors are non-retriable — hostname resolution failures
	// do not benefit from failover to a different server.
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return false
	}

	// 2. Check for transport-level errors (connection refused, reset, etc).
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return true
	}

	// 3. Handle url.Error cases not already caught by the typed checks above.
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

	// 4. Handle bare timeout errors if enabled.
	if retryOnTimeout && ctx != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return true
		}
	}

	return false
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

// normalizeStrategy returns the strategy in lower-case with surrounding
// whitespace removed, so comparisons are case-insensitive.
func normalizeStrategy(s FailoverStrategy) FailoverStrategy {
	return FailoverStrategy(strings.TrimSpace(strings.ToLower(string(s))))
}
