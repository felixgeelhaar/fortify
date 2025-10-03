package retry

import (
	"math"
	"math/rand"
	"time"
)

// BackoffPolicy defines the strategy for calculating retry delays.
type BackoffPolicy int

const (
	// BackoffExponential increases delay exponentially: delay = initialDelay * multiplier^attempt.
	BackoffExponential BackoffPolicy = iota

	// BackoffLinear increases delay linearly: delay = initialDelay * attempt.
	BackoffLinear

	// BackoffConstant uses a constant delay: delay = initialDelay.
	BackoffConstant
)

// calculateBackoff computes the delay before the next retry attempt.
// It applies the configured backoff policy, respects MaxDelay, and optionally adds jitter.
func calculateBackoff(policy BackoffPolicy, attempt int, initialDelay, maxDelay time.Duration, multiplier float64, jitter bool) time.Duration {
	var delay time.Duration

	switch policy {
	case BackoffExponential:
		// delay = initialDelay * multiplier^(attempt-1).
		delay = time.Duration(float64(initialDelay) * math.Pow(multiplier, float64(attempt-1)))

	case BackoffLinear:
		// delay = initialDelay * attempt.
		delay = time.Duration(float64(initialDelay) * float64(attempt))

	case BackoffConstant:
		// delay = initialDelay (constant).
		delay = initialDelay

	default:
		delay = initialDelay
	}

	// Apply max delay cap.
	if maxDelay > 0 && delay > maxDelay {
		delay = maxDelay
	}

	// Add jitter to prevent thundering herd.
	if jitter {
		delay = addJitter(delay)
	}

	return delay
}

// addJitter adds random variance to the delay to prevent synchronized retries.
// Adds 0-10% random variation to the delay.
func addJitter(delay time.Duration) time.Duration {
	if delay == 0 {
		return 0
	}

	// Add random jitter between 0% and 10% of the delay.
	// Using math/rand instead of crypto/rand because:
	// 1. Retry jitter doesn't require cryptographic security
	// 2. The goal is preventing thundering herd, not security
	// 3. math/rand is significantly faster (no syscalls)
	// 4. Predictability of math/rand is acceptable for retry timing
	//nolint:gosec // G404: weak random is intentional and appropriate for retry jitter
	jitterAmount := time.Duration(rand.Float64() * float64(delay) * 0.1)
	return delay + jitterAmount
}
