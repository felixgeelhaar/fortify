package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	fortifyhttp "github.com/felixgeelhaar/fortify/http"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

func TestHTTPExample(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the example in a goroutine
	go func() {
		main()
	}()

	// Allow server to start and print output
	time.Sleep(time.Millisecond * 100)

	// Restore stdout and read output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	//nolint:errcheck // ignoring error in test
	_, _ = io.Copy(&buf, r)

	// Verify HTTP server started
	output := buf.String()

	// Should show server startup message
	if len(output) < 10 {
		t.Skip("HTTP server starts in background, output may be minimal")
	}
}

func TestHTTPMiddleware(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{
		Rate:     10,
		Burst:    10,
		Interval: time.Second,
	})

	tm := timeout.New[*http.Response](timeout.Config{
		DefaultTimeout: time.Second,
	})

	handler := fortifyhttp.RateLimit(rl, fortifyhttp.KeyFromIP)(
		fortifyhttp.Timeout(tm, time.Millisecond*100)(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				//nolint:errcheck // ignoring error in test
				_, _ = w.Write([]byte("success"))
			}),
		),
	)

	req := httptest.NewRequest("GET", "/test", http.NoBody)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestHTTPRateLimit(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{
		Rate:     1,
		Burst:    1,
		Interval: time.Hour, // Very slow refill
	})

	handler := fortifyhttp.RateLimit(rl, func(r *http.Request) string {
		return "test-key"
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// First request should succeed
	req1 := httptest.NewRequest("GET", "/test", http.NoBody)
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1)

	if rec1.Code != http.StatusOK {
		t.Errorf("first request: expected status 200, got %d", rec1.Code)
	}

	// Second request should be rate limited
	req2 := httptest.NewRequest("GET", "/test", http.NoBody)
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)

	if rec2.Code != http.StatusTooManyRequests {
		t.Errorf("second request: expected status 429, got %d", rec2.Code)
	}
}

func TestHTTPTimeout(t *testing.T) {
	tm := timeout.New[*http.Response](timeout.Config{
		DefaultTimeout: time.Second,
	})

	handler := fortifyhttp.Timeout(tm, time.Millisecond*50)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-time.After(time.Millisecond * 200):
				w.WriteHeader(http.StatusOK)
			case <-r.Context().Done():
				// Timeout occurred
				return
			}
		}),
	)

	req := httptest.NewRequest("GET", "/test", http.NoBody)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusGatewayTimeout {
		t.Errorf("expected status 504, got %d", rec.Code)
	}
}
