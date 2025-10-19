package redis_test

import (
	"context"
	"fmt"
	"log"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/felixgeelhaar/fortify/backends/redis"
)

func Example() {
	// Create Redis client
	client := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	// Create distributed rate limiter
	limiter, err := redis.New(redis.Config{
		Client:   client,
		Rate:     100,              // 100 requests
		Burst:    200,              // burst of 200
		Interval: time.Second,      // per second
		KeyPrefix: "myapp:limit:",  // custom key prefix
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Non-blocking check
	if limiter.Allow(ctx, "user-123") {
		fmt.Println("Request allowed")
	} else {
		fmt.Println("Rate limit exceeded")
	}

	// Blocking wait
	if err := limiter.Wait(ctx, "user-123"); err == nil {
		fmt.Println("Token acquired after waiting")
	}

	// Take multiple tokens
	if limiter.Take(ctx, "user-123", 5) {
		fmt.Println("5 tokens consumed")
	}
}

func ExampleNew() {
	client := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	limiter, err := redis.New(redis.Config{
		Client:   client,
		Rate:     50,
		Burst:    100,
		Interval: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = limiter // Use the limiter
}

func ExampleNew_withCallbacks() {
	client := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	limiter, err := redis.New(redis.Config{
		Client:   client,
		Rate:     100,
		Burst:    200,
		Interval: time.Second,
		OnLimit: func(key string) {
			// Called when rate limit is exceeded
			fmt.Printf("Rate limit exceeded for key: %s\n", key)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// This will trigger OnLimit callback if rate limited
	limiter.Allow(ctx, "user-456")
}

func ExampleNew_withKeyFunc() {
	client := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	type contextKey string
	const userIDKey contextKey = "user_id"

	limiter, err := redis.New(redis.Config{
		Client:   client,
		Rate:     100,
		Burst:    200,
		Interval: time.Second,
		// Extract key from context
		KeyFunc: func(ctx context.Context) string {
			if userID, ok := ctx.Value(userIDKey).(string); ok {
				return "user:" + userID
			}
			return "anonymous"
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Key automatically extracted from context
	ctx := context.WithValue(context.Background(), userIDKey, "123")
	limiter.Allow(ctx, "") // Key parameter ignored when KeyFunc is set
}

func ExampleNew_cluster() {
	// Create Redis Cluster client
	client := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs: []string{
			"localhost:7000",
			"localhost:7001",
			"localhost:7002",
		},
	})
	defer client.Close()

	// Works seamlessly with Redis Cluster
	limiter, err := redis.New(redis.Config{
		Client:   client,
		Rate:     1000,
		Burst:    2000,
		Interval: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	limiter.Allow(ctx, "api-endpoint")
}

func ExampleNew_fallback() {
	client := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	// Allow requests when Redis is unavailable (availability over consistency)
	limiter, err := redis.New(redis.Config{
		Client:          client,
		Rate:            100,
		Burst:           200,
		Interval:        time.Second,
		FallbackOnError: true, // Allow on Redis failure
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// If Redis is down, this will return true (fail-open)
	allowed := limiter.Allow(ctx, "user-789")
	fmt.Printf("Allowed: %v\n", allowed)
}
