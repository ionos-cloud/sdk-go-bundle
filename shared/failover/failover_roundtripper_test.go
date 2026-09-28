package failover

import (
	"bytes"
	"context"
	"encoding/pem"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"syscall"
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
		RetryOnTimeout:     true,
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

// A non-retryable method (POST) must be sent through the first endpoint's own
// transport, so its SkipTLSVerify setting is honored. The default base transport
// rejects the server's self-signed certificate, so success proves the endpoint
// transport was used rather than defaultBase.
func TestFailoverRoundTripperNonRetryableMethodUsesSkipTLSTransport(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	fo := Options{Strategy: RoundRobin, ExponentialBackoff: zeroBackoff()}
	rt := NewRoundTripper([]Endpoint{{URL: server.URL, SkipTLSVerify: true}}, fo, http.DefaultTransport)

	req, err := http.NewRequest(http.MethodPost, server.URL+"/path", io.NopCloser(bytes.NewBufferString("data")))
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}
	req.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewBufferString("data")), nil }

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	_ = resp.Body.Close()
}

// Same guarantee via a custom CA: the first endpoint's transport trusts the
// server certificate through CertificateAuthData, while defaultBase does not.
func TestFailoverRoundTripperNonRetryableMethodUsesCustomCATransport(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	caPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw})

	fo := Options{Strategy: RoundRobin, ExponentialBackoff: zeroBackoff()}
	rt := NewRoundTripper([]Endpoint{{URL: server.URL, CertificateAuthData: string(caPEM)}}, fo, http.DefaultTransport)

	req, err := http.NewRequest(http.MethodPost, server.URL+"/path", io.NopCloser(bytes.NewBufferString("data")))
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}
	req.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewBufferString("data")), nil }

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	_ = resp.Body.Close()
}

// A non-retryable request with a non-replayable body (GetBody == nil) must still be
// sent once. The single-attempt path passes the request through unmodified, so it must
// not require replayability the way the retry loop does.
func TestFailoverRoundTripperNonRetryableMethodAllowsNonReplayableBody(t *testing.T) {
	fo := Options{Strategy: RoundRobin, ExponentialBackoff: zeroBackoff()}

	var called bool
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		called = true
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header)}, nil
	})
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, base)

	// POST is not retryable by default; give it a body with no GetBody (non-replayable).
	req, err := http.NewRequest(http.MethodPost, foURLs1Path, io.NopCloser(bytes.NewBufferString("data")))
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}
	req.GetBody = nil

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected non-replayable POST to pass through, got error: %v", err)
	}
	if !called {
		t.Fatal("expected the transport to be called")
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(foErrExp200, resp.StatusCode)
	}
	_ = resp.Body.Close()
}

// A non-retryable request whose incoming host is not one of the endpoints must be
// retargeted to the first configured endpoint, just like a failover attempt.
func TestFailoverRoundTripperNonRetryableMethodRetargetsToFirstEndpoint(t *testing.T) {
	fo := Options{Strategy: RoundRobin, ExponentialBackoff: zeroBackoff()}

	var gotHost string
	base := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		gotHost = r.URL.Host
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header)}, nil
	})
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, base)

	// Incoming request points at a host that is NOT one of the endpoints.
	req, err := http.NewRequest(http.MethodPost, "https://original.example/path", io.NopCloser(bytes.NewBufferString("data")))
	if err != nil {
		t.Fatalf(foErrUnexpReq, err)
	}
	req.GetBody = nil

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf(foErrExpSucc, err)
	}
	if gotHost != foHostS1 {
		t.Fatalf("expected request retargeted to %s, got %s", foHostS1, gotHost)
	}
	_ = resp.Body.Close()
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
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: syscall.ECONNREFUSED}}
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
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: &net.OpError{Op: "dial", Net: "tcp", Err: syscall.ECONNREFUSED}}
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
		t.Fatalf("expected 3 attempts (maxRetries=2), got %d", callCount)
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

func TestFailoverRoundTripperTLSHandshakeErrorNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, tlsHandshakeError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (no failover for TLS handshake error), got %d: %v", len(calls), calls)
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

func TestFailoverRoundTripperDeadlineExceededNotRetriedEvenWhenRetryOnTimeoutEnabled(t *testing.T) {
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
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (only net.OpError is retryable), got %d: %v", len(calls), calls)
	}
}

func TestFailoverRoundTripperEOFNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, io.EOF
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (EOF is not retryable), got %d: %v", len(calls), calls)
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

func TestFailoverRoundTripperUnknownNetOpErrorNotRetried(t *testing.T) {
	fo := Options{
		Strategy:           RoundRobin,
		MaxRetries:         10,
		ExponentialBackoff: zeroBackoff(),
	}

	calls := []string{}
	rt := NewRoundTripper([]Endpoint{{URL: foURLs1}, {URL: foURLs2}}, fo, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		calls = append(calls, r.URL.Host)
		if r.URL.Host == foHostS1 {
			return nil, unknownNetOpError(r)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, _ := http.NewRequest(http.MethodGet, foURLs1Path, nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf(foErrExpErr)
	}
	if len(calls) != 1 {
		t.Fatalf("expected 1 call (unknown net.OpError is not retryable), got %d: %v", len(calls), calls)
	}
}
