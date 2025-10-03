// Package testing provides utilities for chaos engineering and resilience testing.
//
// This package helps test the robustness of resilience patterns by simulating
// various failure scenarios, network conditions, and system behaviors.
package testing

import (
	"context"
	"errors"

	// Using math/rand instead of crypto/rand in chaos testing utilities because:
	// 1. This is a testing package for simulating failures, not production code
	// 2. Random failures don't require cryptographic security - just statistical variance
	// 3. math/rand provides sufficient randomness for chaos engineering
	// 4. crypto/rand would add unnecessary overhead to test utilities.
	// 5. Predictable seeds can actually be useful for reproducing test scenarios.
	"math/rand" //nolint:gosec // G404: weak random is intentional and appropriate for testing utilities
	"sync"
	"sync/atomic"
	"time"
)

// ErrorInjector simulates failures by returning errors based on configured probability.
type ErrorInjector struct {
	err         error        // Error to return
	mu          sync.RWMutex // Protects configuration
	calls       atomic.Int64 // Total calls
	failures    atomic.Int64 // Total failures
	probability float64      // Probability of error (0.0 to 1.0)
}

// NewErrorInjector creates an error injector with the given probability.
// Probability should be between 0.0 (never fail) and 1.0 (always fail).
func NewErrorInjector(probability float64, err error) *ErrorInjector {
	if probability < 0 {
		probability = 0
	}
	if probability > 1 {
		probability = 1
	}
	if err == nil {
		err = errors.New("injected error")
	}

	return &ErrorInjector{
		probability: probability,
		err:         err,
	}
}

// ShouldFail returns true if an error should be injected based on probability.
func (e *ErrorInjector) ShouldFail() bool {
	e.calls.Add(1)
	e.mu.RLock()
	prob := e.probability
	e.mu.RUnlock()

	// Weak random is appropriate here because this is a testing utility
	// that simulates probabilistic failures for chaos engineering
	//nolint:gosec // G404: weak random is intentional for testing
	if rand.Float64() < prob {
		e.failures.Add(1)
		return true
	}
	return false
}

// Error returns the error to inject if ShouldFail returns true.
func (e *ErrorInjector) Error() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.err
}

// Execute runs the function, potentially injecting an error.
func (e *ErrorInjector) Execute(fn func() error) error {
	if e.ShouldFail() {
		return e.Error()
	}
	return fn()
}

// SetProbability updates the failure probability.
func (e *ErrorInjector) SetProbability(p float64) {
	if p < 0 {
		p = 0
	}
	if p > 1 {
		p = 1
	}
	e.mu.Lock()
	e.probability = p
	e.mu.Unlock()
}

// Stats returns statistics about injected errors.
func (e *ErrorInjector) Stats() (calls, failures int64, rate float64) {
	calls = e.calls.Load()
	failures = e.failures.Load()
	if calls > 0 {
		rate = float64(failures) / float64(calls)
	}
	return
}

// Reset resets the error injection statistics.
func (e *ErrorInjector) Reset() {
	e.calls.Store(0)
	e.failures.Store(0)
}

// LatencyInjector simulates network latency by delaying operations.
type LatencyInjector struct {
	min        time.Duration
	max        time.Duration
	mu         sync.RWMutex
	calls      atomic.Int64
	totalDelay atomic.Int64 // in nanoseconds
}

// NewLatencyInjector creates a latency injector with random delays between minLatency and maxLatency.
func NewLatencyInjector(minLatency, maxLatency time.Duration) *LatencyInjector {
	if minLatency < 0 {
		minLatency = 0
	}
	if maxLatency < minLatency {
		maxLatency = minLatency
	}

	return &LatencyInjector{
		min: minLatency,
		max: maxLatency,
	}
}

// Delay introduces a random delay between min and max duration.
func (l *LatencyInjector) Delay(ctx context.Context) error {
	l.calls.Add(1)

	l.mu.RLock()
	minDelay, maxDelay := l.min, l.max
	l.mu.RUnlock()

	// Calculate random delay
	delay := minDelay
	if maxDelay > minDelay {
		// Weak random is appropriate here because this is a testing utility
		// that simulates variable latency for chaos engineering
		//nolint:gosec // G404: weak random is intentional for testing
		delay = minDelay + time.Duration(rand.Int63n(int64(maxDelay-minDelay)))
	}

	l.totalDelay.Add(int64(delay))

	// Wait with context cancellation support
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetLatency updates the latency range.
func (l *LatencyInjector) SetLatency(minLatency, maxLatency time.Duration) {
	if minLatency < 0 {
		minLatency = 0
	}
	if maxLatency < minLatency {
		maxLatency = minLatency
	}
	l.mu.Lock()
	l.min = minLatency
	l.max = maxLatency
	l.mu.Unlock()
}

// Stats returns statistics about injected latency.
func (l *LatencyInjector) Stats() (calls int64, avgLatency time.Duration) {
	calls = l.calls.Load()
	totalDelay := l.totalDelay.Load()
	if calls > 0 {
		avgLatency = time.Duration(totalDelay / calls)
	}
	return
}

// Reset resets the latency statistics.
func (l *LatencyInjector) Reset() {
	l.calls.Store(0)
	l.totalDelay.Store(0)
}

// TimeoutSimulator simulates timeout conditions by canceling contexts.
type TimeoutSimulator struct {
	timeout     time.Duration
	probability float64
	mu          sync.RWMutex
	timeouts    atomic.Int64
	calls       atomic.Int64
}

// NewTimeoutSimulator creates a simulator that cancels contexts based on probability.
func NewTimeoutSimulator(timeout time.Duration, probability float64) *TimeoutSimulator {
	if timeout < 0 {
		timeout = 0
	}
	if probability < 0 {
		probability = 0
	}
	if probability > 1 {
		probability = 1
	}

	return &TimeoutSimulator{
		timeout:     timeout,
		probability: probability,
	}
}

// Context returns a context that may be automatically canceled based on configuration.
func (t *TimeoutSimulator) Context(parent context.Context) (context.Context, context.CancelFunc) {
	t.calls.Add(1)

	t.mu.RLock()
	timeout := t.timeout
	prob := t.probability
	t.mu.RUnlock()

	// Decide if we should simulate timeout
	// Weak random is appropriate here because this is a testing utility
	// that simulates probabilistic timeouts for chaos engineering
	//nolint:gosec // G404: weak random is intentional for testing
	if rand.Float64() < prob {
		t.timeouts.Add(1)
		return context.WithTimeout(parent, timeout)
	}

	// Return a context that won't timeout (but still cancelable)
	return context.WithCancel(parent)
}

// Stats returns statistics about simulated timeouts.
func (t *TimeoutSimulator) Stats() (calls, timeouts int64, rate float64) {
	calls = t.calls.Load()
	timeouts = t.timeouts.Load()
	if calls > 0 {
		rate = float64(timeouts) / float64(calls)
	}
	return
}

// Reset resets the timeout statistics.
func (t *TimeoutSimulator) Reset() {
	t.calls.Store(0)
	t.timeouts.Store(0)
}

// FlakeyService simulates an unreliable service with various failure modes.
type FlakeyService struct {
	errorInjector   *ErrorInjector
	latencyInjector *LatencyInjector
	timeoutSim      *TimeoutSimulator
}

// NewFlakeyService creates a service simulator with error, latency, and timeout injection.
func NewFlakeyService(errorProb float64, minLatency, maxLatency time.Duration) *FlakeyService {
	return &FlakeyService{
		errorInjector:   NewErrorInjector(errorProb, errors.New("service error")),
		latencyInjector: NewLatencyInjector(minLatency, maxLatency),
		timeoutSim:      NewTimeoutSimulator(maxLatency*2, 0.1), // 10% timeout probability
	}
}

// Call simulates calling a flakey service with potential errors and latency.
func (f *FlakeyService) Call(ctx context.Context, fn func() error) error {
	// Inject latency
	if err := f.latencyInjector.Delay(ctx); err != nil {
		return err
	}

	// Check for injected errors
	if f.errorInjector.ShouldFail() {
		return f.errorInjector.Error()
	}

	// Execute actual function
	return fn()
}

// Stats returns combined statistics from all injectors.
func (f *FlakeyService) Stats() map[string]interface{} {
	errorCalls, errorFailures, errorRate := f.errorInjector.Stats()
	latencyCalls, avgLatency := f.latencyInjector.Stats()
	timeoutCalls, timeouts, timeoutRate := f.timeoutSim.Stats()

	return map[string]interface{}{
		"error_calls":    errorCalls,
		"error_failures": errorFailures,
		"error_rate":     errorRate,
		"latency_calls":  latencyCalls,
		"avg_latency":    avgLatency,
		"timeout_calls":  timeoutCalls,
		"timeouts":       timeouts,
		"timeout_rate":   timeoutRate,
	}
}

// Reset resets all statistics.
func (f *FlakeyService) Reset() {
	f.errorInjector.Reset()
	f.latencyInjector.Reset()
	f.timeoutSim.Reset()
}
