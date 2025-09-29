// Package metrics provides Prometheus metrics integration for Fortify resilience patterns.
//
// This package exports metrics for all patterns including circuit breakers, retries,
// rate limiters, timeouts, and bulkheads, enabling production observability.
//
// Example usage:
//
//	import (
//	    "github.com/felixgeelhaar/fortify/metrics"
//	    "github.com/prometheus/client_golang/prometheus"
//	)
//
//	// Register Fortify metrics with default Prometheus registry
//	metrics.MustRegister(prometheus.DefaultRegisterer)
//
//	// Or create a custom registry
//	registry := prometheus.NewRegistry()
//	collector := metrics.NewCollector()
//	registry.MustRegister(collector)
package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector collects Fortify metrics for Prometheus.
type Collector struct {
	// Circuit Breaker metrics
	circuitBreakerState      *prometheus.GaugeVec
	circuitBreakerRequests   *prometheus.CounterVec
	circuitBreakerFailures   *prometheus.CounterVec
	circuitBreakerSuccesses  *prometheus.CounterVec
	circuitBreakerStateTotal *prometheus.CounterVec

	// Retry metrics
	retryAttempts        *prometheus.HistogramVec
	retrySuccesses       *prometheus.CounterVec
	retryFailures        *prometheus.CounterVec
	retryDuration        *prometheus.HistogramVec

	// Rate Limiter metrics
	rateLimitAllowed     *prometheus.CounterVec
	rateLimitDenied      *prometheus.CounterVec
	rateLimitWaitTime    *prometheus.HistogramVec

	// Timeout metrics
	timeoutExecutions    *prometheus.CounterVec
	timeoutExceeded      *prometheus.CounterVec
	timeoutDuration      *prometheus.HistogramVec

	// Bulkhead metrics
	bulkheadActive       *prometheus.GaugeVec
	bulkheadQueued       *prometheus.GaugeVec
	bulkheadRejected     *prometheus.CounterVec
	bulkheadSuccesses    *prometheus.CounterVec
	bulkheadFailures     *prometheus.CounterVec
	bulkheadDuration     *prometheus.HistogramVec

	mu sync.RWMutex
}

var (
	defaultCollector *Collector
	once             sync.Once
)

// DefaultCollector returns the default global metrics collector.
func DefaultCollector() *Collector {
	once.Do(func() {
		defaultCollector = NewCollector()
	})
	return defaultCollector
}

// MustRegister registers the default collector with the given registerer.
// Panics if registration fails.
func MustRegister(registerer prometheus.Registerer) {
	registerer.MustRegister(DefaultCollector())
}

// NewCollector creates a new metrics collector with all Fortify metrics.
func NewCollector() *Collector {
	return &Collector{
		// Circuit Breaker metrics
		circuitBreakerState: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "fortify_circuit_breaker_state",
				Help: "Current state of circuit breaker (0=closed, 1=open, 2=half-open)",
			},
			[]string{"name"},
		),
		circuitBreakerRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_circuit_breaker_requests_total",
				Help: "Total number of requests through circuit breaker",
			},
			[]string{"name", "state"},
		),
		circuitBreakerFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_circuit_breaker_failures_total",
				Help: "Total number of failed requests",
			},
			[]string{"name"},
		),
		circuitBreakerSuccesses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_circuit_breaker_successes_total",
				Help: "Total number of successful requests",
			},
			[]string{"name"},
		),
		circuitBreakerStateTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_circuit_breaker_state_changes_total",
				Help: "Total number of state changes",
			},
			[]string{"name", "from", "to"},
		),

		// Retry metrics
		retryAttempts: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "fortify_retry_attempts",
				Help:    "Number of retry attempts made",
				Buckets: prometheus.LinearBuckets(1, 1, 10),
			},
			[]string{"name"},
		),
		retrySuccesses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_retry_successes_total",
				Help: "Total number of successful retries",
			},
			[]string{"name"},
		),
		retryFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_retry_failures_total",
				Help: "Total number of failed retries",
			},
			[]string{"name"},
		),
		retryDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "fortify_retry_duration_seconds",
				Help:    "Duration of retry operations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"name"},
		),

		// Rate Limiter metrics
		rateLimitAllowed: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_rate_limit_allowed_total",
				Help: "Total number of allowed requests",
			},
			[]string{"name", "key"},
		),
		rateLimitDenied: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_rate_limit_denied_total",
				Help: "Total number of denied requests",
			},
			[]string{"name", "key"},
		),
		rateLimitWaitTime: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "fortify_rate_limit_wait_seconds",
				Help:    "Time spent waiting for rate limit",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"name", "key"},
		),

		// Timeout metrics
		timeoutExecutions: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_timeout_executions_total",
				Help: "Total number of timeout executions",
			},
			[]string{"name"},
		),
		timeoutExceeded: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_timeout_exceeded_total",
				Help: "Total number of exceeded timeouts",
			},
			[]string{"name"},
		),
		timeoutDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "fortify_timeout_duration_seconds",
				Help:    "Duration of timeout operations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"name", "exceeded"},
		),

		// Bulkhead metrics
		bulkheadActive: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "fortify_bulkhead_active",
				Help: "Current number of active requests",
			},
			[]string{"name"},
		),
		bulkheadQueued: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "fortify_bulkhead_queued",
				Help: "Current number of queued requests",
			},
			[]string{"name"},
		),
		bulkheadRejected: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_bulkhead_rejected_total",
				Help: "Total number of rejected requests",
			},
			[]string{"name"},
		),
		bulkheadSuccesses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_bulkhead_successes_total",
				Help: "Total number of successful requests",
			},
			[]string{"name"},
		),
		bulkheadFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fortify_bulkhead_failures_total",
				Help: "Total number of failed requests",
			},
			[]string{"name"},
		),
		bulkheadDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "fortify_bulkhead_duration_seconds",
				Help:    "Duration of bulkhead operations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"name"},
		),
	}
}

// Describe implements prometheus.Collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.circuitBreakerState.Describe(ch)
	c.circuitBreakerRequests.Describe(ch)
	c.circuitBreakerFailures.Describe(ch)
	c.circuitBreakerSuccesses.Describe(ch)
	c.circuitBreakerStateTotal.Describe(ch)

	c.retryAttempts.Describe(ch)
	c.retrySuccesses.Describe(ch)
	c.retryFailures.Describe(ch)
	c.retryDuration.Describe(ch)

	c.rateLimitAllowed.Describe(ch)
	c.rateLimitDenied.Describe(ch)
	c.rateLimitWaitTime.Describe(ch)

	c.timeoutExecutions.Describe(ch)
	c.timeoutExceeded.Describe(ch)
	c.timeoutDuration.Describe(ch)

	c.bulkheadActive.Describe(ch)
	c.bulkheadQueued.Describe(ch)
	c.bulkheadRejected.Describe(ch)
	c.bulkheadSuccesses.Describe(ch)
	c.bulkheadFailures.Describe(ch)
	c.bulkheadDuration.Describe(ch)
}

// Collect implements prometheus.Collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.circuitBreakerState.Collect(ch)
	c.circuitBreakerRequests.Collect(ch)
	c.circuitBreakerFailures.Collect(ch)
	c.circuitBreakerSuccesses.Collect(ch)
	c.circuitBreakerStateTotal.Collect(ch)

	c.retryAttempts.Collect(ch)
	c.retrySuccesses.Collect(ch)
	c.retryFailures.Collect(ch)
	c.retryDuration.Collect(ch)

	c.rateLimitAllowed.Collect(ch)
	c.rateLimitDenied.Collect(ch)
	c.rateLimitWaitTime.Collect(ch)

	c.timeoutExecutions.Collect(ch)
	c.timeoutExceeded.Collect(ch)
	c.timeoutDuration.Collect(ch)

	c.bulkheadActive.Collect(ch)
	c.bulkheadQueued.Collect(ch)
	c.bulkheadRejected.Collect(ch)
	c.bulkheadSuccesses.Collect(ch)
	c.bulkheadFailures.Collect(ch)
	c.bulkheadDuration.Collect(ch)
}

// RecordCircuitBreakerState records the current circuit breaker state.
func (c *Collector) RecordCircuitBreakerState(name string, state float64) {
	c.circuitBreakerState.WithLabelValues(name).Set(state)
}

// RecordCircuitBreakerRequest records a circuit breaker request.
func (c *Collector) RecordCircuitBreakerRequest(name, state string) {
	c.circuitBreakerRequests.WithLabelValues(name, state).Inc()
}

// RecordCircuitBreakerFailure records a circuit breaker failure.
func (c *Collector) RecordCircuitBreakerFailure(name string) {
	c.circuitBreakerFailures.WithLabelValues(name).Inc()
}

// RecordCircuitBreakerSuccess records a circuit breaker success.
func (c *Collector) RecordCircuitBreakerSuccess(name string) {
	c.circuitBreakerSuccesses.WithLabelValues(name).Inc()
}

// RecordCircuitBreakerStateChange records a circuit breaker state change.
func (c *Collector) RecordCircuitBreakerStateChange(name, from, to string) {
	c.circuitBreakerStateTotal.WithLabelValues(name, from, to).Inc()
}

// RecordRetryAttempts records the number of retry attempts.
func (c *Collector) RecordRetryAttempts(name string, attempts float64) {
	c.retryAttempts.WithLabelValues(name).Observe(attempts)
}

// RecordRetrySuccess records a successful retry.
func (c *Collector) RecordRetrySuccess(name string) {
	c.retrySuccesses.WithLabelValues(name).Inc()
}

// RecordRetryFailure records a failed retry.
func (c *Collector) RecordRetryFailure(name string) {
	c.retryFailures.WithLabelValues(name).Inc()
}

// RecordRetryDuration records the duration of a retry operation.
func (c *Collector) RecordRetryDuration(name string, seconds float64) {
	c.retryDuration.WithLabelValues(name).Observe(seconds)
}

// RecordRateLimitAllowed records an allowed rate limit request.
func (c *Collector) RecordRateLimitAllowed(name, key string) {
	c.rateLimitAllowed.WithLabelValues(name, key).Inc()
}

// RecordRateLimitDenied records a denied rate limit request.
func (c *Collector) RecordRateLimitDenied(name, key string) {
	c.rateLimitDenied.WithLabelValues(name, key).Inc()
}

// RecordRateLimitWaitTime records the time spent waiting for rate limit.
func (c *Collector) RecordRateLimitWaitTime(name, key string, seconds float64) {
	c.rateLimitWaitTime.WithLabelValues(name, key).Observe(seconds)
}

// RecordTimeoutExecution records a timeout execution.
func (c *Collector) RecordTimeoutExecution(name string) {
	c.timeoutExecutions.WithLabelValues(name).Inc()
}

// RecordTimeoutExceeded records an exceeded timeout.
func (c *Collector) RecordTimeoutExceeded(name string) {
	c.timeoutExceeded.WithLabelValues(name).Inc()
}

// RecordTimeoutDuration records the duration of a timeout operation.
func (c *Collector) RecordTimeoutDuration(name string, exceeded bool, seconds float64) {
	exceededStr := "false"
	if exceeded {
		exceededStr = "true"
	}
	c.timeoutDuration.WithLabelValues(name, exceededStr).Observe(seconds)
}

// RecordBulkheadActive records the current number of active requests.
func (c *Collector) RecordBulkheadActive(name string, count float64) {
	c.bulkheadActive.WithLabelValues(name).Set(count)
}

// RecordBulkheadQueued records the current number of queued requests.
func (c *Collector) RecordBulkheadQueued(name string, count float64) {
	c.bulkheadQueued.WithLabelValues(name).Set(count)
}

// RecordBulkheadRejected records a rejected bulkhead request.
func (c *Collector) RecordBulkheadRejected(name string) {
	c.bulkheadRejected.WithLabelValues(name).Inc()
}

// RecordBulkheadSuccess records a successful bulkhead request.
func (c *Collector) RecordBulkheadSuccess(name string) {
	c.bulkheadSuccesses.WithLabelValues(name).Inc()
}

// RecordBulkheadFailure records a failed bulkhead request.
func (c *Collector) RecordBulkheadFailure(name string) {
	c.bulkheadFailures.WithLabelValues(name).Inc()
}

// RecordBulkheadDuration records the duration of a bulkhead operation.
func (c *Collector) RecordBulkheadDuration(name string, seconds float64) {
	c.bulkheadDuration.WithLabelValues(name).Observe(seconds)
}