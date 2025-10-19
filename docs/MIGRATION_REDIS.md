# Redis Backend Migration Guide

This guide helps you migrate from Fortify's in-memory rate limiter to the Redis-backed distributed rate limiter.

## Table of Contents

- [When to Use Redis Backend](#when-to-use-redis-backend)
- [When to Use In-Memory](#when-to-use-in-memory)
- [Migration Steps](#migration-steps)
- [Configuration Comparison](#configuration-comparison)
- [Code Changes](#code-changes)
- [Testing Strategy](#testing-strategy)
- [Performance Considerations](#performance-considerations)
- [Production Deployment](#production-deployment)
- [Rollback Plan](#rollback-plan)
- [Troubleshooting](#troubleshooting)

## When to Use Redis Backend

Use the Redis backend when you have:

✅ **Multiple Application Instances**
- Horizontally scaled applications (Kubernetes, ECS, etc.)
- Need consistent rate limits across all instances

✅ **Distributed Systems**
- Microservices architecture
- API Gateway scenarios
- Multi-region deployments with shared limits

✅ **Long-Term State Persistence**
- Rate limits must survive application restarts
- Need audit trail of rate limit events

✅ **Dynamic Scaling**
- Auto-scaling based on load
- Need rate limits to remain consistent during scale events

## When to Use In-Memory

Keep the in-memory limiter when you have:

✅ **Single Instance Deployment**
- Monolithic applications
- No horizontal scaling

✅ **Ultra-Low Latency Requirements**
- Sub-microsecond latency needed
- Network latency unacceptable

✅ **No External Dependencies**
- Minimal infrastructure
- Embedded applications

✅ **Simple Use Cases**
- Basic request throttling
- No need for distributed coordination

## Migration Steps

### Step 1: Add Redis Backend Dependency

```bash
go get github.com/felixgeelhaar/fortify/backends/redis
```

### Step 2: Set Up Redis

Choose your Redis deployment:

**Option A: Standalone Redis**
```bash
docker run -d -p 6379:6379 redis:7-alpine
```

**Option B: Redis Cluster** (production)
```bash
# See Redis Cluster documentation
# https://redis.io/docs/manual/scaling/
```

**Option C: Managed Redis** (recommended for production)
- AWS ElastiCache
- Azure Cache for Redis
- Google Cloud Memorystore
- Redis Enterprise Cloud

### Step 3: Update Code

**Before (in-memory):**

```go
package main

import (
    "github.com/felixgeelhaar/fortify/ratelimit"
)

func setupRateLimiter() ratelimit.RateLimiter {
    return ratelimit.New(ratelimit.Config{
        Rate:     100,
        Burst:    200,
        Interval: time.Second,
        KeyFunc: func(ctx context.Context) string {
            return ctx.Value("user_id").(string)
        },
        Logger:  logger,
        OnLimit: onLimitCallback,
    })
}
```

**After (Redis):**

```go
package main

import (
    "github.com/redis/go-redis/v9"
    redisrl "github.com/felixgeelhaar/fortify/backends/redis"
)

func setupRateLimiter() (ratelimit.RateLimiter, error) {
    // Create Redis client
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       0,
    })

    // Create Redis-backed rate limiter
    return redisrl.New(redisrl.Config{
        Client:   client,  // ← Add Redis client
        Rate:     100,
        Burst:    200,
        Interval: time.Second,
        KeyFunc: func(ctx context.Context) string {
            return ctx.Value("user_id").(string)
        },
        Logger:  logger,
        OnLimit: onLimitCallback,
    })
}
```

### Step 4: Update Imports

```go
// Add new import
import (
    "github.com/redis/go-redis/v9"
    redisrl "github.com/felixgeelhaar/fortify/backends/redis"
)

// Keep interface import unchanged
import (
    "github.com/felixgeelhaar/fortify/ratelimit"
)
```

### Step 5: Test Locally

```bash
# Start Redis
docker run -d -p 6379:6379 redis:7-alpine

# Run your application
go run main.go

# Test rate limiting
for i in {1..15}; do
  curl http://localhost:8080/api/endpoint
done
```

### Step 6: Deploy to Staging

1. Deploy Redis infrastructure
2. Update application with Redis configuration
3. Run integration tests
4. Monitor for errors and performance

### Step 7: Production Rollout

Use a phased rollout:

1. **Canary Deployment** (10% of traffic)
2. **Monitor Metrics** (latency, errors, rate limit accuracy)
3. **Gradual Rollout** (25% → 50% → 100%)
4. **Full Deployment**

## Configuration Comparison

### In-Memory Configuration

```go
ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    KeyFunc:  extractKey,
    Logger:   logger,
    OnLimit:  callback,
})
```

### Redis Configuration

```go
redisrl.New(redisrl.Config{
    // Required: Add Redis client
    Client: redisClient,

    // Same as in-memory
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    KeyFunc:  extractKey,
    Logger:   logger,
    OnLimit:  callback,

    // Redis-specific options
    KeyPrefix:       "myapp:rl:",    // Optional
    BucketTTL:       time.Hour,      // Optional
    FallbackOnError: false,          // Optional
})
```

### New Configuration Options

| Option | Default | Description |
|--------|---------|-------------|
| `Client` | *required* | Redis client (standalone, cluster, or sentinel) |
| `KeyPrefix` | `"fortify:ratelimit:"` | Redis key namespace |
| `BucketTTL` | `1 hour` | Auto-expire idle buckets |
| `FallbackOnError` | `false` | Behavior on Redis failure |

## Code Changes

### Minimal Migration (Drop-In Replacement)

```go
// 1. Add Redis client creation
client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// 2. Change constructor
- limiter := ratelimit.New(config)
+ limiter, err := redisrl.New(redisrl.Config{
+     Client: client,
+     ...config,
+ })
+ if err != nil {
+     log.Fatal(err)
+ }

// 3. Everything else stays the same!
limiter.Allow(ctx, key)
limiter.Wait(ctx, key)
limiter.Take(ctx, key, n)
```

### Recommended Migration (With Error Handling)

```go
func createRateLimiter(redisURL string, logger *slog.Logger) (ratelimit.RateLimiter, error) {
    // Parse Redis URL
    opts, err := redis.ParseURL(redisURL)
    if err != nil {
        return nil, fmt.Errorf("invalid redis URL: %w", err)
    }

    // Create client with connection pool
    client := redis.NewClient(opts)

    // Verify connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("redis connection failed: %w", err)
    }

    // Create rate limiter
    limiter, err := redisrl.New(redisrl.Config{
        Client:          client,
        Rate:            100,
        Burst:           200,
        Interval:        time.Second,
        Logger:          logger,
        BucketTTL:       time.Hour,
        FallbackOnError: false, // Fail-closed for security
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create limiter: %w", err)
    }

    return limiter, nil
}
```

## Testing Strategy

### Unit Tests

No changes needed if using the `RateLimiter` interface:

```go
func TestMyHandler(t *testing.T) {
    // Use miniredis for testing
    mr := miniredis.RunT(t)
    client := redis.NewClient(&redis.Options{
        Addr: mr.Addr(),
    })

    limiter, _ := redisrl.New(redisrl.Config{
        Client: client,
        Rate:   10,
        Burst:  10,
    })

    // Test as before
    testHandler(t, limiter)
}
```

### Integration Tests

Test distributed behavior:

```go
func TestDistributedRateLimiting(t *testing.T) {
    // Start real Redis for integration test
    // Or use miniredis

    // Create multiple limiters (simulating multiple instances)
    limiter1, _ := redisrl.New(config)
    limiter2, _ := redisrl.New(config)

    // Both should share same limits
    for i := 0; i < 10; i++ {
        limiter1.Allow(ctx, "user-123")
    }

    // This should be denied (bucket empty)
    if limiter2.Allow(ctx, "user-123") {
        t.Error("should be rate limited across instances")
    }
}
```

### Load Tests

Compare performance before/after:

```bash
# Before (in-memory)
hey -n 100000 -c 100 http://localhost:8080/api/endpoint

# After (Redis)
hey -n 100000 -c 100 http://localhost:8080/api/endpoint
```

## Performance Considerations

### Latency Impact

Expect these latency increases:

| Deployment | Additional Latency |
|------------|-------------------|
| Local Redis | ~0.5ms |
| Same-AZ Redis | ~1-2ms |
| Cross-AZ Redis | ~5-10ms |
| Cross-Region | ~50-100ms |

**Mitigation:**
- Deploy Redis in same availability zone
- Use Redis Cluster for better distribution
- Monitor p99 latency

### Throughput

Redis backend can handle:
- **50k-100k ops/sec** on single Redis instance
- **500k+ ops/sec** with Redis Cluster

**Tuning:**
```go
client := redis.NewClient(&redis.Options{
    Addr:         "localhost:6379",
    PoolSize:     100,  // Increase for high concurrency
    MinIdleConns: 20,
})
```

### Memory Usage

Each rate limit bucket uses ~100 bytes in Redis:
- 1M buckets ≈ 100 MB
- 10M buckets ≈ 1 GB

**Optimization:**
```go
redisrl.Config{
    BucketTTL: 10 * time.Minute,  // Shorter TTL = less memory
    // ...
}
```

## Production Deployment

### Infrastructure Setup

**1. Redis Deployment**

```yaml
# Kubernetes example
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: redis
spec:
  serviceName: redis
  replicas: 3
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:7-alpine
        ports:
        - containerPort: 6379
        resources:
          limits:
            memory: "2Gi"
            cpu: "1000m"
          requests:
            memory: "1Gi"
            cpu: "500m"
```

**2. Application Configuration**

```go
// Use environment variables
redisURL := os.Getenv("REDIS_URL")
if redisURL == "" {
    redisURL = "redis://localhost:6379"
}

opts, _ := redis.ParseURL(redisURL)
client := redis.NewClient(opts)
```

**3. Connection Pooling**

```go
client := redis.NewClient(&redis.Options{
    Addr:         redisURL,
    PoolSize:     runtime.NumCPU() * 10,
    MinIdleConns: runtime.NumCPU() * 2,
    DialTimeout:  5 * time.Second,
    ReadTimeout:  3 * time.Second,
    WriteTimeout: 3 * time.Second,
})
```

### Monitoring

Track these metrics:

**Application Metrics:**
- Rate limit hits/misses
- Redis connection errors
- Latency percentiles (p50, p95, p99)

**Redis Metrics:**
- Memory usage
- Commands per second
- Evicted keys
- Connection count

**Example with Prometheus:**

```go
import "github.com/felixgeelhaar/fortify/metrics"

// Register metrics
metrics.MustRegister(prometheus.DefaultRegisterer)

// Metrics automatically collected
```

### Health Checks

```go
func healthCheck(redisClient *redis.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
        defer cancel()

        if err := redisClient.Ping(ctx).Err(); err != nil {
            http.Error(w, "Redis unhealthy", http.StatusServiceUnavailable)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }
}
```

## Rollback Plan

If issues arise, rollback quickly:

### Quick Rollback (Emergency)

```go
// Switch back to in-memory (no Redis dependency)
- import redisrl "github.com/felixgeelhaar/fortify/backends/redis"
+ import "github.com/felixgeelhaar/fortify/ratelimit"

- limiter, _ := redisrl.New(redisrl.Config{
-     Client: client,
+ limiter := ratelimit.New(ratelimit.Config{
      Rate:   100,
      Burst:  200,
  })
```

### Gradual Rollback

1. Deploy in-memory version to canary instances
2. Monitor for improvements
3. Gradually roll back remaining instances

### Feature Flag Approach

```go
func createRateLimiter() ratelimit.RateLimiter {
    if os.Getenv("USE_REDIS_RATE_LIMIT") == "true" {
        return createRedisLimiter()
    }
    return createInMemoryLimiter()
}
```

## Troubleshooting

### High Latency

**Symptoms:** Slow API responses

**Solutions:**
1. Check Redis latency: `redis-cli --latency`
2. Verify network connectivity
3. Move Redis closer to app (same AZ)
4. Increase Redis resources

### Memory Issues

**Symptoms:** Redis OOM errors

**Solutions:**
1. Reduce `BucketTTL`
2. Set `maxmemory-policy allkeys-lru`
3. Scale Redis vertically or horizontally
4. Monitor evicted keys

### Rate Limits Not Enforced

**Symptoms:** More requests than expected

**Solutions:**
1. Verify Redis connectivity
2. Check Lua script execution
3. Ensure consistent `KeyFunc` across instances
4. Verify clock synchronization

### Connection Pool Exhausted

**Symptoms:** Connection timeout errors

**Solutions:**
1. Increase `PoolSize`
2. Increase `MinIdleConns`
3. Reduce connection idle timeout
4. Scale application horizontally

## Best Practices

✅ **Use Separate Redis Instance**
- Don't share with application cache
- Easier capacity planning

✅ **Monitor Redis Health**
- Set up alerts for connection failures
- Track memory and CPU usage

✅ **Configure Appropriate TTLs**
- Balance memory vs. bucket reuse
- Typical: 1 hour to 24 hours

✅ **Test Failover Scenarios**
- Redis restart
- Network partition
- High load conditions

✅ **Use Structured Logging**
- Log rate limit events
- Include user/key identifiers

## Next Steps

After successful migration:

1. Monitor metrics for 1-2 weeks
2. Tune configuration based on production load
3. Set up automated alerts
4. Document Redis runbooks
5. Plan for capacity growth

## Support

- 📖 [Redis Backend README](../backends/redis/README.md)
- 📖 [Main Documentation](../README.md)
- 🐛 [Issue Tracker](https://github.com/felixgeelhaar/fortify/issues)
- 💬 [Discussions](https://github.com/felixgeelhaar/fortify/discussions)
