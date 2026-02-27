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

func retryCfg(rt http.RoundTripper, maxRetries int, waitTime, maxWaitTime time.Duration) *Configuration {
	return &Configuration{
		HTTPClient:  &http.Client{Transport: rt},
		MaxRetries:  maxRetries,
		WaitTime:    waitTime,
		MaxWaitTime: maxWaitTime,
	}
}

func TestDoWithApplicationRetry_RetryOn502(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 3 {
		t.Fatalf("expected 3 calls, got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_RetryOn503(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf("expected 2 calls, got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_RetryOn504(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusGatewayTimeout, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf("expected 2 calls, got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_NoRetryOnPostFor5xx(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	body := bytes.NewBufferString("data")
	req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/test", body)
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString("data")), nil
	}

	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 (no retry for POST), got %d", resp.StatusCode)
	}
	if rt.callCount != 1 {
		t.Fatalf("expected 1 call (no retry for POST), got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_RetryOn429WithRetryAfter(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "1"}),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 2*time.Second)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf("expected 2 calls, got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_RetryOn429WithoutRetryAfter(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusTooManyRequests, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 2 {
		t.Fatalf("expected 2 calls, got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_MaxRetriesRespected(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusServiceUnavailable, nil),
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 3, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should return the last 503 response after exhausting retries
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 after max retries, got %d", resp.StatusCode)
	}
	if rt.callCount != 3 {
		t.Fatalf("expected 3 calls (maxRetries=3), got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_TransportError_NoRetry(t *testing.T) {
	rt := &testRoundTripper{
		errors: []error{
			errors.New("dial tcp: connection refused"),
		},
		responses: []*http.Response{
			makeResponse(http.StatusOK, nil), // should never be reached
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	_, _, err := DoWithApplicationRetry(cfg, req)
	if err == nil {
		t.Fatalf("expected transport error, got nil")
	}
	if rt.callCount != 1 {
		t.Fatalf("expected 1 call (no retry on transport error), got %d", rt.callCount)
	}
}

func TestDoWithApplicationRetry_ContextCancellationDuringBackoff(t *testing.T) {
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

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.example.com/test", nil)
	// Long wait time so context cancels during backoff
	_, _, _ = DoWithApplicationRetry(cfg, req)
	// We just verify it doesn't hang — context cancellation should break out of backoff
}

func TestDoWithApplicationRetry_ImmediateReturn_On200(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusOK, nil),
		},
	}
	cfg := retryCfg(rt, 5, 1*time.Millisecond, 10*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rt.callCount != 1 {
		t.Fatalf("expected 1 call, got %d", rt.callCount)
	}
}

func Test_backOff_SleepsForGivenDuration(t *testing.T) {
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

func Test_backOff_ZeroDuration(t *testing.T) {
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

func TestDoWithApplicationRetry_FailoverHandlesNetworkError(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == "s1.example" {
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (s1 fail + s2 ok), got %d", callCount)
	}
}

func TestDoWithApplicationRetry_FailoverHandlesStatusCode(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == "s1.example" {
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (s1 503 drained + s2 ok), got %d", callCount)
	}
}

func TestDoWithApplicationRetry_FailoverExhausted_ReturnsError(t *testing.T) {
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	_, _, err := DoWithApplicationRetry(cfg, req)
	if err == nil {
		t.Fatalf("expected error after failover exhaustion, got nil")
	}
	if callCount != 4 {
		t.Fatalf("expected 4 base transport calls (fo.MaxRetries=4), got %d", callCount)
	}
}

func TestDoWithApplicationRetry_NonFailoverStatusCode_PassesToAppRetry(t *testing.T) {
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (503 pass-through + app retry 200), got %d", callCount)
	}
}

func TestDoWithApplicationRetry_FailoverThenAppRetry(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == "s1.example" {
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 502 = 2 calls; attempt 2: s1 fail + s2 200 = 2 calls
	if callCount != 4 {
		t.Fatalf("expected 4 base transport calls, got %d", callCount)
	}
}

func TestDoWithApplicationRetry_Failover429NotInFailoverCodes_PassesToAppRetry(t *testing.T) {
	callCount := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if callCount == 1 {
			h := make(http.Header)
			h.Set("Retry-After", "0")
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 base transport calls (429 pass-through + app retry 200), got %d", callCount)
	}
}

func TestDoWithApplicationRetry_FailoverThenAppRetryOn429(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		if r.URL.Host == "s1.example" {
			return nil, networkError(r)
		}
		// s2: first call returns 429, second returns 200
		s2Calls++
		if s2Calls == 1 {
			h := make(http.Header)
			h.Set("Retry-After", "0")
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 429 = 2 calls; attempt 2: s1 fail + s2 200 = 2 calls
	if callCount != 4 {
		t.Fatalf("expected 4 base transport calls, got %d", callCount)
	}
}

func TestDoWithApplicationRetry_ThreeServers_FailoverChain(t *testing.T) {
	callCount := 0
	s2Calls := 0
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		switch r.URL.Host {
		case "s1.example":
			return nil, networkError(r)
		case "s2.example":
			s2Calls++
			if s2Calls == 1 {
				// First s2 call: 429 (not in failover codes) — passes through to app retry
				h := make(http.Header)
				h.Set("Retry-After", "0")
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
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
			{URL: "https://s3.example"},
		},
	}
	cfg.HTTPClient = &http.Client{Transport: NewFailoverRoundTripper(cfg, base)}

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/test", nil)
	resp, _, err := DoWithApplicationRetry(cfg, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	// attempt 1: s1 fail + s2 429 = 2 calls
	// attempt 2: s1 fail + s2 503 (drained) + s3 200 = 3 calls
	if callCount != 5 {
		t.Fatalf("expected 5 base transport calls, got %d", callCount)
	}
}
