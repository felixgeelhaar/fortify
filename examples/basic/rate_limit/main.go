// Package main demonstrates basic rate limiter usage.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/ratelimit"
)

func main() {
	// Create a token bucket rate limiter
	// Allows 5 requests per second with burst of 10
	rl := ratelimit.New(&ratelimit.Config{
		Rate:     5,
		Burst:    10,
		Interval: time.Second,
	})

	// Simulate rapid requests
	fmt.Println("Attempting 15 rapid requests...")
	for i := 0; i < 15; i++ {
		if rl.Allow(context.Background(), "user-123") {
			fmt.Printf("Request %d: ✓ Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: ✗ Rate limited\n", i+1)
		}
	}

	// Wait for token refill
	fmt.Println("\nWaiting 1 second for token refill...")
	time.Sleep(time.Second)

	// Try more requests
	fmt.Println("Attempting 5 more requests...")
	for i := 0; i < 5; i++ {
		if rl.Allow(context.Background(), "user-123") {
			fmt.Printf("Request %d: ✓ Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: ✗ Rate limited\n", i+1)
		}
	}

	// Example with Wait (blocking)
	fmt.Println("\n--- Blocking Wait Example ---")
	start := time.Now()
	for i := 0; i < 3; i++ {
		fmt.Printf("Waiting for token %d...\n", i+1)
		if err := rl.Wait(context.Background(), "user-456"); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			elapsed := time.Since(start)
			fmt.Printf("Token %d acquired after %v\n", i+1, elapsed.Round(time.Millisecond*10))
		}
	}
}
