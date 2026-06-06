// Package slog provides small helpers for tagging structured log records
// with the Fortify pattern that produced them.
//
// This package does NOT replace stdlib log/slog. Use stdlib for logger
// construction; use this package for the Pattern enum and the
// WithPattern helper that scopes a logger to a specific pattern.
//
// Sensitive payloads: this package is payload-blind. It never inspects or
// logs operation arguments, results, or wrapped HTTP / gRPC / LLM payloads.
// If you attach your own attributes to a Fortify-scoped logger, keep
// prompts, request bodies, PII, and credentials out. See
// docs/PRODUCTION.md ("Observability and sensitive payloads") for the full
// policy.
//
// Example:
//
//	import (
//	    "log/slog"
//	    "os"
//
//	    fortifyslog "go.klarlabs.de/fortify/slog"
//	)
//
//	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
//	cbLogger := fortifyslog.WithPattern(logger, fortifyslog.PatternCircuitBreaker)
//	cbLogger.Info("state changed", slog.String("from", "closed"), slog.String("to", "open"))
package slog

import (
	"context"
	"log/slog"
)

// Pattern identifies the resilience pattern that produced a log record.
// All Fortify patterns log with `pattern=<value>` set to one of these
// constants when a *slog.Logger is configured.
type Pattern string

const (
	PatternCircuitBreaker Pattern = "circuit_breaker"
	PatternRetry          Pattern = "retry"
	PatternRateLimit      Pattern = "rate_limit"
	PatternTimeout        Pattern = "timeout"
	PatternBulkhead       Pattern = "bulkhead"
	PatternFallback       Pattern = "fallback"
	PatternHedge          Pattern = "hedge"
	PatternAdaptive       Pattern = "adaptive"
)

// WithPattern returns a logger that automatically includes `pattern=<p>` in
// every record. Returns nil if logger is nil.
func WithPattern(logger *slog.Logger, p Pattern) *slog.Logger {
	if logger == nil {
		return nil
	}
	return logger.With(slog.String("pattern", string(p)))
}

// LogContext logs at the given level using the provided context. Convenience
// wrapper over logger.LogAttrs that no-ops if logger is nil.
func LogContext(ctx context.Context, logger *slog.Logger, level slog.Level, message string, attrs ...slog.Attr) {
	if logger == nil {
		return
	}
	logger.LogAttrs(ctx, level, message, attrs...)
}
