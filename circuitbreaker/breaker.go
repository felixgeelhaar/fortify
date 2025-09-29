// Package circuitbreaker provides a generic circuit breaker implementation
// for preventing cascading failures in distributed systems.
//
// A circuit breaker prevents an application from repeatedly trying to execute
// an operation that's likely to fail, allowing it to continue without waiting
// for the fault to be fixed or wasting resources.
//
// The circuit breaker has three states:
//   - Closed: Requests pass through normally. Failures are counted.
//   - Open: Requests fail immediately without execution. After a timeout, transitions to Half-Open.
//   - Half-Open: A limited number of test requests are allowed. Success transitions to Closed, failure to Open.
//
// Example usage:
//
//	cb := circuitbreaker.New[*Response](circuitbreaker.Config{
//	    MaxRequests: 5,
//	    Interval:    10 * time.Second,
//	    Timeout:     60 * time.Second,
//	    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
//	        return counts.ConsecutiveFailures > 5
//	    },
//	})
//
//	result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
//	    return callExternalAPI(ctx)
//	})
package circuitbreaker

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

// CircuitBreaker is a generic interface for circuit breaker pattern implementation.
// It protects against cascading failures by monitoring request outcomes and
// temporarily blocking requests when a failure threshold is reached.
type CircuitBreaker[T any] interface {
	// Execute runs the given function if the circuit breaker allows it.
	// If the circuit is open, it returns ErrCircuitOpen without executing the function.
	// The function receives the context which may be cancelled.
	Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)

	// State returns the current state of the circuit breaker.
	State() State

	// Reset manually resets the circuit breaker to the Closed state and clears all counts.
	Reset()
}

// circuitBreaker is the concrete implementation of CircuitBreaker.
type circuitBreaker[T any] struct {
	expiry        time.Time
	halfOpenStart time.Time
	config        Config
	mu            sync.RWMutex
	counts        Counts
	generation    uint64
	state         State
}

// New creates a new CircuitBreaker with the given configuration.
// The circuit breaker starts in the Closed state.
func New[T any](config Config) CircuitBreaker[T] {
	config.setDefaults()

	cb := &circuitBreaker[T]{
		config:     config,
		state:      StateClosed,
		generation: 0,
		expiry:     time.Now().Add(config.Interval),
	}

	return cb
}

// Execute implements the CircuitBreaker interface.
func (cb *circuitBreaker[T]) Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
	var zero T

	// Check context first
	if err := ctx.Err(); err != nil {
		return zero, err
	}

	// Get current state and generation
	generation, err := cb.beforeRequest()
	if err != nil {
		return zero, err
	}

	// Execute the function
	result, err := fn(ctx)

	// Update state based on result
	cb.afterRequest(generation, cb.config.IsSuccessful(err))

	return result, err
}

// State implements the CircuitBreaker interface.
func (cb *circuitBreaker[T]) State() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	now := time.Now()
	state, _ := cb.currentState(now)
	return state
}

// Reset implements the CircuitBreaker interface.
func (cb *circuitBreaker[T]) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.logStateChange(cb.state, StateClosed)
	cb.state = StateClosed
	cb.generation++
	cb.counts.reset()
	cb.expiry = time.Now().Add(cb.config.Interval)

	if cb.config.OnStateChange != nil {
		// Call outside of lock to prevent potential deadlock
		go cb.config.OnStateChange(cb.state, StateClosed)
	}
}

// beforeRequest checks if the request should be allowed and returns the current generation.
// Returns an error if the circuit is open and not ready for half-open trial requests.
func (cb *circuitBreaker[T]) beforeRequest() (uint64, error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)

	if state == StateOpen {
		return generation, fortifyerrors.ErrCircuitOpen
	}

	if state == StateHalfOpen && cb.counts.Requests >= cb.config.MaxRequests {
		return generation, fortifyerrors.ErrCircuitOpen
	}

	return generation, nil
}

// afterRequest updates the circuit breaker state after a request completes.
// It must be called with the generation returned by beforeRequest.
func (cb *circuitBreaker[T]) afterRequest(generation uint64, success bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	state, currentGen := cb.currentState(now)

	// Ignore results from old generations
	if generation != currentGen {
		return
	}

	if success {
		cb.onSuccess(state, now)
	} else {
		cb.onFailure(state, now)
	}
}

// currentState determines the current state based on time and counts.
// Must be called with lock held.
func (cb *circuitBreaker[T]) currentState(now time.Time) (state State, generation uint64) {
	switch cb.state {
	case StateClosed:
		if cb.config.Interval > 0 && cb.expiry.Before(now) {
			cb.counts.reset()
			cb.expiry = now.Add(cb.config.Interval)
		}
	case StateOpen:
		if cb.expiry.Before(now) {
			cb.setState(StateHalfOpen, now)
		}
	}

	return cb.state, cb.generation
}

// onSuccess handles a successful request outcome.
// Must be called with lock held.
func (cb *circuitBreaker[T]) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onSuccess()
	case StateHalfOpen:
		cb.counts.onSuccess()
		if cb.counts.ConsecutiveSuccesses >= cb.config.MaxRequests {
			cb.setState(StateClosed, now)
		}
	}
}

// onFailure handles a failed request outcome.
// Must be called with lock held.
func (cb *circuitBreaker[T]) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onFailure()
		if cb.config.ReadyToTrip(cb.counts) {
			cb.setState(StateOpen, now)
		}
	case StateHalfOpen:
		cb.setState(StateOpen, now)
	}
}

// setState transitions to a new state and updates related fields.
// Must be called with lock held.
func (cb *circuitBreaker[T]) setState(newState State, now time.Time) {
	if cb.state == newState {
		return
	}

	oldState := cb.state
	cb.state = newState
	cb.generation++
	cb.counts.reset()

	switch newState {
	case StateClosed:
		cb.expiry = now.Add(cb.config.Interval)
	case StateOpen:
		cb.expiry = now.Add(cb.config.Timeout)
	case StateHalfOpen:
		cb.halfOpenStart = now
	}

	cb.logStateChange(oldState, newState)

	if cb.config.OnStateChange != nil {
		// Call outside of lock to prevent potential deadlock
		go cb.config.OnStateChange(oldState, newState)
	}
}

// logStateChange logs state transitions using structured logging.
func (cb *circuitBreaker[T]) logStateChange(from, to State) {
	if cb.config.Logger != nil {
		cb.config.Logger.Info("circuit breaker state changed",
			slog.String("from", from.String()),
			slog.String("to", to.String()),
		)
	}
}

// String returns a string representation of the circuit breaker for debugging.
func (cb *circuitBreaker[T]) String() string {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	return fmt.Sprintf("CircuitBreaker[state=%s, generation=%d, counts=%+v]",
		cb.state, cb.generation, cb.counts)
}
