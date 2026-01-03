package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	fortifyhttp "github.com/felixgeelhaar/fortify/v2/http"
	"github.com/felixgeelhaar/fortify/v2/ratelimit"
	"github.com/felixgeelhaar/fortify/v2/timeout"
)

func TestHTTPExample(t *testing.T) {
	// Skip this test to avoid port conflicts when running with -count > 1
	// The main() function starts a real HTTP server on port 8080, which
	// cannot be bound twice. The middleware functionality is thoroughly
	// tested in TestHTTPMiddleware, TestHTTPRateLimit, and TestHTTPTimeout.
	t.Skip("Skipping main() test to avoid port 8080 conflicts with -count > 1")
}

func TestHTTPMiddleware(t *testing.T) {
	rl := ratelimit.New(&ratelimit.Config{
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
	rl := ratelimit.New(&ratelimit.Config{
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
