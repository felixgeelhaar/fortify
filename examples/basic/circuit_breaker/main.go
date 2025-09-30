// Package main demonstrates basic circuit breaker usage.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
)

func main() {
	// Create a circuit breaker with custom configuration
	cb := circuitbreaker.New[string](circuitbreaker.Config{
		MaxRequests: 3,
		Interval:    time.Second * 10,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			// Open circuit after 3 consecutive failures
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: func(from, to circuitbreaker.State) {
			fmt.Printf("Circuit breaker state changed: %s -> %s\n", from, to)
		},
	})

	// Simulate multiple requests
	for i := 0; i < 10; i++ {
		result, err := cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
			// Simulate a failing service
			if i < 5 {
				return "", errors.New("service unavailable")
			}
			return "success", nil
		})

		if err != nil {
			fmt.Printf("Request %d failed: %v\n", i+1, err)
		} else {
			fmt.Printf("Request %d succeeded: %s\n", i+1, result)
		}

		time.Sleep(time.Millisecond * 100)
	}

	// Wait for circuit to enter half-open state
	fmt.Println("\nWaiting for circuit to enter half-open state...")
	time.Sleep(time.Second * 5)

	// Try a successful request
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "recovered", nil
	})

	if err != nil {
		fmt.Printf("Recovery attempt failed: %v\n", err)
	} else {
		fmt.Printf("Service recovered: %s\n", result)
	}
}
