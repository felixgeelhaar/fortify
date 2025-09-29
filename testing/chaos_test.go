package testing

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestErrorInjector_AlwaysFail(t *testing.T) {
	injector := NewErrorInjector(1.0, errors.New("test error"))

	for i := 0; i < 100; i++ {
		if !injector.ShouldFail() {
			t.Error("Expected ShouldFail to always return true with probability 1.0")
		}
	}

	calls, failures, rate := injector.Stats()
	if calls != 100 {
		t.Errorf("Expected 100 calls, got %d", calls)
	}
	if failures != 100 {
		t.Errorf("Expected 100 failures, got %d", failures)
	}
	if rate != 1.0 {
		t.Errorf("Expected failure rate 1.0, got %f", rate)
	}
}

func TestErrorInjector_NeverFail(t *testing.T) {
	injector := NewErrorInjector(0.0, errors.New("test error"))

	for i := 0; i < 100; i++ {
		if injector.ShouldFail() {
			t.Error("Expected ShouldFail to always return false with probability 0.0")
		}
	}

	calls, failures, rate := injector.Stats()
	if calls != 100 {
		t.Errorf("Expected 100 calls, got %d", calls)
	}
	if failures != 0 {
		t.Errorf("Expected 0 failures, got %d", failures)
	}
	if rate != 0.0 {
		t.Errorf("Expected failure rate 0.0, got %f", rate)
	}
}

func TestErrorInjector_Execute(t *testing.T) {
	injector := NewErrorInjector(1.0, errors.New("injected"))

	err := injector.Execute(func() error {
		return nil
	})

	if err == nil {
		t.Error("Expected error to be injected")
	}
	if err.Error() != "injected" {
		t.Errorf("Expected 'injected' error, got %v", err)
	}
}

func TestErrorInjector_SetProbability(t *testing.T) {
	injector := NewErrorInjector(0.0, errors.New("test"))

	injector.SetProbability(1.0)

	if !injector.ShouldFail() {
		t.Error("Expected ShouldFail to return true after setting probability to 1.0")
	}
}

func TestErrorInjector_Reset(t *testing.T) {
	injector := NewErrorInjector(1.0, errors.New("test"))

	// Generate some failures
	for i := 0; i < 10; i++ {
		injector.ShouldFail()
	}

	injector.Reset()

	calls, failures, _ := injector.Stats()
	if calls != 0 {
		t.Errorf("Expected 0 calls after reset, got %d", calls)
	}
	if failures != 0 {
		t.Errorf("Expected 0 failures after reset, got %d", failures)
	}
}

func TestLatencyInjector_FixedDelay(t *testing.T) {
	delay := 50 * time.Millisecond
	injector := NewLatencyInjector(delay, delay)

	start := time.Now()
	err := injector.Delay(context.Background())
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Allow 10ms tolerance
	if elapsed < delay-10*time.Millisecond || elapsed > delay+10*time.Millisecond {
		t.Errorf("Expected delay ~%v, got %v", delay, elapsed)
	}
}

func TestLatencyInjector_ContextCancellation(t *testing.T) {
	injector := NewLatencyInjector(1*time.Second, 2*time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := injector.Delay(ctx)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func TestLatencyInjector_Stats(t *testing.T) {
	injector := NewLatencyInjector(10*time.Millisecond, 20*time.Millisecond)

	for i := 0; i < 5; i++ {
		_ = injector.Delay(context.Background())
	}

	calls, avgLatency := injector.Stats()
	if calls != 5 {
		t.Errorf("Expected 5 calls, got %d", calls)
	}

	// Average should be between 10-20ms
	if avgLatency < 10*time.Millisecond || avgLatency > 20*time.Millisecond {
		t.Errorf("Expected average latency between 10-20ms, got %v", avgLatency)
	}
}

func TestLatencyInjector_SetLatency(t *testing.T) {
	injector := NewLatencyInjector(10*time.Millisecond, 20*time.Millisecond)

	injector.SetLatency(1*time.Millisecond, 2*time.Millisecond)

	start := time.Now()
	_ = injector.Delay(context.Background())
	elapsed := time.Since(start)

	// Should be much faster now
	if elapsed > 10*time.Millisecond {
		t.Errorf("Expected delay < 10ms after SetLatency, got %v", elapsed)
	}
}

func TestTimeoutSimulator_AlwaysTimeout(t *testing.T) {
	sim := NewTimeoutSimulator(10*time.Millisecond, 1.0)

	for i := 0; i < 10; i++ {
		ctx, cancel := sim.Context(context.Background())
		defer cancel()

		select {
		case <-ctx.Done():
			// Expected timeout
		case <-time.After(50 * time.Millisecond):
			t.Error("Expected context to timeout")
		}
	}

	calls, timeouts, rate := sim.Stats()
	if calls != 10 {
		t.Errorf("Expected 10 calls, got %d", calls)
	}
	if timeouts != 10 {
		t.Errorf("Expected 10 timeouts, got %d", timeouts)
	}
	if rate != 1.0 {
		t.Errorf("Expected timeout rate 1.0, got %f", rate)
	}
}

func TestTimeoutSimulator_NeverTimeout(t *testing.T) {
	sim := NewTimeoutSimulator(10*time.Millisecond, 0.0)

	ctx, cancel := sim.Context(context.Background())
	defer cancel()

	select {
	case <-ctx.Done():
		t.Error("Context should not timeout with probability 0.0")
	case <-time.After(50 * time.Millisecond):
		// Expected - no timeout
	}

	calls, timeouts, _ := sim.Stats()
	if calls != 1 {
		t.Errorf("Expected 1 call, got %d", calls)
	}
	if timeouts != 0 {
		t.Errorf("Expected 0 timeouts, got %d", timeouts)
	}
}

func TestFlakeyService_Success(t *testing.T) {
	service := NewFlakeyService(0.0, 0, 0) // Never fail, no latency

	err := service.Call(context.Background(), func() error {
		return nil
	})

	if err != nil {
		t.Errorf("Expected no error from service with 0 error probability, got %v", err)
	}
}

func TestFlakeyService_AlwaysFail(t *testing.T) {
	service := NewFlakeyService(1.0, 0, 0) // Always fail

	err := service.Call(context.Background(), func() error {
		return nil
	})

	if err == nil {
		t.Error("Expected error from service with 1.0 error probability")
	}
}

func TestFlakeyService_WithLatency(t *testing.T) {
	service := NewFlakeyService(0.0, 50*time.Millisecond, 100*time.Millisecond)

	start := time.Now()
	err := service.Call(context.Background(), func() error {
		return nil
	})
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Should have some latency
	if elapsed < 50*time.Millisecond {
		t.Errorf("Expected at least 50ms latency, got %v", elapsed)
	}
}

func TestFlakeyService_Stats(t *testing.T) {
	service := NewFlakeyService(0.5, 10*time.Millisecond, 20*time.Millisecond)

	// Make several calls
	for i := 0; i < 20; i++ {
		_ = service.Call(context.Background(), func() error {
			return nil
		})
	}

	stats := service.Stats()

	// Check that stats contain expected keys
	expectedKeys := []string{
		"error_calls", "error_failures", "error_rate",
		"latency_calls", "avg_latency",
		"timeout_calls", "timeouts", "timeout_rate",
	}

	for _, key := range expectedKeys {
		if _, ok := stats[key]; !ok {
			t.Errorf("Expected stats to contain key %q", key)
		}
	}

	// Error rate should be close to 0.5
	errorRate := stats["error_rate"].(float64)
	if errorRate < 0.3 || errorRate > 0.7 {
		t.Logf("Warning: error rate %f is outside expected range 0.3-0.7", errorRate)
	}
}

func TestFlakeyService_Reset(t *testing.T) {
	service := NewFlakeyService(1.0, 0, 0)

	// Make some calls
	for i := 0; i < 5; i++ {
		_ = service.Call(context.Background(), func() error {
			return nil
		})
	}

	service.Reset()

	stats := service.Stats()
	if stats["error_calls"].(int64) != 0 {
		t.Error("Expected error_calls to be 0 after reset")
	}
}

func BenchmarkErrorInjector_ShouldFail(b *testing.B) {
	injector := NewErrorInjector(0.5, errors.New("test"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		injector.ShouldFail()
	}
}

func BenchmarkLatencyInjector_Delay(b *testing.B) {
	injector := NewLatencyInjector(1*time.Microsecond, 10*time.Microsecond)
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = injector.Delay(ctx)
	}
}

func BenchmarkFlakeyService_Call(b *testing.B) {
	service := NewFlakeyService(0.1, 1*time.Microsecond, 10*time.Microsecond)
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.Call(ctx, func() error {
			return nil
		})
	}
}