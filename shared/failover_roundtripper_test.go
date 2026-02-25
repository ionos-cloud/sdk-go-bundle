package shared

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type fakeTransport struct {
	calls []string
}

func (f *fakeTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// record host we tried
	host := ""
	urlStr := ""
	if r != nil && r.URL != nil {
		host = r.URL.Host
		urlStr = r.URL.String()
	}
	f.calls = append(f.calls, host)

	// fail on first host, succeed on second
	if host == "s1.example" {
		return nil, &url.Error{Op: "Get", URL: urlStr, Err: errors.New("dial tcp: i/o timeout")}
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("ok")),
		Header:     make(http.Header),
		Request:    r,
	}, nil
}

func TestFailoverRoundTripper_RoundRobin_NetworkError_FailsOverToNextServer(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:       FailoverRoundRobin,
			RetryOnTimeout: false,
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	rt := NewFailoverRoundTripper(cfg, cfg.Failover, ft)

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

func TestFailoverRoundTripper_DoesNotRetry_WhenMethodNotRetryable(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:         FailoverRoundRobin,
			RetryableMethods: []string{http.MethodGet},
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	rt := NewFailoverRoundTripper(cfg, cfg.Failover, ft)

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

func TestFailoverRoundTripper_RetriesOnNoSuchHost(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:       FailoverRoundRobin,
			RetryOnTimeout: false,
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	ft := &fakeTransport{}
	// Override fake transport behavior: treat first host as "no such host".
	ft2 := &struct{ *fakeTransport }{ft}
	_ = ft2

	// Inline transport to simulate DNS error.
	rt := NewFailoverRoundTripper(cfg, cfg.Failover, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		host := ""
		urlStr := ""
		if r != nil && r.URL != nil {
			host = r.URL.Host
			urlStr = r.URL.String()
		}
		ft.calls = append(ft.calls, host)
		if host == "s1.example" {
			return nil, &url.Error{Op: "Get", URL: urlStr, Err: errors.New("lookup s1.example: no such host")}
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header), Request: r}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://s1.example/some/path", nil)
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
		t.Fatalf("expected 2 calls due to failover, got %d: %+v", len(ft.calls), ft.calls)
	}
	if ft.calls[0] != "s1.example" || ft.calls[1] != "s2.example" {
		t.Fatalf("unexpected call order: %+v", ft.calls)
	}
}

func TestFailoverRoundTripper_FailoverOnStatusCodes(t *testing.T) {
	cfg := &Configuration{
		Failover: &FailoverOptions{
			Strategy:              FailoverRoundRobin,
			FailoverOnStatusCodes: []int{http.StatusServiceUnavailable},
		},
		Servers: ServerConfigurations{
			{URL: "https://s1.example"},
			{URL: "https://s2.example"},
		},
	}

	calls := []string{}
	rt := NewFailoverRoundTripper(cfg, cfg.Failover, roundTripperFunc(func(r *http.Request) (*http.Response, error) {
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

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
