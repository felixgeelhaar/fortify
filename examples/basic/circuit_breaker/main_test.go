package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
)

func TestCircuitBreakerExample(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the example
	main()

	// Restore stdout and read output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	//nolint:errcheck // ignoring error in test
	_, _ = io.Copy(&buf, r)

	// Verify circuit breaker behavior is demonstrated
	output := buf.String()

	// Should show state changes
	if output == "" {
		t.Error("expected output from circuit breaker example")
	}
}

func TestCircuitBreakerStateTransitions(t *testing.T) {
	var stateChanged atomic.Bool

	cb := circuitbreaker.New[string](circuitbreaker.Config{
		MaxRequests: 1,
		Interval:    time.Second * 10,
		Timeout:     time.Second,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			// Open after 2 consecutive failures
			return counts.ConsecutiveFailures >= 2
		},
		OnStateChange: func(from, to circuitbreaker.State) {
			stateChanged.Store(true)
		},
	})

	// Trigger failures to open circuit
	// Need at least 2 failures to trip based on our ReadyToTrip function
	for i := 0; i < 3; i++ {
		//nolint:errcheck // ignoring error in test
		_, _ = cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
			return "", errors.New("failure")
		})
	}

	// Give callback time to execute
	time.Sleep(time.Millisecond * 10)

	// Verify circuit state changed
	if !stateChanged.Load() {
		t.Error("expected circuit breaker to change state after consecutive failures")
	}
}
