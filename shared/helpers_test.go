package shared

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// roundTripperFunc adapts a function to the http.RoundTripper interface.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// fakeTransport records call hosts and fails on s1.example with a network error.
type fakeTransport struct {
	calls []string
}

func (f *fakeTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	host := ""
	urlStr := ""
	if r != nil && r.URL != nil {
		host = r.URL.Host
		urlStr = r.URL.String()
	}
	f.calls = append(f.calls, host)

	if host == "s1.example" {
		return nil, &url.Error{Op: "Get", URL: urlStr, Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("i/o timeout")}}
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("ok")),
		Header:     make(http.Header),
		Request:    r,
	}, nil
}

// testRoundTripper is a configurable RoundTripper for testing DoWithApplicationRetry.
type testRoundTripper struct {
	responses []*http.Response
	errors    []error
	callCount int
}

func (t *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	idx := t.callCount
	t.callCount++
	if idx < len(t.errors) && t.errors[idx] != nil {
		return nil, t.errors[idx]
	}
	if idx < len(t.responses) {
		return t.responses[idx], nil
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString("ok")), Header: make(http.Header)}, nil
}

func makeResponse(statusCode int, headers map[string]string) *http.Response {
	h := make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return &http.Response{
		StatusCode: statusCode,
		Status:     strconv.Itoa(statusCode),
		Body:       io.NopCloser(bytes.NewBufferString("")),
		Header:     h,
	}
}

func connRefusedError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}}
}

func connResetError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: &net.OpError{Op: "read", Net: "tcp", Err: errors.New("connection reset by peer")}}
}

func ioTimeoutError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("i/o timeout")}}
}

func dnsTemporaryError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: &net.OpError{Op: "dial", Net: "tcp",
			Err: &net.DNSError{Err: "server misbehaving", Name: r.URL.Host, IsTemporary: true}}}
}

func tlsCertError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: errors.New("tls: failed to verify certificate")}
}

func redirectError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: errors.New("stopped after 10 redirects")}
}

func deadlineExceededError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: context.DeadlineExceeded}
}

func dnsNotFoundError(r *http.Request) error {
	return &url.Error{Op: "Get", URL: r.URL.String(),
		Err: &net.OpError{Op: "dial", Net: "tcp",
			Err: &net.DNSError{Err: "no such host", Name: r.URL.Host, IsNotFound: true}}}
}

func zeroBackoff() *ExponentialBackoffOptions {
	mult := float64(1)
	jitter := float64(0)
	return &ExponentialBackoffOptions{
		InitialInterval:     1 * time.Nanosecond,
		MaxInterval:         1 * time.Nanosecond,
		Multiplier:          &mult,
		RandomizationFactor: &jitter,
	}
}
