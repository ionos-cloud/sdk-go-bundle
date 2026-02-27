// IONOS Shared Libraries – Application-level retry

package shared

import (
	"context"
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
func DoWithApplicationRetry(
	client *http.Client,
	request *http.Request,
	maxRetries int,
	waitTime time.Duration,
	maxWaitTime time.Duration,
) (*http.Response, time.Duration, error) {
	var resp *http.Response
	var httpRequestTime time.Duration
	var err error

	for attempt := range maxRetries {
		// We need to clone the request with every retry because Body closes after the request.
		clonedRequest, cloneErr := cloneRequestForRetry(request)
		if cloneErr != nil {
			return nil, httpRequestTime, cloneErr
		}

		logRequest(request, attempt+1)

		httpRequestStartTime := time.Now()
		clonedRequest.Close = true
		resp, err = client.Do(clonedRequest)
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
			backoffTime = waitTime

		case http.StatusTooManyRequests:
			if retryAfterSeconds := resp.Header.Get("Retry-After"); retryAfterSeconds != "" {
				retryWait, parseErr := time.ParseDuration(retryAfterSeconds + "s")
				if parseErr != nil {
					return resp, httpRequestTime, parseErr
				}
				backoffTime = retryWait
			} else {
				backoffTime = waitTime
			}
		default:
			return resp, httpRequestTime, err
		}

		if attempt == maxRetries-1 {
			if SdkLogLevel.Satisfies(Debug) {
				SdkLogger.Printf(" Number of maximum retries exceeded (%d retries)\n", maxRetries)
			}
			return resp, httpRequestTime, err
		}

		// Drain body before retrying so the connection can be reused.
		drainBody(resp)

		if backoffTime > maxWaitTime {
			backoffTime = maxWaitTime
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
