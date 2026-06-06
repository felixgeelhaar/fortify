// Command ratelimit-redis demonstrates wiring the Redis-backed Store into
// fortify's rate limiter.
//
// Run a local Redis (e.g., `docker run -p 6379:6379 redis:7-alpine`) and
// then `go run .` in this directory. The program rate-limits 50 calls
// against a 10-token-per-second bucket and prints how many were allowed.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	redisstore "go.klarlabs.de/fortify/examples/ratelimit-redis"
	"go.klarlabs.de/fortify/ratelimit"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer func() { _ = client.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis ping failed; is Redis running on :6379? %v", err)
	}

	store := redisstore.New(client, redisstore.WithPrefix("demo:"), redisstore.WithTTL(time.Hour))

	rl := ratelimit.New(ratelimit.Config{
		Rate:     10,
		Burst:    20,
		Interval: time.Second,
		Store:    store,
		FailOpen: true,
	})
	defer func() { _ = rl.Close() }()

	allowed, denied := 0, 0
	// Attach bucket params so the Lua script gets the right rate/interval.
	bucketCtx := redisstore.WithBucketParams(context.Background(), 10, time.Second)

	for i := 0; i < 50; i++ {
		if rl.Allow(bucketCtx, "user-1") {
			allowed++
		} else {
			denied++
		}
	}
	fmt.Printf("allowed=%d denied=%d (rate=10/s burst=20)\n", allowed, denied)
}
