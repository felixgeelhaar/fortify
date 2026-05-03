package adaptive_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/felixgeelhaar/fortify/adaptive"
)

// Example demonstrates basic adaptive limiter usage.
func Example() {
	a := adaptive.New[string](adaptive.Config{
		InitialLimit:     5,
		MinLimit:         1,
		MaxLimit:         50,
		SuccessThreshold: 10,
	})

	result, err := a.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "ok", nil
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("result:", result, "limit:", a.Limit())
	// Output: result: ok limit: 5
}

// Example_aimd shows the additive-increase / multiplicative-decrease behavior.
func Example_aimd() {
	a := adaptive.New[int](adaptive.Config{
		InitialLimit:     10,
		MinLimit:         1,
		MaxLimit:         100,
		SuccessThreshold: 5, // 5 successes → +1 to limit
	})

	// 5 successes raise the limit by 1 (10 → 11)
	for i := 0; i < 5; i++ {
		_, _ = a.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	fmt.Println("after 5 successes, limit:", a.Limit())

	// One failure halves the limit (11 → 5)
	_, _ = a.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, errors.New("downstream overloaded")
	})
	fmt.Println("after 1 failure, limit:", a.Limit())

	// Output:
	// after 5 successes, limit: 11
	// after 1 failure, limit: 5
}
