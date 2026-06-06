# Redis-backed rate limiter Store

Reference implementation of `ratelimit.Store` using Redis for cross-process token-bucket coordination. This example is a **separate Go module** so the core fortify module stays free of `go-redis` as a dependency.

## Why a separate module?

Fortify's promise is "zero core dependencies". Adding `github.com/redis/go-redis/v9` to fortify's `go.mod` would force every consumer of any fortify package to pull go-redis transitively. Splitting Redis support into its own module preserves the promise; users opt in by importing this module separately.

## Running the demo

Start a Redis (any version ≥ 6 works):

```bash
docker run --rm -p 6379:6379 redis:7-alpine
```

Then from this directory:

```bash
go run ./cmd/demo
```

Expected output (with default 10/s rate, 20 burst, 50 requests):

```
allowed=20 denied=30 (rate=10/s burst=20)
```

## Adopting in your project

```bash
go get go.klarlabs.de/fortify/examples/ratelimit-redis
```

```go
import (
    redisstore "go.klarlabs.de/fortify/examples/ratelimit-redis"
    "go.klarlabs.de/fortify/ratelimit"
    "github.com/redis/go-redis/v9"
)

client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
store := redisstore.New(client,
    redisstore.WithPrefix("rl:"),
    redisstore.WithTTL(1*time.Hour),
)

rl := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    Store:    store,
    FailOpen: true, // allow on Redis errors; capped at Burst tokens
})

// IMPORTANT: attach bucket params to ctx for the Lua script.
ctx = redisstore.WithBucketParams(ctx, 100, time.Second)
rl.Allow(ctx, "user-id")
```

## How the atomic update works

The Lua script (`atomicScriptSrc` in `store.go`):

1. Calls `TIME` for the Redis server timestamp (avoids client clock skew).
2. Reads existing `tokens` and `last` from a Hash at the bucket key.
3. Refills tokens based on elapsed time, capped at `Burst`.
4. Checks if requested tokens are available; if so, decrements.
5. Writes back tokens + timestamp; sets `PEXPIRE` for the TTL.
6. Returns `(tokens_remaining, last_refill_ns, allowed)`.

Tokens are stored as integers scaled by 10⁶ to avoid Lua's float quirks while preserving sub-token precision.

## Caveats

- **Custom update semantics**: `AtomicUpdate` does not invoke the user-supplied closure. Redis Lua cannot execute arbitrary Go. The Lua script encodes fortify's standard token-bucket protocol directly. If you need custom replenishment logic (sliding window, leaky bucket, etc.), implement your own Store.
- **Time source**: server-side TIME means one extra round-trip per call vs. local time. This avoids skew across application instances.
- **TTL**: dormant keys expire after the configured TTL (default 1h). For very long-lived users, increase TTL.
- **Failures**: on Redis errors, returns `ratelimit.ErrStorageUnavailable` wrapped. The rate limiter's `FailOpen` flag controls whether requests are admitted during outages.

## Production hardening

This example is a starting point. For production, consider:

- Connection pooling and circuit-breaker around the Redis client itself
- Retry on `READONLY` errors during failover (Redis Sentinel / Cluster)
- Sharding the Lua script across cluster slots if using Redis Cluster (use `{tag}` in keys)
- Metrics on Lua script latency (the `Metrics.OnStoreLatency` hook in fortify already covers this)
- Periodic test of the script's idempotency under load
