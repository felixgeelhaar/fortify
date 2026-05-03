// Command demo exercises the Redis-backed distributed circuit breaker.
//
// Usage:
//
//	docker run --rm -p 6379:6379 redis:7-alpine
//	go run ./cmd/demo
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	redisbreaker "github.com/felixgeelhaar/fortify/examples/circuitbreaker-redis"
	"github.com/felixgeelhaar/fortify/circuitbreaker"
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

	cb, err := redisbreaker.New[int](redisbreaker.Config{
		Client:           client,
		Key:              "demo:downstream-A",
		FailureThreshold: 3,
		OpenTimeout:      2 * time.Second,
		MaxRequests:      1,
		OnStateChange: func(from, to circuitbreaker.State) {
			fmt.Printf("[transition] %s -> %s\n", from, to)
		},
	})
	if err != nil {
		log.Fatalf("breaker: %v", err)
	}
	defer func() { _ = cb.Close() }()
	cb.Reset()

	// Trip the breaker with 3 consecutive failures.
	for i := 0; i < 3; i++ {
		_, _ = cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("downstream is down")
		})
	}
	fmt.Println("state after 3 failures:", cb.State())

	// Should now reject.
	_, err = cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})
	fmt.Println("call while open:", err)

	// Wait for OpenTimeout, then a successful trial closes the breaker.
	time.Sleep(2100 * time.Millisecond)
	v, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})
	fmt.Println("call after timeout:", v, err)
	fmt.Println("final state:", cb.State())
}
