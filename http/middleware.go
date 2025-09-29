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
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
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
					return &http.Response{StatusCode: rec.statusCode}, fortifyerrors.ErrCircuitOpen
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
			done := make(chan struct{})
			var timedOut bool

			_, err := tm.Execute(r.Context(), duration, func(ctx context.Context) (*http.Response, error) {
				defer close(done)

				rec := &responseRecorder{
					ResponseWriter: w,
					statusCode:     http.StatusOK,
				}
				next.ServeHTTP(rec, r.WithContext(ctx))
				return &http.Response{StatusCode: rec.statusCode}, nil
			})

			if err != nil {
				// Check if handler already completed
				select {
				case <-done:
					// Handler completed, don't write timeout error
					return
				default:
					// Handler still running or timed out
					timedOut = true
				}

				if timedOut {
					http.Error(w, "Gateway Timeout", http.StatusGatewayTimeout)
					return
				}
			}
		})
	}
}

// KeyFromIP extracts the client IP address as the rate limiting key.
func KeyFromIP(r *http.Request) string {
	ip := r.RemoteAddr
	// Strip port if present
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

// KeyFromHeader returns a KeyExtractor that extracts the key from an HTTP header.
func KeyFromHeader(header string) KeyExtractor {
	return func(r *http.Request) string {
		return r.Header.Get(header)
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
