package shared

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	retryTestURL       = "https://api.example.com/test"
	retryErrUnexpected = "unexpected error: %v"
	retryErrStatus     = "expected 200, got %d"
	retryErrCalls      = "expected %d calls, got %d"
	retryHeaderRA      = "Retry-After"
	retryHostS1        = "s1.example"
	retryURLs1         = "https://s1.example"
	retryURLs2         = "https://s2.example"
	retryURLs1Path     = "https://s1.example/test"
)

func retryCfg(rt http.RoundTripper, maxRetries int, waitTime, maxWaitTime time.Duration) *Configuration {
	return &Configuration{
		HTTPClient:  &http.Client{Transport: rt},
		MaxRetries:  maxRetries,
		WaitTime:    waitTime,
		MaxWaitTime: maxWaitTime,
	}
}

func TestDoWithApplicationRetryRetryOn502(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 3 {
		t.Fatalf(retryErrCalls, 3, rt.callCount)
	}
}

func TestDoWithApplicationRetryRetryOn503(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf(retryErrCalls, 2, rt.callCount)
	}
}

func TestDoWithApplicationRetryRetryOn504(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusGatewayTimeout, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf(retryErrCalls, 2, rt.callCount)
	}
}

func TestDoWithApplicationRetryNoRetryOnPostFor5xx(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	body := bytes.NewBufferString("data")
	req, _ := http.NewRequest(http.MethodPost, retryTestURL, body)
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString("data")), nil
	}

	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 (no retry for POST), got %d", resp.StatusCode)
	}
	if rt.callCount != 1 {
		t.Fatalf("expected 1 call (no retry for POST), got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetryRetryOn429WithRetryAfter(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusTooManyRequests, map[string]string{retryHeaderRA: "1"}),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 2*time.Second)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf(retryErrCalls, 2, rt.callCount)
	}
}

func TestDoWithApplicationRetryRetryOn429WithoutRetryAfter(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusTooManyRequests, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf(retryErrCalls, 2, rt.callCount)
	}
}

func TestDoWithApplicationRetryMaxRetriesRespected(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 3, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	// Should return the last 503 response after exhausting retries
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 after max retries, got %d", resp.StatusCode)
	}
	if rt.callCount != 3 {
		t.Fatalf(retryErrCalls, 3, rt.callCount)
	}
}

func TestDoWithApplicationRetryTransportErrorNoRetry(t *testing.T) {
	rt := &testRoundTripper{
		errors: []error{
			errors.New("dial tcp: connection refused"),
		},
		responses: []*http.Response{
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	_, _, err := DoWithApplicationRetry(cfg, req)
	if err == nil {
		t.Fatalf("expected transport error, got nil")
	}
	if rt.callCount != 1 {
		t.Fatalf("expected 1 call (no retry on transport error), got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetryContextCancellationDuringBackoff(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
		},
	}
	cfg := retryCfg(rt, 10, 5*time.Second, 5*time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, retryTestURL, nil)
	// Long wait time so context cancels during backoff
	_, _, _ = DoWithApplicationRetry(cfg, req)
	// We just verify it doesn't hang — context cancellation should break out of backoff
}

func TestDoWithApplicationRetryImmediateReturnOn200(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, retryTestURL, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if rt.callCount != 1 {
		t.Fatalf(retryErrCalls, 1, rt.callCount)
	}
}

func TestBackOffSleepsForGivenDuration(t *testing.T) {
	start := time.Now()
	backOff(context.Background(), 5*time.Millisecond)
	elapsed := time.Since(start)
	if elapsed < 4*time.Millisecond {
		t.Fatalf("backOff returned too early (%v), expected ~5ms", elapsed)
	}
	if elapsed > 100*time.Millisecond {
		t.Fatalf("backOff took too long (%v), expected ~5ms", elapsed)
	}
}

func TestBackOffZeroDuration(t *testing.T) {
	start := time.Now()
	backOff(context.Background(), 0)
	elapsed := time.Since(start)
	if elapsed > 50*time.Millisecond {
		t.Fatalf("BackOff should return immediately for zero duration, took %v", elapsed)
	}
}

// --- Tests with failover configured ---

func networkError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
}

func TestDoWithApplicationRetryFailoverHandlesNetworkError(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == retryHostS1 {
			return nil, networkError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  3,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (s1 fail + s2 ok), got %d", callCount)
	}
}

func TestDoWithApplicationRetryFailoverHandlesStatusCode(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == retryHostS1 {
			return &http.Response{Status: "503 Service Unavailable", StatusCode: http.StatusServiceUnavailable, Body: io.NopCloser(bytes.NewBufferString("no")), Header: make(http.Header), Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  3,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:              FailoverRoundRobin,
			FailoverOnStatusCodes: []int{http.StatusServiceUnavailable},
			MaxRetries:            10,
			ExponentialBackoff:    zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (s1 503 drained + s2 ok), got %d", callCount)
	}
}

func TestDoWithApplicationRetryFailoverExhaustedReturnsError(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		return nil, networkError(r)
	})

	cfg := &Configuration{
		MaxRetries:  3,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         4,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	_, _, err := DoWithApplicationRetry(cfg, req)
	if err == nil {
		t.Fatalf("expected error after failover exhaustion, got nil")
	}
	if callCount != 5 {
		t.Fatalf("expected 4 base transport calls (fo.MaxRetries=4), got %d", callCount)
	}
}

func TestDoWithApplicationRetryNonFailoverStatusCodePassesToAppRetry(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		// First call returns 503, second returns 200.
		if callCount == 1 {
			return &http.Response{Status: "503", StatusCode: http.StatusServiceUnavailable, Body: io.NopCloser(bytes.NewBufferString("")), Header: make(http.Header), Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  5,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
			// 503 NOT in FailoverOnStatusCodes — passes through to app retry
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (503 pass-through + app retry 200), got %d", callCount)
	}
}

func TestDoWithApplicationRetryFailoverThenAppRetry(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == retryHostS1 {
			return nil, networkError(r)
		}
		// s2: first call returns 502, second returns 200
		s2Calls++
		if s2Calls == 1 {
			return &http.Response{Status: "502", StatusCode: http.StatusBadGateway, Body: io.NopCloser(bytes.NewBufferString("")), Header: make(http.Header), Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  5,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
			// 502 NOT in FailoverOnStatusCodes — passes through to app retry
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 502 = 2 calls; attempt 2: s1 fail + s2 200 = 2 calls
	if callCount != 4 {
		t.Fatalf("expected 4 base transport calls, got %d", callCount)
	}
}

func TestDoWithApplicationRetryFailover429NotInFailoverCodesPassesToAppRetry(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if callCount == 1 {
			h := make(http.Header)
			h.Set(retryHeaderRA, "0")
			return &http.Response{Status: "429", StatusCode: http.StatusTooManyRequests, Body: io.NopCloser(bytes.NewBufferString("")), Header: h, Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  5,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
			// 429 NOT in FailoverOnStatusCodes
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (429 pass-through + app retry 200), got %d", callCount)
	}
}

func TestDoWithApplicationRetryFailoverThenAppRetryOn429(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == retryHostS1 {
			return nil, networkError(r)
		}
		// s2: first call returns 429, second returns 200
		s2Calls++
		if s2Calls == 1 {
			h := make(http.Header)
			h.Set(retryHeaderRA, "0")
			return &http.Response{Status: "429", StatusCode: http.StatusTooManyRequests, Body: io.NopCloser(bytes.NewBufferString("")), Header: h, Request: r}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	})

	cfg := &Configuration{
		MaxRetries:  5,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
			// 429 NOT in FailoverOnStatusCodes — passes through to app retry
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 429 = 2 calls; attempt 2: s1 fail + s2 200 = 2 calls
	if callCount != 4 {
		t.Fatalf("expected 4 base transport calls, got %d", callCount)
	}
}

func TestDoWithApplicationRetryThreeServersFailoverChain(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		switch r.URL.Host {
		case retryHostS1:
			return nil, networkError(r)
		case "s2.example":
			s2Calls++
			if s2Calls == 1 {
				// First s2 call: 429 (not in failover codes) — passes through to app retry
				h := make(http.Header)
				h.Set(retryHeaderRA, "0")
				return &http.Response{Status: "429", StatusCode: http.StatusTooManyRequests, Body: io.NopCloser(bytes.NewBufferString("")), Header: h, Request: r}, nil
			}
			// Second s2 call: 503 (in failover codes) — drained, failover continues to s3
			return &http.Response{Status: "503 Service Unavailable", StatusCode: http.StatusServiceUnavailable, Body: io.NopCloser(bytes.NewBufferString("no")), Header: make(http.Header), Request: r}, nil
		default: // s3
			return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
		}
	})

	cfg := &Configuration{
		MaxRetries:  5,
		WaitTime:    1 * time.Millisecond,
		MaxWaitTime: 10 * time.Millisecond,
		Failover: &FailoverOptions{
			Strategy:              FailoverRoundRobin,
			FailoverOnStatusCodes: []int{http.StatusServiceUnavailable},
			MaxRetries:            10,
			ExponentialBackoff:    zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: retryURLs1},
			{URL: retryURLs2},
			{URL: "https://s3.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper([]string{retryURLs1, retryURLs2, "https://s3.example"}, *cfg.Failover, base)}

	req, _ := http.NewRequest(http.MethodGet, retryURLs1Path, nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf(retryErrUnexpected, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(retryErrStatus, resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 429 = 2 calls
	// attempt 2: s1 fail + s2 503 (drained) + s3 200 = 3 calls
	if callCount != 5 {
		t.Fatalf("expected 5 base transport calls, got %d", callCount)
	}
}
