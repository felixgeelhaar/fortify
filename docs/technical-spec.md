# Go Resilience Package - Technical Specification (2025)

## Overview

A production-ready resilience package for Go services that provides circuit breakers, retries, rate limiting, timeouts, and bulkheads. Designed for monorepo use with zero external dependencies (core package), full observability support, and adherence to 2025 Go best practices.

**Target Go Version:** 1.23+ (generics, structured logging, native fuzzing)

## Design Principles

1. **Zero Dependencies (Core):** Standard library only for core functionality
2. **Context-First:** All operations respect context cancellation and propagation
3. **Type-Safe:** Leverage Go 1.23+ generics for compile-time safety
4. **Observable:** Native `log/slog` and OpenTelemetry W3C TraceContext support
5. **Testable:** Easy to mock, race-tested, with fuzzing support
6. **Composable:** Patterns chain and combine elegantly
7. **Production-Ready:** Thread-safe, benchmarked, battle-tested patterns

## Package Structure

```
resilience/
├── circuitbreaker/
│   ├── breaker.go
│   ├── state.go
│   └── config.go
├── retry/
│   ├── retry.go
│   ├── backoff.go
│   └── config.go
├── ratelimit/
│   ├── limiter.go
│   ├── tokenbucket.go
│   └── config.go
├── timeout/
│   ├── timeout.go
│   └── config.go
├── bulkhead/
│   ├── bulkhead.go
│   └── config.go
├── middleware/
│   ├── http.go
│   ├── grpc.go
│   └── chain.go
├── otel/
│   ├── propagation.go
│   ├── tracing.go
│   └── metrics.go
├── slog/
│   ├── handlers.go
│   └── attributes.go
├── errors/
│   └── errors.go
└── examples/
    ├── http/
    ├── grpc/
    ├── otel/
    └── composition/
```

## Core Components

### 1. Circuit Breaker

**Purpose:** Prevent cascading failures by failing fast when dependencies are unhealthy.

**Features:**
- States: Closed, Open, Half-Open
- Configurable failure thresholds
- Automatic recovery with exponential backoff
- OpenTelemetry span creation
- Structured logging with `log/slog`

**API Design:**

```go
package circuitbreaker

import (
    "context"
    "log/slog"
    "time"
    "go.opentelemetry.io/otel/trace"
)

type CircuitBreaker[T any] interface {
    Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
    State() State
    Reset()
}

type Config struct {
    MaxRequests       uint32
    Interval          time.Duration
    Timeout           time.Duration
    ReadyToTrip       func(counts Counts) bool
    OnStateChange     func(from, to State)
    IsSuccessful      func(err error) bool

    // Observability (optional)
    Logger            *slog.Logger
    Tracer            trace.Tracer
}

type State int

const (
    StateClosed State = iota
    StateOpen
    StateHalfOpen
)

type Counts struct {
    Requests             uint32
    TotalSuccesses       uint32
    TotalFailures        uint32
    ConsecutiveSuccesses uint32
    ConsecutiveFailures  uint32
}

// Example usage
func NewCircuitBreaker[T any](config Config) CircuitBreaker[T]
```

**Example:**

```go
cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    MaxRequests: 5,
    Interval:    10 * time.Second,
    Timeout:     60 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures > 5
    },
    Logger: slog.Default(),
    Tracer: otel.Tracer("my-service"),
})

result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return callExternalAPI(ctx)
})
```

### 2. Retry

**Purpose:** Automatically retry failed operations with intelligent backoff strategies.

**Features:**
- Multiple backoff strategies: exponential, linear, constant, jitter
- Maximum attempts and timeout limits
- Error classification (retryable vs non-retryable)
- Context cancellation support
- Trace span per attempt

**API Design:**

```go
package retry

import (
    "context"
    "log/slog"
    "time"
    "go.opentelemetry.io/otel/trace"
)

type Retry[T any] interface {
    Do(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
}

type BackoffPolicy int

const (
    BackoffExponential BackoffPolicy = iota
    BackoffLinear
    BackoffConstant
)

type Config struct {
    MaxAttempts       int
    InitialDelay      time.Duration
    MaxDelay          time.Duration
    Multiplier        float64
    BackoffPolicy     BackoffPolicy
    Jitter            bool

    // Error classification
    RetryableErrors    []error              // errors.Is comparison
    NonRetryableErrors []error              // errors.Is comparison
    IsRetryable        func(error) bool     // Custom classifier

    // Callbacks
    OnRetry           func(attempt int, err error)

    // Observability
    Logger            *slog.Logger
    Tracer            trace.Tracer
}

func New[T any](config Config) Retry[T]
```

**Example:**

```go
r := retry.New[*User](retry.Config{
    MaxAttempts:   3,
    InitialDelay:  100 * time.Millisecond,
    MaxDelay:      5 * time.Second,
    Multiplier:    2.0,
    BackoffPolicy: retry.BackoffExponential,
    Jitter:        true,
    RetryableErrors: []error{
        ErrNetworkTimeout,
        ErrServiceUnavailable,
    },
    Logger: slog.Default(),
})

user, err := r.Do(ctx, func(ctx context.Context) (*User, error) {
    return fetchUser(ctx, userID)
})
```

### 3. Rate Limiter

**Purpose:** Control request rates to prevent resource exhaustion.

**Features:**
- Token bucket algorithm
- Per-key rate limiting (user, IP, tenant)
- Sliding window support
- Distributed rate limiting (Redis backend - optional addon)
- Wait vs immediate rejection

**API Design:**

```go
package ratelimit

import (
    "context"
    "time"
    "log/slog"
)

type RateLimiter interface {
    Allow(ctx context.Context, key string) bool
    Wait(ctx context.Context, key string) error
    Take(ctx context.Context, key string, tokens int) bool
}

type Config struct {
    Rate       int                                    // tokens per interval
    Burst      int                                    // bucket size
    Interval   time.Duration
    KeyFunc    func(ctx context.Context) string      // extract key from context

    // Observability
    Logger     *slog.Logger
    OnLimit    func(key string)
}

func New(config Config) RateLimiter
```

**Example:**

```go
limiter := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    150,
    Interval: time.Second,
    KeyFunc: func(ctx context.Context) string {
        return ctx.Value("user_id").(string)
    },
    Logger: slog.Default(),
})

if limiter.Allow(ctx, "user123") {
    // Process request
} else {
    // Return 429 Too Many Requests
}
```

### 4. Timeout

**Purpose:** Enforce operation time limits with graceful degradation.

**Features:**
- Per-operation timeouts
- Deadline propagation via context
- Configurable default timeout
- Timeout callbacks for cleanup

**API Design:**

```go
package timeout

import (
    "context"
    "time"
    "log/slog"
)

type Timeout[T any] interface {
    Execute(ctx context.Context, timeout time.Duration, fn func(context.Context) (T, error)) (T, error)
}

type Config struct {
    DefaultTimeout time.Duration
    OnTimeout      func()
    Logger         *slog.Logger
}

func New[T any](config Config) Timeout[T]
```

**Example:**

```go
t := timeout.New[[]byte](timeout.Config{
    DefaultTimeout: 5 * time.Second,
    Logger:         slog.Default(),
})

data, err := t.Execute(ctx, 2*time.Second, func(ctx context.Context) ([]byte, error) {
    return fetchData(ctx)
})
```

### 5. Bulkhead

**Purpose:** Isolate resources to prevent resource exhaustion.

**Features:**
- Semaphore-based concurrency limiting
- Queue support for waiting requests
- Per-resource isolation
- Metrics on utilization

**API Design:**

```go
package bulkhead

import (
    "context"
    "time"
    "log/slog"
)

type Bulkhead[T any] interface {
    Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
}

type Config struct {
    MaxConcurrent int
    MaxQueue      int
    QueueTimeout  time.Duration

    // Observability
    Logger        *slog.Logger
    OnRejected    func()
}

func New[T any](config Config) Bulkhead[T]
```

## Pattern Composition

### Middleware Chain

Compose multiple resilience patterns elegantly:

```go
package middleware

import "context"

type Handler[T any] func(context.Context) (T, error)

type Middleware[T any] func(Handler[T]) Handler[T]

func Chain[T any](middlewares ...Middleware[T]) Middleware[T] {
    return func(next Handler[T]) Handler[T] {
        for i := len(middlewares) - 1; i >= 0; i-- {
            next = middlewares[i](next)
        }
        return next
    }
}
```

**Example Usage:**

```go
// Create individual patterns
cb := circuitbreaker.New[*Response](cbConfig)
retry := retry.New[*Response](retryConfig)
timeout := timeout.New[*Response](timeoutConfig)

// Compose them
handler := middleware.Chain(
    circuitbreaker.AsMiddleware(cb),
    retry.AsMiddleware(retry),
    timeout.AsMiddleware(timeout),
)(func(ctx context.Context) (*Response, error) {
    return callAPI(ctx)
})

// Execute with all patterns
result, err := handler(ctx)
```

### HTTP Middleware

Standard `net/http` middleware support:

```go
package middleware

import "net/http"

func CircuitBreakerMiddleware(cb circuitbreaker.CircuitBreaker[any], logger *slog.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            _, err := cb.Execute(r.Context(), func(ctx context.Context) (any, error) {
                // Capture response through custom ResponseWriter
                next.ServeHTTP(w, r.WithContext(ctx))
                return nil, nil
            })
            if err != nil {
                logger.ErrorContext(r.Context(), "circuit breaker open",
                    slog.String("path", r.URL.Path))
                http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
            }
        })
    }
}

func RateLimitMiddleware(limiter ratelimit.RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            key := limiter.KeyFunc(r.Context())
            if !limiter.Allow(r.Context(), key) {
                http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}

func TimeoutMiddleware(duration time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), duration)
            defer cancel()
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

### gRPC Interceptors

Full gRPC support with proper context propagation:

```go
package middleware

import (
    "google.golang.org/grpc"
    "go.opentelemetry.io/otel/trace"
)

func UnaryCircuitBreakerInterceptor[T any](cb circuitbreaker.CircuitBreaker[T]) grpc.UnaryClientInterceptor {
    return func(ctx context.Context, method string, req, reply interface{},
                cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
        _, err := cb.Execute(ctx, func(ctx context.Context) (interface{}, error) {
            return nil, invoker(ctx, method, req, reply, cc, opts...)
        })
        return err
    }
}

func UnaryRetryInterceptor[T any](r retry.Retry[T]) grpc.UnaryClientInterceptor {
    return func(ctx context.Context, method string, req, reply interface{},
                cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
        _, err := r.Do(ctx, func(ctx context.Context) (interface{}, error) {
            return nil, invoker(ctx, method, req, reply, cc, opts...)
        })
        return err
    }
}

func StreamCircuitBreakerInterceptor[T any](cb circuitbreaker.CircuitBreaker[T]) grpc.StreamClientInterceptor {
    return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
                method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
        result, err := cb.Execute(ctx, func(ctx context.Context) (grpc.ClientStream, error) {
            return streamer(ctx, desc, cc, method, opts...)
        })
        if err != nil {
            return nil, err
        }
        return result, nil
    }
}
```

## Observability

### OpenTelemetry Integration

Full W3C TraceContext support with automatic span creation:

```go
package otel

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/codes"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/trace"
)

// TracedCircuitBreaker wraps operations with tracing
type TracedCircuitBreaker[T any] struct {
    cb     circuitbreaker.CircuitBreaker[T]
    tracer trace.Tracer
}

func (t *TracedCircuitBreaker[T]) Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
    ctx, span := t.tracer.Start(ctx, "circuit_breaker",
        trace.WithAttributes(
            attribute.String("pattern", "circuit_breaker"),
        ))
    defer span.End()

    result, err := t.cb.Execute(ctx, fn)

    if err != nil {
        span.SetStatus(codes.Error, err.Error())
        span.RecordError(err)
    }

    span.SetAttributes(
        attribute.String("state", t.cb.State().String()),
    )

    return result, err
}

// Configure W3C TraceContext propagation
func SetupPropagation() {
    otel.SetTextMapPropagator(
        propagation.NewCompositeTextMapPropagator(
            propagation.TraceContext{},
            propagation.Baggage{},
        ),
    )
}
```

### Structured Logging with log/slog

Native integration with Go's structured logging:

```go
package slog

import (
    "context"
    "log/slog"
    "go.opentelemetry.io/otel/trace"
)

// ContextHandler enriches logs with trace context
type ContextHandler struct {
    handler slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
    // Add trace context to logs
    if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
        r.AddAttrs(
            slog.String("trace_id", span.SpanContext().TraceID().String()),
            slog.String("span_id", span.SpanContext().SpanID().String()),
        )
    }
    return h.handler.Handle(ctx, r)
}

// Usage
logger := slog.New(&ContextHandler{
    handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }),
})
```

### Metrics Interface

Pluggable metrics for any backend:

```go
package metrics

type MetricsRecorder interface {
    RecordCircuitBreakerState(name string, state State)
    RecordRetryAttempt(name string, attempt int, success bool)
    RecordRateLimitHit(name string)
    RecordBulkheadUtilization(name string, active, queued int)
    RecordTimeout(name string)
}

// Prometheus implementation
type PrometheusRecorder struct {
    circuitBreakerState *prometheus.GaugeVec
    retryAttempts       *prometheus.CounterVec
    rateLimitHits       *prometheus.CounterVec
    bulkheadActive      *prometheus.GaugeVec
}

// OpenTelemetry implementation
type OTelRecorder struct {
    meter metric.Meter
}
```

## Error Handling

### Standard Error Types

```go
package errors

import (
    "errors"
    "fmt"
)

var (
    ErrCircuitOpen        = errors.New("circuit breaker is open")
    ErrRateLimitExceeded  = errors.New("rate limit exceeded")
    ErrBulkheadFull       = errors.New("bulkhead at capacity")
    ErrTimeout            = errors.New("operation timeout")
    ErrMaxAttemptsReached = errors.New("max retry attempts reached")
)

// Retryable error interface
type RetryableError interface {
    error
    Retryable() bool
}

type retryableError struct {
    err error
}

func (e *retryableError) Error() string {
    return e.err.Error()
}

func (e *retryableError) Retryable() bool {
    return true
}

func (e *retryableError) Unwrap() error {
    return e.err
}

func AsRetryable(err error) error {
    return &retryableError{err: err}
}

// Check if error is retryable
func IsRetryable(err error) bool {
    var re RetryableError
    return errors.As(err, &re) && re.Retryable()
}
```

## Configuration Patterns

### Builder Pattern

```go
cb := circuitbreaker.NewBuilder[*Response]().
    WithMaxRequests(5).
    WithTimeout(60 * time.Second).
    WithInterval(10 * time.Second).
    WithReadyToTrip(func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures > 5
    }).
    WithLogger(slog.Default()).
    WithTracer(otel.Tracer("my-service")).
    Build()
```

### Functional Options

```go
cb := circuitbreaker.New[*Response](
    circuitbreaker.WithMaxRequests(5),
    circuitbreaker.WithTimeout(60 * time.Second),
    circuitbreaker.WithInterval(10 * time.Second),
    circuitbreaker.WithLogger(slog.Default()),
    circuitbreaker.WithTracer(otel.Tracer("my-service")),
)
```

## Testing Strategy

### Unit Tests

```go
func TestCircuitBreakerOpensAfterFailures(t *testing.T) {
    cb := circuitbreaker.New[int](circuitbreaker.Config{
        MaxRequests: 1,
        Interval:    time.Second,
        Timeout:     time.Second,
        ReadyToTrip: func(counts circuitbreaker.Counts) bool {
            return counts.ConsecutiveFailures >= 3
        },
    })

    // Trigger failures
    for i := 0; i < 3; i++ {
        _, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
            return 0, errors.New("failure")
        })
        require.Error(t, err)
    }

    // Circuit should be open
    assert.Equal(t, circuitbreaker.StateOpen, cb.State())
}
```

### Benchmark Tests

```go
func BenchmarkCircuitBreakerSuccess(b *testing.B) {
    cb := circuitbreaker.New[int](circuitbreaker.Config{
        MaxRequests: 100,
        Interval:    time.Second,
    })

    ctx := context.Background()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
            return 42, nil
        })
    }
}
```

### Fuzz Tests (Go 1.18+)

```go
func FuzzRetryBackoff(f *testing.F) {
    f.Add(int64(100), float64(2.0), 5)

    f.Fuzz(func(t *testing.T, initialDelay int64, multiplier float64, maxAttempts int) {
        if initialDelay <= 0 || multiplier <= 0 || maxAttempts <= 0 {
            t.Skip()
        }

        r := retry.New[int](retry.Config{
            InitialDelay: time.Duration(initialDelay),
            Multiplier:   multiplier,
            MaxAttempts:  maxAttempts,
        })

        // Should not panic
        _, _ = r.Do(context.Background(), func(ctx context.Context) (int, error) {
            return 0, errors.New("test")
        })
    })
}
```

### Race Detection

```bash
go test -race ./...
```

## Performance Targets

- **Overhead:** < 1µs for fast-path operations (circuit closed, no retry)
- **Allocations:** Zero allocations in hot paths where possible
- **Memory:** < 100 bytes per pattern instance
- **Throughput:** > 1M ops/sec on modern hardware

### Performance Notes

- **Generic Performance:** When using generics with interfaces, consider using pointers for better performance
- **Context Overhead:** Context propagation adds ~50ns per hop
- **Logging:** Use leveled logging (debug/info) to control overhead
- **Tracing:** Sampling recommended for high-throughput services (1-10%)

## Migration Guide

### From sony/gobreaker

```go
// Old
breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
    Name:        "MyService",
    MaxRequests: 5,
    Interval:    time.Second * 10,
    Timeout:     time.Second * 60,
})

result, err := breaker.Execute(func() (interface{}, error) {
    return callService()
})

// New
cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    MaxRequests: 5,
    Interval:    10 * time.Second,
    Timeout:     60 * time.Second,
    Logger:      slog.Default(),
})

result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return callService(ctx)
})
```

## Best Practices

### 1. Always Use Context

```go
// ❌ Bad
func callAPI() (*Response, error) {
    return http.Get("https://api.example.com")
}

// ✅ Good
func callAPI(ctx context.Context) (*Response, error) {
    req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
    return http.DefaultClient.Do(req)
}
```

### 2. Configure Timeouts Appropriately

```go
// Stack timeouts: overall < circuit breaker < retry < individual operation
timeout.Execute(ctx, 10*time.Second, func(ctx context.Context) {
    cb.Execute(ctx, func(ctx context.Context) {
        retry.Do(ctx, func(ctx context.Context) {
            callAPI(ctx) // 1s timeout
        })
    })
})
```

### 3. Use Structured Logging

```go
logger.InfoContext(ctx, "circuit breaker state changed",
    slog.String("from", oldState.String()),
    slog.String("to", newState.String()),
    slog.String("service", "payment-api"),
)
```

### 4. Classify Errors Properly

```go
retry.Config{
    RetryableErrors: []error{
        context.DeadlineExceeded,
        ErrServiceUnavailable,
        ErrNetworkTimeout,
    },
    NonRetryableErrors: []error{
        ErrBadRequest,
        ErrUnauthorized,
        ErrNotFound,
    },
}
```

### 5. Monitor Circuit Breaker States

```go
cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        metrics.RecordStateChange("payment-api", from, to)
        if to == circuitbreaker.StateOpen {
            alerts.SendAlert("Circuit breaker opened for payment-api")
        }
    },
})
```

## Documentation Requirements

- [ ] Comprehensive godoc for all exported types
- [ ] Usage examples for each pattern
- [ ] Migration guide from popular alternatives
- [ ] Performance characteristics with benchmarks
- [ ] Best practices guide
- [ ] OpenTelemetry integration examples
- [ ] Troubleshooting guide

## Future Enhancements (v2+)

- Adaptive circuit breakers (ML-based thresholds)
- Distributed bulkhead with Redis coordination
- Advanced rate limiting (sliding log, leaky bucket variants)
- Fallback pattern support
- Health check integration
- Dynamic configuration reloading via context
- Prometheus metrics built-in
- Grafana dashboard templates

## Success Metrics

- < 1µs overhead for fast-path operations
- Zero allocations in hot paths where possible
- 100% test coverage for core logic
- Adoption by 10+ teams within first quarter
- 90+ score on Go Report Card
- < 10 open issues after first month

## License

Apache 2.0 or MIT (dual license for maximum adoption)

## Contributing

See CONTRIBUTING.md for:

- Code style guide (gofmt, golangci-lint)
- Testing requirements (unit, integration, benchmarks, fuzzing)
- Documentation standards
- Review process
- Release procedures

## References

- [Go 1.23 Release Notes](https://go.dev/doc/go1.23)
- [log/slog Documentation](https://pkg.go.dev/log/slog)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
- [W3C Trace Context](https://www.w3.org/TR/trace-context/)
- [Microservices Resilience Patterns](https://learn.microsoft.com/en-us/azure/architecture/patterns/)