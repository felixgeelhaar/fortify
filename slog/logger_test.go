package slog

import (
	"bytes"
	"context"
	"encoding/json"
	stdslog "log/slog"
	"testing"
)

func TestPatternConstants(t *testing.T) {
	cases := []struct {
		p    Pattern
		want string
	}{
		{PatternCircuitBreaker, "circuit_breaker"},
		{PatternRetry, "retry"},
		{PatternRateLimit, "rate_limit"},
		{PatternTimeout, "timeout"},
		{PatternBulkhead, "bulkhead"},
		{PatternFallback, "fallback"},
		{PatternHedge, "hedge"},
		{PatternAdaptive, "adaptive"},
	}
	for _, c := range cases {
		if string(c.p) != c.want {
			t.Errorf("%v = %q, want %q", c.p, string(c.p), c.want)
		}
	}
}

func TestWithPattern_AddsPatternAttribute(t *testing.T) {
	var buf bytes.Buffer
	logger := stdslog.New(stdslog.NewJSONHandler(&buf, nil))

	cb := WithPattern(logger, PatternCircuitBreaker)
	cb.Info("hello")

	var rec map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &rec); err != nil {
		t.Fatalf("parse log: %v", err)
	}
	if rec["pattern"] != "circuit_breaker" {
		t.Fatalf("pattern = %v, want circuit_breaker", rec["pattern"])
	}
}

func TestWithPattern_NilLoggerReturnsNil(t *testing.T) {
	if got := WithPattern(nil, PatternRetry); got != nil {
		t.Fatalf("WithPattern(nil) = %v, want nil", got)
	}
}

func TestLogContext_NoOpOnNilLogger(t *testing.T) {
	// Should not panic.
	LogContext(context.Background(), nil, stdslog.LevelInfo, "msg",
		stdslog.String("key", "value"),
	)
}

func TestLogContext_WritesAttrs(t *testing.T) {
	var buf bytes.Buffer
	logger := stdslog.New(stdslog.NewJSONHandler(&buf, nil))

	LogContext(context.Background(), logger, stdslog.LevelWarn, "warn-msg",
		stdslog.String("k", "v"),
	)

	var rec map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &rec); err != nil {
		t.Fatalf("parse log: %v", err)
	}
	if rec["msg"] != "warn-msg" {
		t.Fatalf("msg = %v, want warn-msg", rec["msg"])
	}
	if rec["k"] != "v" {
		t.Fatalf("k = %v, want v", rec["k"])
	}
	if rec["level"] != "WARN" {
		t.Fatalf("level = %v, want WARN", rec["level"])
	}
}
