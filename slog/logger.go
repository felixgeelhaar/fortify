// Package slog provides structured logging utilities for Fortify resilience patterns.
//
// This package extends Go's standard log/slog package with pattern-specific
// logging helpers that add consistent metadata for circuit breakers, retries,
// rate limiters, timeouts, and bulkheads.
//
// Example usage:
//
//	logger := slog.NewLogger(slog.NewJSONHandler(os.Stdout, nil))
//
//	// Log a circuit breaker event
//	LogPatternEvent(logger, PatternCircuitBreaker, "state_change",
//	    slog.String("from", "closed"),
//	    slog.String("to", "open"),
//	)
package slog

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// Pattern represents a resilience pattern type for logging.
type Pattern string

const (
	// PatternCircuitBreaker identifies circuit breaker pattern logs.
	PatternCircuitBreaker Pattern = "circuit_breaker"

	// PatternRetry identifies retry pattern logs.
	PatternRetry Pattern = "retry"

	// PatternRateLimit identifies rate limiter pattern logs.
	PatternRateLimit Pattern = "rate_limit"

	// PatternTimeout identifies timeout pattern logs.
	PatternTimeout Pattern = "timeout"

	// PatternBulkhead identifies bulkhead pattern logs.
	PatternBulkhead Pattern = "bulkhead"
)

// NewLogger creates a new structured logger with the given handler.
// If handler is nil, creates a default JSON handler writing to stdout.
func NewLogger(handler slog.Handler) *slog.Logger {
	if handler == nil {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}
	return slog.New(handler)
}

// NewTextLogger creates a logger with human-readable text output.
func NewTextLogger(w io.Writer, level slog.Level) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}
	handler := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(handler)
}

// NewJSONLogger creates a logger with JSON output.
func NewJSONLogger(w io.Writer, level slog.Level) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(handler)
}

// LogPatternEvent logs a pattern-specific event with structured metadata.
func LogPatternEvent(logger *slog.Logger, pattern Pattern, event string, attrs ...slog.Attr) {
	if logger == nil {
		return
	}

	args := make([]slog.Attr, 0, len(attrs)+2)
	args = append(args, slog.String("pattern", string(pattern)))
	args = append(args, slog.String("event", event))
	args = append(args, attrs...)

	logger.LogAttrs(context.Background(), slog.LevelInfo, "pattern event", args...)
}

// LogPatternError logs a pattern-specific error with structured metadata.
func LogPatternError(logger *slog.Logger, pattern Pattern, message string, attrs ...slog.Attr) {
	if logger == nil {
		return
	}

	args := make([]slog.Attr, 0, len(attrs)+1)
	args = append(args, slog.String("pattern", string(pattern)))
	args = append(args, attrs...)

	logger.LogAttrs(context.Background(), slog.LevelError, message, args...)
}

// LogPatternMetrics logs pattern metrics with structured metadata.
func LogPatternMetrics(logger *slog.Logger, pattern Pattern, metric string, attrs ...slog.Attr) {
	if logger == nil {
		return
	}

	args := make([]slog.Attr, 0, len(attrs)+2)
	args = append(args, slog.String("pattern", string(pattern)))
	args = append(args, slog.String("metric", metric))
	args = append(args, attrs...)

	logger.LogAttrs(context.Background(), slog.LevelInfo, "pattern metrics", args...)
}

// LogWithContext logs a message with context values extracted.
func LogWithContext(ctx context.Context, logger *slog.Logger, level slog.Level, message string, attrs ...slog.Attr) {
	if logger == nil {
		return
	}

	logger.LogAttrs(ctx, level, message, attrs...)
}

// WithPattern returns a logger that automatically includes the pattern field.
func WithPattern(logger *slog.Logger, pattern Pattern) *slog.Logger {
	if logger == nil {
		return nil
	}
	return logger.With(slog.String("pattern", string(pattern)))
}
