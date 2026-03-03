package failover

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

const (
	foURLs1         = "https://s1.example"
	foURLs2         = "https://s2.example"
	foURLs1Path     = "https://s1.example/path"
	foURLs1SomePath = "https://s1.example/some/path"
	foHostS1        = "s1.example"
	foHostS2        = "s2.example"
	foErrUnexpReq   = "unexpected error creating request: %v"
	foErrExpSucc    = "expected success, got error: %v"
	foErrExp200     = "expected 200, got %d"
	foErrExpErr     = "expected error, got success"
	foErrUnexp      = "unexpected error: %v"
	foErrExpOrder   = "expected [s1.example s2.example], got %v"
)

func TestFailoverRoundTripperRoundRobinNetworkErrorFailsOverToNextServer(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		RetryOnTimeout:     false,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	ft := &fakeTransport{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, ft)

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/some/path?x=1", nil)
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 response, got %+v", resp)
	}

	if len(ft.calls) != 2 {
		t.Fatalf("expected 2 transport calls, got %d: %+v", len(ft.calls), ft.calls)
	}
	if ft.calls[0] != foHostS1 {
		t.Fatalf("expected first call to s1.example, got %q", ft.calls[0])
	}
	if ft.calls[1] != foHostS2 {
		t.Fatalf("expected second call to s2.example, got %q", ft.calls[1])
	}
}

func TestFailoverRoundTripperRoundRobinConnectionResetFailsOverToNextServer(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, connResetError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != foHostS1 || calls[1] != foHostS2 {
		t.Fatalf(foErrExpOrder, calls)
	}
}

func TestFailoverRoundTripperRoundRobinIOTimeoutFailsOverToNextServer(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, ioTimeoutError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != foHostS1 || calls[1] != foHostS2 {
		t.Fatalf(foErrExpOrder, calls)
	}
}

func TestFailoverRoundTripperDoesNotRetryWhenMethodNotRetryable(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		RetryableMethods:   []string{http.MethodGet},
		ExponentialBackoff: zeroBackoff(),
	}

	ft := &fakeTransport{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, ft)

	// POST is not retryable per config
	req, err := http.NewRequest(http.MethodPost, foURLs1SomePath, io.NopCloser(bytes.NewBufferString("x")))
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
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
	if ft.calls[0] != foHostS1 {
		t.Fatalf("expected call to s1.example, got %q", ft.calls[0])
	}
}

func TestFailoverRoundTripperPostNotRetriedByDefault(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		ExponentialBackoff: zeroBackoff(),
	}

	ft := &fakeTransport{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, ft)

	body := bytes.NewBufferString("data")
	req, err := http.NewRequest(http.MethodPost, foURLs1Path, io.NopCloser(body))
	if err != nil {
		t.Fatalf(foErrUnexp, err)
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

func TestFailoverRoundTripperFailoverOnStatusCodes(t *testing.T) {
	fo := Options{
		Strategy:              RoundRobin,
		FailoverOnStatusCodes: []int{http.StatusServiceUnavailable},
		ExponentialBackoff:    zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return &http.Response{Status: "503 Service Unavailable", StatusCode: http.StatusServiceUnavailable, Body: io.NopCloser(bytes.NewBufferString("no")), Header: make(http.Header), Request: r}, nil
		}
		return &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, foURLs1SomePath, nil)
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != foHostS1 || calls[1] != foHostS2 {
		t.Fatalf("unexpected call order: %+v", calls)
	}
}

func TestFailoverRoundTripperPassThroughWhenFailoverDisabled(t *testing.T) {
	// Empty strategy disables failover.
	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, Options{}, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	if err != nil {
		t.Fatalf(foErrUnexp, err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	// Should be exactly 1 call (no retry)
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (pass-through), got %d", len(calls))
	}
}

func TestFailoverRoundTripperPassThroughSingleServer(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	if err != nil {
		t.Fatalf(foErrUnexp, err)
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (single server pass-through), got %d", len(calls))
	}
}

func TestFailoverRoundTripperContextCancellation(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	callCount := 0
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		// Cancel context after first call
		cancel()
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
	}))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, foURLs1Path, nil)
	if err != nil {
		t.Fatalf(foErrUnexp, err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
}

func TestFailoverRoundTripperDNSErrorNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		RetryOnTimeout:     false,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, dnsNotFoundError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, foURLs1SomePath, nil)
	if err != nil {
		t.Fatalf(foErrUnexp, err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for DNS error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperMaxRetriesExhausted(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         2,
		ExponentialBackoff: zeroBackoff(),
	}

	callCount := 0
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		callCount++
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
	}))

	req, err := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	if err != nil {
		t.Fatalf(foErrUnexp, err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error after exhausting retries")
	}
	if callCount != 3 {
		t.Fatalf("expected 2 attempts (maxRetries=2), got %d", callCount)
	}
}

func TestFailoverRoundTripperTLSCertificateErrorNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, tlsCertError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for TLS error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperRedirectErrorNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, redirectError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for redirect error), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperDeadlineExceededNotRetriedWhenRetryOnTimeoutDisabled(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		RetryOnTimeout:     false,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, deadlineExceededError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover when RetryOnTimeout disabled), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperDeadlineExceededRetriedWhenRetryOnTimeoutEnabled(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		RetryOnTimeout:     true,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, deadlineExceededError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	if len(calls) != 2 || calls[0] != foHostS1 || calls[1] != foHostS2 {
		t.Fatalf(foErrExpOrder, calls)
	}
}

func TestFailoverRoundTripperDNSTemporaryNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, dnsTemporaryError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (DNS errors are never retried), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperContextCanceledNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, context.Canceled
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (context cancellation is never retried), got %d: %v", len(calls), calls)
	}
}
