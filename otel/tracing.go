// Package otel provides OpenTelemetry integration for Fortify resilience patterns.
//
// This package enables distributed tracing and metrics collection for circuit breakers,
// retries, rate limiters, timeouts, and bulkheads using OpenTelemetry standards.
//
// Example usage:
//
//	import (
//	    "go.opentelemetry.io/otel"
//	    "go.opentelemetry.io/otel/sdk/trace"
//	)
//
//	tp := trace.NewTracerProvider(...)
//	tracer := otel.NewTracer(tp, "my-service")
//
//	// Start a span for a circuit breaker operation
//	ctx, span := tracer.StartSpan(ctx, otel.PatternCircuitBreaker, "execute")
//	defer span.End()
package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Pattern represents a resilience pattern type for tracing.
type Pattern string

const (
	// PatternCircuitBreaker identifies circuit breaker pattern traces.
	PatternCircuitBreaker Pattern = "circuit_breaker"

	// PatternRetry identifies retry pattern traces.
	PatternRetry Pattern = "retry"

	// PatternRateLimit identifies rate limiter pattern traces.
	PatternRateLimit Pattern = "rate_limit"

	// PatternTimeout identifies timeout pattern traces.
	PatternTimeout Pattern = "timeout"

	// PatternBulkhead identifies bulkhead pattern traces.
	PatternBulkhead Pattern = "bulkhead"
)

// Tracer wraps an OpenTelemetry tracer with pattern-specific helpers.
type Tracer struct {
	tracer      trace.Tracer
	serviceName string
}

// NewTracer creates a new Tracer with the given TracerProvider and service name.
// If provider is nil, returns a tracer that uses the global TracerProvider.
func NewTracer(provider trace.TracerProvider, serviceName string) *Tracer {
	if provider == nil {
		provider = trace.NewNoopTracerProvider()
	}

	return &Tracer{
		tracer:      provider.Tracer(serviceName),
		serviceName: serviceName,
	}
}

// StartSpan starts a new span for the given pattern and operation.
// The span name is automatically prefixed with the pattern name.
func (t *Tracer) StartSpan(ctx context.Context, pattern Pattern, operation string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("%s.%s", pattern, operation)

	ctx, span := t.tracer.Start(ctx, spanName,
		trace.WithAttributes(
			attribute.String("fortify.pattern", string(pattern)),
			attribute.String("service.name", t.serviceName),
		),
	)

	return ctx, span
}

// RecordError records an error on the span with a descriptive message.
func (t *Tracer) RecordError(span trace.Span, err error, message string) {
	if span == nil || err == nil {
		return
	}

	span.RecordError(err, trace.WithAttributes(
		attribute.String("error.message", message),
	))
	span.SetStatus(codes.Error, message)
}

// AddEvent adds an event to the span with optional attributes.
func (t *Tracer) AddEvent(span trace.Span, name string, attrs ...attribute.KeyValue) {
	if span == nil {
		return
	}

	span.AddEvent(name, trace.WithAttributes(attrs...))
}

// SetAttributes sets attributes on the span.
func (t *Tracer) SetAttributes(span trace.Span, attrs ...attribute.KeyValue) {
	if span == nil {
		return
	}

	span.SetAttributes(attrs...)
}

// SetStatus sets the status of the span.
func (t *Tracer) SetStatus(span trace.Span, code codes.Code, description string) {
	if span == nil {
		return
	}

	span.SetStatus(code, description)
}
