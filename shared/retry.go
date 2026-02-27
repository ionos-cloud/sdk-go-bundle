// IONOS Shared Libraries – Application-level retry

package shared

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

// DoWithApplicationRetry executes an HTTP request with application-level retry
// logic. This is the shared equivalent of the per-product callAPI method.
//
// Retry behaviour:
//   - 502/503/504: retries with backoff (skips POST)
//   - 429: retries honoring Retry-After header, falls back to waitTime
//   - Other status codes or transport errors: returns immediately
//   - Respects context cancellation during backoff
//
// # Interaction with FailoverRoundTripper
//
// client.Do invokes the configured http.RoundTripper, which may be a
// FailoverRoundTripper. When a multi-server FailoverStrategy is active,
// transport-level network errors and status codes listed in
// FailoverOnStatusCodes are handled by the round tripper before this
// function sees the response: it either receives a successful response or
// err != nil after all failover attempts are exhausted.
//
// Status codes NOT in FailoverOnStatusCodes (including the 502/503/504/429
// that this function retries on) pass through normally.
//
// Worst-case total attempts: Configuration.MaxRetries × FailoverOptions.MaxRetries.
func DoWithApplicationRetry(
	cfg *Configuration,
	request *http.Request,
) (*http.Response, time.Duration, error) {
	if cfg == nil {
		return nil, 0, errors.New("nil configuration")
	}

	var resp *http.Response
	var httpRequestTime time.Duration
	var err error

	for attempt := range cfg.MaxRetries {
		// We need to clone the request with every retry because Body closes after the request.
		clonedRequest, cloneErr := cloneRequestForRetry(request)
		if cloneErr != nil {
			return nil, httpRequestTime, cloneErr
		}

		logRequest(request, attempt+1)

		httpRequestStartTime := time.Now()
		clonedRequest.Close = true
		resp, err = cfg.HTTPClient.Do(clonedRequest)
		httpRequestTime = time.Since(httpRequestStartTime)
		if err != nil {
			return resp, httpRequestTime, err
		}

		logResponse(resp)

		var backoffTime time.Duration

		switch resp.StatusCode {
		case http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
			http.StatusBadGateway:
			if request.Method == http.MethodPost {
				return resp, httpRequestTime, err
			}
			backoffTime = cfg.WaitTime

		case http.StatusTooManyRequests:
			if retryAfterSeconds := resp.Header.Get("Retry-After"); retryAfterSeconds != "" {
				retryWait, parseErr := time.ParseDuration(retryAfterSeconds + "s")
				if parseErr != nil {
					return resp, httpRequestTime, parseErr
				}
				backoffTime = retryWait
			} else {
				backoffTime = cfg.WaitTime
			}
		default:
			return resp, httpRequestTime, err
		}

		if attempt == cfg.MaxRetries-1 {
			if SdkLogLevel.Satisfies(Debug) {
				SdkLogger.Printf(" Number of maximum retries exceeded (%d retries)\n", cfg.MaxRetries)
			}
			return resp, httpRequestTime, err
		}

		// Drain body before retrying so the connection can be reused.
		drainBody(resp)

		if backoffTime > cfg.MaxWaitTime {
			backoffTime = cfg.MaxWaitTime
		}
		backOff(request.Context(), backoffTime)
	}

	return resp, httpRequestTime, err
}

// drainBody discards and closes the response body so the underlying
// connection can be returned to the pool.
func drainBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

// logRequest dumps the outgoing request at Debug level.
// The Authorization header is stripped unless Trace is enabled.
func logRequest(request *http.Request, retryCount int) {
	if !SdkLogLevel.Satisfies(Debug) {
		return
	}
	logReq := request.Clone(request.Context())
	if !SdkLogLevel.Satisfies(Trace) {
		logReq.Header.Del("Authorization")
	}
	dump, err := httputil.DumpRequestOut(logReq, true)
	if err == nil {
		SdkLogger.Printf(" DumpRequestOut : %s\n", string(dump))
	} else {
		SdkLogger.Printf(" DumpRequestOut err: %+v", err)
	}
	SdkLogger.Printf("\n try no: %d\n", retryCount)
}

// logResponse dumps the server response at Debug level.
func logResponse(resp *http.Response) {
	if !SdkLogLevel.Satisfies(Debug) {
		return
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		SdkLogger.Printf("\n DumpResponse : %s\n", string(dump))
	} else {
		SdkLogger.Printf(" DumpResponse err %+v", err)
	}
}

// backOff sleeps for the given duration and respects context cancellation.
func backOff(ctx context.Context, t time.Duration) {
	if SdkLogLevel.Satisfies(Debug) {
		SdkLogger.Printf(" Sleeping %s before retrying request\n", t.String())
	}
	if t <= 0 {
		return
	}
	timer := time.NewTimer(t)
	defer timer.Stop()

	select {
	case <-ctx.Done():
	case <-timer.C:
	}
}
