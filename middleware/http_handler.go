package middleware

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

// HTTPHandlerConfig configures the HTTPHandler preset.
type HTTPHandlerConfig struct {
	// Timeout caps each request's processing time. Required.
	Timeout time.Duration

	// RateLimit (optional) gates incoming requests. nil = no rate limiting.
	RateLimit ratelimit.RateLimiter

	// RateLimitKeyFn extracts the rate-limit key from the request. Required
	// when RateLimit is set; ignored otherwise.
	RateLimitKeyFn func(*http.Request) string

	// CBFailureThreshold is the consecutive-failure count that trips the
	// breaker Open. Defaults to 10.
	CBFailureThreshold uint32

	// CBOpenTimeout is how long the breaker stays Open before trial.
	// Defaults to 30s.
	CBOpenTimeout time.Duration
}

func (c *HTTPHandlerConfig) setDefaults() {
	if c.CBFailureThreshold == 0 {
		c.CBFailureThreshold = 10
	}
	if c.CBOpenTimeout <= 0 {
		c.CBOpenTimeout = 30 * time.Second
	}
}

// HTTPHandler returns an http.Handler that wraps `inner` with:
//
//	RateLimit (optional) → CircuitBreaker → Timeout → inner
//
// Status codes returned to the client:
//
//   - 429 Too Many Requests — rate limit exceeded
//   - 503 Service Unavailable — circuit breaker open
//   - 504 Gateway Timeout — request exceeded Timeout
//
// The breaker considers any 5xx response a failure (consistent with
// fortify/http.CircuitBreaker). Breaker state is shared across all requests
// served by this handler instance.
//
// Returns an error if Timeout is zero/negative or if RateLimit is non-nil
// but RateLimitKeyFn is nil.
func HTTPHandler(inner http.Handler, cfg HTTPHandlerConfig) (http.Handler, error) {
	if cfg.Timeout <= 0 {
		return nil, errors.New("middleware.HTTPHandler: Timeout must be positive")
	}
	if cfg.RateLimit != nil && cfg.RateLimitKeyFn == nil {
		return nil, errors.New("middleware.HTTPHandler: RateLimitKeyFn required when RateLimit set")
	}
	cfg.setDefaults()

	cb := circuitbreaker.New[int](circuitbreaker.Config{
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     cfg.CBOpenTimeout,
		ReadyToTrip: func(c circuitbreaker.Counts) bool {
			return c.ConsecutiveFailures >= cfg.CBFailureThreshold
		},
	})

	tm := timeout.New[int](timeout.Config{DefaultTimeout: cfg.Timeout})

	return &httpHandler{
		inner:   inner,
		cb:      cb,
		tm:      tm,
		timeout: cfg.Timeout,
		rl:      cfg.RateLimit,
		keyFn:   cfg.RateLimitKeyFn,
	}, nil
}

type httpHandler struct {
	inner   http.Handler
	cb      circuitbreaker.CircuitBreaker[int]
	tm      timeout.Timeout[int]
	timeout time.Duration
	rl      ratelimit.RateLimiter
	keyFn   func(*http.Request) string
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Rate limit gate (outermost).
	if h.rl != nil {
		if !h.rl.Allow(r.Context(), h.keyFn(r)) {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
	}

	rec := newServerRecorder(w)

	// Wrap with CB + Timeout. The "result" type is the recorded status.
	_, err := h.cb.Execute(r.Context(), func(ctx context.Context) (int, error) {
		_, terr := h.tm.Execute(ctx, h.timeout, func(tctx context.Context) (int, error) {
			h.inner.ServeHTTP(rec, r.WithContext(tctx))
			rec.mu.Lock()
			sc := rec.statusCode
			rec.mu.Unlock()
			if sc >= 500 {
				return sc, ferrors.ErrCircuitOpen
			}
			return sc, nil
		})
		return rec.code(), terr
	})

	if err == nil {
		return
	}

	rec.mu.Lock()
	already := rec.written
	rec.mu.Unlock()
	if already {
		// Downstream already wrote; we can't add a status. Just return.
		return
	}

	switch {
	case errors.Is(err, ferrors.ErrCircuitOpen):
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	case errors.Is(err, context.DeadlineExceeded):
		http.Error(w, "Gateway Timeout", http.StatusGatewayTimeout)
	default:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// serverRecorder mirrors http.responseRecorder used by the per-pattern HTTP
// middleware but lives here to avoid an import cycle and to keep this preset
// self-contained.
type serverRecorder struct {
	http.ResponseWriter
	statusCode int
	written    bool
	mu         sync.Mutex
}

func newServerRecorder(w http.ResponseWriter) *serverRecorder {
	return &serverRecorder{ResponseWriter: w, statusCode: http.StatusOK}
}

func (r *serverRecorder) WriteHeader(statusCode int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.written {
		return
	}
	r.statusCode = statusCode
	r.written = true
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *serverRecorder) Write(b []byte) (int, error) {
	r.mu.Lock()
	if !r.written {
		r.statusCode = http.StatusOK
		r.written = true
		r.ResponseWriter.WriteHeader(http.StatusOK)
	}
	r.mu.Unlock()
	return r.ResponseWriter.Write(b)
}

func (r *serverRecorder) code() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.statusCode
}
