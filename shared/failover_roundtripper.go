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
// Notes:
//   - Network errors trigger retries with exponential backoff, cycling to the
//     next server.
//   - Status codes in FailoverOnStatusCodes also trigger retries: the response
//     body is drained and the request is sent to the next server. Response
//     headers (e.g. Retry-After) are not inspected at this layer.
//   - If the request carries a body, the request must have GetBody set (the SDK
//     generates requests in a way that supports this).
//   - For non-idempotent requests (POST, PATCH, etc.), enabling failover on
//     timeouts can produce duplicates when RetryOnTimeout is set. Context
//     cancellation always stops retries immediately.
//
// The request URL is rewritten by swapping scheme/host with each server URL,
// preserving path and query.
//
// The server URLs and options are snapshotted at construction time; subsequent
// changes to the original slices or struct do not affect this instance.
type FailoverRoundTripper struct {
	serverURLs []string
	fo         FailoverOptions
	base       http.RoundTripper
}

// NewFailoverRoundTripper creates a FailoverRoundTripper that cycles through
// serverURLs according to fo when a request fails. All slices are copied so
// callers can safely mutate the originals. When the strategy is "none" or
// empty, RoundTrip passes through to base with no retry overhead.
func NewFailoverRoundTripper(serverURLs []string, fo FailoverOptions, base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}

	// Snapshot slices so external mutations do not affect this instance.
	urlsCopy := make([]string, len(serverURLs))
	copy(urlsCopy, serverURLs)

	methodsCopy := make([]string, len(fo.RetryableMethods))
	copy(methodsCopy, fo.RetryableMethods)
	fo.RetryableMethods = methodsCopy

	statusCodesCopy := make([]int, len(fo.FailoverOnStatusCodes))
	copy(statusCodesCopy, fo.FailoverOnStatusCodes)
	fo.FailoverOnStatusCodes = statusCodesCopy

	return &FailoverRoundTripper{
		serverURLs: urlsCopy,
		fo:         fo,
		base:       base,
	}
}

// RoundTrip implements http.RoundTripper. It validates preconditions and
// dispatches to a strategy-specific method based on the configured
// FailoverStrategy:
//
//   - FailoverRoundRobin: delegates to orderedRoundTrip, which cycles through
//     servers sequentially.
//   - FailoverNone / empty: passes through to the base
//     transport with no retry logic.
//   - Unknown strategy: returns an error immediately (configuration error).
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

	numServers := len(t.serverURLs)

	switch normalizeStrategy(t.fo.Strategy) {
	case normalizeStrategy(FailoverRoundRobin):
		if numServers == 0 {
			return t.base.RoundTrip(req)
		}
		order := func(attempt int) int { return attempt % numServers }
		return t.orderedRoundTrip(req, order)
	case FailoverNone, "": // Explicit "none" or empty strategy disables failover and passes through to base.
		return t.base.RoundTrip(req)
	default: // unknown strategy is a configuration error that we surface immediately.
		return nil, fmt.Errorf("unknown failover strategy: %s", t.fo.Strategy)
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
//     and context cancellation are never retried. Timeout errors are only
//     retried when FailoverOptions.RetryOnTimeout is set.
//   - Status codes in FailoverOnStatusCodes: drains the response body and
//     retries against the next server. Response headers (e.g. Retry-After) are
//     not inspected at this layer.
//   - All attempts exhausted: returns an error (not an HTTP response). This
//     means callAPI receives err != nil and returns immediately — its own
//     status-code-based retry logic (for 502/503/504/429) is never reached.
func (t *FailoverRoundTripper) orderedRoundTrip(req *http.Request, order serverOrder) (*http.Response, error) {
	fo := &t.fo

	if !isRetryableMethod(fo, req.Method) {
		return t.base.RoundTrip(req)
	}

	bo := fo.ExponentialBackoff.NewExponentialBackoff()

	maxRetries := fo.MaxRetries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}
	// maxRetries in config means "number of retries", so total attempts = maxRetries + 1
	var lastErr error
	totalAttempts := maxRetries + 1
	for attempt := range totalAttempts {
		serverIndex := order(attempt)
		serverURL := t.serverURLs[serverIndex]

		LogDebug("[Failover] attempt=%d, serverIndex=%d, serverURL=%s, method=%s", attempt, serverIndex, serverURL, req.Method)

		resp, err := t.doFailoverAttempt(req, serverURL)
		if err != nil {
			if !isNetworkErrorRT(req.Context(), err, fo.RetryOnTimeout) {
				LogDebug("[Failover] attempt=%d failed with non-retriable error on Servers[%d]: %v", attempt, serverIndex, err)
				return nil, err
			}

			LogDebug("[Failover] attempt=%d failed with retriable error on Servers[%d]: %v", attempt, serverIndex, err)

			lastErr = err
			// Don't sleep if this was the last attempt
			if attempt < totalAttempts-1 {
				backOff(req.Context(), bo.NextBackOff())
			}
			continue
		}

		if !shouldFailoverOnStatus(fo, resp.StatusCode) {
			LogDebug("[Failover] attempt=%d ends failover loop with status=%d on Servers[%d]=%s", attempt, resp.StatusCode, serverIndex, serverURL)
			return resp, nil
		}

		LogDebug("[Failover] attempt=%d, status=%d triggers failover to next server", attempt, resp.StatusCode)
		drainBody(resp)
		// Don't sleep if this was the last attempt
		if attempt < totalAttempts-1 {
			backOff(req.Context(), bo.NextBackOff())
		}
		lastErr = fmt.Errorf("failover status: %s", resp.Status)
	}

	return nil, lastErr
}

// doFailoverAttempt prepares and executes a single failover attempt against
// the given server URL.
func (t *FailoverRoundTripper) doFailoverAttempt(req *http.Request, serverURL string) (*http.Response, error) {
	targetURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, fmt.Errorf("invalid server URL %q: %w", serverURL, err)
	}

	attemptReq, err := cloneRequestForRetry(req)
	if err != nil {
		return nil, fmt.Errorf("failed to clone request for retry: %w", err)
	}

	attemptReq.URL.Scheme = targetURL.Scheme
	attemptReq.URL.Host = targetURL.Host
	attemptReq.Host = targetURL.Host

	LogDebug("[Failover] method=%s url=%s", attemptReq.Method, attemptReq.URL.String())
	return t.base.RoundTrip(attemptReq)
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

	// 4. Handle context cancellation as non-retriable. This is a safety measure
	// to prevent failover retries from continuing indefinitely after the caller has
	// given up.
	if ctx != nil && errors.Is(err, context.Canceled) {
		return false
	}

	// 5. Handle bare timeout errors if enabled.
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
