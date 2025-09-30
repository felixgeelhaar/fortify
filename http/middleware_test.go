package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

func TestCircuitBreakerMiddleware(t *testing.T) {
	t.Run("allows requests when circuit closed", func(t *testing.T) {
		cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
			MaxRequests: 10,
			Interval:    time.Second,
		})

		handler := CircuitBreaker(cb)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("success"))
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if rec.Body.String() != "success" {
			t.Errorf("body = %s, want success", rec.Body.String())
		}
	})

	t.Run("returns 503 when circuit open", func(t *testing.T) {
		cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
			MaxRequests: 1,
			Interval:    time.Second,
			ReadyToTrip: func(counts circuitbreaker.Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		handler := CircuitBreaker(cb)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		// First request fails, opens circuit
		req1 := httptest.NewRequest("GET", "/test", nil)
		rec1 := httptest.NewRecorder()
		handler.ServeHTTP(rec1, req1)

		// Second request should be rejected (circuit open)
		req2 := httptest.NewRequest("GET", "/test", nil)
		rec2 := httptest.NewRecorder()
		handler.ServeHTTP(rec2, req2)

		if rec2.Code != http.StatusServiceUnavailable {
			t.Errorf("status = %d, want %d", rec2.Code, http.StatusServiceUnavailable)
		}
	})
}

func TestRateLimitMiddleware(t *testing.T) {
	t.Run("allows requests within rate limit", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     10,
			Interval: time.Second,
		})

		handler := RateLimit(rl, func(r *http.Request) string {
			return "test-key"
		})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("returns 429 when rate limit exceeded", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     1,
			Burst:    1,
			Interval: time.Hour, // Very slow refill
		})

		handler := RateLimit(rl, func(r *http.Request) string {
			return "test-key"
		})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		// First request succeeds
		req1 := httptest.NewRequest("GET", "/test", nil)
		rec1 := httptest.NewRecorder()
		handler.ServeHTTP(rec1, req1)

		// Second request should be rate limited
		req2 := httptest.NewRequest("GET", "/test", nil)
		rec2 := httptest.NewRecorder()
		handler.ServeHTTP(rec2, req2)

		if rec2.Code != http.StatusTooManyRequests {
			t.Errorf("status = %d, want %d", rec2.Code, http.StatusTooManyRequests)
		}
	})
}

func TestTimeoutMiddleware(t *testing.T) {
	t.Run("completes request within timeout", func(t *testing.T) {
		tm := timeout.New[*http.Response](timeout.Config{
			DefaultTimeout: time.Second,
		})

		handler := Timeout(tm, 100*time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("returns 504 when request times out", func(t *testing.T) {
		tm := timeout.New[*http.Response](timeout.Config{
			DefaultTimeout: time.Second,
		})

		responseSent := make(chan bool, 1)
		handler := Timeout(tm, 50*time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respect context cancellation to avoid race
			select {
			case <-r.Context().Done():
				responseSent <- false
				return
			case <-time.After(100 * time.Millisecond):
				w.WriteHeader(http.StatusOK)
				responseSent <- true
			}
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		// Wait for handler goroutine to complete
		sent := <-responseSent

		if sent {
			t.Error("handler should have been cancelled")
		}

		if rec.Code != http.StatusGatewayTimeout {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusGatewayTimeout)
		}
	})
}

func TestMiddlewareChaining(t *testing.T) {
	t.Run("chains multiple middlewares", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     100,
			Interval: time.Second,
		})
		tm := timeout.New[*http.Response](timeout.Config{
			DefaultTimeout: time.Second,
		})

		handler := RateLimit(rl, func(r *http.Request) string {
			return "test-key"
		})(Timeout(tm, 100*time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("success"))
		})))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if rec.Body.String() != "success" {
			t.Errorf("body = %s, want success", rec.Body.String())
		}
	})
}

func TestKeyExtractors(t *testing.T) {
	t.Run("extracts key from IP address", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "192.168.1.1:12345"

		key := KeyFromIP(req)
		if key != "192.168.1.1" {
			t.Errorf("key = %s, want 192.168.1.1", key)
		}
	})

	t.Run("extracts key from header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-User-ID", "user-123")

		extractor := KeyFromHeader("X-User-ID")
		key := extractor(req)
		if key != "user-123" {
			t.Errorf("key = %s, want user-123", key)
		}
	})
}
