# Distributed circuit breaker (Redis-backed)

Reference implementation of `circuitbreaker.CircuitBreaker[T]` whose state is shared across multiple service replicas via Redis.

This example is a **separate Go module** (own `go.mod`) so the core fortify module stays free of `go-redis`.

## When to use

Use this when **multiple replicas** of your service must agree on whether a shared downstream is healthy. Examples:

- A fleet of API workers all calling the same payment provider
- A horizontally-scaled batch processor calling a single SaaS endpoint
- Any setup where each replica's per-process breaker would not see enough failure signal alone

For per-replica protection (no cross-process coordination needed), use the in-process `fortify/circuitbreaker` — it is sub-microsecond and zero-allocation. This adapter trades that performance for cross-instance agreement.

## How it works

State (Closed/Open/HalfOpen, generation, expiry, counters) lives in a single Redis Hash. Two Lua scripts run atomically:

- **`admit`** — read state, auto-transition Open→HalfOpen on timeout, gate the request, return state + generation + allow/deny + retry-after
- **`record`** — given the generation observed at admission, apply the call's outcome (advance counters, possibly transition state)

Generations protect against stale records: if state has rotated under a slow caller, the record is discarded.

## Running the demo

```bash
docker run --rm -p 6379:6379 redis:7-alpine
go run ./cmd/demo
```

Expected output:

```
[transition] closed -> open
state after 3 failures: open
call while open: circuit breaker is open (state=open, retry_after=2s)
[transition] open -> half-open
[transition] half-open -> closed
call after timeout: 42 <nil>
final state: closed
```

## Trade-offs vs. in-process CB

| Concern              | In-process            | Redis-backed                       |
| -------------------- | --------------------- | ---------------------------------- |
| Latency per Execute  | ~70 ns (Closed)       | bound by Redis RTT (~1ms typical)  |
| Cross-replica agreement | No                 | Yes                                |
| Custom predicate     | Arbitrary Go closure  | Threshold only (Lua-bound)         |
| State change callback| Cross-instance via dispatcher goroutine | Local-instance only |
| Failure mode         | n/a                   | Configurable FailOpen / FailClosed |

## Limitations

- **Predicate restriction**: only `consecutive_failures >= threshold` is supported. Adding sliding-window or failure-ratio predicates requires extending the Lua script.
- **State change callbacks fire only locally**: the instance that observes the transition fires `OnStateChange`. Sibling instances learn about it on their next admission check. For cross-instance notification, subscribe to a Redis pubsub channel and broadcast.
- **`Reset` is racy**: deletes the Hash key. In-flight calls may observe the rebooted state mid-flight.
- **Half-open trials counter** uses `half_open_admitted` decrement on success. If a trial caller crashes between admission and recording, the counter never decrements — a pathological corner case. A production hardening would use a per-trial token list with TTL.

## Adopting in your project

```bash
go get go.klarlabs.de/fortify/examples/circuitbreaker-redis
```

```go
import (
    redisbreaker "go.klarlabs.de/fortify/examples/circuitbreaker-redis"
    "github.com/redis/go-redis/v9"
)

client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
cb, err := redisbreaker.New[Response](redisbreaker.Config{
    Client:           client,
    Key:              "fortify:cb:payment-api",
    FailureThreshold: 5,
    OpenTimeout:      30 * time.Second,
    MaxRequests:      3,
    FailMode:         redisbreaker.FailOpen,
})

result, err := cb.Execute(ctx, func(ctx context.Context) (Response, error) {
    return paymentClient.Charge(ctx, payload)
})
```

The returned `*Breaker[T]` satisfies `circuitbreaker.CircuitBreaker[T]`, so it composes with `middleware.Chain[T].WithCircuitBreaker(...)` exactly like the in-process implementation.
