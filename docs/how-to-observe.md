# How to wire observability

Fortify ships three integrations: structured logging via `slog`, distributed tracing via OpenTelemetry, and Prometheus metrics. All three are opt-in.

## Structured logging (`slog`)

Every pattern accepts a `*slog.Logger` in its `Config`. Log lines include pattern name, key fields, and `slog.GroupValue` for state transitions.

```go
import (
    "log/slog"
    "os"

    "github.com/felixgeelhaar/fortify/circuitbreaker"
)

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

cb := circuitbreaker.New[Response](circuitbreaker.Config{
    MaxRequests: 5,
    Logger:      logger,
})
```

Add correlation IDs to context; Fortify propagates them via `LogContext` calls (`ErrorContext`, `WarnContext`, etc.).

## OpenTelemetry tracing (`otel`)

```go
import (
    fortifyotel "github.com/felixgeelhaar/fortify/otel"
    "go.opentelemetry.io/otel"
)

tracer := fortifyotel.NewTracer(otel.GetTracerProvider(), "my-service")

ctx, span := tracer.StartSpan(ctx, fortifyotel.PatternCircuitBreaker, "execute")
defer span.End()

tracer.SetAttributes(span,
    attribute.String("fortify.cb.state", cb.State().String()),
)
```

Span attribute conventions:

| Attribute              | Set by                  |
| ---------------------- | ----------------------- |
| `fortify.cb.state`     | Circuit breaker         |
| `fortify.retry.attempt`| Retry                   |
| `fortify.rl.key`       | Rate limit (sanitized)  |
| `fortify.timeout.ms`   | Timeout                 |
| `fortify.bulkhead.queued` | Bulkhead             |

## Prometheus metrics (`metrics`)

```go
import (
    "github.com/felixgeelhaar/fortify/metrics"
    "github.com/prometheus/client_golang/prometheus"
)

metrics.MustRegister(prometheus.DefaultRegisterer)
collector := metrics.DefaultCollector()
```

Bind to your patterns:

```go
cb := circuitbreaker.New[Response](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        collector.RecordCircuitBreakerStateChange("api-client", from.String(), to.String())
    },
})
```

Available metrics (each labelled with the pattern instance name):

- **Circuit breaker:** state gauge, request counter (by outcome), state-change counter
- **Retry:** attempts histogram, success/failure counter, total duration histogram
- **Rate limit:** allow/deny counter, store-latency histogram
- **Timeout:** execution histogram, exceeded counter
- **Bulkhead:** active gauge, queued gauge, rejected counter

### Cardinality discipline

Prometheus labels are unbounded by default — be careful with key-derived labels. Fortify's metrics use `instance name` (your choice) as the only configurable label. The rate limiter does **not** plumb client-supplied keys to Prometheus labels by default; if you wire `OnAllow`/`OnDeny` callbacks to do so, sanitize and bound them yourself.

## Rate limiter callbacks

For finer-grained observability, set `OnAllow` and `OnLimit` callbacks on `ratelimit.Config`:

```go
rl := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    OnAllow: func(ctx context.Context, key string) {
        // Called when a request is allowed.
    },
    OnLimit: func(ctx context.Context, key string) {
        // Called when a request is rate limited.
    },
})
```

Or implement the `Metrics` interface to receive all four signals (`OnAllow`, `OnDeny`, `OnError`, `OnStoreLatency`):

```go
type myMetrics struct{}

func (m *myMetrics) OnAllow(ctx context.Context, key string)        { /* ... */ }
func (m *myMetrics) OnDeny(ctx context.Context, key string)         { /* ... */ }
func (m *myMetrics) OnError(ctx context.Context, key string, err error) { /* ... */ }
func (m *myMetrics) OnStoreLatency(ctx context.Context, op string, d time.Duration) { /* ... */ }

rl := ratelimit.New(ratelimit.Config{Metrics: &myMetrics{}})
```

Callbacks run synchronously on the request path. Keep them cheap (counter increments, channel sends); offload heavy work to a goroutine inside the callback.
