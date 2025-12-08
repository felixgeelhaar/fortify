# Migration Guide

This guide helps you migrate from other Go resilience libraries to Fortify.

## Table of Contents

- [From sony/gobreaker](#from-sonygobreaker)
- [From avast/retry-go](#from-avastretry-go)
- [From golang.org/x/time/rate](#from-golangorgxtimerate)
- [From eapache/go-resiliency](#from-eapachego-resiliency)
- [Feature Comparison](#feature-comparison)
- [Migration Checklist](#migration-checklist)

## From sony/gobreaker

### Overview

sony/gobreaker is a popular circuit breaker implementation. Fortify provides similar functionality with additional features like generics, observability, and middleware composition.

### Basic Migration

**Before (sony/gobreaker):**
```go
import "github.com/sony/gobreaker"

cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
    Name:        "my-service",
    MaxRequests: 3,
    Interval:    time.Second * 60,
    Timeout:     time.Second * 30,
    ReadyToTrip: func(counts gobreaker.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 3 && failureRatio >= 0.6
    },
    OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
        log.Printf("CB %s: %s -> %s", name, from, to)
    },
})

result, err := cb.Execute(func() (interface{}, error) {
    return apiCall()
})
```

**After (Fortify):**
```go
import "github.com/felixgeelhaar/fortify/circuitbreaker"

cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    MaxRequests: 3,
    Interval:    time.Minute,
    Timeout:     30 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 3 && failureRatio >= 0.6
    },
    OnStateChange: func(from, to circuitbreaker.State) {
        slog.Info("circuit breaker state change",
            slog.String("from", from.String()),
            slog.String("to", to.String()),
        )
    },
})

result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return apiCall(ctx)
})
```

### Key Differences

1. **Type Safety**: Fortify uses generics - specify return type `[*Response]`
2. **Context Support**: All operations accept `context.Context`
3. **Name Field**: Fortify doesn't require a name field (use tags in observability)
4. **State Constants**: Use `circuitbreaker.State` instead of `gobreaker.State`

### State Mapping

| sony/gobreaker | Fortify |
|---------------|---------|
| `gobreaker.StateClosed` | `circuitbreaker.StateClosed` |
| `gobreaker.StateOpen` | `circuitbreaker.StateOpen` |
| `gobreaker.StateHalfOpen` | `circuitbreaker.StateHalfOpen` |

### Counts Struct

Both libraries use similar `Counts` structs:

```go
// Both have:
type Counts struct {
    Requests             uint32
    TotalSuccesses       uint32
    TotalFailures        uint32
    ConsecutiveSuccesses uint32
    ConsecutiveFailures  uint32
}
```

## From avast/retry-go

### Overview

avast/retry-go provides functional retry logic. Fortify offers similar patterns with better type safety and integrated observability.

### Basic Migration

**Before (avast/retry-go):**
```go
import "github.com/avast/retry-go"

err := retry.Do(
    func() error {
        return apiCall()
    },
    retry.Attempts(3),
    retry.Delay(100*time.Millisecond),
    retry.DelayType(retry.BackOffDelay),
    retry.OnRetry(func(n uint, err error) {
        log.Printf("Retry %d: %v", n, err)
    }),
)
```

**After (Fortify):**
```go
import "github.com/felixgeelhaar/fortify/retry"

r := retry.New[any](&retry.Config{
    MaxAttempts:   3,
    InitialDelay:  100 * time.Millisecond,
    BackoffPolicy: retry.BackoffExponential,
    Multiplier:    2.0,
    OnRetry: func(attempt int, err error) {
        slog.Info("retrying operation",
            slog.Int("attempt", attempt),
            slog.String("error", err.Error()),
        )
    },
})

_, err := r.Do(ctx, func(ctx context.Context) (any, error) {
    return nil, apiCall(ctx)
})
```

### Key Differences

1. **Reusable Instance**: Fortify uses reusable retry instances
2. **Generic Return Type**: Specify return type for type safety
3. **Context Support**: All operations require context
4. **Backoff Policies**: Explicitly configured backoff strategies

### Backoff Policy Mapping

| avast/retry-go | Fortify |
|---------------|---------|
| `retry.BackOffDelay` | `retry.BackoffExponential` |
| `retry.FixedDelay` | `retry.BackoffConstant` |
| No direct equivalent | `retry.BackoffLinear` |

### Error Classification

**Before (avast/retry-go):**
```go
err := retry.Do(
    func() error {
        return apiCall()
    },
    retry.RetryIf(func(err error) bool {
        return errors.Is(err, ErrTemporary)
    }),
)
```

**After (Fortify):**
```go
r := retry.New[any](&retry.Config{
    IsRetryable: func(err error) bool {
        return errors.Is(err, ErrTemporary)
    },
})

_, err := r.Do(ctx, func(ctx context.Context) (any, error) {
    return nil, apiCall(ctx)
})
```

## From golang.org/x/time/rate

### Overview

golang.org/x/time/rate provides token bucket rate limiting. Fortify offers similar functionality with additional features.

### Basic Migration

**Before (golang.org/x/time/rate):**
```go
import "golang.org/x/time/rate"

// 100 requests per second with burst of 150
limiter := rate.NewLimiter(100, 150)

if !limiter.Allow() {
    http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
    return
}

// Or with wait
if err := limiter.Wait(ctx); err != nil {
    http.Error(w, "Rate limit wait failed", http.StatusTooManyRequests)
    return
}
```

**After (Fortify):**
```go
import "github.com/felixgeelhaar/fortify/ratelimit"

limiter := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    150,
    Interval: time.Second,
})

if !limiter.Allow(ctx, key) {
    http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
    return
}

// Or with wait
if err := limiter.Wait(ctx, key); err != nil {
    http.Error(w, "Rate limit wait failed", http.StatusTooManyRequests)
    return
}
```

### Key Differences

1. **Per-Key Limiting**: Fortify supports per-key rate limiting out of the box
2. **Configurable Interval**: Explicitly set the rate interval
3. **Callbacks**: Add `OnLimit` callback for monitoring
4. **Context Support**: All methods accept context

### Per-Key Rate Limiting

**golang.org/x/time/rate** (requires manual map):
```go
limiters := make(map[string]*rate.Limiter)
mu := sync.RWMutex{}

func getLimiter(key string) *rate.Limiter {
    mu.RLock()
    limiter, exists := limiters[key]
    mu.RUnlock()

    if !exists {
        mu.Lock()
        limiter = rate.NewLimiter(100, 150)
        limiters[key] = limiter
        mu.Unlock()
    }

    return limiter
}
```

**Fortify** (built-in):
```go
limiter := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    150,
    Interval: time.Second,
})

// Automatically handles per-key buckets
limiter.Allow(ctx, userID)
```

### Reserve Method

golang.org/x/time/rate has a `Reserve()` method. In Fortify, use `Take()`:

**Before:**
```go
r := limiter.Reserve()
if !r.OK() {
    // Rate limit exceeded
}
time.Sleep(r.Delay())
```

**After:**
```go
// Take 1 token (same as Allow)
if !limiter.Take(ctx, key, 1) {
    // Rate limit exceeded
}

// For multi-token operations
if !limiter.Take(ctx, key, 5) {
    // Not enough tokens
}
```

## From eapache/go-resiliency

### Overview

eapache/go-resiliency provides multiple resilience patterns. Fortify offers similar patterns with modern Go features.

### Circuit Breaker Migration

**Before (eapache/go-resiliency):**
```go
import "github.com/eapache/go-resiliency/breaker"

b := breaker.New(3, 1, 5*time.Second)

err := b.Run(func() error {
    return apiCall()
})

if err == breaker.ErrBreakerOpen {
    // Handle open circuit
}
```

**After (Fortify):**
```go
import (
    "github.com/felixgeelhaar/fortify/circuitbreaker"
    fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

cb := circuitbreaker.New[any](circuitbreaker.Config{
    MaxRequests: 1,
    Interval:    5 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures >= 3
    },
})

_, err := cb.Execute(ctx, func(ctx context.Context) (any, error) {
    return nil, apiCall(ctx)
})

if errors.Is(err, fortifyerrors.ErrCircuitOpen) {
    // Handle open circuit
}
```

### Retrier Migration

**Before (eapache/go-resiliency):**
```go
import "github.com/eapache/go-resiliency/retrier"

r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)

err := r.Run(func() error {
    return apiCall()
})
```

**After (Fortify):**
```go
import "github.com/felixgeelhaar/fortify/retry"

r := retry.New[any](&retry.Config{
    MaxAttempts:   3,
    InitialDelay:  100 * time.Millisecond,
    BackoffPolicy: retry.BackoffConstant,
})

_, err := r.Do(ctx, func(ctx context.Context) (any, error) {
    return nil, apiCall(ctx)
})
```

### Deadline/Timeout Migration

**Before (eapache/go-resiliency):**
```go
import "github.com/eapache/go-resiliency/deadline"

d := deadline.New(5 * time.Second)

err := d.Run(func(stopper <-chan struct{}) error {
    return apiCall()
})

if err == deadline.ErrTimedOut {
    // Handle timeout
}
```

**After (Fortify):**
```go
import "github.com/felixgeelhaar/fortify/timeout"

tm := timeout.New[any](timeout.Config{
    DefaultTimeout: 5 * time.Second,
})

_, err := tm.Execute(ctx, 5*time.Second, func(ctx context.Context) (any, error) {
    return nil, apiCall(ctx)
})

if errors.Is(err, context.DeadlineExceeded) {
    // Handle timeout
}
```

## Feature Comparison

### Circuit Breaker

| Feature | sony/gobreaker | eapache/breaker | Fortify |
|---------|---------------|-----------------|---------|
| Type Safety | ❌ | ❌ | ✅ (Generics) |
| Context Support | ❌ | ❌ | ✅ |
| State Callbacks | ✅ | ❌ | ✅ |
| Custom Success Logic | ❌ | ❌ | ✅ |
| Observability | Limited | ❌ | ✅ (slog, metrics) |
| Middleware | ❌ | ❌ | ✅ |

### Retry

| Feature | avast/retry-go | eapache/retrier | Fortify |
|---------|---------------|-----------------|---------|
| Type Safety | ❌ | ❌ | ✅ (Generics) |
| Context Support | ❌ | ❌ | ✅ |
| Backoff Policies | Limited | Limited | ✅ (3 types + jitter) |
| Error Classification | ✅ | ❌ | ✅ (Multiple options) |
| Retry Callbacks | ✅ | ❌ | ✅ |
| Observability | Limited | ❌ | ✅ |

### Rate Limiting

| Feature | golang.org/x/time/rate | Fortify |
|---------|------------------------|---------|
| Token Bucket | ✅ | ✅ |
| Per-Key Limiting | Manual | ✅ (Built-in) |
| Context Support | Limited | ✅ |
| Callbacks | ❌ | ✅ |
| Key Functions | ❌ | ✅ |
| Observability | ❌ | ✅ |

## Migration Checklist

### Pre-Migration

- [ ] Review current resilience patterns in use
- [ ] Identify all libraries to be replaced
- [ ] Map current configuration to Fortify equivalents
- [ ] Plan migration strategy (gradual vs big bang)
- [ ] Set up observability for both old and new implementations

### During Migration

- [ ] Add Fortify dependency: `go get github.com/felixgeelhaar/fortify`
- [ ] Implement Fortify patterns alongside existing ones
- [ ] Update function signatures to accept `context.Context`
- [ ] Add type parameters for generic patterns
- [ ] Configure callbacks for observability
- [ ] Update error handling to use Fortify errors
- [ ] Test thoroughly with existing test suite
- [ ] Compare behavior with old implementation

### Post-Migration

- [ ] Remove old library dependencies
- [ ] Update documentation
- [ ] Monitor metrics for both implementations during transition
- [ ] Validate performance characteristics
- [ ] Clean up unused code
- [ ] Update team documentation and runbooks

## Common Migration Patterns

### Pattern 1: Gradual Replacement

Replace one pattern at a time, starting with lowest risk:

1. **Start with Rate Limiting** (lowest impact)
2. **Add Timeout** (easy to test)
3. **Migrate Retry Logic** (moderate complexity)
4. **Replace Circuit Breakers** (highest impact)

### Pattern 2: Shadow Mode

Run old and new implementations in parallel:

```go
// Execute both implementations
oldResult, oldErr := oldCircuitBreaker.Execute(func() (interface{}, error) {
    return apiCall()
})

newResult, newErr := newCircuitBreaker.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return apiCall(ctx)
})

// Compare results and log differences
if !reflect.DeepEqual(oldErr, newErr) {
    logger.Warn("circuit breaker migration mismatch",
        slog.Any("old_error", oldErr),
        slog.Any("new_error", newErr),
    )
}

// Use old result initially, switch to new after validation
return oldResult, oldErr
```

### Pattern 3: Feature Flag

Use feature flags to control rollout:

```go
if featureFlags.IsEnabled("fortify_circuit_breaker") {
    return fortifyCircuitBreaker.Execute(ctx, operation)
} else {
    return legacyCircuitBreaker.Execute(operation)
}
```

## Troubleshooting Migration

### Type Errors

**Problem:** `cannot use operation (type func() error) as type func(context.Context) (T, error)`

**Solution:** Update function signatures to accept context and return typed values:

```go
// Before
func operation() error {
    return apiCall()
}

// After
func operation(ctx context.Context) (*Response, error) {
    return apiCall(ctx)
}
```

### Context Cancellation

**Problem:** Operations don't respect context cancellation

**Solution:** Ensure operations check context:

```go
func operation(ctx context.Context) (*Response, error) {
    // Check context before expensive operations
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }

    return apiCall(ctx)
}
```

### Error Handling

**Problem:** Error comparison doesn't work

**Solution:** Use `errors.Is()` instead of equality:

```go
// Before
if err == gobreaker.ErrOpenState {
    // Handle
}

// After
if errors.Is(err, fortifyerrors.ErrCircuitOpen) {
    // Handle
}
```

### Performance Differences

**Problem:** Different performance characteristics

**Solution:** Benchmark both implementations:

```go
func BenchmarkOldCircuitBreaker(b *testing.B) {
    cb := oldlib.NewCircuitBreaker(settings)
    for i := 0; i < b.N; i++ {
        cb.Execute(operation)
    }
}

func BenchmarkFortifyCircuitBreaker(b *testing.B) {
    cb := circuitbreaker.New[any](config)
    ctx := context.Background()
    for i := 0; i < b.N; i++ {
        cb.Execute(ctx, operation)
    }
}
```

## Getting Help

- [GitHub Issues](https://github.com/felixgeelhaar/fortify/issues)
- [Documentation](../README.md)
- [Examples](../examples/)
- [Production Guide](./PRODUCTION.md)
- [Error Handling Guide](./ERROR_HANDLING.md)

## Migration Support

If you need help migrating:

1. Review the [examples](../examples/) directory
2. Check [existing issues](https://github.com/felixgeelhaar/fortify/issues) for similar migrations
3. Open a [new issue](https://github.com/felixgeelhaar/fortify/issues/new) with:
   - Current library and version
   - Your use case
   - Specific migration challenges
   - Code examples (if possible)
