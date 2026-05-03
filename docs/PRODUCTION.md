# Production Deployment Guide

This guide covers best practices for deploying Fortify resilience patterns in production environments.

## Table of Contents

- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Observability](#observability)
- [Performance Tuning](#performance-tuning)
- [Security](#security)
- [Monitoring](#monitoring)
- [Troubleshooting](#troubleshooting)

## Quick Start

### Prerequisites

- Go 1.23+ (for generics support)
- Understanding of your service's SLOs and error budgets
- Observability infrastructure (logging, metrics, tracing)

### Basic Setup

```go
import (
    "github.com/felixgeelhaar/fortify/circuitbreaker"
    "github.com/felixgeelhaar/fortify/retry"
    "github.com/felixgeelhaar/fortify/middleware"
)

// Circuit Breaker for external API
cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    MaxRequests: 5,
    Interval:    time.Minute,
    Timeout:     30 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 10 && failureRatio >= 0.6
    },
    Logger: slog.Default(),
})

// Retry with exponential backoff
r := retry.New[*Response](&retry.Config{
    MaxAttempts:   3,
    InitialDelay:  100 * time.Millisecond,
    MaxDelay:      5 * time.Second,
    Multiplier:    2.0,
    BackoffPolicy: retry.BackoffExponential,
    Jitter:        true,
    Logger:        slog.Default(),
})
```

## Configuration

### Circuit Breaker Settings

**Production-Ready Configuration:**

```go
cb := circuitbreaker.New[T](circuitbreaker.Config{
    // Allow 10 requests in half-open state
    MaxRequests: 10,

    // Clear counts every 60 seconds in closed state
    Interval: time.Minute,

    // Open -> Half-Open after 30 seconds
    Timeout: 30 * time.Second,

    // Trip on 60% failure rate with minimum 20 requests
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        if counts.Requests < 20 {
            return false
        }
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return failureRatio >= 0.6
    },

    // Log state changes for monitoring
    OnStateChange: func(from, to circuitbreaker.State) {
        slog.Warn("circuit breaker state change",
            slog.String("from", from.String()),
            slog.String("to", to.String()),
        )
        // Emit metric here
    },

    Logger: slog.Default(),
})
```

**Key Considerations:**

- `ReadyToTrip`: Balance sensitivity vs stability. Too sensitive = unnecessary trips, too lenient = service degradation
- `Interval`: Should align with your service's error budget window
- `Timeout`: Consider downstream service recovery time
- `MaxRequests`: Higher values give better signal in half-open, but risk more failures

### Retry Settings

**Production-Ready Configuration:**

```go
r := retry.New[T](&retry.Config{
    // Maximum 3 total attempts (1 initial + 2 retries)
    MaxAttempts: 3,

    // Start with 100ms delay
    InitialDelay: 100 * time.Millisecond,

    // Cap at 5 seconds to prevent excessive waits
    MaxDelay: 5 * time.Second,

    // Exponential backoff with 2x multiplier
    Multiplier: 2.0,
    BackoffPolicy: retry.BackoffExponential,

    // Add jitter to prevent thundering herd
    Jitter: true,

    // Only retry on specific errors
    IsRetryable: func(err error) bool {
        return errors.Is(err, ErrTemporary) ||
               errors.Is(err, ErrRateLimited)
    },

    // Log retries for monitoring
    OnRetry: func(attempt int, err error) {
        slog.Info("retrying operation",
            slog.Int("attempt", attempt),
            slog.String("error", err.Error()),
        )
    },

    Logger: slog.Default(),
})
```

**Key Considerations:**

- `MaxAttempts`: Balance user experience vs success probability
- `Jitter`: Always enable in production to prevent thundering herd
- `IsRetryable`: Only retry errors that are likely to succeed on retry
- `MaxDelay`: Prevent indefinite waits that degrade UX

### Rate Limiter Settings

**Production-Ready Configuration:**

```go
rl := ratelimit.New(ratelimit.Config{
    // 100 requests per second per key
    Rate: 100,

    // Allow bursts up to 150 requests
    Burst: 150,

    // Rate calculation interval
    Interval: time.Second,

    // Extract key from request context
    KeyFunc: func(ctx context.Context) string {
        // Use user ID, API key, or IP address
        userID := ctx.Value("user_id").(string)
        return userID
    },

    // Log rate limit events
    OnLimit: func(key string) {
        slog.Warn("rate limit exceeded",
            slog.String("key", key),
        )
        // Emit metric here
    },

    Logger: slog.Default(),
})
```

**Key Considerations:**

- `Rate` and `Burst`: Set based on service capacity and fair use policy
- `KeyFunc`: Choose appropriate key (user, tenant, IP) for your use case
- `Interval`: Typically 1 second for straightforward rate limiting

### Timeout Settings

**Production-Ready Configuration:**

```go
tm := timeout.New[T](timeout.Config{
    // Default timeout for all operations
    DefaultTimeout: 5 * time.Second,

    // Log timeout events
    OnTimeout: func() {
        slog.Warn("operation timed out")
        // Emit metric here
    },

    Logger: slog.Default(),
})

// Use with specific timeout per operation
result, err := tm.Execute(ctx, 10*time.Second, func(ctx context.Context) (T, error) {
    // Operation that may take up to 10 seconds
})
```

**Key Considerations:**

- `DefaultTimeout`: Set based on P99 latency of downstream services
- Per-operation timeouts: Override default for specific slow operations
- Always respect parent context deadlines

### Bulkhead Settings

**Production-Ready Configuration:**

```go
bh := bulkhead.New[T](bulkhead.Config{
    // Maximum 50 concurrent executions
    MaxConcurrent: 50,

    // Queue up to 100 requests when at capacity
    MaxQueue: 100,

    // Timeout for queued requests
    QueueTimeout: 10 * time.Second,

    // Log rejections
    OnRejected: func() {
        slog.Warn("bulkhead rejected request")
        // Emit metric here
    },

    Logger: slog.Default(),
})
```

**Key Considerations:**

- `MaxConcurrent`: Set based on service capacity and resource limits
- `MaxQueue`: Prevent unbounded queue growth
- `QueueTimeout`: Balance queueing vs failing fast

## Observability

### Structured Logging

**Best Practices:**

```go
// Use structured logging with context
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))

cb := circuitbreaker.New[T](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        logger.Warn("circuit breaker state change",
            slog.String("from", from.String()),
            slog.String("to", to.String()),
            slog.String("service", "api-client"),
            slog.Time("timestamp", time.Now()),
        )
    },
    Logger: logger,
})
```

### Metrics Integration

**Prometheus Example:**

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    circuitBreakerState = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "circuit_breaker_state",
            Help: "Circuit breaker state (0=closed, 1=open, 2=half-open)",
        },
        []string{"service"},
    )

    retryAttempts = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "retry_attempts_total",
            Help: "Total retry attempts",
        },
        []string{"service", "success"},
    )
)

cb := circuitbreaker.New[T](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        var state float64
        switch to {
        case circuitbreaker.StateClosed:
            state = 0
        case circuitbreaker.StateOpen:
            state = 1
        case circuitbreaker.StateHalfOpen:
            state = 2
        }
        circuitBreakerState.WithLabelValues("api-client").Set(state)
    },
})
```

### Distributed Tracing

**OpenTelemetry Integration:**

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

// Add span attributes in callbacks
cb := circuitbreaker.New[T](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        span := trace.SpanFromContext(ctx)
        span.AddEvent("circuit_breaker_state_change",
            trace.WithAttributes(
                attribute.String("from_state", from.String()),
                attribute.String("to_state", to.String()),
            ),
        )
    },
})
```

## Performance Tuning

### Memory Optimization

**Rate Limiter:**
- Each key creates a token bucket (~100 bytes)
- Implement key expiration for inactive buckets
- Monitor memory usage with many unique keys

**Circuit Breaker:**
- Minimal memory footprint (~200 bytes per instance)
- Stateless in Open state (discards requests immediately)
- Counts cleared on interval in Closed state

### Goroutine Management

**Bulkhead Worker Pool:**
```go
// Worker pool is created on-demand
bh := bulkhead.New[T](bulkhead.Config{
    MaxConcurrent: 50,  // Creates worker pool
    MaxQueue: 100,      // No additional goroutines
})
defer bh.Close()  // Important: cleanup workers
```

**Best Practices:**
- Always call `Close()` on bulkheads when done
- Use context cancellation to prevent goroutine leaks
- Monitor goroutine count in production

### Latency Optimization

**Fast Path:**
- Circuit breaker closed state: ~100ns overhead
- Rate limiter with available tokens: ~200ns overhead
- Timeout with no timeout: ~50ns overhead

**Optimization Tips:**
1. Minimize callback complexity (called on hot path)
2. Use pointer receivers for large structs
3. Reuse pattern instances (don't create per-request)
4. Profile with `go test -bench=. -benchmem`

## Security

### Input Validation

```go
// Validate configuration at startup
func validateConfig(cfg circuitbreaker.Config) error {
    if cfg.MaxRequests == 0 {
        return errors.New("MaxRequests must be > 0")
    }
    if cfg.Timeout <= 0 {
        return errors.New("Timeout must be > 0")
    }
    return nil
}
```

### Error Information Disclosure

```go
// Don't expose internal errors to clients
cb := circuitbreaker.New[T](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        // Log detailed info internally
        logger.Warn("circuit open", slog.Any("counts", counts))

        // Return generic error to client
        // error will be fortifyerrors.ErrCircuitOpen
    },
})
```

### Resource Limits

```go
// Prevent resource exhaustion
bh := bulkhead.New[T](bulkhead.Config{
    MaxConcurrent: 100,
    MaxQueue: 500,
    QueueTimeout: 30 * time.Second,  // Prevent indefinite queueing
})

rl := ratelimit.New(ratelimit.Config{
    Rate: 1000,
    Burst: 1500,  // Limit burst size
    Interval: time.Second,
})
```

## Monitoring

### Key Metrics

**Circuit Breaker:**
- State (closed/open/half-open)
- Total requests
- Success/failure counts
- State transitions
- Time in each state

**Retry:**
- Total retries
- Retry success rate
- Retry latency
- Max attempts reached count

**Rate Limiter:**
- Requests allowed
- Requests rejected
- Active buckets
- Token refill rate

**Timeout:**
- Timeout occurrences
- Timeout duration distribution
- Operations exceeding timeout

**Bulkhead:**
- Concurrent executions
- Queue length
- Rejections (capacity/queue full)
- Queue wait time

### Alerting Thresholds

**Critical Alerts:**
```
- circuit_breaker_state{state="open"} == 1 for 5m
- rate_limit_rejections > 1000 in 1m
- bulkhead_rejections > 100 in 1m
- timeout_rate > 10% in 5m
```

**Warning Alerts:**
```
- circuit_breaker_transitions > 10 in 5m
- retry_attempts > 50% of requests
- bulkhead_queue_length > 80% capacity
```

## Troubleshooting

### Circuit Breaker Not Opening

**Symptoms:** Circuit stays closed despite failures

**Checklist:**
1. Verify `ReadyToTrip` logic is correct
2. Check if errors are being returned (not swallowed)
3. Ensure `IsSuccessful` is properly configured
4. Verify `Interval` isn't clearing counts too frequently
5. Check minimum request threshold in `ReadyToTrip`

### Excessive Retries

**Symptoms:** Too many retry attempts

**Checklist:**
1. Verify `IsRetryable` only retries appropriate errors
2. Check `MaxAttempts` isn't too high
3. Ensure errors are properly classified
4. Verify context cancellation is working
5. Check for retry loops (nested retries)

### Rate Limiting False Positives

**Symptoms:** Legitimate requests being rate limited

**Checklist:**
1. Verify `KeyFunc` generates correct keys
2. Check `Rate` and `Burst` configuration
3. Ensure `Interval` is appropriate
4. Verify token bucket refill logic
5. Check for clock skew issues

### Timeout Issues

**Symptoms:** Operations timing out unexpectedly

**Checklist:**
1. Verify timeout duration is appropriate
2. Check if operation respects context cancellation
3. Ensure no goroutine leaks
4. Verify parent context deadline
5. Check for blocking operations

### Bulkhead Rejections

**Symptoms:** Requests rejected at capacity

**Checklist:**
1. Verify `MaxConcurrent` is appropriate
2. Check `MaxQueue` configuration
3. Ensure `QueueTimeout` isn't too short
4. Verify operations are completing
5. Check for goroutine leaks in workers
6. Ensure `Close()` is called properly

### Memory Leaks

**Symptoms:** Growing memory usage

**Checklist:**
1. Verify bulkhead `Close()` is called
2. Check rate limiter with many unique keys
3. Ensure contexts are properly cancelled
4. Verify no circular references in callbacks
5. Profile with `pprof` heap analysis

## Production Checklist

Before deploying to production:

- [ ] Configuration validated and tested
- [ ] Observability integrated (logging, metrics, tracing)
- [ ] Alerts configured with appropriate thresholds
- [ ] Error handling tested with chaos engineering
- [ ] Performance benchmarked under load
- [ ] Resource limits configured
- [ ] Security review completed
- [ ] Documentation updated
- [ ] Runbook created for on-call engineers
- [ ] Rollback plan documented
- [ ] Load testing completed
- [ ] Failure injection testing performed

## Best Practices Summary

1. **Start Conservative:** Begin with lenient settings and tighten based on observed behavior
2. **Monitor Everything:** Comprehensive observability is essential for tuning
3. **Test Failure Scenarios:** Use chaos engineering to validate resilience
4. **Respect Context:** Always honor context cancellation
5. **Log Appropriately:** Log state changes and errors, not every request
6. **Emit Metrics:** Track all pattern operations for visibility
7. **Version Configuration:** Use infrastructure-as-code for pattern config
8. **Document Decisions:** Record why specific thresholds were chosen
9. **Review Regularly:** Revisit configuration as traffic patterns change
10. **Plan for Growth:** Consider scalability in initial design

## Additional Resources

- [Technical Specification](./technical-spec.md)
- [Performance Testing Guide](./PERFORMANCE_TESTING.md)
- [API Documentation](https://pkg.go.dev/github.com/felixgeelhaar/fortify)
- [Examples](../examples/)
- [GitHub Issues](https://github.com/felixgeelhaar/fortify/issues)
