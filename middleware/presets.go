package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/felixgeelhaar/fortify/retry"
	"github.com/felixgeelhaar/fortify/timeout"
)

// Presets are opinionated, configurable defaults for the most common
// resilience shapes. They exist so you don't have to re-derive the same
// chain on every call site. Prefer building your own Chain when the preset
// doesn't fit; presets are starting points, not silver bullets.

// HTTPClientConfig configures the HTTPClient preset.
type HTTPClientConfig struct {
	// Timeout caps individual call latency. Required (zero rejected).
	Timeout time.Duration

	// MaxRetries bounds total attempts including the first. Defaults to 3.
	MaxRetries int

	// RetryInitialDelay is the first backoff delay. Defaults to 100ms.
	RetryInitialDelay time.Duration

	// CBFailureThreshold is the consecutive-failure count that trips the
	// breaker Open. Defaults to 5.
	CBFailureThreshold uint32

	// CBOpenTimeout is how long the breaker stays Open before trial.
	// Defaults to 30s.
	CBOpenTimeout time.Duration
}

func (c *HTTPClientConfig) setDefaults() {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 3
	}
	if c.RetryInitialDelay <= 0 {
		c.RetryInitialDelay = 100 * time.Millisecond
	}
	if c.CBFailureThreshold == 0 {
		c.CBFailureThreshold = 5
	}
	if c.CBOpenTimeout <= 0 {
		c.CBOpenTimeout = 30 * time.Second
	}
}

// HTTPClient returns a preconfigured chain for outbound HTTP calls:
//
//	CircuitBreaker → Retry → Timeout → operation
//
// The breaker trips after CBFailureThreshold consecutive failures and stays
// Open for CBOpenTimeout. Retry uses exponential backoff with jitter and
// will not retry on context cancellation. Timeout caps each individual
// attempt (not the total chain).
//
// The breaker and retry are owned by the returned chain; do not share them
// across unrelated workloads.
//
// Returns an error if Timeout is zero or negative.
func HTTPClient(cfg HTTPClientConfig) (*Chain[*http.Response], error) {
	if cfg.Timeout <= 0 {
		return nil, errors.New("middleware.HTTPClient: Timeout must be positive")
	}
	cfg.setDefaults()

	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     cfg.CBOpenTimeout,
		ReadyToTrip: func(c circuitbreaker.Counts) bool {
			return c.ConsecutiveFailures >= cfg.CBFailureThreshold
		},
	})

	r := retry.New[*http.Response](retry.Config{
		MaxAttempts:   cfg.MaxRetries,
		InitialDelay:  cfg.RetryInitialDelay,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		Jitter:        true,
		IsRetryable: func(err error) bool {
			// Don't retry on circuit-open or context cancellation;
			// retry everything else.
			return !errors.Is(err, ferrors.ErrCircuitOpen) &&
				!errors.Is(err, context.Canceled) &&
				!errors.Is(err, context.DeadlineExceeded)
		},
	})

	tm := timeout.New[*http.Response](timeout.Config{
		DefaultTimeout: cfg.Timeout,
	})

	return New[*http.Response]().
		WithCircuitBreaker(cb).
		WithRetry(r).
		WithTimeout(tm, cfg.Timeout), nil
}

// DatabaseQueryConfig configures the DatabaseQuery preset.
type DatabaseQueryConfig struct {
	// QueryTimeout caps individual query latency. Required.
	QueryTimeout time.Duration

	// MaxRetries bounds total attempts. Defaults to 2 (one retry).
	// Database queries should retry sparingly; many errors are non-transient.
	MaxRetries int

	// CBFailureThreshold trips the breaker after this many consecutive
	// failures. Defaults to 10 (databases tolerate more before tripping
	// because failures are usually fatal, not transient).
	CBFailureThreshold uint32
}

func (c *DatabaseQueryConfig) setDefaults() {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 2
	}
	if c.CBFailureThreshold == 0 {
		c.CBFailureThreshold = 10
	}
}

// DatabaseQuery returns a chain tuned for database operations:
//
//	CircuitBreaker → Retry (rare, short backoff) → Timeout → operation
//
// Retry is conservative because most database errors are non-transient.
// The breaker trips later than HTTPClient because false-positive trips on
// query errors are particularly disruptive. Each attempt is bounded by
// QueryTimeout.
//
// The result type is `any` so the same chain can wrap heterogeneous query
// shapes; cast at the call site.
//
// Returns an error if QueryTimeout is zero or negative.
func DatabaseQuery(cfg DatabaseQueryConfig) (*Chain[any], error) {
	if cfg.QueryTimeout <= 0 {
		return nil, errors.New("middleware.DatabaseQuery: QueryTimeout must be positive")
	}
	cfg.setDefaults()

	cb := circuitbreaker.New[any](circuitbreaker.Config{
		MaxRequests: 2,
		Interval:    60 * time.Second,
		Timeout:     20 * time.Second,
		ReadyToTrip: func(c circuitbreaker.Counts) bool {
			return c.ConsecutiveFailures >= cfg.CBFailureThreshold
		},
	})

	r := retry.New[any](retry.Config{
		MaxAttempts:   cfg.MaxRetries,
		InitialDelay:  50 * time.Millisecond,
		MaxDelay:      500 * time.Millisecond,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		Jitter:        true,
	})

	tm := timeout.New[any](timeout.Config{
		DefaultTimeout: cfg.QueryTimeout,
	})

	return New[any]().
		WithCircuitBreaker(cb).
		WithRetry(r).
		WithTimeout(tm, cfg.QueryTimeout), nil
}

// RPCDownstreamConfig configures the RPCDownstream preset.
type RPCDownstreamConfig struct {
	// CallTimeout caps each RPC. Required.
	CallTimeout time.Duration

	// MaxRetries bounds total attempts. Defaults to 3.
	MaxRetries int

	// CBFailureThreshold trips the breaker. Defaults to 5.
	CBFailureThreshold uint32
}

func (c *RPCDownstreamConfig) setDefaults() {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 3
	}
	if c.CBFailureThreshold == 0 {
		c.CBFailureThreshold = 5
	}
}

// RPCDownstream returns a chain tuned for RPC calls to a single downstream
// service:
//
//	CircuitBreaker → Retry → Timeout → operation
//
// Use one chain per downstream so the breaker and retry state are scoped
// correctly. Sharing a single chain across multiple downstreams trips
// the breaker on aggregate failures, hiding which downstream is unhealthy.
//
// Returns an error if CallTimeout is zero or negative.
func RPCDownstream[T any](cfg RPCDownstreamConfig) (*Chain[T], error) {
	if cfg.CallTimeout <= 0 {
		return nil, errors.New("middleware.RPCDownstream: CallTimeout must be positive")
	}
	cfg.setDefaults()

	cb := circuitbreaker.New[T](circuitbreaker.Config{
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(c circuitbreaker.Counts) bool {
			return c.ConsecutiveFailures >= cfg.CBFailureThreshold
		},
	})

	r := retry.New[T](retry.Config{
		MaxAttempts:   cfg.MaxRetries,
		InitialDelay:  100 * time.Millisecond,
		MaxDelay:      2 * time.Second,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		Jitter:        true,
	})

	tm := timeout.New[T](timeout.Config{
		DefaultTimeout: cfg.CallTimeout,
	})

	return New[T]().
		WithCircuitBreaker(cb).
		WithRetry(r).
		WithTimeout(tm, cfg.CallTimeout), nil
}
