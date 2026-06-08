package retry

import (
	"math"
	"math/rand/v2"
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

// maxBackoffDuration caps any backoff computation to a value safely
// representable as time.Duration without wrapping. Without this guard, a
// pathological Multiplier or attempt count produces +Inf via math.Pow, which
// converts to a negative time.Duration and breaks select{} timers.
const maxBackoffDuration = 24 * time.Hour

// calculateBackoff computes the delay before the next retry attempt.
// It applies the configured backoff policy, respects MaxDelay, and optionally adds jitter.
func calculateBackoff(policy BackoffPolicy, attempt int, initialDelay, maxDelay time.Duration, multiplier float64, jitter bool) time.Duration {
	var delayF float64

	switch policy {
	case BackoffExponential:
		// delay = initialDelay * multiplier^(attempt-1).
		delayF = float64(initialDelay) * math.Pow(multiplier, float64(attempt-1))

	case BackoffLinear:
		// delay = initialDelay * attempt.
		delayF = float64(initialDelay) * float64(attempt)

	case BackoffConstant:
		// delay = initialDelay (constant).
		delayF = float64(initialDelay)

	default:
		delayF = float64(initialDelay)
	}

	// Guard against NaN/Inf and float -> Duration overflow.
	cap := float64(maxBackoffDuration)
	if math.IsNaN(delayF) || math.IsInf(delayF, 0) || delayF > cap || delayF < 0 {
		delayF = cap
	}
	delay := time.Duration(delayF)

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

	// math/rand/v2 has lock-free per-goroutine state, avoiding the global
	// mutex contention of math/rand under parallel callers. Predictability
	// is acceptable for retry jitter.
	jitterAmount := time.Duration(rand.Float64() * float64(delay) * 0.1)
	return delay + jitterAmount
}
