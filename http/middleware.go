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
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/felixgeelhaar/fortify/v2/circuitbreaker"
	"github.com/felixgeelhaar/fortify/v2/ferrors"
	"github.com/felixgeelhaar/fortify/v2/ratelimit"
	"github.com/felixgeelhaar/fortify/v2/timeout"
	"golang.org/x/text/unicode/norm"
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

// SanitizeKey sanitizes a rate limiting key by:
//   - Normalizing Unicode to NFC form (prevents equivalent-string bypass attacks)
//   - Removing control characters and non-printable characters
//   - Truncating to the specified maximum length (by rune count, not bytes)
//
// This helps prevent memory exhaustion, injection attacks, and rate limit bypass
// via Unicode equivalence exploitation.
//
// Performance: Uses fast path for ASCII-only strings (common case for IPs, API keys).
func SanitizeKey(key string, maxLen int) string {
	if maxLen <= 0 {
		maxLen = DefaultMaxKeyLength
	}

	// Fast path: check if key is clean ASCII (printable, no control chars)
	// This is the common case for IP addresses, API keys, user IDs, etc.
	if isCleanASCII(key) && len(key) <= maxLen {
		return key
	}

	// Fast path for ASCII-only strings that just need truncation
	if isASCII(key) {
		return sanitizeASCII(key, maxLen)
	}

	// Slow path: full Unicode handling for non-ASCII strings
	return sanitizeUnicode(key, maxLen)
}

// isCleanASCII returns true if the string contains only printable ASCII characters (32-126).
// This is an extremely fast check that handles the common case.
func isCleanASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

// isASCII returns true if the string contains only ASCII characters (0-127).
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

// sanitizeASCII handles ASCII-only strings efficiently without Unicode overhead.
func sanitizeASCII(key string, maxLen int) string {
	// Check if we need to remove any control characters
	needsSanitization := false
	for i := 0; i < len(key); i++ {
		c := key[i]
		if c < 32 || c == 127 {
			needsSanitization = true
			break
		}
	}

	if !needsSanitization {
		// Just truncate if needed
		if len(key) > maxLen {
			return key[:maxLen]
		}
		return key
	}

	// Remove control characters
	result := make([]byte, 0, len(key))
	for i := 0; i < len(key); i++ {
		c := key[i]
		if c >= 32 && c != 127 {
			result = append(result, c)
		}
	}

	if len(result) > maxLen {
		result = result[:maxLen]
	}

	return string(result)
}

// sanitizeUnicode handles non-ASCII strings with full Unicode normalization.
func sanitizeUnicode(key string, maxLen int) string {
	// Normalize to NFC (Canonical Composition) to prevent Unicode bypass attacks
	// e.g., "user123" vs "user\u0031\u0032\u0033" will now produce the same key
	key = norm.NFC.String(key)

	// Remove control characters and non-printable characters
	sanitized := strings.Map(func(r rune) rune {
		if unicode.IsControl(r) || !unicode.IsPrint(r) {
			return -1 // Remove control and non-printable characters
		}
		return r
	}, key)

	// Truncate by rune count (not byte count) to avoid splitting multi-byte UTF-8 characters
	if utf8.RuneCountInString(sanitized) > maxLen {
		runes := []rune(sanitized)
		sanitized = string(runes[:maxLen])
	}

	return sanitized
}

// KeyFromIP extracts the client IP address as the rate limiting key.
// It properly handles both IPv4 and IPv6 addresses, stripping zone identifiers
// from IPv6 addresses to prevent rate limit bypass attacks.
//
// Security: IPv6 zone identifiers (e.g., fe80::1%eth0) are stripped to prevent
// attackers from bypassing rate limits by varying the zone identifier while
// using the same IP address.
func KeyFromIP(r *http.Request) string {
	ip := r.RemoteAddr

	// Use net.SplitHostPort for proper IPv4/IPv6 handling
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		// RemoteAddr might not have a port (unusual but possible)
		host = ip
	}

	// Strip IPv6 zone identifier to prevent bypass attacks
	// e.g., "fe80::1%eth0" and "fe80::1%eth1" should be treated as the same IP
	if idx := strings.Index(host, "%"); idx != -1 {
		host = host[:idx]
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
//
// Panics if header name is empty or contains invalid HTTP header characters.
// This validation happens at construction time to fail fast on misconfiguration.
func KeyFromHeader(header string) KeyExtractor {
	validateHeaderName(header)
	return func(r *http.Request) string {
		value := r.Header.Get(header)
		return SanitizeKey(value, DefaultMaxKeyLength)
	}
}

// KeyFromHeaderWithMaxLen returns a KeyExtractor that extracts the key from an HTTP header
// with a custom maximum key length.
//
// Panics if header name is empty or contains invalid HTTP header characters.
// This validation happens at construction time to fail fast on misconfiguration.
func KeyFromHeaderWithMaxLen(header string, maxLen int) KeyExtractor {
	validateHeaderName(header)
	return func(r *http.Request) string {
		value := r.Header.Get(header)
		return SanitizeKey(value, maxLen)
	}
}

// validateHeaderName validates that a header name contains only valid HTTP header characters.
// Panics if the header name is invalid, as this indicates a programming error.
//
// Valid HTTP header characters are ASCII 33-126 (printable, non-whitespace) except colon.
// See RFC 7230 Section 3.2.6: https://tools.ietf.org/html/rfc7230#section-3.2.6
func validateHeaderName(header string) {
	if header == "" {
		panic("fortify/http: KeyFromHeader requires non-empty header name")
	}

	for i, ch := range header {
		// Valid header token characters: ASCII 33-126 except delimiters
		// Delimiters include: ( ) < > @ , ; : \ " / [ ] ? = { }
		// For simplicity, we only reject the most dangerous: control chars, space, and colon
		if ch < 33 || ch > 126 || ch == ':' {
			panic(fmt.Sprintf("fortify/http: invalid header name %q at position %d", header, i))
		}
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
