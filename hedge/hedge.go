// Package hedge provides hedged-request execution for tail-latency reduction.
//
// A hedged request fires the primary attempt immediately. If it has not
// returned within HedgeDelay, a second attempt is fired in parallel (and
// optionally a third, fourth, ...) up to MaxAttempts. The first successful
// result wins; remaining in-flight attempts are cancelled via context.
//
// Hedging trades extra work for lower tail latency. Use only on idempotent
// operations: under hedging, multiple attempts may run to completion before
// the cancellation propagates, so each attempt's side effects must be safe
// to repeat.
//
// Example usage:
//
//	h := hedge.New[*Response](hedge.Config{
//	    MaxAttempts: 3,
//	    HedgeDelay:  50 * time.Millisecond,
//	})
//
//	resp, err := h.Execute(ctx, func(ctx context.Context) (*Response, error) {
//	    return client.Do(ctx, req)
//	})
package hedge

import (
	"context"
	"log/slog"
	"runtime/debug"
	"time"
)

// Hedge is a generic interface for hedged-request execution.
type Hedge[T any] interface {
	// Execute runs the given function and, on tail latency, runs additional
	// parallel attempts up to MaxAttempts. The first successful result wins;
	// remaining attempts are cancelled. If all attempts fail, the first
	// error is returned.
	Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
}

// hedge is the concrete implementation of Hedge.
type hedge[T any] struct {
	config Config
}

// New creates a new Hedge instance with the given configuration.
//
//nolint:gocritic // hugeParam: Config passed by value for API consistency across all patterns
func New[T any](config Config) Hedge[T] {
	config.setDefaults()
	return &hedge[T]{config: config}
}

// result carries one attempt's outcome to the coordinator.
type result[T any] struct {
	value T
	err   error
}

// Execute implements the Hedge interface.
func (h *hedge[T]) Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
	var zero T

	// Check parent context up front.
	if err := ctx.Err(); err != nil {
		return zero, err
	}

	// All attempts share a derived context so we can cancel losers when a
	// winner returns.
	hedgeCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Buffered to MaxAttempts so completing attempts never block on a
	// goroutine leak after the winner returns and we drop the receiver.
	results := make(chan result[T], h.config.MaxAttempts)

	fire := func() {
		v, e := fn(hedgeCtx)
		select {
		case results <- result[T]{value: v, err: e}:
		case <-hedgeCtx.Done():
		}
	}

	// Fire primary immediately.
	go fire()
	fired := 1

	// Reusable timer; gates the next hedge fire. Stopped on winner.
	timer := time.NewTimer(h.config.HedgeDelay)
	defer timer.Stop()

	var firstErr error
	collected := 0

	for {
		select {
		case r := <-results:
			collected++
			if r.err == nil {
				// Winner: cancel siblings.
				return r.value, nil
			}
			if firstErr == nil {
				firstErr = r.err
			}
			// All fired attempts have now reported errors.
			if collected == fired && fired == h.config.MaxAttempts {
				return zero, firstErr
			}

		case <-timer.C:
			if fired < h.config.MaxAttempts {
				fired++
				go fire()
				h.logHedge(ctx, fired)
				if h.config.OnHedge != nil {
					attempt := fired
					h.safeCallback(func() { h.config.OnHedge(attempt) })
				}
				timer.Reset(h.config.HedgeDelay)
			}
			// else: all attempts fired; wait for results.

		case <-ctx.Done():
			return zero, ctx.Err()
		}
	}
}

// logHedge logs a hedge attempt firing using structured logging.
func (h *hedge[T]) logHedge(ctx context.Context, attempt int) {
	if h.config.Logger != nil {
		h.config.Logger.DebugContext(ctx, "hedge attempt fired",
			slog.Int("attempt", attempt),
			slog.Int("max_attempts", h.config.MaxAttempts),
			slog.Duration("hedge_delay", h.config.HedgeDelay),
		)
	}
}

// safeCallback executes a callback with panic recovery.
func (h *hedge[T]) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if h.config.Logger != nil {
				h.config.Logger.Error("hedge callback panic",
					slog.String("pattern", "hedge"),
					slog.Any("panic", r),
					slog.String("stack", string(debug.Stack())),
				)
			}
		}
	}()
	fn()
}
