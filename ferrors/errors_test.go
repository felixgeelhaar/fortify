package ferrors

import (
	"errors"
	"log/slog"
	"testing"
	"time"
)

func TestSentinelErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "ErrCircuitOpen",
			err:  ErrCircuitOpen,
			want: "circuit breaker is open",
		},
		{
			name: "ErrRateLimitExceeded",
			err:  ErrRateLimitExceeded,
			want: "rate limit exceeded",
		},
		{
			name: "ErrBulkheadFull",
			err:  ErrBulkheadFull,
			want: "bulkhead at capacity",
		},
		{
			name: "ErrTimeout",
			err:  ErrTimeout,
			want: "operation timeout",
		},
		{
			name: "ErrMaxAttemptsReached",
			err:  ErrMaxAttemptsReached,
			want: "max retry attempts reached",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("error message = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetryableError(t *testing.T) {
	baseErr := errors.New("network timeout")
	retryable := AsRetryable(baseErr)

	t.Run("implements RetryableError interface", func(t *testing.T) {
		var re RetryableError
		if !errors.As(retryable, &re) {
			t.Error("AsRetryable() did not create a RetryableError")
		}
	})

	t.Run("Retryable returns true", func(t *testing.T) {
		var re RetryableError
		if errors.As(retryable, &re) {
			if !re.Retryable() {
				t.Error("Retryable() = false, want true")
			}
		}
	})

	t.Run("preserves original error message", func(t *testing.T) {
		if retryable.Error() != baseErr.Error() {
			t.Errorf("Error() = %v, want %v", retryable.Error(), baseErr.Error())
		}
	})

	t.Run("unwraps to original error", func(t *testing.T) {
		if !errors.Is(retryable, baseErr) {
			t.Error("retryable error does not unwrap to original error")
		}
	})
}

func TestIsRetryable(t *testing.T) {
	tests := []struct {
		err  error
		name string
		want bool
	}{
		{
			name: "retryable error returns true",
			err:  AsRetryable(errors.New("test error")),
			want: true,
		},
		{
			name: "non-retryable error returns false",
			err:  errors.New("test error"),
			want: false,
		},
		{
			name: "nil error returns false",
			err:  nil,
			want: false,
		},
		{
			name: "wrapped retryable error returns true",
			err:  errors.Join(AsRetryable(errors.New("inner")), errors.New("outer")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRetryable(tt.err); got != tt.want {
				t.Errorf("IsRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsRetryableWithNil(t *testing.T) {
	// Test that AsRetryable handles nil gracefully
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("AsRetryable(nil) panicked: %v", r)
		}
	}()

	retryable := AsRetryable(nil)
	if retryable == nil {
		t.Error("AsRetryable(nil) returned nil, expected wrapped error")
	}
}

func TestRetryableErrorUnwrap(t *testing.T) {
	baseErr := errors.New("base error")
	retryable := AsRetryable(baseErr)

	unwrapped := errors.Unwrap(retryable)
	if unwrapped != baseErr {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, baseErr)
	}
}

func TestCircuitOpenError_LogValue(t *testing.T) {
	t.Run("populated", func(t *testing.T) {
		e := &CircuitOpenError{
			Name:                 "downstream",
			State:                "open",
			RetryAfter:           250 * time.Millisecond,
			TotalRequests:        7,
			TotalFailures:        5,
			ConsecutiveFailures:  3,
			ConsecutiveSuccesses: 0,
		}
		v := e.LogValue()
		if v.Kind() != slog.KindGroup {
			t.Fatalf("LogValue kind = %v, want Group", v.Kind())
		}
		want := map[string]any{
			"error":                 "circuit_open",
			"name":                  "downstream",
			"state":                 "open",
			"retry_after":           250 * time.Millisecond,
			"total_requests":        uint64(7),
			"total_failures":        uint64(5),
			"consecutive_failures":  uint64(3),
			"consecutive_successes": uint64(0),
		}
		assertGroupAttrs(t, v, want)
	})

	t.Run("nil receiver returns empty group", func(t *testing.T) {
		var e *CircuitOpenError
		v := e.LogValue()
		if v.Kind() != slog.KindGroup {
			t.Fatalf("LogValue kind = %v, want Group", v.Kind())
		}
		if got := v.Group(); len(got) != 0 {
			t.Errorf("nil LogValue group = %v, want empty", got)
		}
	})

	t.Run("omits empty optional fields", func(t *testing.T) {
		e := &CircuitOpenError{}
		v := e.LogValue()
		got := groupAttrMap(v)
		if _, ok := got["name"]; ok {
			t.Errorf("expected no 'name' attr, got %v", got["name"])
		}
		if _, ok := got["state"]; ok {
			t.Errorf("expected no 'state' attr, got %v", got["state"])
		}
		if _, ok := got["retry_after"]; ok {
			t.Errorf("expected no 'retry_after' attr, got %v", got["retry_after"])
		}
	})
}

func TestRateLimitError_LogValue(t *testing.T) {
	e := &RateLimitError{Key: "user-42", RetryAfter: time.Second}
	v := e.LogValue()
	assertGroupAttrs(t, v, map[string]any{
		"error":       "rate_limit_exceeded",
		"key":         "user-42",
		"retry_after": time.Second,
	})

	var nilErr *RateLimitError
	if got := nilErr.LogValue(); got.Kind() != slog.KindGroup || len(got.Group()) != 0 {
		t.Errorf("nil LogValue = %v, want empty group", got)
	}
}

func TestTimeoutError_LogValue(t *testing.T) {
	e := &TimeoutError{Timeout: 2 * time.Second}
	v := e.LogValue()
	assertGroupAttrs(t, v, map[string]any{
		"error":   "timeout",
		"timeout": 2 * time.Second,
	})

	bare := &TimeoutError{}
	got := groupAttrMap(bare.LogValue())
	if _, ok := got["timeout"]; ok {
		t.Errorf("expected no 'timeout' attr when zero, got %v", got["timeout"])
	}
}

func groupAttrMap(v slog.Value) map[string]any {
	out := map[string]any{}
	for _, a := range v.Group() {
		out[a.Key] = a.Value.Any()
	}
	return out
}

func assertGroupAttrs(t *testing.T, v slog.Value, want map[string]any) {
	t.Helper()
	got := groupAttrMap(v)
	for k, w := range want {
		g, ok := got[k]
		if !ok {
			t.Errorf("missing attr %q in %v", k, got)
			continue
		}
		if g != w {
			t.Errorf("attr %q = %v (%T), want %v (%T)", k, g, g, w, w)
		}
	}
}
