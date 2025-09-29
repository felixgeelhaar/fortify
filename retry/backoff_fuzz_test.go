package retry

import (
	"testing"
	"time"
)

func FuzzBackoffCalculation(f *testing.F) {
	// Seed corpus with typical values
	f.Add(int64(100), float64(2.0), 5, int64(5000), true)
	f.Add(int64(10), float64(1.5), 3, int64(1000), false)
	f.Add(int64(50), float64(3.0), 10, int64(10000), true)

	f.Fuzz(func(t *testing.T, initialDelayMs int64, multiplier float64, attempt int, maxDelayMs int64, jitter bool) {
		// Skip invalid inputs
		if initialDelayMs <= 0 || multiplier <= 0 || attempt <= 0 || maxDelayMs < 0 {
			t.Skip()
		}

		initialDelay := time.Duration(initialDelayMs) * time.Millisecond
		maxDelay := time.Duration(maxDelayMs) * time.Millisecond

		// Test all backoff policies
		policies := []BackoffPolicy{BackoffExponential, BackoffLinear, BackoffConstant}

		for _, policy := range policies {
			// Should not panic
			delay := calculateBackoff(policy, attempt, initialDelay, maxDelay, multiplier, jitter)

			// Delay should be non-negative
			if delay < 0 {
				t.Errorf("negative delay: %v", delay)
			}

			// If maxDelay is set and non-zero, delay should not exceed it (accounting for jitter)
			if maxDelay > 0 && delay > maxDelay*11/10 { // Allow 10% over for jitter
				t.Errorf("delay %v exceeds maxDelay %v (policy: %v, attempt: %d)", delay, maxDelay, policy, attempt)
			}
		}
	})
}
