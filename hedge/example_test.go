package hedge_test

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"go.klarlabs.de/fortify/hedge"
)

// Example demonstrates basic hedged-request usage.
func Example() {
	h := hedge.New[string](hedge.Config{
		MaxAttempts: 3,
		HedgeDelay:  50 * time.Millisecond,
	})

	// Simulate a downstream where the first attempt is slow but later
	// attempts return quickly.
	var calls atomic.Int32
	result, err := h.Execute(context.Background(), func(ctx context.Context) (string, error) {
		n := calls.Add(1)
		if n == 1 {
			// Primary: slow.
			select {
			case <-time.After(200 * time.Millisecond):
				return "", errors.New("primary timed out")
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
		// Hedge: fast.
		return "ok", nil
	})

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
	// Output: result: ok
}

// Example_onHedge shows how to observe hedge fires for metrics or tracing.
func Example_onHedge() {
	var hedgeFired atomic.Int32

	h := hedge.New[int](hedge.Config{
		MaxAttempts: 2,
		HedgeDelay:  10 * time.Millisecond,
		OnHedge: func(attempt int) {
			hedgeFired.Add(1)
		},
	})

	_, _ = h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		// Always return quickly so the hedge usually doesn't fire.
		return 1, nil
	})

	fmt.Printf("hedges fired: %d\n", hedgeFired.Load())
	// Output: hedges fired: 0
}
