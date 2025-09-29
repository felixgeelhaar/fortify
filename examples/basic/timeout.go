// Package main demonstrates basic timeout usage.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/timeout"
)

func main() {
	// Create timeout manager
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 5,
		OnTimeout: func(duration time.Duration) {
			fmt.Printf("Operation timed out after %v\n", duration)
		},
	})

	// Example 1: Operation completes within timeout
	fmt.Println("--- Example 1: Operation completes within timeout ---")
	result, err := tm.Execute(context.Background(), time.Second, func(ctx context.Context) (string, error) {
		time.Sleep(time.Millisecond * 500)
		return "completed successfully", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	// Example 2: Operation times out
	fmt.Println("\n--- Example 2: Operation times out ---")
	_, err = tm.Execute(context.Background(), time.Millisecond*500, func(ctx context.Context) (string, error) {
		time.Sleep(time.Second * 2)
		return "should not reach here", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Example 3: Respecting context cancellation
	fmt.Println("\n--- Example 3: Respecting context cancellation ---")
	result, err = tm.Execute(context.Background(), time.Second, func(ctx context.Context) (string, error) {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			default:
				fmt.Printf("Working... step %d/10\n", i+1)
				time.Sleep(time.Millisecond * 100)
			}
		}
		return "completed all steps", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	// Example 4: Using default timeout
	fmt.Println("\n--- Example 4: Using default timeout ---")
	result, err = tm.ExecuteWithDefault(context.Background(), func(ctx context.Context) (string, error) {
		time.Sleep(time.Millisecond * 100)
		return "using default timeout", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}
