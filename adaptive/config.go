package adaptive

import (
	"errors"
	"log/slog"
	"time"
)

// ErrLimitExceeded is returned when the current adaptive concurrency limit is
// reached. The limit may auto-recover over time as in-flight operations drain
// and successes accumulate.
var ErrLimitExceeded = errors.New("adaptive: concurrency limit exceeded")

// Algorithm selects the limit-tuning strategy.
type Algorithm int

const (
	// AlgorithmAIMD is the default: additive increase on success, multiplicative
	// decrease on failure. Predictable and lock-free; ignores latency.
	AlgorithmAIMD Algorithm = iota

	// AlgorithmVegas is RTT-aware. Tracks the minimum observed call latency
	// (no-load baseline) and the EMA of recent latencies; estimates the
	// "queue depth" induced by the current concurrency limit; raises the
	// limit when the queue is shallow (< VegasAlpha) and lowers when deep
	// (> VegasBeta). Reacts to rising latency BEFORE failures appear.
	//
	// Slightly more overhead than AIMD because each call records two atomic
	// time samples. Use when latency is your primary saturation signal.
	AlgorithmVegas

	// AlgorithmGradient2 is a smoothed gradient-of-RTT controller (Netflix
	// concurrency-limits naming). Computes a gradient g = minRTT / longEMA
	// (clamped to [0.5, 1.0]) and a target newLimit = max(MinLimit, g*cur + queue)
	// where queue is √currentLimit. Reacts more aggressively than Vegas
	// under variable load: small RTT inflation already shrinks the limit
	// proportionally, instead of waiting for a fixed water mark.
	//
	// Use Gradient2 when load is bursty and Vegas's α/β step adjustments
	// are too coarse. AIMD remains the default; Gradient2 has the highest
	// per-call overhead (two atomic time samples + one sqrt + one float
	// multiply).
	AlgorithmGradient2
)

const (
	defaultInitialLimit     = 10
	defaultMinLimit         = 1
	defaultMaxLimit         = 200
	defaultSuccessThreshold = 10
	defaultVegasAlpha       = 3
	defaultVegasBeta        = 6
)

// Config holds the configuration for an adaptive concurrency limiter.
type Config struct {
	// OnLimitChange is called when the limit changes, with the old and new
	// values. Useful for metrics. Wrapped with panic recovery; runs on the
	// caller's goroutine.
	OnLimitChange func(old, new int)

	// Logger is used for structured logging of limit changes. nil = silent.
	Logger *slog.Logger

	// Algorithm picks the tuning strategy. Defaults to AlgorithmAIMD.
	Algorithm Algorithm

	// InitialLimit is the starting concurrency cap. Defaults to 10.
	// Clamped into [MinLimit, MaxLimit].
	InitialLimit int

	// MinLimit is the floor that multiplicative decrease will not go below.
	// Defaults to 1.
	MinLimit int

	// MaxLimit is the ceiling that additive increase will not exceed.
	// Defaults to 200.
	MaxLimit int

	// SuccessThreshold (AIMD only): consecutive successes required before
	// the limit is incremented. Defaults to 10. Clamped to >= 1.
	SuccessThreshold int

	// VegasAlpha (Vegas only): low-water mark for queue-depth estimate.
	// When the estimated queue is below VegasAlpha, the limit grows by 1.
	// Defaults to 3.
	VegasAlpha int

	// VegasBeta (Vegas only): high-water mark for queue-depth estimate.
	// When the estimated queue exceeds VegasBeta, the limit shrinks by 1.
	// Must be > VegasAlpha. Defaults to 6.
	VegasBeta int

	// VegasMinSamples (Vegas only): minimum number of RTT samples observed
	// before adjustments fire. Defaults to 10. Prevents knee-jerk reactions
	// to a few unrepresentative samples on cold start.
	VegasMinSamples int

	// GradientMinSamples (Gradient2 only): minimum number of RTT samples
	// before the controller starts adjusting. Defaults to 10.
	GradientMinSamples int

	// GradientSmoothing (Gradient2 only): EMA coefficient for the long-term
	// RTT used as the saturation signal. Range (0, 1]; smaller = smoother.
	// Defaults to 0.2 (halflife ≈ 3 samples).
	GradientSmoothing float64
}

// _ documents the unused time import dependency for Vegas tuning constants.
var _ = time.Nanosecond

// setDefaults applies default values and clamps invalid configurations.
func (c *Config) setDefaults() {
	if c.MinLimit <= 0 {
		c.MinLimit = defaultMinLimit
	}
	if c.MaxLimit <= 0 {
		c.MaxLimit = defaultMaxLimit
	}
	if c.MaxLimit < c.MinLimit {
		c.MaxLimit = c.MinLimit
	}
	if c.InitialLimit <= 0 {
		c.InitialLimit = defaultInitialLimit
	}
	if c.InitialLimit < c.MinLimit {
		c.InitialLimit = c.MinLimit
	}
	if c.InitialLimit > c.MaxLimit {
		c.InitialLimit = c.MaxLimit
	}
	if c.SuccessThreshold <= 0 {
		c.SuccessThreshold = defaultSuccessThreshold
	}
	if c.VegasAlpha <= 0 {
		c.VegasAlpha = defaultVegasAlpha
	}
	if c.VegasBeta <= c.VegasAlpha {
		c.VegasBeta = c.VegasAlpha * 2
	}
	if c.VegasBeta <= 0 {
		c.VegasBeta = defaultVegasBeta
	}
	if c.VegasMinSamples <= 0 {
		c.VegasMinSamples = 10
	}
	if c.GradientMinSamples <= 0 {
		c.GradientMinSamples = 10
	}
	if c.GradientSmoothing <= 0 || c.GradientSmoothing > 1 {
		c.GradientSmoothing = 0.2
	}
}
