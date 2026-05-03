package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/middleware"
	"github.com/felixgeelhaar/fortify/ratelimit"
)

func TestHTTPHandler_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		middleware.HTTPHandlerConfig{Timeout: 0})
	if err == nil {
		t.Fatal("want error, got nil")
	}
}

func TestHTTPHandler_RejectsRateLimitWithoutKeyFn(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{Rate: 10, Burst: 10, Interval: time.Second})
	defer func() { _ = rl.Close() }()
	_, err := middleware.HTTPHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		middleware.HTTPHandlerConfig{Timeout: 1 * time.Second, RateLimit: rl},
	)
	if err == nil {
		t.Fatal("want error for missing keyFn, got nil")
	}
}

func TestHTTPHandler_PassesThrough(t *testing.T) {
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	h, err := middleware.HTTPHandler(inner, middleware.HTTPHandlerConfig{Timeout: 1 * time.Second})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if body := rr.Body.String(); body != "ok" {
		t.Fatalf("body = %q, want ok", body)
	}
}

func TestHTTPHandler_RateLimitReturns429(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{Rate: 1, Burst: 1, Interval: time.Hour})
	defer func() { _ = rl.Close() }()

	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	h, err := middleware.HTTPHandler(inner, middleware.HTTPHandlerConfig{
		Timeout:        1 * time.Second,
		RateLimit:      rl,
		RateLimitKeyFn: func(*http.Request) string { return "k" },
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	// First admitted.
	rr1 := httptest.NewRecorder()
	h.ServeHTTP(rr1, httptest.NewRequest(http.MethodGet, "/", nil))
	if rr1.Code != http.StatusOK {
		t.Fatalf("first call status = %d, want 200", rr1.Code)
	}
	// Second rejected.
	rr2 := httptest.NewRecorder()
	h.ServeHTTP(rr2, httptest.NewRequest(http.MethodGet, "/", nil))
	if rr2.Code != http.StatusTooManyRequests {
		t.Fatalf("second call status = %d, want 429", rr2.Code)
	}
}

func TestHTTPHandler_TripsCircuitOn5xxStorm(t *testing.T) {
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	})
	h, err := middleware.HTTPHandler(inner, middleware.HTTPHandlerConfig{
		Timeout:            1 * time.Second,
		CBFailureThreshold: 3,
		CBOpenTimeout:      30 * time.Second,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	// Trigger threshold failures.
	for i := 0; i < 3; i++ {
		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/", nil))
		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500 during fail run, got %d", rr.Code)
		}
	}
	// Next call should hit Open breaker → 503.
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/", nil))
	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("breaker did not trip; got %d, want 503", rr.Code)
	}
}
