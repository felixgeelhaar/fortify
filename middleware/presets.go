package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/felixgeelhaar/fortify/budget"
	"github.com/felixgeelhaar/fortify/bulkhead"
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

// LLMCallConfig configures the LLMCall preset. The preset is tuned for
// outbound calls to an LLM provider (OpenAI, Anthropic, generic
// OpenAI-compatible endpoints). It is scoped to one provider+model pair;
// build one chain per provider/model combination so the breaker and
// concurrency limits are scoped correctly.
//
// Sensitive payloads: this preset never inspects the function's argument
// or result. The Charge callback you supply runs against the result, so
// keep prompt content and completion text out of any error fields you
// derive there. See docs/PRODUCTION.md ("Observability and sensitive
// payloads").
type LLMCallConfig[T any] struct {
	// Provider is a short identifier for the upstream (e.g. "openai",
	// "anthropic", "ollama"). Used as a label for logging only by the
	// caller; included here so a future metrics integration can pick it
	// up automatically.
	Provider string

	// Model is the specific model identifier (e.g. "gpt-5",
	// "claude-opus-4-7"). Combined with Provider it scopes the chain.
	Model string

	// CallTimeout caps the total wall-clock duration of a single
	// non-streaming attempt. Required (zero rejected). For streaming
	// callers, use streamtimeout outside this preset.
	CallTimeout time.Duration

	// MaxRetries bounds total attempts including the first. Defaults to 3.
	MaxRetries int

	// RetryInitialDelay is the first backoff delay. Defaults to 250ms,
	// which is conservative enough that a brief provider hiccup does not
	// double-spend tokens immediately.
	RetryInitialDelay time.Duration

	// AssumeIdempotent controls whether retries are attempted on errors
	// that may have produced a partial server-side effect. LLM calls are
	// generally NOT idempotent (the same prompt yields different
	// completions, and a partial completion already cost tokens), so
	// the default is false: retries fire only when the failure is known
	// to be transport-level (RetryableError-marked), or matches one of
	// the explicit retryable sentinels.
	AssumeIdempotent bool

	// CBFailureThreshold is the consecutive-failure count that trips
	// the breaker Open. Defaults to 5.
	CBFailureThreshold uint32

	// CBOpenTimeout is how long the breaker stays Open before trial.
	// Defaults to 30s.
	CBOpenTimeout time.Duration

	// MaxConcurrent caps in-flight calls to this provider/model from
	// this process. Defaults to 64. Set to a small number for paid APIs
	// to limit the blast radius of a runaway agent loop.
	MaxConcurrent int

	// Budget configures the cost ceiling. At least one of
	// Budget.Max.Tokens, Budget.Max.USDMicros, or Budget.Max.Calls must
	// be positive. The ceiling is shared across all retries of a single
	// chain instance.
	Budget BudgetConfig[T]
}

// BudgetConfig is the LLMCall-facing wrapper around budget.Config so
// the preset can supply consistent defaults and force-on slog.LogValuer
// without exposing the raw budget Config to callers who only care about
// the ceiling.
type BudgetConfig[T any] struct {
	// Max is the budget ceiling. Required (at least one positive field).
	Max budget.Cost
	// Charge converts a successful or failed call into a Cost.
	// May be nil if the only enforced dimension is Calls.
	Charge budget.Charge[T]
	// OnExceeded fires once on the first breach.
	OnExceeded func(consumed budget.Cost)
}

func (c *LLMCallConfig[T]) setDefaults() {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 3
	}
	if c.RetryInitialDelay <= 0 {
		c.RetryInitialDelay = 250 * time.Millisecond
	}
	if c.CBFailureThreshold == 0 {
		c.CBFailureThreshold = 5
	}
	if c.CBOpenTimeout <= 0 {
		c.CBOpenTimeout = 30 * time.Second
	}
	if c.MaxConcurrent <= 0 {
		c.MaxConcurrent = 64
	}
}

// LLMCall returns a chain tuned for outbound LLM provider calls:
//
//	Bulkhead → CircuitBreaker → Retry → Budget → Timeout → operation
//
// The budget sits inside Retry so every attempt is charged, capping the
// total cost of a retry storm during a provider incident. The budget's
// *BudgetExceededError is non-retryable: once the ceiling is hit, the
// chain returns the error verbatim instead of attempting another round.
//
// Defaults assume non-idempotent semantics for LLM calls (see
// AssumeIdempotent docs). Override IsRetryable indirectly via
// AssumeIdempotent=true to match retryablehttp/HTTPClient behaviour.
//
// Returns an error if CallTimeout is zero or negative, or if the Budget
// is misconfigured.
func LLMCall[T any](cfg LLMCallConfig[T]) (*Chain[T], error) {
	if cfg.CallTimeout <= 0 {
		return nil, errors.New("middleware.LLMCall: CallTimeout must be positive")
	}
	cfg.setDefaults()

	b, err := budget.New[T](budget.Config[T]{
		Max:        cfg.Budget.Max,
		Charge:     cfg.Budget.Charge,
		OnExceeded: cfg.Budget.OnExceeded,
	})
	if err != nil {
		return nil, err
	}

	bh := bulkhead.New[T](bulkhead.Config{
		MaxConcurrent: cfg.MaxConcurrent,
	})

	cb := circuitbreaker.New[T](circuitbreaker.Config{
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     cfg.CBOpenTimeout,
		ReadyToTrip: func(c circuitbreaker.Counts) bool {
			return c.ConsecutiveFailures >= cfg.CBFailureThreshold
		},
	})

	isRetryable := llmIsRetryable(cfg.AssumeIdempotent)
	r := retry.New[T](retry.Config{
		MaxAttempts:   cfg.MaxRetries,
		InitialDelay:  cfg.RetryInitialDelay,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		Jitter:        true,
		IsRetryable:   isRetryable,
	})

	tm := timeout.New[T](timeout.Config{
		DefaultTimeout: cfg.CallTimeout,
	})

	return New[T]().
		WithBulkhead(bh).
		WithCircuitBreaker(cb).
		WithRetry(r).
		WithBudget(b).
		WithTimeout(tm, cfg.CallTimeout), nil
}

// llmIsRetryable returns the retry predicate used by LLMCall. When
// assumeIdempotent is false (the default), retries only fire on errors
// explicitly marked retryable via ferrors.AsRetryable, plus rate-limit
// errors that carry a Retry-After hint. Non-retryable in any mode:
// budget exceeded, circuit open, context cancelled, context deadline.
func llmIsRetryable(assumeIdempotent bool) func(error) bool {
	return func(err error) bool {
		if err == nil {
			return false
		}
		// Hard non-retryables.
		if errors.Is(err, budget.ErrBudgetExceeded) {
			return false
		}
		if errors.Is(err, ferrors.ErrCircuitOpen) {
			return false
		}
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return false
		}
		// Always retry rate-limit responses; the upstream told us when
		// to come back.
		if errors.Is(err, ferrors.ErrRateLimitExceeded) {
			return true
		}
		// Caller-marked retryable wins regardless of idempotency setting.
		if ferrors.IsRetryable(err) {
			return true
		}
		// Idempotent mode: be permissive (matches HTTPClient defaults).
		// Non-idempotent (default): only the explicit signals above retry.
		return assumeIdempotent
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
