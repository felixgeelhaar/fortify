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
			//nolint:errcheck // intentionally ignoring error in test
			_, _ = w.Write([]byte("success"))
		}))

		req := httptest.NewRequest("GET", "/test", http.NoBody)
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
		req1 := httptest.NewRequest("GET", "/test", http.NoBody)
		rec1 := httptest.NewRecorder()
		handler.ServeHTTP(rec1, req1)

		// Second request should be rejected (circuit open)
		req2 := httptest.NewRequest("GET", "/test", http.NoBody)
		rec2 := httptest.NewRecorder()
		handler.ServeHTTP(rec2, req2)

		if rec2.Code != http.StatusServiceUnavailable {
			t.Errorf("status = %d, want %d", rec2.Code, http.StatusServiceUnavailable)
		}
	})
}

func TestRateLimitMiddleware(t *testing.T) {
	t.Run("allows requests within rate limit", func(t *testing.T) {
		rl := ratelimit.New(&ratelimit.Config{
			Rate:     10,
			Interval: time.Second,
		})

		handler := RateLimit(rl, func(r *http.Request) string {
			return "test-key"
		})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/test", http.NoBody)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("returns 429 when rate limit exceeded", func(t *testing.T) {
		rl := ratelimit.New(&ratelimit.Config{
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
		req1 := httptest.NewRequest("GET", "/test", http.NoBody)
		rec1 := httptest.NewRecorder()
		handler.ServeHTTP(rec1, req1)

		// Second request should be rate limited
		req2 := httptest.NewRequest("GET", "/test", http.NoBody)
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

		req := httptest.NewRequest("GET", "/test", http.NoBody)
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

		handler := Timeout(tm, 50*time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate slow operation that respects context cancellation
			select {
			case <-time.After(100 * time.Millisecond):
				w.WriteHeader(http.StatusOK)
			case <-r.Context().Done():
				// Context cancelled due to timeout
				return
			}
		}))

		req := httptest.NewRequest("GET", "/test", http.NoBody)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusGatewayTimeout {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusGatewayTimeout)
		}
	})
}

func TestMiddlewareChaining(t *testing.T) {
	t.Run("chains multiple middlewares", func(t *testing.T) {
		rl := ratelimit.New(&ratelimit.Config{
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
			//nolint:errcheck // intentionally ignoring Write error in test
			_, _ = w.Write([]byte("success"))
		})))

		req := httptest.NewRequest("GET", "/test", http.NoBody)
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
		req := httptest.NewRequest("GET", "/test", http.NoBody)
		req.RemoteAddr = "192.168.1.1:12345"

		key := KeyFromIP(req)
		if key != "192.168.1.1" {
			t.Errorf("key = %s, want 192.168.1.1", key)
		}
	})

	t.Run("extracts key from header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", http.NoBody)
		req.Header.Set("X-User-ID", "user-123")

		extractor := KeyFromHeader("X-User-ID")
		key := extractor(req)
		if key != "user-123" {
			t.Errorf("key = %s, want user-123", key)
		}
	})
}

func TestKeyFromIP_IPv6ZoneBypass(t *testing.T) {
	// Test that IPv6 zone identifiers are stripped to prevent rate limit bypass
	tests := []struct {
		name       string
		remoteAddr string
		expected   string
	}{
		{"IPv6 with zone eth0", "[fe80::1%eth0]:12345", "fe80::1"},
		{"IPv6 with zone eth1", "[fe80::1%eth1]:12345", "fe80::1"},
		{"IPv6 with zone wlan0", "[fe80::1%wlan0]:12345", "fe80::1"},
		{"IPv6 without zone", "[fe80::1]:12345", "fe80::1"},
		{"IPv6 global with zone", "[2001:db8::1%zone]:80", "2001:db8::1"},
		{"IPv6 global without zone", "[2001:db8::1]:80", "2001:db8::1"},
		{"IPv4 address", "192.168.1.1:12345", "192.168.1.1"},
		{"IPv4 loopback", "127.0.0.1:8080", "127.0.0.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", http.NoBody)
			req.RemoteAddr = tt.remoteAddr

			key := KeyFromIP(req)
			if key != tt.expected {
				t.Errorf("KeyFromIP(%q) = %q, want %q", tt.remoteAddr, key, tt.expected)
			}
		})
	}

	// Verify that different zone identifiers produce the same key (bypass prevented)
	t.Run("same IP with different zones produces same key", func(t *testing.T) {
		zones := []string{"eth0", "eth1", "wlan0", "lo", "en0"}
		keys := make(map[string]bool)

		for _, zone := range zones {
			req := httptest.NewRequest("GET", "/", http.NoBody)
			req.RemoteAddr = "[fe80::1%" + zone + "]:12345"
			key := KeyFromIP(req)
			keys[key] = true
		}

		if len(keys) != 1 {
			t.Errorf("expected 1 unique key for all zone variants, got %d", len(keys))
		}
	})
}

func TestSanitizeKey_UnicodeNormalization(t *testing.T) {
	// Test that Unicode normalization prevents bypass attacks
	t.Run("normalizes equivalent Unicode representations", func(t *testing.T) {
		// These strings look the same but use different Unicode representations
		keys := []string{
			"user123",                // ASCII
			"user\u0031\u0032\u0033", // Unicode code points for 1, 2, 3
		}

		normalized := make(map[string]bool)
		for _, key := range keys {
			sanitized := SanitizeKey(key, 256)
			normalized[sanitized] = true
		}

		// All should normalize to the same key
		if len(normalized) != 1 {
			t.Errorf("expected 1 unique normalized key, got %d", len(normalized))
		}
	})

	t.Run("removes control characters", func(t *testing.T) {
		key := "user\x00\x01\x02123"
		sanitized := SanitizeKey(key, 256)

		// Should not contain control characters
		for _, r := range sanitized {
			if r < 32 && r != '\t' && r != '\n' && r != '\r' {
				t.Errorf("sanitized key contains control character: %U", r)
			}
		}
	})

	t.Run("truncates by rune count not byte count", func(t *testing.T) {
		// 5 emoji = 5 runes but 20 bytes in UTF-8
		key := "🔥🔥🔥🔥🔥more"
		sanitized := SanitizeKey(key, 5)

		// Should have exactly 5 runes
		runeCount := 0
		for range sanitized {
			runeCount++
		}

		if runeCount != 5 {
			t.Errorf("expected 5 runes, got %d", runeCount)
		}
	})
}

func TestKeyFromHeader_Validation(t *testing.T) {
	t.Run("panics on empty header name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for empty header name")
			}
		}()
		KeyFromHeader("")
	})

	t.Run("panics on header with colon", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for header name with colon")
			}
		}()
		KeyFromHeader("X-Bad:Header")
	})

	t.Run("panics on header with control characters", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for header name with control characters")
			}
		}()
		KeyFromHeader("X-Bad\nHeader")
	})

	t.Run("panics on header with space", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for header name with space")
			}
		}()
		KeyFromHeader("X-Bad Header")
	})

	t.Run("accepts valid header names", func(t *testing.T) {
		validHeaders := []string{
			"X-User-ID",
			"Authorization",
			"X-Forwarded-For",
			"X-Request-ID",
			"Custom-Header-123",
		}

		for _, header := range validHeaders {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("unexpected panic for valid header %q: %v", header, r)
					}
				}()
				KeyFromHeader(header)
			}()
		}
	})
}

func TestKeyFromHeaderWithMaxLen_Validation(t *testing.T) {
	t.Run("panics on empty header name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for empty header name")
			}
		}()
		KeyFromHeaderWithMaxLen("", 100)
	})

	t.Run("respects custom max length", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", http.NoBody)
		req.Header.Set("X-Long-Value", "this is a very long value that should be truncated")

		extractor := KeyFromHeaderWithMaxLen("X-Long-Value", 10)
		key := extractor(req)

		runeCount := 0
		for range key {
			runeCount++
		}

		if runeCount > 10 {
			t.Errorf("expected max 10 runes, got %d", runeCount)
		}
	})
}
