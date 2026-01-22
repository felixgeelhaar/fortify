<div align="center">
  <img src="assets/fortify.png" alt="Fortify Logo" width="200"/>
  <h1>Fortify</h1>
</div>

[![Go Reference](https://pkg.go.dev/badge/github.com/felixgeelhaar/fortify.svg)](https://pkg.go.dev/github.com/felixgeelhaar/fortify)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixgeelhaar/fortify)](https://goreportcard.com/report/github.com/felixgeelhaar/fortify)
[![CI Status](https://github.com/felixgeelhaar/fortify/workflows/CI/badge.svg)](https://github.com/felixgeelhaar/fortify/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-84.9%25-brightgreen)](https://github.com/felixgeelhaar/fortify)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/felixgeelhaar/fortify)](https://github.com/felixgeelhaar/fortify)
[![Release](https://img.shields.io/github/v/release/felixgeelhaar/fortify)](https://github.com/felixgeelhaar/fortify/releases)

Fortify is a production-grade resilience and fault-tolerance library for Go 1.23+. It provides a comprehensive suite of battle-tested patterns including circuit breakers, retries, rate limiting, timeouts, and bulkheads with zero external dependencies for core functionality.

## Features

- **🔒 Type-Safe**: Built with Go 1.23+ generics for compile-time safety
- **⚡ High Performance**: <1µs overhead with zero allocations in hot paths
- **🎯 Zero Dependencies**: Core patterns have no external dependencies
- **🔍 Observable**: Built-in support for structured logging (slog) and OpenTelemetry
- **📊 Prometheus Metrics**: Export metrics for all resilience patterns
- **🌐 Framework Integration**: First-class support for HTTP and gRPC
- **🧩 Composable**: Fluent API for combining multiple patterns
- **🧪 Well Tested**: >95% test coverage with race detection
- **🌪️ Chaos Engineering**: Built-in testing utilities for resilience validation
- **📈 Performance Testing**: Automated regression detection and benchmarking
- **📊 Production Ready**: Battle-tested patterns with comprehensive examples

## Installation

```bash
go get github.com/felixgeelhaar/fortify
```

**Requirements:** Go 1.23 or higher

## Quick Start

```go
package main

import (
    "context"
    "time"

    "github.com/felixgeelhaar/fortify/circuitbreaker"
    "github.com/felixgeelhaar/fortify/retry"
)

func main() {
    // Create a circuit breaker
    cb := circuitbreaker.New[string](circuitbreaker.Config{
        MaxRequests: 100,
        Interval:    time.Second * 10,
        ReadyToTrip: func(counts circuitbreaker.Counts) bool {
            return counts.ConsecutiveFailures >= 3
        },
    })

    // Create a retry strategy
    r := retry.New[string](retry.Config{
        MaxAttempts:   3,
        InitialDelay:  time.Millisecond * 100,
        BackoffPolicy: retry.BackoffExponential,
    })

    // Use them together
    result, err := cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
        return r.Do(ctx, func(ctx context.Context) (string, error) {
            return callExternalService(ctx)
        })
    })
}
```

## Patterns

### Circuit Breaker

Prevents cascading failures by temporarily stopping requests to failing services.

```go
import "github.com/felixgeelhaar/fortify/circuitbreaker"

cb := circuitbreaker.New[Response](circuitbreaker.Config{
    MaxRequests: 100,
    Interval:    time.Second * 60,
    Timeout:     time.Second * 30, // Half-open timeout
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 10 && failureRatio >= 0.5
    },
    OnStateChange: func(from, to circuitbreaker.State) {
        log.Printf("Circuit breaker: %s -> %s", from, to)
    },
})

result, err := cb.Execute(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})
```

**States:** Closed → Open → Half-Open → Closed

**Use Cases:**
- Protecting against cascading failures
- Preventing resource exhaustion
- Fast failure for unhealthy dependencies

### Retry

Automatically retries failed operations with configurable backoff strategies.

```go
import "github.com/felixgeelhaar/fortify/retry"

r := retry.New[Response](retry.Config{
    MaxAttempts:   5,
    InitialDelay:  time.Millisecond * 100,
    MaxDelay:      time.Second * 10,
    BackoffPolicy: retry.BackoffExponential,
    Multiplier:    2.0,
    Jitter:        true,
    ShouldRetry: func(err error) bool {
        return isTransientError(err)
    },
})

result, err := r.Do(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})
```

**Backoff Policies:**
- `BackoffConstant`: Fixed delay between retries
- `BackoffLinear`: Linearly increasing delay
- `BackoffExponential`: Exponentially increasing delay

**Use Cases:**
- Handling transient network failures
- Dealing with rate-limited APIs
- Recovering from temporary service unavailability

### Rate Limiting

Controls the rate of operations using a token bucket algorithm with a pluggable storage backend.

```go
import "github.com/felixgeelhaar/fortify/ratelimit"

rl := ratelimit.New(&ratelimit.Config{
    Rate:     100,               // 100 requests
    Burst:    200,               // burst of 200
    Interval: time.Second,       // per second
})

// Non-blocking check
if rl.Allow(ctx, "user-123") {
    handleRequest()
}

// Blocking wait
if err := rl.Wait(ctx, "user-123"); err == nil {
    handleRequest()
}
```

**Use Cases:**
- Protecting APIs from abuse
- Ensuring fair resource usage
- Implementing user quotas

#### Per-Key Rate Limiting Patterns

The rate limiter supports flexible per-key limiting out of the box. Use different keys to implement various rate limiting strategies:

```go
import "github.com/felixgeelhaar/fortify/ratelimit"

rl := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
})

// Global rate limiting - all requests share one bucket
func handleGlobalLimit(ctx context.Context) bool {
    return rl.Allow(ctx, "global")
}

// Per-method rate limiting - separate limits per API endpoint
func handleMethodLimit(ctx context.Context, method string) bool {
    return rl.Allow(ctx, method) // e.g., "GET:/api/users", "POST:/api/orders"
}

// Per-client rate limiting - separate limits per user/client
func handleClientLimit(ctx context.Context, clientID string) bool {
    return rl.Allow(ctx, clientID) // e.g., "user-123", "api-key-abc"
}

// Combined: per-client-per-method limiting
func handleCombinedLimit(ctx context.Context, clientID, method string) bool {
    key := clientID + ":" + method // e.g., "user-123:POST:/api/orders"
    return rl.Allow(ctx, key)
}
```

**Dynamic Key Extraction with KeyFunc:**

Use `KeyFunc` to extract rate limit keys from context automatically:

```go
rl := ratelimit.New(&ratelimit.Config{
    Rate:     10,
    Burst:    20,
    Interval: time.Second,
    KeyFunc: func(ctx context.Context) string {
        // Extract user ID from context (set by auth middleware)
        if userID, ok := ctx.Value("userID").(string); ok {
            return userID
        }
        // Fallback to IP address for unauthenticated requests
        if ip, ok := ctx.Value("clientIP").(string); ok {
            return "ip:" + ip
        }
        return "anonymous"
    },
})

// Key is extracted automatically from context
if rl.Allow(ctx, "") { // key parameter ignored when KeyFunc is set
    handleRequest()
}
```

#### Event Callbacks for Observability

Monitor rate limiting events using the `OnAllow` and `OnLimit` callbacks:

```go
rl := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    OnAllow: func(ctx context.Context, key string) {
        // Called when a request is allowed
        metrics.IncrementCounter("rate_limit_allowed", key)
    },
    OnLimit: func(ctx context.Context, key string) {
        // Called when a request is rate limited
        log.Printf("Rate limited: %s", key)
        metrics.IncrementCounter("rate_limit_exceeded", key)
    },
    Logger: slog.Default(), // Structured logging for all events
})
```

For comprehensive observability, implement the `Metrics` interface:

```go
type MyMetrics struct{}

func (m *MyMetrics) OnAllow(ctx context.Context, key string) {
    // Request was allowed
}

func (m *MyMetrics) OnDeny(ctx context.Context, key string) {
    // Request was rate limited
}

func (m *MyMetrics) OnError(ctx context.Context, key string, err error) {
    // Storage error occurred
}

func (m *MyMetrics) OnStoreLatency(ctx context.Context, op string, d time.Duration) {
    // Track storage operation latency
}

rl := ratelimit.New(&ratelimit.Config{
    Rate:    100,
    Burst:   200,
    Metrics: &MyMetrics{},
})
```

#### Custom Storage Backends

For distributed systems, implement the `Store` interface to share rate limits across multiple application instances:

```go
import "github.com/felixgeelhaar/fortify/ratelimit"

// Implement the Store interface for your backend (Redis, DynamoDB, etc.)
type RedisStore struct {
    client redis.UniversalClient
    prefix string
    ttl    time.Duration
}

func (r *RedisStore) AtomicUpdate(ctx context.Context, key string,
    updateFn func(*ratelimit.BucketState) *ratelimit.BucketState) (*ratelimit.BucketState, error) {
    // Use WATCH/MULTI/EXEC or Lua scripts for atomic operations
    // See examples for full implementation
}

func (r *RedisStore) Delete(ctx context.Context, key string) error { /* ... */ }
func (r *RedisStore) Close() error { /* ... */ }

// Use custom store
rl := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    Store:    &RedisStore{client: redisClient, prefix: "rl:", ttl: time.Hour},
    FailOpen: true,  // Allow requests if storage fails
})
```

**Store Interface:**
- `AtomicUpdate`: Atomic read-modify-write for token bucket state
- `Delete`: Remove a bucket from storage
- `Close`: Release resources

**Configuration:**
- `Store`: Custom storage backend (defaults to in-memory)
- `FailOpen`: Allow requests when storage fails (availability over consistency)

### Timeout

Enforces time limits on operations with context propagation.

```go
import "github.com/felixgeelhaar/fortify/timeout"

tm := timeout.New[Response](timeout.Config{
    DefaultTimeout: time.Second * 30,
    OnTimeout: func(duration time.Duration) {
        log.Printf("Operation timed out after %v", duration)
    },
})

// Use specific timeout
result, err := tm.Execute(ctx, 5*time.Second, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})

// Use default timeout
result, err := tm.ExecuteWithDefault(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})
```

**Use Cases:**
- Enforcing SLA response times
- Preventing resource leaks
- Setting operation deadlines

### Bulkhead

Limits concurrent operations to prevent resource exhaustion.

```go
import "github.com/felixgeelhaar/fortify/bulkhead"

bh := bulkhead.New[Response](bulkhead.Config{
    MaxConcurrent: 10,                  // Max concurrent operations
    MaxQueue:      20,                  // Max queued operations
    QueueTimeout:  time.Second * 5,     // Queue wait timeout
    OnRejected: func() {
        log.Println("Request rejected: bulkhead full")
    },
})

result, err := bh.Execute(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})

// Get statistics
stats := bh.Stats()
log.Printf("Active: %d, Queued: %d, Rejected: %d",
    stats.ActiveRequests, stats.QueuedRequests, stats.RejectedRequests)
```

**Use Cases:**
- Preventing resource exhaustion
- Isolating critical operations
- Managing concurrent access

### Fallback

Provides graceful degradation with automatic fallback on errors.

```go
import "github.com/felixgeelhaar/fortify/fallback"

fb := fallback.New[Response](fallback.Config{
    Primary: func(ctx context.Context) (Response, error) {
        return primaryService.Call(ctx)
    },
    Fallback: func(ctx context.Context, err error) (Response, error) {
        log.Printf("Primary failed: %v, using fallback", err)
        return fallbackService.Call(ctx)
    },
    ShouldFallback: func(err error) bool {
        return isServiceError(err) // Only fallback on service errors
    },
    OnFallback: func(err error) {
        metrics.IncFallbackCount()
    },
})

result, err := fb.Execute(ctx)
```

**Use Cases:**
- Graceful service degradation
- Multi-tier service architectures
- Cache fallback strategies

## Middleware Composition

Combine multiple patterns into a single execution chain:

```go
import "github.com/felixgeelhaar/fortify/middleware"

chain := middleware.New[Response]().
    WithBulkhead(bh).
    WithRateLimit(rl, "user-key").
    WithTimeout(tm, 5*time.Second).
    WithCircuitBreaker(cb).
    WithRetry(r)

result, err := chain.Execute(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})
```

**Order matters:**
1. Bulkhead - Limit concurrency first
2. Rate Limit - Check quotas
3. Timeout - Enforce time limits
4. Circuit Breaker - Check service health
5. Retry - Retry on failures

## HTTP Middleware

Integrate resilience patterns with standard `http.Handler`:

```go
import (
    "net/http"
    fortifyhttp "github.com/felixgeelhaar/fortify/http"
)

// Create patterns
cb := circuitbreaker.New[*http.Response](/* config */)
rl := ratelimit.New(/* config */)
tm := timeout.New[*http.Response](/* config */)

// Apply middleware
handler := fortifyhttp.RateLimit(rl, fortifyhttp.KeyFromIP)(
    fortifyhttp.Timeout(tm, 5*time.Second)(
        fortifyhttp.CircuitBreaker(cb)(
            http.HandlerFunc(myHandler),
        ),
    ),
)

http.Handle("/api", handler)
```

**Key Extractors:**
- `KeyFromIP` - Extract client IP
- `KeyFromHeader(name)` - Extract from HTTP header

**Status Codes:**
- `503 Service Unavailable` - Circuit breaker open
- `429 Too Many Requests` - Rate limit exceeded
- `504 Gateway Timeout` - Request timeout

## gRPC Interceptors

Integrate with gRPC services:

```go
import (
    fortifygrpc "github.com/felixgeelhaar/fortify/grpc"
    "google.golang.org/grpc"
)

// Unary interceptors
server := grpc.NewServer(
    grpc.UnaryInterceptor(
        fortifygrpc.UnaryCircuitBreakerInterceptor(cb),
    ),
    grpc.StreamInterceptor(
        fortifygrpc.StreamRateLimitInterceptor(rl,
            fortifygrpc.StreamKeyFromMetadata("x-api-key")),
    ),
)
```

**Interceptors:**
- `UnaryCircuitBreakerInterceptor`
- `UnaryRateLimitInterceptor`
- `UnaryTimeoutInterceptor`
- `StreamCircuitBreakerInterceptor`
- `StreamRateLimitInterceptor`
- `StreamTimeoutInterceptor`

## Observability

### Structured Logging

```go
import (
    "log/slog"
    fortifyslog "github.com/felixgeelhaar/fortify/slog"
)

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

fortifyslog.LogPatternEvent(logger, fortifyslog.PatternCircuitBreaker, "state_change",
    slog.String("from", "closed"),
    slog.String("to", "open"),
)
```

### OpenTelemetry Tracing

```go
import (
    fortifyotel "github.com/felixgeelhaar/fortify/otel"
    "go.opentelemetry.io/otel/sdk/trace"
)

provider := trace.NewTracerProvider(/* config */)
tracer := fortifyotel.NewTracer(provider, "my-service")

ctx, span := tracer.StartSpan(ctx, fortifyotel.PatternCircuitBreaker, "execute")
defer span.End()

tracer.SetAttributes(span,
    attribute.Int("requests", 100),
    attribute.String("state", "closed"),
)
```

### Prometheus Metrics

Export detailed metrics for all resilience patterns:

```go
import (
    "github.com/felixgeelhaar/fortify/metrics"
    "github.com/prometheus/client_golang/prometheus"
)

// Register Fortify metrics with Prometheus
metrics.MustRegister(prometheus.DefaultRegisterer)

// Use the default collector
collector := metrics.DefaultCollector()

// Record circuit breaker metrics
collector.RecordCircuitBreakerRequest("api-client", "closed")
collector.RecordCircuitBreakerSuccess("api-client")

// Record retry metrics
collector.RecordRetryAttempts("database-query", 2)
collector.RecordRetrySuccess("database-query")
```

**Available Metrics:**
- Circuit Breaker: state, requests, successes, failures, state changes
- Retry: attempts, duration, successes, failures
- Rate Limit: allowed/denied requests, wait times
- Timeout: executions, exceeded, durations
- Bulkhead: active/queued requests, rejections, durations

## Performance

Fortify is designed for production use with minimal overhead:

| Pattern          | Overhead | Allocations | Notes |
|-----------------|----------|-------------|-------|
| Circuit Breaker | ~30ns    | 0           | State check only |
| Retry           | ~25ns    | 0           | Per attempt overhead |
| Rate Limiter    | ~200ns   | 3           | Full token bucket with atomic update |
| Timeout         | ~50ns    | 0           | Context wrapping |
| Bulkhead        | ~39ns    | 0           | Semaphore check |

*Benchmarks on Apple M1, Go 1.23*

**Rate Limiter Details:**
- `Allow()`: ~200ns, 74B, 3 allocs (token bucket state + slog attributes)
- `Take()`: ~197ns, 65B, 3 allocs
- `BucketCount()`: ~3ns, 0 allocs (O(1) atomic read)
- Concurrent operations: ~395ns with contention (8 goroutines)

## Examples

Comprehensive examples are available in the [`examples/`](./examples/) directory:

- **Basic**: Individual pattern usage
  - [Circuit Breaker](./examples/basic/circuit_breaker.go)
  - [Retry](./examples/basic/retry.go)
  - [Rate Limiting](./examples/basic/rate_limit.go)
  - [Timeout](./examples/basic/timeout.go)
  - [Bulkhead](./examples/basic/bulkhead.go)

- **HTTP**: Web server integration
  - [HTTP Middleware Server](./examples/http/server.go)

- **Composition**: Advanced patterns
  - [Middleware Chain](./examples/composition/chain.go)

Run examples:
```bash
go run examples/basic/circuit_breaker.go
go run examples/http/server.go
go run examples/composition/chain.go
```

## Best Practices

### Pattern Selection

- **Circuit Breaker**: Use for external dependencies that can fail
- **Retry**: Use for transient failures (network issues, rate limits)
- **Rate Limiter**: Use to protect your API from overload
- **Timeout**: Use to enforce SLAs and prevent resource leaks
- **Bulkhead**: Use to isolate critical operations

### Configuration

- **Circuit Breaker**: Tune `ReadyToTrip` based on your error budget
- **Retry**: Use exponential backoff with jitter for distributed systems
- **Rate Limiter**: Set burst capacity for handling traffic spikes
- **Timeout**: Set timeouts based on p99 latency + buffer
- **Bulkhead**: Size based on available resources and expected load

### Composition Order

Recommended order for combining patterns:

1. **Bulkhead** - Limit concurrency to prevent resource exhaustion
2. **Rate Limit** - Check quotas before processing
3. **Timeout** - Set operation deadline
4. **Circuit Breaker** - Check service health
5. **Retry** - Handle transient failures

### Observability

- Always configure `OnStateChange`, `OnRetry`, `OnTimeout`, and `OnRejected` callbacks
- Use structured logging for better debugging
- Integrate OpenTelemetry for distributed tracing
- Monitor pattern metrics in production

## Testing

### Unit and Integration Tests

Run tests with race detection:

```bash
# All tests
go test -v -race ./...

# Specific package
go test -v -race ./circuitbreaker

# With coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Chaos Engineering

Test resilience with built-in chaos utilities:

```go
import fortifytesting "github.com/felixgeelhaar/fortify/testing"

// Inject errors with configurable probability
injector := fortifytesting.NewErrorInjector(0.3, errors.New("service unavailable"))

// Add network latency
latency := fortifytesting.NewLatencyInjector(10*time.Millisecond, 50*time.Millisecond)

// Simulate timeouts
timeout := fortifytesting.NewTimeoutSimulator(100*time.Millisecond, 0.5)

// Create flakey service combining all
service := fortifytesting.NewFlakeyService(0.3, 10*time.Millisecond, 30*time.Millisecond)
```

**Chaos Utilities:**
- `ErrorInjector`: Simulate failures with probability
- `LatencyInjector`: Add realistic network delays
- `TimeoutSimulator`: Create timeout scenarios
- `FlakeyService`: Combine errors, latency, and timeouts

### Performance Regression Testing

Automated benchmark tracking and regression detection:

```bash
# Run benchmarks with automation
./scripts/benchmark.sh run

# Generate performance baseline
./scripts/benchmark.sh generate-baseline

# Check for regressions
./scripts/benchmark.sh check

# Complete workflow
./scripts/benchmark.sh all
```

**Features:**
- Automatic regression detection (time, allocations, memory)
- Configurable thresholds (10% time, 20% allocs, 15% memory)
- Historical tracking with JSON storage
- CI/CD integration with GitHub Actions
- Detailed performance reports

See [Performance Testing Guide](./docs/PERFORMANCE_TESTING.md) for details.

### Benchmarks

Run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass with race detection
5. Submit a pull request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

Fortify is inspired by resilience libraries from other ecosystems:

- [Hystrix](https://github.com/Netflix/Hystrix) (Java/Netflix)
- [resilience4j](https://github.com/resilience4j/resilience4j) (Java)
- [Polly](https://github.com/App-vNext/Polly) (.NET)

## Support

- 📖 [Documentation](https://pkg.go.dev/github.com/felixgeelhaar/fortify)
- 🐛 [Issue Tracker](https://github.com/felixgeelhaar/fortify/issues)
- 💬 [Discussions](https://github.com/felixgeelhaar/fortify/discussions)

---

Built with ❤️ by [Felix Geelhaar](https://github.com/felixgeelhaar)