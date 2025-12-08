// Package http provides HTTP middleware for integrating Fortify resilience patterns
// with standard http.Handler implementations.
//
// This package offers middleware functions that wrap HTTP handlers with circuit breakers,
// rate limiters, timeouts, and other resilience patterns, making it easy to add
// fault tolerance to HTTP services.
//
// Example usage:
//
//	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{...})
//	rl := ratelimit.New(ratelimit.Config{...})
//
//	handler := http.CircuitBreaker(cb)(
//	    http.RateLimit(rl, http.KeyFromIP)(
//	        http.HandlerFunc(myHandler),
//	    ),
//	)
package http

import (
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

const (
	// DefaultMaxKeyLength is the default maximum length for rate limiting keys.
	// Keys longer than this are truncated to prevent memory exhaustion.
	DefaultMaxKeyLength = 256
)

// KeyExtractor extracts a rate limiting key from an HTTP request.
type KeyExtractor func(*http.Request) string

// CircuitBreaker wraps an HTTP handler with circuit breaker protection.
// Returns 503 Service Unavailable when the circuit is open.
func CircuitBreaker(cb circuitbreaker.CircuitBreaker[*http.Response]) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp, err := cb.Execute(r.Context(), func(ctx context.Context) (*http.Response, error) {
				// Create a response recorder to capture the handler's response
				rec := &responseRecorder{
					ResponseWriter: w,
					statusCode:     http.StatusOK,
				}
				next.ServeHTTP(rec, r.WithContext(ctx))

				// Return error if status code indicates failure
				if rec.statusCode >= 500 {
					return &http.Response{StatusCode: rec.statusCode}, ferrors.ErrCircuitOpen
				}

				return &http.Response{StatusCode: rec.statusCode}, nil
			})

			if err != nil {
				// Circuit is open or other error occurred
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
				return
			}

			// Response already written by handler
			_ = resp
		})
	}
}

// RateLimit wraps an HTTP handler with rate limiting.
// Returns 429 Too Many Requests when the rate limit is exceeded.
func RateLimit(rl ratelimit.RateLimiter, keyExtractor KeyExtractor) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := keyExtractor(r)

			if !rl.Allow(r.Context(), key) {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Timeout wraps an HTTP handler with timeout enforcement.
// Returns 504 Gateway Timeout when the request times out.
func Timeout(tm timeout.Timeout[*http.Response], duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rec := &responseRecorder{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			_, err := tm.Execute(r.Context(), duration, func(ctx context.Context) (*http.Response, error) {
				next.ServeHTTP(rec, r.WithContext(ctx))

				// Check if context was cancelled (timeout occurred)
				if ctx.Err() != nil {
					return nil, ctx.Err()
				}

				return &http.Response{StatusCode: rec.statusCode}, nil
			})

			if err != nil {
				// Timeout occurred - write error response if handler hasn't written anything substantial
				if !rec.written || rec.statusCode == http.StatusOK {
					rec.WriteHeader(http.StatusGatewayTimeout)
					//nolint:errcheck // intentionally ignoring error in middleware
					_, _ = rec.Write([]byte("Gateway Timeout"))
				}
			}
		})
	}
}

// SanitizeKey sanitizes a rate limiting key by removing control characters
// and truncating to the specified maximum length.
// This helps prevent memory exhaustion and injection attacks.
func SanitizeKey(key string, maxLen int) string {
	if maxLen <= 0 {
		maxLen = DefaultMaxKeyLength
	}

	// Remove control characters (keeps printable characters and spaces)
	sanitized := strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1 // Remove control characters
		}
		return r
	}, key)

	// Truncate to max length
	if len(sanitized) > maxLen {
		sanitized = sanitized[:maxLen]
	}

	return sanitized
}

// KeyFromIP extracts the client IP address as the rate limiting key.
// It properly handles both IPv4 and IPv6 addresses, including IPv6 with zone identifiers.
func KeyFromIP(r *http.Request) string {
	ip := r.RemoteAddr

	// Use net.SplitHostPort for proper IPv4/IPv6 handling
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		// RemoteAddr might not have a port (unusual but possible)
		host = ip
	}

	// Validate it looks like an IP address
	if parsedIP := net.ParseIP(host); parsedIP != nil {
		return parsedIP.String() // Normalized IP string
	}

	// Fallback: sanitize whatever we got
	return SanitizeKey(host, DefaultMaxKeyLength)
}

// KeyFromHeader returns a KeyExtractor that extracts the key from an HTTP header.
// The extracted value is sanitized to prevent injection attacks and memory exhaustion.
func KeyFromHeader(header string) KeyExtractor {
	return func(r *http.Request) string {
		value := r.Header.Get(header)
		return SanitizeKey(value, DefaultMaxKeyLength)
	}
}

// KeyFromHeaderWithMaxLen returns a KeyExtractor that extracts the key from an HTTP header
// with a custom maximum key length.
func KeyFromHeaderWithMaxLen(header string, maxLen int) KeyExtractor {
	return func(r *http.Request) string {
		value := r.Header.Get(header)
		return SanitizeKey(value, maxLen)
	}
}

// responseRecorder captures the HTTP response status code.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	written    bool
	mu         sync.Mutex
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.written {
		r.statusCode = statusCode
		r.written = true
		r.ResponseWriter.WriteHeader(statusCode)
	}
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.mu.Lock()
	if !r.written {
		r.statusCode = http.StatusOK
		r.written = true
		r.ResponseWriter.WriteHeader(http.StatusOK)
	}
	r.mu.Unlock()

	return r.ResponseWriter.Write(b)
}
