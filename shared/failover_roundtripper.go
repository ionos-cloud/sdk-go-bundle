/*
 * IONOS Shared Libraries – Failover RoundTripper
 */

package shared

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FailoverRoundTripper is an http.RoundTripper wrapper that retries the request
// against multiple configured servers when the underlying transport returns a
// network-level error.
//
// It is controlled via Configuration.FailoverStrategy.
//
// Notes:
//   - This wrapper only retries on transport/network errors (not HTTP 4xx/5xx).
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

// NewFailoverRoundTripper creates a new FailoverRoundTripper with the given configuration and base RoundTripper.
func NewFailoverRoundTripper(cfg *Configuration, base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &FailoverRoundTripper{
		cfg:  cfg,
		base: base,
	}
}

// RoundTrip - implements roundtrip failover logic based on the configured strategy and servers
func (t *FailoverRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	if t == nil || t.cfg == nil {
		if t != nil && t.base != nil {
			return t.base.RoundTrip(req)
		}
		return http.DefaultTransport.RoundTrip(req)
	}

	strategyName := strings.TrimSpace(strings.ToLower(string(t.cfg.FailoverStrategy)))
	servers := len(t.cfg.Servers)
	if strategyName == "" || strategyName == string(FailoverNone) || servers <= 1 {
		return t.base.RoundTrip(req)
	}

	if strategyName != strings.ToLower(string(FailoverRoundRobin)) {
		return t.base.RoundTrip(req)
	}

	// Check if method is allowed for failover retries.
	if !isRetryableMethod(t.cfg, req.Method) {
		return t.base.RoundTrip(req)
	}

	var lastErr error
	for i := range servers {
		// Always start from the first server in the list.
		serverURL := t.cfg.Servers[i].URL

		targetURL, err := url.Parse(serverURL)
		if err != nil {
			lastErr = err
			continue
		}

		attemptReq, err := cloneRequestForRetry(req)
		if err != nil {
			return nil, err
		}

		// Update both URL and the Host header field.
		attemptReq.URL.Scheme = targetURL.Scheme
		attemptReq.URL.Host = targetURL.Host
		attemptReq.Host = targetURL.Host

		resp, err := t.base.RoundTrip(attemptReq)
		if err == nil {
			return resp, nil
		}

		lastErr = err
		if !isNetworkErrorRT(attemptReq.Context(), err, t.cfg.RetryOnTimeout) {
			return nil, err
		}

		// Ensure we don't spin too hot in case of immediate failures.
		tinyBackoff(attemptReq.Context())
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

func isRetryableMethod(cfg *Configuration, method string) bool {
	m := strings.ToUpper(strings.TrimSpace(method))
	if cfg == nil {
		return defaultRetryableMethods[m]
	}

	// If not configured, use defaults.
	if len(cfg.RetryableMethods) == 0 {
		return defaultRetryableMethods[m]
	}

	for _, v := range cfg.RetryableMethods {
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

	// 1. Check for standard DNS resolution errors (typed).
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		if dnsErr.IsNotFound || dnsErr.IsTemporary || dnsErr.IsTimeout {
			return true
		}
	}

	// 2. Check for other transport-level errors (connection refused, reset, etc).
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return true
	}

	// 3. Fallback for wrapped url.Errors.
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		// String fallback for platforms where DNSError might not be perfectly populated.
		lowErr := strings.ToLower(urlErr.Error())
		if strings.Contains(lowErr, "no such host") || strings.Contains(lowErr, "connection refused") {
			return true
		}
		// If it's a URL error and we haven't matched a specific reason,
		// generally we treat transport-level problems as retryable.
		return true
	}

	// 4. Handle timeouts if enabled.
	if retryOnTimeout && ctx != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return true
		}
	}

	return false
}

// todo replace with a more robust backoff strategy exponential with jitter
func tinyBackoff(ctx context.Context) {
	// 10ms is enough to avoid busy-looping while staying responsive.
	t := 10 * time.Millisecond
	if ctx == nil {
		time.Sleep(t)
		return
	}
	select {
	case <-time.After(t):
	case <-ctx.Done():
	}
}
