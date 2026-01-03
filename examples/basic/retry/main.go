// Package main demonstrates basic retry usage.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/v2/retry"
)

func main() {
	// Create retry with exponential backoff
	r := retry.New[int](retry.Config{
		MaxAttempts:   5,
		InitialDelay:  time.Millisecond * 100,
		MaxDelay:      time.Second * 5,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		OnRetry: func(attempt int, err error) {
			fmt.Printf("Attempt %d failed: %v\n", attempt, err)
		},
	})

	attempt := 0
	result, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		fmt.Printf("Executing attempt %d...\n", attempt)

		// Succeed on the 4th attempt
		if attempt < 4 {
			return 0, errors.New("temporary failure")
		}
		return 42, nil
	})

	if err != nil {
		fmt.Printf("\nFinal result: Failed after all retries: %v\n", err)
	} else {
		fmt.Printf("\nFinal result: Success with value %d\n", result)
	}

	// Example with context cancellation
	fmt.Println("\n--- Context Cancellation Example ---")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	attempt = 0
	_, err = r.Do(ctx, func(ctx context.Context) (int, error) {
		attempt++
		fmt.Printf("Executing attempt %d...\n", attempt)
		time.Sleep(time.Millisecond * 200)
		return 0, errors.New("still failing")
	})

	if err != nil {
		fmt.Printf("Operation cancelled: %v\n", err)
	}
}
