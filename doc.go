// Package fortify provides production-grade resilience and fault-tolerance
// patterns for Go 1.23+.
//
// Fortify offers a comprehensive suite of battle-tested patterns including
// circuit breakers, retries, rate limiting, timeouts, and bulkheads with
// zero external dependencies for core functionality.
//
// # Features
//
//   - 🔒 Type-Safe: Built with Go 1.23+ generics for compile-time safety
//   - ⚡ High Performance: <1µs overhead with zero allocations in hot paths
//   - 🎯 Zero Dependencies: Core patterns have no external dependencies
//   - 🔍 Observable: Built-in support for structured logging (slog) and OpenTelemetry
//   - 📊 Prometheus Metrics: Export metrics for all resilience patterns
//   - 🌐 Framework Integration: First-class support for HTTP and gRPC
//   - 🧩 Composable: Fluent API for combining multiple patterns
//   - 🧪 Well Tested: >95% test coverage with race detection
//
// # Installation
//
//	go get github.com/felixgeelhaar/fortify
//
// Requirements: Go 1.23 or higher
//
// # Quick Start
//
//	import (
//	    "context"
//	    "time"
//
//	    "github.com/felixgeelhaar/fortify/circuitbreaker"
//	    "github.com/felixgeelhaar/fortify/retry"
//	)
//
//	func main() {
//	    // Create a circuit breaker
//	    cb := circuitbreaker.New[string](circuitbreaker.Config{
//	        MaxRequests: 100,
//	        Interval:    time.Second * 10,
//	        ReadyToTrip: func(counts circuitbreaker.Counts) bool {
//	            return counts.ConsecutiveFailures >= 3
//	        },
//	    })
//
//	    // Create a retry strategy
//	    r := retry.New[string](retry.Config{
//	        MaxAttempts:   3,
//	        InitialDelay:  time.Millisecond * 100,
//	        BackoffPolicy: retry.BackoffExponential,
//	    })
//
//	    // Use them together
//	    result, err := cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
//	        return r.Do(ctx, func(ctx context.Context) (string, error) {
//	            return callExternalService(ctx)
//	        })
//	    })
//	}
//
// # Core Patterns
//
// Circuit Breaker: Prevents cascading failures by temporarily stopping requests to failing services
//
//	import "github.com/felixgeelhaar/fortify/circuitbreaker"
//
//	cb := circuitbreaker.New[Response](circuitbreaker.Config{
//	    MaxRequests: 100,
//	    Interval:    time.Second * 60,
//	    Timeout:     time.Second * 30,
//	})
//
// Retry: Handles transient failures with intelligent backoff strategies
//
//	import "github.com/felixgeelhaar/fortify/retry"
//
//	r := retry.New[Response](retry.Config{
//	    MaxAttempts:   3,
//	    InitialDelay:  time.Millisecond * 100,
//	    BackoffPolicy: retry.BackoffExponential,
//	    Multiplier:    2.0,
//	})
//
// Rate Limiter: Token bucket algorithm for controlling request rates
//
//	import "github.com/felixgeelhaar/fortify/ratelimit"
//
//	rl := ratelimit.New(ratelimit.Config{
//	    Rate:     100,
//	    Burst:    10,
//	    Interval: time.Second,
//	})
//
// Timeout: Context-based timeout enforcement
//
//	import "github.com/felixgeelhaar/fortify/timeout"
//
//	t := timeout.New[Response](timeout.Config{
//	    DefaultTimeout: time.Second * 5,
//	})
//
// Bulkhead: Limits concurrent executions to prevent resource exhaustion
//
//	import "github.com/felixgeelhaar/fortify/bulkhead"
//
//	bh := bulkhead.New[Response](bulkhead.Config{
//	    MaxConcurrent: 10,
//	    MaxQueue:      100,
//	})
//
// Fallback: Provides graceful degradation with fallback values
//
//	import "github.com/felixgeelhaar/fortify/fallback"
//
//	fb := fallback.New[Response](fallback.Config[Response]{
//	    Fallback: func(ctx context.Context, err error) (Response, error) {
//	        return cachedResponse, nil
//	    },
//	})
//
// # Middleware Composition
//
// Combine multiple patterns using the middleware chain:
//
//	import "github.com/felixgeelhaar/fortify/middleware"
//
//	chain := middleware.New[Response]().
//	    WithCircuitBreaker(cb).
//	    WithRetry(r).
//	    WithTimeout(t, 5*time.Second).
//	    WithRateLimit(rl, "user-123")
//
//	result, err := chain.Execute(ctx, func(ctx context.Context) (Response, error) {
//	    return apiClient.Call(ctx)
//	})
//
// # HTTP Integration
//
// Use resilience patterns as HTTP middleware:
//
//	import (
//	    "net/http"
//	    "github.com/felixgeelhaar/fortify/http"
//	)
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/api", apiHandler)
//
//	// Wrap with circuit breaker
//	handler := fortifyhttp.CircuitBreaker(cb, mux)
//
//	http.ListenAndServe(":8080", handler)
//
// # gRPC Integration
//
// Add resilience to gRPC services:
//
//	import (
//	    "github.com/felixgeelhaar/fortify/grpc"
//	    "google.golang.org/grpc"
//	)
//
//	server := grpc.NewServer(
//	    grpc.UnaryInterceptor(
//	        fortifygrpc.UnaryCircuitBreakerInterceptor(cb),
//	    ),
//	)
//
// # Observability
//
// Built-in support for logging, tracing, and metrics:
//
//	import (
//	    "github.com/felixgeelhaar/fortify/otel"
//	    "github.com/felixgeelhaar/fortify/slog"
//	    "github.com/felixgeelhaar/fortify/metrics"
//	)
//
//	// OpenTelemetry tracing
//	cb := circuitbreaker.New[Response](circuitbreaker.Config{
//	    // ... config
//	})
//	traced := otel.WithTracing(cb, "circuit-breaker")
//
//	// Structured logging
//	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
//	cb := circuitbreaker.New[Response](circuitbreaker.Config{
//	    Logger: logger,
//	    // ... config
//	})
//
//	// Prometheus metrics
//	collector := metrics.NewCollector()
//	prometheus.MustRegister(collector)
//
// # Package Organization
//
// Core Patterns:
//   - circuitbreaker: Circuit breaker pattern implementation
//   - retry: Retry pattern with backoff strategies
//   - ratelimit: Token bucket rate limiting
//   - timeout: Context-based timeout enforcement
//   - bulkhead: Concurrency limiting with semaphores
//   - fallback: Graceful degradation pattern
//
// Integration:
//   - http: HTTP middleware for resilience patterns
//   - grpc: gRPC interceptors for client and server
//   - middleware: Composable middleware chains
//
// Observability:
//   - otel: OpenTelemetry tracing integration
//   - slog: Structured logging handlers
//   - metrics: Prometheus metrics exporters
//
// Utilities:
//   - ferrors: Standard error types and utilities
//   - testing: Testing utilities and chaos engineering tools
//
// # Error Handling
//
// Fortify provides standard error types for pattern-specific failures:
//
//	import "github.com/felixgeelhaar/fortify/ferrors"
//
//	result, err := cb.Execute(ctx, operation)
//	if errors.Is(err, ferrors.ErrCircuitOpen) {
//	    // Circuit breaker is open
//	}
//	if errors.Is(err, ferrors.ErrRateLimitExceeded) {
//	    // Rate limit exceeded
//	}
//	if errors.Is(err, ferrors.ErrBulkheadFull) {
//	    // Bulkhead at capacity
//	}
//
// # Testing
//
// Fortify includes utilities for chaos engineering and resilience testing:
//
//	import "github.com/felixgeelhaar/fortify/testing"
//
//	// Error injection for testing
//	injector := testing.NewErrorInjector(0.3) // 30% failure rate
//	result, err := injector.Execute(ctx, func(ctx context.Context) (Response, error) {
//	    return apiCall(ctx)
//	})
//
//	// Latency injection
//	latency := testing.NewLatencyInjector(
//	    time.Millisecond*100,
//	    time.Millisecond*500,
//	)
//
// # Performance
//
// Fast-path overhead (per operation):
//   - Circuit Breaker (closed): ~100ns
//   - Rate Limiter (tokens available): ~200ns
//   - Timeout (no timeout): ~50ns
//   - Retry (first attempt success): ~150ns
//   - Bulkhead (capacity available): ~300ns
//
// Memory footprint:
//   - Circuit Breaker: ~200 bytes per instance
//   - Rate Limiter: ~100 bytes per key bucket
//   - Timeout: ~80 bytes per instance
//   - Retry: ~120 bytes per instance
//   - Bulkhead: ~250 bytes + worker pool
//
// # Documentation
//
// For detailed documentation, examples, and API reference:
//   - Package Documentation: https://pkg.go.dev/github.com/felixgeelhaar/fortify
//   - GitHub Repository: https://github.com/felixgeelhaar/fortify
//   - Production Guide: https://github.com/felixgeelhaar/fortify/blob/main/docs/PRODUCTION.md
//   - Error Handling: https://github.com/felixgeelhaar/fortify/blob/main/docs/ERROR_HANDLING.md
//   - Migration Guide: https://github.com/felixgeelhaar/fortify/blob/main/docs/MIGRATION.md
//
// # License
//
// Fortify is released under the MIT License.
// See LICENSE file for details.
package fortify
