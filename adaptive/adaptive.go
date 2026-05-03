// Package adaptive provides an auto-tuning concurrency limiter that adjusts
// its in-flight cap in response to observed success and failure outcomes.
//
// Unlike a static bulkhead (which has a fixed concurrency cap), the adaptive
// limiter starts at InitialLimit and follows an AIMD policy:
//
//   - On every SuccessThreshold consecutive successes, the limit increases
//     by one (additive increase), up to MaxLimit.
//   - On any failure, the limit halves (multiplicative decrease), down to
//     MinLimit.
//
// AIMD is well-understood from TCP congestion control. It probes for higher
// throughput when the downstream is healthy and backs off quickly under
// pressure. It is not as sophisticated as RTT-based algorithms (e.g.,
// Vegas, Gradient2), but it is predictable and requires no clock-domain
// coordination.
//
// Use this when:
//
//   - You don't know the right concurrency cap up front and want it
//     discovered at runtime.
//   - Downstream capacity changes over time (autoscaling, regional failover).
//   - You want fast back-off during incidents without manual intervention.
//
// Use a static bulkhead instead when:
//
//   - You have a hard external constraint (DB connection pool size, license
//     count) — overshoot is unacceptable.
//   - Failures are not a reliable signal of overload (e.g., business-logic
//     errors that have nothing to do with capacity).
//
// Example usage:
//
//	a := adaptive.New[*Response](adaptive.Config{
//	    InitialLimit:     10,
//	    MinLimit:         2,
//	    MaxLimit:         100,
//	    SuccessThreshold: 20,
//	})
//
//	resp, err := a.Execute(ctx, func(ctx context.Context) (*Response, error) {
//	    return downstream.Call(ctx)
//	})
package adaptive

import (
	"context"
	"log/slog"
	"math"
	"runtime/debug"
	"sync/atomic"
	"time"
)

// Limiter is a generic adaptive concurrency limiter.
type Limiter[T any] interface {
	// Execute runs the given function if the current concurrency limit
	// allows. Returns ErrLimitExceeded if the limit is reached.
	Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)

	// Limit returns the current concurrency cap.
	Limit() int

	// InFlight returns the number of operations currently in flight.
	InFlight() int
}

// limiter is the concrete implementation of Limiter.
type limiter[T any] struct {
	config    Config
	limit     atomic.Int32
	inFlight  atomic.Int32
	successes atomic.Int32

	// Vegas-only RTT tracking. Stored as nanoseconds. Zero indicates "not
	// yet observed".
	minRTT     atomic.Int64
	emaRTT     atomic.Int64 // EMA of recent RTTs, halflife ≈ 8 samples
	rttSamples atomic.Int64
}

// New creates a new adaptive concurrency limiter.
//
//nolint:gocritic // hugeParam: Config passed by value for API consistency across all patterns
func New[T any](config Config) Limiter[T] {
	config.setDefaults()
	l := &limiter[T]{config: config}
	l.limit.Store(int32(config.InitialLimit))
	return l
}

// Execute implements the Limiter interface.
func (l *limiter[T]) Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
	var zero T

	if err := ctx.Err(); err != nil {
		return zero, err
	}

	if !l.tryAcquire() {
		return zero, ErrLimitExceeded
	}
	defer l.inFlight.Add(-1)

	// Time the call only when an RTT-aware algorithm is active.
	var start time.Time
	rttAware := l.config.Algorithm == AlgorithmVegas || l.config.Algorithm == AlgorithmGradient2
	if rttAware {
		start = time.Now()
	}

	result, err := fn(ctx)
	if err != nil {
		l.onFailure(ctx)
		return result, err
	}

	switch l.config.Algorithm {
	case AlgorithmVegas:
		l.vegasOnSuccess(ctx, time.Since(start))
	case AlgorithmGradient2:
		l.gradient2OnSuccess(ctx, time.Since(start))
	default:
		l.aimdOnSuccess(ctx)
	}
	return result, nil
}

// Limit implements the Limiter interface.
func (l *limiter[T]) Limit() int { return int(l.limit.Load()) }

// InFlight implements the Limiter interface.
func (l *limiter[T]) InFlight() int { return int(l.inFlight.Load()) }

// tryAcquire atomically increments the in-flight counter if doing so would
// not exceed the current limit. Returns false if at capacity.
func (l *limiter[T]) tryAcquire() bool {
	for {
		cur := l.inFlight.Load()
		lim := l.limit.Load()
		if cur >= lim {
			return false
		}
		if l.inFlight.CompareAndSwap(cur, cur+1) {
			return true
		}
		// CAS lost; retry.
	}
}

// aimdOnSuccess advances the success counter; on reaching SuccessThreshold,
// performs an additive increase of the limit (up to MaxLimit).
func (l *limiter[T]) aimdOnSuccess(ctx context.Context) {
	s := l.successes.Add(1)
	if s < int32(l.config.SuccessThreshold) {
		return
	}
	// Reset the success counter; raise the limit by one if room remains.
	l.successes.Store(0)
	for {
		cur := l.limit.Load()
		if cur >= int32(l.config.MaxLimit) {
			return
		}
		if l.limit.CompareAndSwap(cur, cur+1) {
			l.notify(ctx, int(cur), int(cur+1))
			return
		}
	}
}

// vegasOnSuccess records a latency sample, updates baseline + EMA, and
// applies a queue-depth-based limit adjustment.
//
// Algorithm:
//
//	expected = currentLimit * minRTT / observedRTT
//	queue    = currentLimit - expected
//	if queue < VegasAlpha → limit++
//	if queue > VegasBeta  → limit--
//
// In words: if the current RTT is close to the no-load minimum (queue is
// shallow), we have headroom; raise the limit. If RTT has stretched (queue
// is deep), we are inducing wait time downstream; lower the limit.
func (l *limiter[T]) vegasOnSuccess(ctx context.Context, rtt time.Duration) {
	rttNs := rtt.Nanoseconds()
	if rttNs <= 0 {
		return
	}

	// Update baseline minRTT via CAS retry loop.
	for {
		cur := l.minRTT.Load()
		if cur != 0 && rttNs >= cur {
			break
		}
		if l.minRTT.CompareAndSwap(cur, rttNs) {
			break
		}
	}

	// Update EMA RTT. Halflife ≈ 8 samples (alpha = 1/8).
	for {
		cur := l.emaRTT.Load()
		var next int64
		if cur == 0 {
			next = rttNs
		} else {
			// next = cur*7/8 + rttNs/8
			next = (cur*7 + rttNs) / 8
		}
		if l.emaRTT.CompareAndSwap(cur, next) {
			break
		}
	}

	samples := l.rttSamples.Add(1)
	if samples < int64(l.config.VegasMinSamples) {
		return
	}

	minRTT := l.minRTT.Load()
	emaRTT := l.emaRTT.Load()
	if minRTT == 0 || emaRTT == 0 {
		return
	}

	cur := l.limit.Load()
	// expected = cur * minRTT / emaRTT
	// queue    = cur - expected = cur * (1 - minRTT/emaRTT) = cur * (emaRTT - minRTT) / emaRTT
	queue := int32(int64(cur) * (emaRTT - minRTT) / emaRTT)

	switch {
	case queue < int32(l.config.VegasAlpha):
		// Underutilized: try a step up.
		if cur < int32(l.config.MaxLimit) {
			if l.limit.CompareAndSwap(cur, cur+1) {
				l.notify(ctx, int(cur), int(cur+1))
			}
		}
	case queue > int32(l.config.VegasBeta):
		// Saturating: step down. Use additive decrease, not halving;
		// Vegas reacts earlier than failure-driven AIMD so it can afford
		// gentler corrections.
		if cur > int32(l.config.MinLimit) {
			if l.limit.CompareAndSwap(cur, cur-1) {
				l.notify(ctx, int(cur), int(cur-1))
			}
		}
	}
}

// gradient2OnSuccess records a sample, updates the long EMA, and adjusts
// the limit according to the smoothed RTT gradient.
//
// gradient = clamp(minRTT / longEMA, 0.5, 1.0)
// queueSize = sqrt(currentLimit)
// newLimit = floor(currentLimit * gradient + queueSize)
//
// When RTT is at the no-load minimum (gradient ≈ 1) the limit grows by
// √limit per sample. As RTT inflates the gradient shrinks proportionally,
// pulling the limit down. Clamping to 0.5 prevents overcorrection on
// outliers; clamping to 1.0 prevents overshoot.
func (l *limiter[T]) gradient2OnSuccess(ctx context.Context, rtt time.Duration) {
	rttNs := rtt.Nanoseconds()
	if rttNs <= 0 {
		return
	}

	// minRTT update (CAS).
	for {
		cur := l.minRTT.Load()
		if cur != 0 && rttNs >= cur {
			break
		}
		if l.minRTT.CompareAndSwap(cur, rttNs) {
			break
		}
	}

	// Long EMA update with configurable smoothing α.
	alpha := l.config.GradientSmoothing
	for {
		cur := l.emaRTT.Load()
		var next int64
		if cur == 0 {
			next = rttNs
		} else {
			next = int64(float64(cur)*(1-alpha) + float64(rttNs)*alpha)
		}
		if l.emaRTT.CompareAndSwap(cur, next) {
			break
		}
	}

	samples := l.rttSamples.Add(1)
	if samples < int64(l.config.GradientMinSamples) {
		return
	}

	minRTT := l.minRTT.Load()
	emaRTT := l.emaRTT.Load()
	if minRTT == 0 || emaRTT == 0 {
		return
	}

	gradient := float64(minRTT) / float64(emaRTT)
	if gradient < 0.5 {
		gradient = 0.5
	}
	if gradient > 1.0 {
		gradient = 1.0
	}

	cur := l.limit.Load()
	queueSize := math.Sqrt(float64(cur))
	target := int32(math.Floor(float64(cur)*gradient + queueSize))

	if target < int32(l.config.MinLimit) {
		target = int32(l.config.MinLimit)
	}
	if target > int32(l.config.MaxLimit) {
		target = int32(l.config.MaxLimit)
	}
	if target == cur {
		return
	}
	if l.limit.CompareAndSwap(cur, target) {
		l.notify(ctx, int(cur), int(target))
	}
}

// onFailure performs a multiplicative decrease (halving) of the limit,
// bounded below by MinLimit.
func (l *limiter[T]) onFailure(ctx context.Context) {
	l.successes.Store(0)
	for {
		cur := l.limit.Load()
		next := cur / 2
		if next < int32(l.config.MinLimit) {
			next = int32(l.config.MinLimit)
		}
		if next == cur {
			return
		}
		if l.limit.CompareAndSwap(cur, next) {
			l.notify(ctx, int(cur), int(next))
			return
		}
	}
}

// notify logs and fires OnLimitChange.
func (l *limiter[T]) notify(ctx context.Context, oldLimit, newLimit int) {
	if l.config.Logger != nil {
		l.config.Logger.InfoContext(ctx, "adaptive limit changed",
			slog.Int("old", oldLimit),
			slog.Int("new", newLimit),
		)
	}
	if l.config.OnLimitChange != nil {
		l.safeCallback(func() { l.config.OnLimitChange(oldLimit, newLimit) })
	}
}

// safeCallback executes a callback with panic recovery.
func (l *limiter[T]) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if l.config.Logger != nil {
				l.config.Logger.Error("adaptive callback panic",
					slog.String("pattern", "adaptive"),
					slog.Any("panic", r),
					slog.String("stack", string(debug.Stack())),
				)
			}
		}
	}()
	fn()
}
