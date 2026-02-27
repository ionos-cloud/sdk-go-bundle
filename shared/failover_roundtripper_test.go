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
)

func TestFailoverRoundTripper_RoundRobin_NetworkError_FailsOverToNextServer(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			RetryOnTimeout:     false,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
		},
		MaxRetries: 10,
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	rt := NewFailoverRoundTripper(cfg, ft)

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/some/path?x=1", nil)
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 response, got %+v", resp)
	}

	if len(ft.calls) != 2 {
		t.Fatalf("expected 2 transport calls, got %d: %+v", len(ft.calls), ft.calls)
	}
	if ft.calls[0] != "s1.example" {
		t.Fatalf("expected first call to s1.example, got %q", ft.calls[0])
	}
	if ft.calls[1] != "s2.example" {
		t.Fatalf("expected second call to s2.example, got %q", ft.calls[1])
	}
}

func TestFailoverRoundTripper_RoundRobin_ConnectionReset_FailsOverToNextServer(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, connResetError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != "s1.example" || calls[1] != "s2.example" {
		t.Fatalf("expected [s1.example s2.example], got %v", calls)
	}
}

func TestFailoverRoundTripper_RoundRobin_IOTimeout_FailsOverToNextServer(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, ioTimeoutError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != "s1.example" || calls[1] != "s2.example" {
		t.Fatalf("expected [s1.example s2.example], got %v", calls)
	}
}

func TestFailoverRoundTripper_DoesNotRetry_WhenMethodNotRetryable(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			RetryableMethods:   []string{http.MethodGet},
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	rt := NewFailoverRoundTripper(cfg, ft)

	// POST is not retryable per config
	req, err := http.NewRequest(http.MethodPost, "https://s1.example/some/path", io.NopCloser(bytes.NewBufferString("x")))
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}
	// make body replayable for completeness
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString("x")), nil
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error on first server (no retry), got nil")
	}

	if len(ft.calls) != 1 {
		t.Fatalf("expected 1 transport call, got %d: %+v", len(ft.calls), ft.calls)
	}
	if ft.calls[0] != "s1.example" {
		t.Fatalf("expected call to s1.example, got %q", ft.calls[0])
	}
}

func TestFailoverRoundTripper_PostNotRetriedByDefault(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	rt := NewFailoverRoundTripper(cfg, ft)

	body := bytes.NewBufferString("data")
	req, err := http.NewRequest(http.MethodPost, "https://s1.example/path", io.NopCloser(body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString("data")), nil
	}

	// POST is not in defaultRetryableMethods, so it should pass through
	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error from first server, got nil")
	}
	if len(ft.calls) != 1 {
		t.Fatalf("expected 1 transport call (no retry for POST), got %d", len(ft.calls))
	}
}

func TestFailoverRoundTripper_FailoverOnStatusCodes(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:              FailoverRoundRobin,
			FailoverOnStatusCodes: []int{http.StatusServiceUnavailable},
			ExponentialBackoff:    zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return &http.Response{Status: "503 Service Unavailable", StatusCode: http.StatusServiceUnavailable, Body: io.NopCloser(bytes.NewBufferString("no")), Header: make(http.Header), Request: r}, nil
		}
		return &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/some/path", nil)
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != "s1.example" || calls[1] != "s2.example" {
		t.Fatalf("unexpected call order: %+v", calls)
	}
}

func TestFailoverRoundTripper_PassThrough_WhenFailoverDisabled(t *testing.T) {
	// Failover nil
	cfg := &Configuration{
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	// Should be exactly 1 call (no retry)
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (pass-through), got %d", len(calls))
	}
}

func TestFailoverRoundTripper_PassThrough_SingleServer(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (single server pass-through), got %d", len(calls))
	}
}

func TestFailoverRoundTripper_ContextCancellation(t *testing.T) {
	cfg := &Configuration{
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

	ctx, cancel := context.WithCancel(context.Background())
	callCount := 0
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		// Cancel context after first call
		cancel()
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
	}))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://s1.example/path", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestFailoverRoundTripper_DNSError_NotRetried(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			RetryOnTimeout:     false,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, dnsNotFoundError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/some/path", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for DNS error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripper_MaxRetriesExhausted(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			MaxRetries:         2,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	callCount := 0
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error after exhausting retries")
	}
	if callCount != 2 {
		t.Fatalf("expected 2 attempts (maxRetries=2), got %d", callCount)
	}
}

func TestFailoverRoundTripper_TLSCertificateError_NotRetried(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, tlsCertError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for TLS error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripper_RedirectError_NotRetried(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, redirectError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for redirect error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripper_DeadlineExceeded_NotRetried_WhenRetryOnTimeoutDisabled(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			RetryOnTimeout:     false,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, deadlineExceededError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover when RetryOnTimeout disabled), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripper_DeadlineExceeded_Retried_WhenRetryOnTimeoutEnabled(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:           FailoverRoundRobin,
			RetryOnTimeout:     true,
			MaxRetries:         10,
			ExponentialBackoff: zeroBackoff(),
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, deadlineExceededError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != "s1.example" || calls[1] != "s2.example" {
		t.Fatalf("expected [s1.example s2.example], got %v", calls)
	}
}

func TestFailoverRoundTripper_DNSTemporary_NotRetried(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, dnsTemporaryError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (DNS errors are never retried), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripper_ContextCanceled_NotRetried(t *testing.T) {
	cfg := &Configuration{
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

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == "s1.example" {
			return nil, context.Canceled
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, "https://s1.example/path", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error, got success")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (context cancellation is never retried), got %d: %v", len(calls), calls)
	}
}

