# Fortify Redis Backend

Redis-backed distributed rate limiting for Fortify. This package provides the same `RateLimiter` interface as the core `fortify/ratelimit` package, but stores state in Redis for distributed systems.

## Features

- **Distributed Rate Limiting**: Share rate limits across multiple application instances
- **Atomic Operations**: Lua scripts ensure race-condition-free token bucket implementation
- **Production-Grade**: Supports Redis Cluster, Redis Sentinel, and connection pooling
- **Drop-In Replacement**: Same interface as in-memory rate limiter
- **Full Observability**: Integrates with slog logging and OpenTelemetry
- **Flexible Error Handling**: Configurable fail-open or fail-closed behavior
- **Auto-Cleanup**: Automatic bucket expiration with configurable TTL

## Installation

```bash
go get github.com/felixgeelhaar/fortify/backends/redis
```

**Requirements:**
- Go 1.23 or higher
- Redis 6.0 or higher (for Lua script support)

## Quick Start

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/redis/go-redis/v9"
    redisrl "github.com/felixgeelhaar/fortify/backends/redis"
)

func main() {
    // Create Redis client
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    defer client.Close()

    // Create distributed rate limiter
    limiter, err := redisrl.New(redisrl.Config{
        Client:   client,
        Rate:     100,              // 100 requests
        Burst:    200,              // burst of 200
        Interval: time.Second,      // per second
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Non-blocking check
    if limiter.Allow(ctx, "user-123") {
        // Process request
    } else {
        // Return 429 Too Many Requests
    }
}
```

## Configuration

### Basic Configuration

```go
config := redisrl.Config{
    Client:   redisClient,      // Required: Redis client
    Rate:     100,               // Tokens per interval (default: 100)
    Burst:    200,               // Max burst size (default: same as Rate)
    Interval: time.Second,       // Refill interval (default: 1 second)
}
```

### Advanced Configuration

```go
config := redisrl.Config{
    Client:   redisClient,
    Rate:     100,
    Burst:    200,
    Interval: time.Second,

    // Custom key prefix for namespacing
    KeyPrefix: "myapp:ratelimit:",  // Default: "fortify:ratelimit:"

    // Extract key from context
    KeyFunc: func(ctx context.Context) string {
        if userID, ok := ctx.Value("user_id").(string); ok {
            return "user:" + userID
        }
        return "anonymous"
    },

    // Called when rate limit is exceeded
    OnLimit: func(key string) {
        metrics.RecordRateLimit(key)
    },

    // Structured logging
    Logger: slog.Default(),

    // Auto-cleanup idle buckets
    BucketTTL: time.Hour,  // Default: 1 hour

    // Error handling strategy
    FallbackOnError: false,  // Default: false (fail-closed)
}
```

## Usage Examples

### HTTP API Rate Limiting

```go
import (
    "net/http"
    redisrl "github.com/felixgeelhaar/fortify/backends/redis"
)

func rateLimitMiddleware(limiter ratelimit.RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Extract user ID from request
            userID := r.Header.Get("X-User-ID")

            if !limiter.Allow(r.Context(), userID) {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
```

### Per-User Rate Limiting

```go
// Configure rate limiter to extract user from context
limiter, _ := redisrl.New(redisrl.Config{
    Client:   redisClient,
    Rate:     1000,
    Burst:    2000,
    Interval: time.Minute,
    KeyFunc: func(ctx context.Context) string {
        // Extract user ID from context
        if userID, ok := ctx.Value("user_id").(string); ok {
            return "user:" + userID
        }
        return "anonymous"
    },
})

// In your handler
ctx := context.WithValue(r.Context(), "user_id", "user-123")
if limiter.Allow(ctx, "") {  // Key extracted from context
    // Process request
}
```

### Multiple Token Consumption

```go
// API endpoint that consumes tokens based on operation cost
limiter, _ := redisrl.New(redisrl.Config{
    Client:   redisClient,
    Rate:     1000,
    Burst:    5000,
    Interval: time.Minute,
})

// Expensive operation costs 10 tokens
if limiter.Take(ctx, "user-123", 10) {
    // Process expensive operation
} else {
    // Quota exceeded
}
```

### Blocking Wait

```go
// Wait for token availability (blocks until token is available)
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := limiter.Wait(ctx, "user-123"); err != nil {
    if err == context.DeadlineExceeded {
        // Timeout waiting for token
    }
    return
}

// Token acquired, process request
```

## Redis Deployment Options

### Standalone Redis

```go
client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

limiter, _ := redisrl.New(redisrl.Config{
    Client: client,
    // ... other config
})
```

### Redis Cluster

```go
client := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{
        "redis-node-1:6379",
        "redis-node-2:6379",
        "redis-node-3:6379",
    },
})

limiter, _ := redisrl.New(redisrl.Config{
    Client: client,
    // ... other config
})
```

### Redis Sentinel (High Availability)

```go
client := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "mymaster",
    SentinelAddrs: []string{
        "sentinel-1:26379",
        "sentinel-2:26379",
        "sentinel-3:26379",
    },
})

limiter, _ := redisrl.New(redisrl.Config{
    Client: client,
    // ... other config
})
```

## Error Handling

### Fail-Closed (Default)

Deny requests when Redis is unavailable (consistency over availability):

```go
limiter, _ := redisrl.New(redisrl.Config{
    Client:          redisClient,
    FallbackOnError: false,  // Default
    // ... other config
})

// If Redis is down, Allow() returns false
```

### Fail-Open

Allow requests when Redis is unavailable (availability over consistency):

```go
limiter, _ := redisrl.New(redisrl.Config{
    Client:          redisClient,
    FallbackOnError: true,  // Allow on Redis failure
    // ... other config
})

// If Redis is down, Allow() returns true
```

### Error Logging

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

limiter, _ := redisrl.New(redisrl.Config{
    Client: redisClient,
    Logger: logger,  // Logs rate limit events and errors
    // ... other config
})
```

## Performance

### Benchmarks

Typical performance characteristics (local Redis):

| Operation | Latency (p50) | Latency (p99) | Throughput |
|-----------|---------------|---------------|------------|
| Allow()   | ~1ms          | ~3ms          | ~50k ops/s |
| Take()    | ~1ms          | ~3ms          | ~50k ops/s |
| Wait()    | varies        | varies        | N/A        |

*Benchmarks on Apple M1, Go 1.23, Redis 7.0*

### Optimization Tips

1. **Connection Pooling**: Use appropriate pool size for your workload
```go
client := redis.NewClient(&redis.Options{
    Addr:         "localhost:6379",
    PoolSize:     10,  // Default is 10 * runtime.NumCPU()
    MinIdleConns: 5,
})
```

2. **Batch Operations**: Use pipelines for multiple rate limit checks
```go
pipe := client.Pipeline()
// Add multiple operations to pipeline
_, err := pipe.Exec(ctx)
```

3. **Key Prefix**: Use short, efficient key prefixes
```go
config := redisrl.Config{
    KeyPrefix: "rl:",  // Short prefix saves memory
    // ... other config
}
```

## Production Deployment

### Redis Configuration

Recommended Redis configuration for rate limiting:

```redis
# Persistence - Optional for rate limiting
save ""
appendonly no

# Memory management
maxmemory 1gb
maxmemory-policy allkeys-lru

# Performance
tcp-keepalive 300
timeout 0

# Lua scripting
lua-time-limit 5000
```

### Monitoring

Monitor these Redis metrics:

- `connected_clients`: Connection pool health
- `evicted_keys`: Memory pressure indicator
- `used_memory`: Memory usage
- `ops_per_sec`: Operation throughput
- `latency`: Response time

### Scaling

For high-throughput scenarios:

1. **Redis Cluster**: Distribute load across multiple nodes
2. **Read Replicas**: For read-heavy workloads (not typical for rate limiting)
3. **Multiple Redis Instances**: Partition by key range

## Troubleshooting

### High Latency

**Symptom**: Slow `Allow()` operations

**Solutions**:
- Check Redis latency: `redis-cli --latency`
- Verify network connectivity
- Increase connection pool size
- Consider Redis Cluster for distribution

### Memory Issues

**Symptom**: Redis OOM errors

**Solutions**:
- Decrease `BucketTTL` to expire buckets faster
- Set `maxmemory-policy allkeys-lru`
- Monitor `evicted_keys` metric
- Increase Redis memory or add nodes

### Rate Limits Not Working

**Symptom**: More requests allowed than expected

**Solutions**:
- Verify Lua script execution: `redis-cli SCRIPT EXISTS <sha>`
- Check for clock skew across servers
- Ensure consistent configuration across instances
- Verify key resolution (KeyFunc)

### Connection Errors

**Symptom**: Redis connection failures

**Solutions**:
- Check Redis availability
- Verify network connectivity
- Configure appropriate timeouts
- Use `FallbackOnError` for graceful degradation

## Migration from In-Memory

See [Migration Guide](../../docs/MIGRATION_REDIS.md) for detailed instructions.

**Quick migration:**

```go
// Before (in-memory)
import "github.com/felixgeelhaar/fortify/ratelimit"

limiter := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
})

// After (Redis)
import redisrl "github.com/felixgeelhaar/fortify/backends/redis"

limiter, _ := redisrl.New(redisrl.Config{
    Client:   redisClient,  // Add Redis client
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
})

// Same interface - no other changes needed!
```

## Best Practices

1. **Use KeyFunc for Context Extraction**
   - Keeps key resolution logic centralized
   - Avoids manual key construction

2. **Set Appropriate BucketTTL**
   - Balance between memory usage and bucket reuse
   - Typical values: 1 hour to 24 hours

3. **Monitor Redis Health**
   - Set up alerts for connection failures
   - Track memory usage and eviction rates

4. **Choose Error Handling Strategy**
   - Fail-closed for critical rate limits (security)
   - Fail-open for user-facing features (UX)

5. **Use Connection Pooling**
   - Size pool based on concurrency needs
   - Monitor connection saturation

## Examples

See complete examples in [`examples/backends/redis/`](../../examples/backends/redis/).

## License

MIT License - see [LICENSE](../../LICENSE) file for details.

## Support

- 📖 [Main Documentation](../../README.md)
- 🐛 [Issue Tracker](https://github.com/felixgeelhaar/fortify/issues)
- 💬 [Discussions](https://github.com/felixgeelhaar/fortify/discussions)
