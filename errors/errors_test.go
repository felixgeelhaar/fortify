package errors

import (
	"errors"
	"testing"
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
