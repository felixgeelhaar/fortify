package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestCollector(t *testing.T) {
	collector := NewCollector()
	registry := prometheus.NewRegistry()

	if err := registry.Register(collector); err != nil {
		t.Fatalf("Failed to register collector: %v", err)
	}

	// Test circuit breaker metrics
	collector.RecordCircuitBreakerState("test-cb", 0)
	collector.RecordCircuitBreakerRequest("test-cb", "closed")
	collector.RecordCircuitBreakerSuccess("test-cb")
	collector.RecordCircuitBreakerStateChange("test-cb", "closed", "open")

	// Test retry metrics
	collector.RecordRetryAttempts("test-retry", 3)
	collector.RecordRetrySuccess("test-retry")
	collector.RecordRetryDuration("test-retry", 0.5)

	// Test rate limit metrics
	collector.RecordRateLimitAllowed("test-rl", "key1")
	collector.RecordRateLimitDenied("test-rl", "key1")
	collector.RecordRateLimitWaitTime("test-rl", "key1", 0.1)

	// Test timeout metrics
	collector.RecordTimeoutExecution("test-timeout")
	collector.RecordTimeoutExceeded("test-timeout")
	collector.RecordTimeoutDuration("test-timeout", true, 1.0)

	// Test bulkhead metrics
	collector.RecordBulkheadActive("test-bh", 5)
	collector.RecordBulkheadQueued("test-bh", 2)
	collector.RecordBulkheadRejected("test-bh")
	collector.RecordBulkheadSuccess("test-bh")
	collector.RecordBulkheadDuration("test-bh", 0.3)

	// Verify metrics are collected
	count, err := testutil.GatherAndCount(registry)
	if err != nil {
		t.Fatalf("Failed to gather metrics: %v", err)
	}

	if count == 0 {
		t.Error("Expected metrics to be collected, but got 0")
	}
}

func TestDefaultCollector(t *testing.T) {
	collector := DefaultCollector()
	if collector == nil {
		t.Fatal("Expected default collector, got nil")
	}

	// Should return same instance
	collector2 := DefaultCollector()
	if collector != collector2 {
		t.Error("Expected same default collector instance")
	}
}

func TestMustRegister(t *testing.T) {
	registry := prometheus.NewRegistry()

	// Should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustRegister panicked: %v", r)
		}
	}()

	MustRegister(registry)
}

func BenchmarkRecordCircuitBreakerRequest(b *testing.B) {
	collector := NewCollector()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		collector.RecordCircuitBreakerRequest("bench-cb", "closed")
	}
}

func BenchmarkRecordRetryAttempts(b *testing.B) {
	collector := NewCollector()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		collector.RecordRetryAttempts("bench-retry", 3)
	}
}

func BenchmarkRecordRateLimitAllowed(b *testing.B) {
	collector := NewCollector()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		collector.RecordRateLimitAllowed("bench-rl", "key1")
	}
}
