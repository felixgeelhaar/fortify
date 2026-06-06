// Package streamtimeout enforces three independent deadlines on a
// streaming operation: time-to-first-byte, idle-between-chunks, and
// total wall-clock duration.
//
// The standard timeout primitive collapses these dimensions into a
// single deadline, which fits unary calls but breaks streaming work:
// a long-but-healthy completion looks identical to a stalled one. A
// streaming LLM response, an SSE feed, or a chunked HTTP body needs
// the three timers tracked separately.
//
// The caller's function receives a Mark callback. Calling Mark()
// resets the idle timer and, on its first invocation, satisfies the
// FirstByteTimeout. The caller is expected to invoke Mark on every
// inbound chunk; failure to do so is observationally identical to a
// stall and will cause the idle timer to fire.
//
// Example with an io.Reader-style stream:
//
//	st := streamtimeout.New[Response](streamtimeout.Config{
//	    FirstByteTimeout: 5 * time.Second,
//	    IdleTimeout:      2 * time.Second,
//	    TotalTimeout:     60 * time.Second,
//	})
//	out, err := st.Execute(ctx, func(ctx context.Context, mark streamtimeout.Mark) (Response, error) {
//	    for {
//	        chunk, err := r.ReadChunk(ctx)
//	        mark()
//	        if err != nil { ... }
//	    }
//	})
package streamtimeout

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"go.klarlabs.de/fortify/ferrors"
)

// Stage names the deadline that fired in a *StreamTimeoutError.
type Stage string

const (
	// StageFirstByte indicates the FirstByteTimeout fired before Mark
	// was called for the first time.
	StageFirstByte Stage = "first_byte"
	// StageIdle indicates the IdleTimeout fired between chunks.
	StageIdle Stage = "idle"
	// StageTotal indicates the TotalTimeout fired regardless of activity.
	StageTotal Stage = "total"
)

// Mark is the callback the operation invokes on each inbound chunk.
// The first call satisfies FirstByteTimeout; every call resets the
// idle timer.
type Mark func()

// Config configures a StreamTimeout. Only positive durations are
// enforced; zero disables the corresponding watchdog. At least one
// must be positive.
type Config struct {
	// FirstByteTimeout caps the time from Execute start until the
	// caller's first Mark invocation.
	FirstByteTimeout time.Duration
	// IdleTimeout caps the gap between consecutive Mark invocations.
	IdleTimeout time.Duration
	// TotalTimeout caps the wall-clock duration of the entire call.
	TotalTimeout time.Duration

	// Logger receives diagnostic warnings. Nil disables logging.
	Logger *slog.Logger
}

// StreamTimeout is the runtime instance produced by New.
type StreamTimeout[T any] struct {
	cfg Config
}

// New constructs a StreamTimeout. Returns an error if no deadline is
// configured (a zero-config instance would be a no-op).
func New[T any](cfg Config) (*StreamTimeout[T], error) {
	if cfg.FirstByteTimeout <= 0 && cfg.IdleTimeout <= 0 && cfg.TotalTimeout <= 0 {
		return nil, errors.New("streamtimeout.New: at least one of FirstByteTimeout, IdleTimeout, TotalTimeout must be positive")
	}
	return &StreamTimeout[T]{cfg: cfg}, nil
}

// Execute runs fn with a Mark callback wired to the configured
// deadlines. If any deadline fires, the context passed to fn is
// cancelled and Execute returns *StreamTimeoutError. The error
// composes with errors.Is(err, ferrors.ErrTimeout) and
// errors.Is(err, context.DeadlineExceeded) for callers that match
// generic timeout sentinels.
//
// The caller's fn MUST respect ctx cancellation; otherwise the
// timeout is advisory and Execute will block until fn returns.
func (s *StreamTimeout[T]) Execute(
	ctx context.Context,
	fn func(ctx context.Context, mark Mark) (T, error),
) (T, error) {
	var zero T

	parent := ctx
	cctx, cancel := context.WithCancel(parent)
	defer cancel()

	var (
		mu       sync.Mutex
		firedAt  *time.Time
		stage    Stage
		seenMark atomic.Bool
		startAt  = time.Now()
	)

	// Each watchdog records the first stage to fire and cancels the
	// context. Subsequent watchdogs are no-ops.
	fire := func(st Stage) {
		mu.Lock()
		defer mu.Unlock()
		if firedAt != nil {
			return
		}
		now := time.Now()
		firedAt = &now
		stage = st
		cancel()
	}

	// Channel signalling first-mark, used to cancel the first-byte
	// watchdog cleanly.
	firstMark := make(chan struct{})
	var firstMarkOnce sync.Once

	// Idle ticker is only started after first mark so it doesn't race
	// the first-byte watchdog.
	var idleTimer *time.Timer
	var idleStop chan struct{}

	mark := Mark(func() {
		if seenMark.CompareAndSwap(false, true) {
			firstMarkOnce.Do(func() { close(firstMark) })
			if s.cfg.IdleTimeout > 0 {
				idleStop = make(chan struct{})
				idleTimer = time.NewTimer(s.cfg.IdleTimeout)
				go func(stop chan struct{}, t *time.Timer) {
					for {
						select {
						case <-stop:
							if !t.Stop() {
								select {
								case <-t.C:
								default:
								}
							}
							return
						case <-cctx.Done():
							return
						case <-t.C:
							fire(StageIdle)
							return
						}
					}
				}(idleStop, idleTimer)
			}
			return
		}
		// Subsequent marks reset idle.
		if idleTimer != nil {
			if !idleTimer.Stop() {
				select {
				case <-idleTimer.C:
				default:
				}
			}
			idleTimer.Reset(s.cfg.IdleTimeout)
		}
	})

	// First-byte watchdog.
	if s.cfg.FirstByteTimeout > 0 {
		fbTimer := time.NewTimer(s.cfg.FirstByteTimeout)
		go func() {
			defer fbTimer.Stop()
			select {
			case <-firstMark:
				return
			case <-cctx.Done():
				return
			case <-fbTimer.C:
				fire(StageFirstByte)
			}
		}()
	}

	// Total watchdog.
	if s.cfg.TotalTimeout > 0 {
		totalTimer := time.NewTimer(s.cfg.TotalTimeout)
		go func() {
			defer totalTimer.Stop()
			select {
			case <-cctx.Done():
				return
			case <-totalTimer.C:
				fire(StageTotal)
			}
		}()
	}

	result, err := fn(cctx, mark)

	// Stop idle goroutine if running.
	if idleStop != nil {
		close(idleStop)
	}

	mu.Lock()
	st := stage
	fired := firedAt != nil
	mu.Unlock()

	if fired {
		// If the parent context cancelled first, propagate that instead.
		if pErr := parent.Err(); pErr != nil {
			return zero, pErr
		}
		s.logTimeout(st, time.Since(startAt))
		return zero, &StreamTimeoutError{
			Stage:            st,
			FirstByteTimeout: s.cfg.FirstByteTimeout,
			IdleTimeout:      s.cfg.IdleTimeout,
			TotalTimeout:     s.cfg.TotalTimeout,
			Elapsed:          time.Since(startAt),
		}
	}

	return result, err
}

func (s *StreamTimeout[T]) logTimeout(stage Stage, elapsed time.Duration) {
	if s.cfg.Logger == nil {
		return
	}
	s.cfg.Logger.Warn("streaming operation timeout",
		slog.String("pattern", "streamtimeout"),
		slog.String("stage", string(stage)),
		slog.Duration("elapsed", elapsed),
	)
}

// StreamTimeoutError reports which of the three deadlines fired and
// the configured limits at the moment of failure. errors.Is(err,
// ferrors.ErrTimeout) and errors.Is(err, context.DeadlineExceeded)
// both keep matching for callers that handle generic timeouts.
type StreamTimeoutError struct {
	// Stage is the deadline that fired.
	Stage Stage
	// FirstByteTimeout, IdleTimeout, TotalTimeout mirror Config.
	FirstByteTimeout time.Duration
	IdleTimeout      time.Duration
	TotalTimeout     time.Duration
	// Elapsed is the wall-clock time from Execute start to fire.
	Elapsed time.Duration
}

// Error implements error.
func (e *StreamTimeoutError) Error() string {
	switch e.Stage {
	case StageFirstByte:
		return fmt.Sprintf("stream first-byte timeout after %s (limit %s)", e.Elapsed, e.FirstByteTimeout)
	case StageIdle:
		return fmt.Sprintf("stream idle timeout (no chunk in %s)", e.IdleTimeout)
	case StageTotal:
		return fmt.Sprintf("stream total timeout after %s (limit %s)", e.Elapsed, e.TotalTimeout)
	default:
		return fmt.Sprintf("stream timeout (stage=%s)", e.Stage)
	}
}

// Unwrap exposes the generic timeout sentinel and
// context.DeadlineExceeded so existing callers that match either keep
// working.
func (e *StreamTimeoutError) Unwrap() []error {
	return []error{ferrors.ErrTimeout, context.DeadlineExceeded}
}

// LogValue implements slog.LogValuer.
func (e *StreamTimeoutError) LogValue() slog.Value {
	if e == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("error", "stream_timeout"),
		slog.String("stage", string(e.Stage)),
		slog.Duration("elapsed", e.Elapsed),
		slog.Duration("first_byte_timeout", e.FirstByteTimeout),
		slog.Duration("idle_timeout", e.IdleTimeout),
		slog.Duration("total_timeout", e.TotalTimeout),
	)
}
