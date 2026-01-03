// Package main demonstrates basic bulkhead usage.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/felixgeelhaar/fortify/v2/bulkhead"
)

func main() {
	// Create bulkhead with concurrency limit of 3
	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 3,
		MaxQueue:      2,
		QueueTimeout:  time.Second,
		OnRejected: func() {
			fmt.Println("Request rejected: bulkhead full")
		},
	})

	// Launch 10 concurrent requests
	fmt.Println("Launching 10 concurrent requests (max 3 concurrent, 2 queued)...")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			result, err := bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
				fmt.Printf("Request %d: executing (simulating 500ms work)\n", id)
				time.Sleep(time.Millisecond * 500)
				return id * 10, nil
			})

			if err != nil {
				fmt.Printf("Request %d: ✗ %v\n", id, err)
			} else {
				fmt.Printf("Request %d: ✓ result=%d\n", id, result)
			}
		}(i + 1)

		// Stagger requests slightly
		time.Sleep(time.Millisecond * 50)
	}

	wg.Wait()

	// Example with context cancellation
	fmt.Println("\n--- Context Cancellation Example ---")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	result, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
		fmt.Println("Starting long-running operation...")
		select {
		case <-ctx.Done():
			fmt.Println("Operation cancelled")
			return 0, ctx.Err()
		case <-time.After(time.Second):
			return 42, nil
		}
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
