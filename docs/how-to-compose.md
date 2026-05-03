# How to compose patterns

Patterns are most useful in combination. Fortify's `middleware.Chain` is a fluent builder for layering them in the right order.

## Recommended chain

```go
import "github.com/felixgeelhaar/fortify/middleware"

chain := middleware.New[Response]().
    WithBulkhead(bh).                      // outermost
    WithRateLimit(rl, "user-key").
    WithTimeout(tm, 5*time.Second).
    WithCircuitBreaker(cb).
    WithRetry(r)                            // innermost (closest to operation)

result, err := chain.Execute(ctx, func(ctx context.Context) (Response, error) {
    return makeRequest(ctx)
})
```

## Why this order

Reading outer → inner:

1. **Bulkhead first** — shed load before doing any other work. Cheapest rejection.
2. **Rate limit next** — enforce quota before consuming downstream budget.
3. **Timeout outside circuit breaker** — guarantees a deadline even if the breaker is in Half-Open and stalls on a trial request.
4. **Circuit breaker before retry** — retry on transient failure, but if the breaker has tripped Open, fail fast instead of retrying through a known-bad dependency.
5. **Retry innermost** — only retries the operation itself, not the surrounding patterns. A retry won't trigger a second bulkhead or rate-limit charge.

## Order pitfalls

### Retry outside CB (wrong)

```go
chain := middleware.New[Response]().
    WithRetry(r).            // BAD: retries count against the CB
    WithCircuitBreaker(cb).
```

Each retry is a separate `cb.Execute` call. A flaky downstream causes the breaker to trip on the *retries*, not the originating failure rate.

### Bulkhead inside retry (wrong)

```go
chain := middleware.New[Response]().
    WithRetry(r).
    WithBulkhead(bh).        // BAD: each retry queues a new slot
```

Retries enqueue separately, defeating the bulkhead's "shed load on saturation" intent.

### Two timeouts (redundant)

If the surrounding context already has a deadline, the inner `Timeout` only adds value when its duration is shorter than the parent context's remaining time. Otherwise the parent cancels first and timeout's structured error is bypassed. (Fortify's `timeout.Execute` detects this and propagates the parent error verbatim.)

## Direct composition (without `middleware.Chain`)

If you don't need the fluent API:

```go
result, err := bh.Execute(ctx, func(ctx context.Context) (Response, error) {
    if !rl.Allow(ctx, key) {
        var zero Response
        return zero, ratelimit.ErrLimitExceeded
    }
    return cb.Execute(ctx, func(ctx context.Context) (Response, error) {
        return r.Execute(ctx, func(ctx context.Context) (Response, error) {
            return tm.Execute(ctx, 5*time.Second, makeRequest)
        })
    })
})
```

`middleware.Chain` is sugar over this. Use whichever is clearer at the call site.

## Pre-built bundles

Currently there are no pre-built bundles. If you find yourself recreating the same chain across services, consider extracting a helper in your codebase that returns a configured `middleware.Chain`.
