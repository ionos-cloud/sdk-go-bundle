package shared

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestDoWithApplicationRetry_RetryOn502(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusBadGateway, nil),
			makeResponse(http.StatusOK, nil),
		},
	}
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	body := bytes.NewBufferString("data")
	req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/test", body)
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString("data")), nil
	}

	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 2*time.Second)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 3, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	_, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
	client := &http.Client{Transport: rt}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.example.com/test", nil)
	// Long wait time so context cancels during backoff
	_, _, _ = DoWithApplicationRetry(client, req, 10, 5*time.Second, 5*time.Second)
	// We just verify it doesn't hang — context cancellation should break out of backoff
}

func TestDoWithApplicationRetry_ImmediateReturn_On200(t *testing.T) {
	rt := &testRoundTripper{
		responses: []*http.Response{
			makeResponse(http.StatusOK, nil),
		},
	}
	client := &http.Client{Transport: rt}

	req, _ := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	resp, _, err := DoWithApplicationRetry(client, req, 5, 1*time.Millisecond, 10*time.Millisecond)
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
