package otel

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestStartSpan(t *testing.T) {
	t.Run("creates span with pattern attributes", func(t *testing.T) {
		exporter := tracetest.NewInMemoryExporter()
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSyncer(exporter),
		)
		tracer := NewTracer(tp, "test-service")

		_, span := tracer.StartSpan(context.Background(), PatternCircuitBreaker, "test_operation")
		span.End()

		spans := exporter.GetSpans()
		if len(spans) == 0 {
			t.Fatal("no spans recorded")
		}

		spanData := spans[0]
		if spanData.Name != "circuit_breaker.test_operation" {
			t.Errorf("span name = %s, want circuit_breaker.test_operation", spanData.Name)
		}

		// Check for pattern attribute
		hasPattern := false
		for _, attr := range spanData.Attributes {
			if attr.Key == "fortify.pattern" && attr.Value.AsString() == "circuit_breaker" {
				hasPattern = true
				break
			}
		}
		if !hasPattern {
			t.Error("span should have fortify.pattern attribute")
		}
	})

	t.Run("returns valid context and span", func(t *testing.T) {
		exporter := tracetest.NewInMemoryExporter()
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSyncer(exporter),
		)
		tracer := NewTracer(tp, "test-service")

		ctx, span := tracer.StartSpan(context.Background(), PatternRetry, "attempt")
		defer span.End()

		if ctx == nil {
			t.Error("context should not be nil")
		}
		if span == nil {
			t.Error("span should not be nil")
		}
	})
}

func TestRecordError(t *testing.T) {
	t.Run("records error on span", func(t *testing.T) {
		exporter := tracetest.NewInMemoryExporter()
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSyncer(exporter),
		)
		tracer := NewTracer(tp, "test-service")

		_, span := tracer.StartSpan(context.Background(), PatternTimeout, "execute")
		tracer.RecordError(span, context.DeadlineExceeded, "operation timed out")
		span.End()

		spans := exporter.GetSpans()
		if len(spans) == 0 {
			t.Fatal("no spans recorded")
		}

		spanData := spans[0]
		if spanData.Status.Code != codes.Error {
			t.Errorf("span status = %v, want Error", spanData.Status.Code)
		}
	})
}

func TestAddEvent(t *testing.T) {
	t.Run("adds event to span", func(t *testing.T) {
		exporter := tracetest.NewInMemoryExporter()
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSyncer(exporter),
		)
		tracer := NewTracer(tp, "test-service")

		_, span := tracer.StartSpan(context.Background(), PatternCircuitBreaker, "execute")
		tracer.AddEvent(span, "state_change", attribute.String("from", "closed"), attribute.String("to", "open"))
		span.End()

		spans := exporter.GetSpans()
		if len(spans) == 0 {
			t.Fatal("no spans recorded")
		}

		spanData := spans[0]
		if len(spanData.Events) == 0 {
			t.Error("span should have events")
		}
	})
}

func TestSetAttributes(t *testing.T) {
	t.Run("sets attributes on span", func(t *testing.T) {
		exporter := tracetest.NewInMemoryExporter()
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSyncer(exporter),
		)
		tracer := NewTracer(tp, "test-service")

		_, span := tracer.StartSpan(context.Background(), PatternRateLimit, "check")
		tracer.SetAttributes(span,
			attribute.Int("tokens", 5),
			attribute.Int("remaining", 95),
		)
		span.End()

		spans := exporter.GetSpans()
		if len(spans) == 0 {
			t.Fatal("no spans recorded")
		}

		spanData := spans[0]
		hasTokens := false
		for _, attr := range spanData.Attributes {
			if attr.Key == "tokens" && attr.Value.AsInt64() == 5 {
				hasTokens = true
				break
			}
		}
		if !hasTokens {
			t.Error("span should have tokens attribute")
		}
	})
}

func TestNoOpTracer(t *testing.T) {
	t.Run("creates no-op tracer when provider is nil", func(t *testing.T) {
		tracer := NewTracer(nil, "test-service")
		if tracer == nil {
			t.Error("tracer should not be nil")
		}

		_, span := tracer.StartSpan(context.Background(), PatternBulkhead, "execute")
		defer span.End()

		// Should not panic
		tracer.RecordError(span, context.Canceled, "test")
		tracer.AddEvent(span, "test")
		tracer.SetAttributes(span, attribute.String("key", "value"))
	})
}

func TestPatternTypes(t *testing.T) {
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