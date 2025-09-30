// Package main demonstrates HTTP middleware usage.
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	fortifyhttp "github.com/felixgeelhaar/fortify/http"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

func main() {
	// Create resilience patterns
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second * 10,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: func(from, to circuitbreaker.State) {
			log.Printf("Circuit breaker: %s -> %s", from, to)
		},
	})

	rl := ratelimit.New(ratelimit.Config{
		Rate:     10, // 10 requests
		Burst:    20, // burst of 20
		Interval: time.Second,
	})

	tm := timeout.New[*http.Response](timeout.Config{
		DefaultTimeout: time.Second * 30,
		OnTimeout: func() {
			log.Println("Request timed out")
		},
	})

	// Create handlers with middleware
	mux := http.NewServeMux()

	// Protected endpoint with all patterns
	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate occasional failures and slow responses
		if rand.Intn(10) < 2 {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Simulate variable response time
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Protected resource accessed at %s\n", time.Now().Format(time.RFC3339))
	})

	// Apply middleware: rate limit -> timeout -> circuit breaker -> handler
	protectedWithMiddleware := fortifyhttp.RateLimit(rl, fortifyhttp.KeyFromIP)(
		fortifyhttp.Timeout(tm, 2*time.Second)(
			fortifyhttp.CircuitBreaker(cb)(protectedHandler),
		),
	)

	mux.Handle("/protected", protectedWithMiddleware)

	// Rate limited endpoint by API key
	apiHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API accessed with key: %s at %s\n", apiKey, time.Now().Format(time.RFC3339))
	})

	apiWithRateLimit := fortifyhttp.RateLimit(rl, fortifyhttp.KeyFromHeader("X-API-Key"))(apiHandler)
	mux.Handle("/api", apiWithRateLimit)

	// Health check endpoint (no middleware)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// Slow endpoint with timeout
	slowHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			// Context was cancelled (timeout occurred)
			return
		case <-time.After(5 * time.Second):
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Slow operation completed")
		}
	})

	slowWithTimeout := fortifyhttp.Timeout(tm, 1*time.Second)(slowHandler)
	mux.Handle("/slow", slowWithTimeout)

	// Start server
	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	log.Println("Endpoints:")
	log.Println("  GET /protected  - Protected with circuit breaker, timeout, and rate limit")
	log.Println("  GET /api        - Rate limited by X-API-Key header")
	log.Println("  GET /health     - Health check (no middleware)")
	log.Println("  GET /slow       - Slow endpoint with 1s timeout")
	log.Println("\nExample requests:")
	log.Println("  curl http://localhost:8080/protected")
	log.Println("  curl -H 'X-API-Key: test-key' http://localhost:8080/api")
	log.Println("  curl http://localhost:8080/slow")

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

// Graceful shutdown example
func gracefulShutdown(server *http.Server) {
	// Wait for interrupt signal
	// (in a real application, you'd use signal.Notify)
	time.Sleep(time.Minute)

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
