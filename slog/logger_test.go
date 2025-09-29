package slog

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"
	"time"
)

func TestLoggerWithPattern(t *testing.T) {
	t.Run("logs circuit breaker state change", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		LogPatternEvent(logger, PatternCircuitBreaker, "state_change",
			slog.String("from", "closed"),
			slog.String("to", "open"),
			slog.Int("failures", 5),
		)

		output := buf.String()
		if !strings.Contains(output, "pattern") {
			t.Error("output should contain pattern field")
		}
		if !strings.Contains(output, "circuit_breaker") {
			t.Error("output should contain circuit_breaker pattern")
		}
		if !strings.Contains(output, "state_change") {
			t.Error("output should contain event type")
		}
	})

	t.Run("logs retry attempt", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		LogPatternEvent(logger, PatternRetry, "attempt",
			slog.Int("attempt", 2),
			slog.Duration("backoff", 100*time.Millisecond),
		)

		output := buf.String()
		if !strings.Contains(output, "retry") {
			t.Error("output should contain retry pattern")
		}
		if !strings.Contains(output, "attempt") {
			t.Error("output should contain event type")
		}
	})
}

func TestLoggerWithContext(t *testing.T) {
	t.Run("adds context values to log", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		ctx := context.WithValue(context.Background(), "request_id", "req-123")
		LogWithContext(ctx, logger, slog.LevelInfo, "test message",
			slog.String("key", "value"),
		)

		output := buf.String()
		if !strings.Contains(output, "test message") {
			t.Error("output should contain message")
		}
	})
}

func TestLoggerPatternTypes(t *testing.T) {
	patterns := []struct {
		pattern  Pattern
		expected string
	}{
		{PatternCircuitBreaker, "circuit_breaker"},
		{PatternRetry, "retry"},
		{PatternRateLimit, "rate_limit"},
		{PatternTimeout, "timeout"},
		{PatternBulkhead, "bulkhead"},
	}

	for _, tc := range patterns {
		t.Run(string(tc.pattern), func(t *testing.T) {
			if string(tc.pattern) != tc.expected {
				t.Errorf("pattern = %s, want %s", tc.pattern, tc.expected)
			}
		})
	}
}

func TestNewLoggerWithDefaults(t *testing.T) {
	t.Run("creates logger with default settings", func(t *testing.T) {
		logger := NewLogger(nil)
		if logger == nil {
			t.Error("logger should not be nil")
		}
	})

	t.Run("uses provided options", func(t *testing.T) {
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		logger := NewLogger(handler)

		logger.Debug("test debug message")
		if !strings.Contains(buf.String(), "test debug message") {
			t.Error("should log debug messages with debug level handler")
		}
	})
}

func TestLogPatternError(t *testing.T) {
	t.Run("logs pattern error", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelError,
		}))

		LogPatternError(logger, PatternCircuitBreaker, "execution failed",
			slog.String("error", "connection timeout"),
		)

		output := buf.String()
		if !strings.Contains(output, "circuit_breaker") {
			t.Error("output should contain pattern")
		}
		if !strings.Contains(output, "execution failed") {
			t.Error("output should contain message")
		}
	})
}

func TestLogPatternMetrics(t *testing.T) {
	t.Run("logs pattern metrics", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		LogPatternMetrics(logger, PatternRateLimit, "tokens_consumed",
			slog.Int("tokens", 5),
			slog.Int("remaining", 95),
		)

		output := buf.String()
		if !strings.Contains(output, "rate_limit") {
			t.Error("output should contain pattern")
		}
		if !strings.Contains(output, "tokens_consumed") {
			t.Error("output should contain metric name")
		}
	})
}

func TestNewTextLogger(t *testing.T) {
	t.Run("creates text logger with custom writer", func(t *testing.T) {
		var buf bytes.Buffer
		logger := NewTextLogger(&buf, slog.LevelInfo)
		if logger == nil {
			t.Error("logger should not be nil")
		}

		logger.Info("test message")
		output := buf.String()
		if !strings.Contains(output, "test message") {
			t.Error("output should contain message")
		}
	})

	t.Run("uses stdout when writer is nil", func(t *testing.T) {
		logger := NewTextLogger(nil, slog.LevelDebug)
		if logger == nil {
			t.Error("logger should not be nil")
		}
	})
}

func TestNewJSONLogger(t *testing.T) {
	t.Run("creates JSON logger with custom writer", func(t *testing.T) {
		var buf bytes.Buffer
		logger := NewJSONLogger(&buf, slog.LevelInfo)
		if logger == nil {
			t.Error("logger should not be nil")
		}

		logger.Info("test message")
		output := buf.String()
		if !strings.Contains(output, "test message") {
			t.Error("output should contain message")
		}
		// JSON output should contain quotes
		if !strings.Contains(output, `"msg"`) {
			t.Error("output should be JSON formatted")
		}
	})

	t.Run("uses stdout when writer is nil", func(t *testing.T) {
		logger := NewJSONLogger(nil, slog.LevelWarn)
		if logger == nil {
			t.Error("logger should not be nil")
		}
	})
}

func TestWithPattern(t *testing.T) {
	t.Run("returns logger with pattern field", func(t *testing.T) {
		var buf bytes.Buffer
		baseLogger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		patternLogger := WithPattern(baseLogger, PatternTimeout)
		if patternLogger == nil {
			t.Error("pattern logger should not be nil")
		}

		patternLogger.Info("test event")
		output := buf.String()
		if !strings.Contains(output, "timeout") {
			t.Error("output should contain pattern field")
		}
	})

	t.Run("returns nil when base logger is nil", func(t *testing.T) {
		patternLogger := WithPattern(nil, PatternRetry)
		if patternLogger != nil {
			t.Error("should return nil when base logger is nil")
		}
	})
}

func TestLogPatternEvent_NilLogger(t *testing.T) {
	// Should not panic with nil logger
	LogPatternEvent(nil, PatternCircuitBreaker, "test", slog.String("key", "value"))
}

func TestLogPatternError_NilLogger(t *testing.T) {
	// Should not panic with nil logger
	LogPatternError(nil, PatternRetry, "test error", slog.String("error", "details"))
}

func TestLogPatternMetrics_NilLogger(t *testing.T) {
	// Should not panic with nil logger
	LogPatternMetrics(nil, PatternBulkhead, "metric", slog.Int("value", 42))
}

func TestLogWithContext_NilLogger(t *testing.T) {
	// Should not panic with nil logger
	LogWithContext(context.Background(), nil, slog.LevelInfo, "test", slog.String("key", "value"))
}
