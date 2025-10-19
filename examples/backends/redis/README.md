# Distributed Rate Limiting Example

This example demonstrates how to use Fortify's Redis-backed rate limiter for distributed rate limiting across multiple application instances.

## Overview

The example includes:
- **3 application instances** running the same HTTP API
- **Shared Redis instance** for distributed rate limiting
- **Docker Compose** setup for easy deployment
- **Rate limiting** across all instances (10 requests/minute per user)

## Quick Start

### Prerequisites

- Docker and Docker Compose
- `curl` and `jq` (for testing)

### Start the Services

```bash
make start
```

This starts:
- Redis server on port 6379
- App instance 1 on port 8081
- App instance 2 on port 8082
- App instance 3 on port 8083

### Test Distributed Rate Limiting

```bash
make test
```

This sends 15 requests distributed across the 3 instances with the same user ID. You'll see:
- First 10 requests succeed (within rate limit)
- Remaining 5 requests are rate limited with 429 status

**Key observation**: Rate limiting works across all instances because they share the same Redis backend!

### View Logs

```bash
# All services
make logs

# Specific instance
make logs-app1
make logs-app2
make logs-app3

# Redis
make logs-redis
```

### Stop Services

```bash
make stop
```

## Endpoints

Each application instance exposes:

- `GET /health` - Health check (not rate limited)
- `GET /api/data` - Sample API endpoint (rate limited)
- `GET /api/status` - Rate limit status

## Configuration

Rate limiting is configured in `main.go`:

```go
const (
    requestsPerMinute = 10  // Max requests per minute
    burstSize         = 15  // Burst capacity
)
```

### Environment Variables

- `PORT` - HTTP server port (default: 8080)
- `REDIS_URL` - Redis connection URL (default: localhost:6379)
- `REDIS_PASSWORD` - Redis password (optional)

## Testing Examples

### Test Single Instance

```bash
# Send 15 requests to one instance
for i in {1..15}; do
  curl -H "X-User-ID: user-123" http://localhost:8081/api/data
done
```

### Test Multiple Users

```bash
# Different users have independent rate limits
curl -H "X-User-ID: alice" http://localhost:8081/api/data
curl -H "X-User-ID: bob" http://localhost:8082/api/data
```

### Test Distributed

```bash
# Same user across different instances shares rate limit
curl -H "X-User-ID: user-123" http://localhost:8081/api/data
curl -H "X-User-ID: user-123" http://localhost:8082/api/data
curl -H "X-User-ID: user-123" http://localhost:8083/api/data
```

## Monitoring

### Check Health

```bash
make health
```

### View Redis Keys

```bash
make redis-keys
```

### Monitor Redis Commands

```bash
make redis-monitor
```

### Redis CLI

```bash
make redis-cli

# Inside Redis CLI, inspect a rate limit bucket
HGETALL example:ratelimit:user-123
```

## Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   App 1     │     │   App 2     │     │   App 3     │
│  :8081      │     │  :8082      │     │  :8083      │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       │                   │                   │
       └───────────────────┴───────────────────┘
                           │
                           │ Shared Rate Limits
                           ▼
                    ┌─────────────┐
                    │    Redis    │
                    │   :6379     │
                    └─────────────┘
```

All instances share the same Redis instance, ensuring consistent rate limiting across the distributed system.

## How It Works

1. **Request arrives** at any instance (e.g., App 1)
2. **User ID extracted** from `X-User-ID` header
3. **Lua script executed** atomically in Redis:
   - Calculate token refill based on elapsed time
   - Check if tokens available
   - Consume token if available
4. **Response returned** based on result:
   - `200 OK` if allowed
   - `429 Too Many Requests` if rate limited

## Key Features Demonstrated

- ✅ **Atomic operations** via Lua scripts
- ✅ **Distributed rate limiting** across multiple instances
- ✅ **Per-user rate limits** with independent buckets
- ✅ **Token bucket algorithm** with burst support
- ✅ **Graceful error handling** with fail-closed behavior
- ✅ **Structured logging** with JSON output
- ✅ **Health checks** for monitoring
- ✅ **Auto-cleanup** of idle buckets via TTL

## Production Considerations

### Scaling

For production deployments:

1. **Redis Cluster**: Use Redis Cluster for horizontal scaling
2. **Connection Pooling**: Tune pool size based on load
3. **Monitoring**: Track Redis latency and memory usage
4. **Alerting**: Set up alerts for rate limit violations

### High Availability

Consider:

1. **Redis Sentinel**: For automatic failover
2. **Fail-open vs Fail-closed**: Choose based on requirements
3. **Circuit Breakers**: Protect against Redis failures
4. **Backup Redis**: Read replicas for disaster recovery

## Troubleshooting

### Services Won't Start

```bash
# Check service status
docker-compose ps

# View logs
docker-compose logs
```

### Rate Limiting Not Working

```bash
# Check Redis connectivity
make redis-cli
PING

# View rate limit keys
make redis-keys

# Monitor Redis commands
make redis-monitor
```

### High Memory Usage

```bash
# Check Redis memory
make redis-cli
INFO memory

# Reduce bucket TTL in main.go
BucketTTL: 10 * time.Minute  // Instead of 1 hour
```

## Cleanup

```bash
# Stop and remove all containers and volumes
make clean
```

## Next Steps

- See [Redis Backend README](../../backends/redis/README.md) for detailed documentation
- Read [Migration Guide](../../docs/MIGRATION_REDIS.md) for migrating from in-memory
- Check [Main README](../../README.md) for other resilience patterns

## License

MIT License - see [LICENSE](../../LICENSE) file for details.
