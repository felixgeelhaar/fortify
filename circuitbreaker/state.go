package circuitbreaker

import "fmt"

// State represents the current state of the circuit breaker.
type State int

const (
	// StateClosed allows requests to pass through normally.
	// The circuit breaker monitors requests and failures.
	StateClosed State = iota

	// StateOpen blocks all requests immediately without executing them.
	// After the timeout period, the circuit transitions to Half-Open.
	StateOpen

	// StateHalfOpen allows a limited number of trial requests to pass through.
	// If they succeed, the circuit transitions back to Closed.
	// If they fail, the circuit transitions back to Open.
	StateHalfOpen
)

// String returns the string representation of the State.
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return fmt.Sprintf("unknown(%d)", s)
	}
}

// Counts holds statistics about circuit breaker requests and outcomes.
// All fields are updated atomically and can be safely accessed concurrently.
type Counts struct {
	// Requests is the total number of requests in the current interval.
	Requests uint32

	// TotalSuccesses is the total number of successful requests in the current interval.
	TotalSuccesses uint32

	// TotalFailures is the total number of failed requests in the current interval.
	TotalFailures uint32

	// ConsecutiveSuccesses is the number of consecutive successful requests.
	// Reset to 0 on any failure.
	ConsecutiveSuccesses uint32

	// ConsecutiveFailures is the number of consecutive failed requests.
	// Reset to 0 on any success.
	ConsecutiveFailures uint32
}

// onSuccess updates counts after a successful request.
func (c *Counts) onSuccess() {
	c.Requests++
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

// onFailure updates counts after a failed request.
func (c *Counts) onFailure() {
	c.Requests++
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

// reset resets all counts to zero.
func (c *Counts) reset() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}