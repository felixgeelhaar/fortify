# Fortify

Composable resilience patterns for Go.

Fortify gives you circuit breaker, retry, rate limit, timeout, bulkhead, fallback, hedge, and adaptive concurrency under one roof, with first-class observability via OpenTelemetry, Prometheus, and `slog`. Zero core dependencies.

## Quick links

- **[Concepts](concepts.md)** — what each pattern does and when to reach for it
- **[How-to: compose](how-to-compose.md)** — chain ordering rationale + pitfalls
- **[How-to: observe](how-to-observe.md)** — `slog`, OpenTelemetry, Prometheus wiring
- **[How-to: rate limit](how-to-rate-limit.md)** — recipes for per-key, KeyFunc, custom Store
- **[How-to: test](how-to-test.md)** — chaos utilities and regression testing
- **[Integrations](integrations.md)** — HTTP middleware + RoundTripper, gRPC interceptors

## Install

```bash
go get github.com/felixgeelhaar/fortify
```

## 60-second example

```go
import (
    "context"
    "time"

    "github.com/felixgeelhaar/fortify/circuitbreaker"
    "github.com/felixgeelhaar/fortify/retry"
)

cb := circuitbreaker.New[string](circuitbreaker.Config{
    MaxRequests: 5,
    Interval:    10 * time.Second,
})
defer cb.Close()

r := retry.New[string](retry.Config{
    MaxAttempts:   3,
    InitialDelay:  100 * time.Millisecond,
    BackoffPolicy: retry.BackoffExponential,
    Jitter:        true,
})

result, err := cb.Execute(ctx, func(ctx context.Context) (string, error) {
    return r.Execute(ctx, callExternalService)
})
```

## Why Fortify

Most Go resilience libraries cover a single pattern. Stitching together a circuit breaker (`sony/gobreaker`), a retry policy (`hashicorp/go-retryablehttp`), and a rate limiter (`golang.org/x/time/rate`) gives you three different APIs, three different observability stories, and ad-hoc composition.

Fortify is the resilience library for teams that want all of it under one roof, with consistent ergonomics and observability built in. See [`COMPARISON.md`](COMPARISON.md) for the side-by-side.

## Project status

- See [GOVERNANCE.md](https://github.com/felixgeelhaar/fortify/blob/main/GOVERNANCE.md) for maintainership and semver policy
- See [SECURITY.md](https://github.com/felixgeelhaar/fortify/blob/main/SECURITY.md) for vulnerability disclosure
- See [ADOPTERS.md](https://github.com/felixgeelhaar/fortify/blob/main/ADOPTERS.md) — PRs welcome to add yourself

MIT-licensed. Currently solo-maintained.
