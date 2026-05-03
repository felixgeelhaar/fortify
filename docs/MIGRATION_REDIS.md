# Custom Storage Backend Migration Guide

This guide helps you implement custom storage backends for Fortify's rate limiter using the `Store` interface.

## Table of Contents

- [Overview](#overview)
- [Store Interface](#store-interface)
- [Implementation Examples](#implementation-examples)
  - [Redis Implementation](#redis-implementation)
  - [DynamoDB Implementation](#dynamodb-implementation)
- [Migration from In-Memory](#migration-from-in-memory)
- [Configuration](#configuration)
- [Testing Strategy](#testing-strategy)
- [Production Deployment](#production-deployment)
- [Troubleshooting](#troubleshooting)

## Overview

Fortify's rate limiter uses a pluggable `Store` interface for state management. By default, an in-memory store is used. For distributed rate limiting across multiple application instances, implement a custom `Store` backed by Redis, DynamoDB, or another distributed backend.

### When to Use a Custom Store

Use a custom distributed store when you have:

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

### When to Use In-Memory (Default)

Keep the default in-memory store when you have:

✅ **Single Instance Deployment**
- Monolithic applications
- No horizontal scaling

✅ **Ultra-Low Latency Requirements**
- Sub-microsecond latency needed
- Network latency unacceptable

## Store Interface

The `Store` interface defines three methods:

```go
type BucketState struct {
    Tokens     float64   // Current available tokens
    LastRefill time.Time // Last refill timestamp
}

type Store interface {
    // AtomicUpdate atomically reads, modifies, and writes bucket state.
    // The updateFn receives current state (nil if new) and returns new state.
    // Implementation must ensure atomic read-modify-write.
    AtomicUpdate(ctx context.Context, key string, updateFn func(*BucketState) *BucketState) (*BucketState, error)

    // Delete removes a bucket from the store.
    Delete(ctx context.Context, key string) error

    // Close releases resources held by the store.
    Close() error
}
```

### Atomicity Requirements

The `AtomicUpdate` method must guarantee:

1. **Atomic read-modify-write** - The entire operation must be indivisible
2. **No race conditions** - Concurrent calls for the same key must be serialized
3. **Correct token bucket algorithm** - The `updateFn` contains the algorithm logic

For distributed backends, achieve atomicity using:
- **Redis**: WATCH/MULTI/EXEC or Lua scripts
- **DynamoDB**: Conditional writes with version attributes
- **PostgreSQL**: SELECT FOR UPDATE transactions

## Implementation Examples

### Redis Implementation

```go
package redisstore

import (
    "context"
    "encoding/json"
    "time"

    "github.com/felixgeelhaar/fortify/ratelimit"
    "github.com/redis/go-redis/v9"
)

type RedisStore struct {
    client redis.UniversalClient
    prefix string
    ttl    time.Duration
}

func New(client redis.UniversalClient, prefix string, ttl time.Duration) *RedisStore {
    return &RedisStore{
        client: client,
        prefix: prefix,
        ttl:    ttl,
    }
}

func (r *RedisStore) AtomicUpdate(ctx context.Context, key string,
    updateFn func(*ratelimit.BucketState) *ratelimit.BucketState) (*ratelimit.BucketState, error) {

    redisKey := r.prefix + key
    var finalState *ratelimit.BucketState

    // Use WATCH for optimistic locking with retries
    err := r.client.Watch(ctx, func(tx *redis.Tx) error {
        // Get current state
        var state *ratelimit.BucketState
        data, err := tx.Get(ctx, redisKey).Bytes()
        if err == nil {
            state = &ratelimit.BucketState{}
            if err := json.Unmarshal(data, state); err != nil {
                return err
            }
        } else if err != redis.Nil {
            return err
        }

        // Apply update function (token bucket algorithm runs here)
        newState := updateFn(state)
        finalState = newState

        if newState == nil {
            return nil // No change needed
        }

        // Store atomically
        newData, err := json.Marshal(newState)
        if err != nil {
            return err
        }

        _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
            pipe.Set(ctx, redisKey, newData, r.ttl)
            return nil
        })
        return err
    }, redisKey)

    if err != nil {
        return nil, err
    }

    return finalState, nil
}

func (r *RedisStore) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, r.prefix+key).Err()
}

func (r *RedisStore) Close() error {
    return r.client.Close()
}
```

### DynamoDB Implementation

```go
package dynamostore

import (
    "context"
    "strconv"
    "time"

    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
    "github.com/felixgeelhaar/fortify/ratelimit"
)

type DynamoStore struct {
    client    *dynamodb.Client
    tableName string
    ttl       time.Duration
}

type bucketItem struct {
    Key        string  `dynamodbav:"pk"`
    Tokens     float64 `dynamodbav:"tokens"`
    LastRefill int64   `dynamodbav:"last_refill"`
    Version    int64   `dynamodbav:"version"`
    TTL        int64   `dynamodbav:"ttl"`
}

func New(client *dynamodb.Client, tableName string, ttl time.Duration) *DynamoStore {
    return &DynamoStore{
        client:    client,
        tableName: tableName,
        ttl:       ttl,
    }
}

func (d *DynamoStore) AtomicUpdate(ctx context.Context, key string,
    updateFn func(*ratelimit.BucketState) *ratelimit.BucketState) (*ratelimit.BucketState, error) {

    // Retry loop for conditional write conflicts
    for attempts := 0; attempts < 10; attempts++ {
        // Get current item
        result, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
            TableName: &d.tableName,
            Key: map[string]types.AttributeValue{
                "pk": &types.AttributeValueMemberS{Value: key},
            },
        })
        if err != nil {
            return nil, err
        }

        var state *ratelimit.BucketState
        var version int64 = 0

        if result.Item != nil {
            var item bucketItem
            if err := attributevalue.UnmarshalMap(result.Item, &item); err != nil {
                return nil, err
            }
            state = &ratelimit.BucketState{
                Tokens:     item.Tokens,
                LastRefill: time.Unix(0, item.LastRefill),
            }
            version = item.Version
        }

        // Apply update function
        newState := updateFn(state)
        if newState == nil {
            return state, nil
        }

        // Conditional write with version check
        newItem := bucketItem{
            Key:        key,
            Tokens:     newState.Tokens,
            LastRefill: newState.LastRefill.UnixNano(),
            Version:    version + 1,
            TTL:        time.Now().Add(d.ttl).Unix(),
        }

        item, err := attributevalue.MarshalMap(newItem)
        if err != nil {
            return nil, err
        }

        conditionExpr := "attribute_not_exists(pk) OR version = :v"
        _, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
            TableName:           &d.tableName,
            Item:                item,
            ConditionExpression: &conditionExpr,
            ExpressionAttributeValues: map[string]types.AttributeValue{
                ":v": &types.AttributeValueMemberN{Value: strconv.FormatInt(version, 10)},
            },
        })

        if err == nil {
            return newState, nil
        }

        // Check if it's a condition check failure (retry)
        var ccf *types.ConditionalCheckFailedException
        if !errors.As(err, &ccf) {
            return nil, err
        }
        // Retry on conflict
    }

    return nil, errors.New("max retries exceeded")
}

func (d *DynamoStore) Delete(ctx context.Context, key string) error {
    _, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
        TableName: &d.tableName,
        Key: map[string]types.AttributeValue{
            "pk": &types.AttributeValueMemberS{Value: key},
        },
    })
    return err
}

func (d *DynamoStore) Close() error {
    return nil
}
```

## Migration from In-Memory

### Step 1: Implement Your Store

Choose your backend and implement the `Store` interface (see examples above).

### Step 2: Update Configuration

**Before (default in-memory):**

```go
limiter := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
})
```

**After (custom store):**

```go
// Create your custom store
redisStore := redisstore.New(redisClient, "ratelimit:", time.Hour)

limiter := ratelimit.New(ratelimit.Config{
    Rate:     100,
    Burst:    200,
    Interval: time.Second,
    Store:    redisStore,
    FailOpen: true,  // Allow requests if storage fails
})
```

### Step 3: Handle Errors

Configure `FailOpen` based on your requirements:

- `FailOpen: false` (default) - Deny requests when storage fails (consistency)
- `FailOpen: true` - Allow requests when storage fails (availability)

## Configuration

| Option | Default | Description |
|--------|---------|-------------|
| `Store` | `MemoryStore` | Storage backend implementation |
| `FailOpen` | `false` | Behavior on storage failure |
| `Rate` | `100` | Tokens added per interval |
| `Burst` | `Rate` | Maximum bucket capacity |
| `Interval` | `1s` | Refill interval |

## Testing Strategy

### Unit Tests with Mock Store

```go
type mockStore struct {
    states map[string]*ratelimit.BucketState
    mu     sync.Mutex
}

func (m *mockStore) AtomicUpdate(ctx context.Context, key string,
    updateFn func(*ratelimit.BucketState) *ratelimit.BucketState) (*ratelimit.BucketState, error) {
    m.mu.Lock()
    defer m.mu.Unlock()

    state := m.states[key]
    newState := updateFn(state)
    if newState != nil {
        m.states[key] = newState
    }
    return newState, nil
}

func TestWithMockStore(t *testing.T) {
    store := &mockStore{states: make(map[string]*ratelimit.BucketState)}
    limiter := ratelimit.New(ratelimit.Config{
        Rate:  10,
        Burst: 10,
        Store: store,
    })

    // Test as normal
}
```

### Integration Tests

```go
func TestDistributedRateLimiting(t *testing.T) {
    // Create shared store
    store := redisstore.New(redisClient, "test:", time.Hour)

    // Create multiple limiters (simulating multiple instances)
    limiter1 := ratelimit.New(ratelimit.Config{Rate: 10, Burst: 10, Store: store})
    limiter2 := ratelimit.New(ratelimit.Config{Rate: 10, Burst: 10, Store: store})

    ctx := context.Background()

    // Exhaust quota from limiter1
    for i := 0; i < 10; i++ {
        limiter1.Allow(ctx, "user-123")
    }

    // limiter2 should see empty bucket
    if limiter2.Allow(ctx, "user-123") {
        t.Error("should be rate limited across instances")
    }
}
```

## Production Deployment

### Redis Best Practices

1. **Use Redis Cluster** for high availability
2. **Deploy in same availability zone** to minimize latency
3. **Configure appropriate TTL** for bucket expiration
4. **Monitor Redis metrics** (memory, connections, latency)

### DynamoDB Best Practices

1. **Enable TTL** on the table for automatic cleanup
2. **Use on-demand capacity** for variable workloads
3. **Configure appropriate WCU/RCU** for provisioned capacity
4. **Enable point-in-time recovery** for data protection

### Health Checks

```go
func healthCheck(store ratelimit.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
        defer cancel()

        // Test store operation
        _, err := store.AtomicUpdate(ctx, "_health", func(s *ratelimit.BucketState) *ratelimit.BucketState {
            return s // No-op
        })

        if err != nil {
            http.Error(w, "Store unhealthy", http.StatusServiceUnavailable)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}
```

## Troubleshooting

### High Latency

**Symptoms:** Slow API responses

**Solutions:**
1. Move storage closer to application (same AZ)
2. Increase connection pool size
3. Check network connectivity
4. Monitor storage backend metrics

### Rate Limits Not Enforced

**Symptoms:** More requests than expected

**Solutions:**
1. Verify storage connectivity
2. Check atomicity implementation
3. Ensure consistent key generation
4. Verify clock synchronization

### Storage Errors

**Symptoms:** Requests failing or being allowed unexpectedly

**Solutions:**
1. Check `FailOpen` configuration
2. Verify storage backend health
3. Increase connection timeouts
4. Add retry logic to store implementation

## Support

- 📖 [Main Documentation](../README.md)
- 🐛 [Issue Tracker](https://github.com/felixgeelhaar/fortify/issues)
- 💬 [Discussions](https://github.com/felixgeelhaar/fortify/discussions)
