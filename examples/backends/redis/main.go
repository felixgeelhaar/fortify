// Package main demonstrates distributed rate limiting using Redis backend.
//
// This example shows how to use Fortify's Redis-backed rate limiter in a
// production-like HTTP API scenario with multiple application instances
// sharing rate limits across a distributed system.
//
// Run with Docker Compose:
//
//	docker-compose up
//
// Or run manually:
//
//	# Start Redis
//	docker run -p 6379:6379 redis:7-alpine
//
//	# Run the server
//	go run main.go
//
// Test the rate limiter:
//
//	# Make requests (first 10 should succeed, 11th should be rate limited)
//	for i in {1..15}; do curl -H "X-User-ID: user-123" http://localhost:8080/api/data; done
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/felixgeelhaar/fortify/ratelimit"
	redisrl "github.com/felixgeelhaar/fortify/backends/redis"
)

const (
	defaultPort     = "8080"
	defaultRedisURL = "localhost:6379"

	// Rate limit configuration
	requestsPerMinute = 10
	burstSize         = 15
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Get configuration from environment
	port := getEnv("PORT", defaultPort)
	redisURL := getEnv("REDIS_URL", defaultRedisURL)

	// Create Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:         redisURL,
		Password:     getEnv("REDIS_PASSWORD", ""),
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Verify Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", redisURL, err)
	}

	slog.Info("Connected to Redis", "addr", redisURL)

	// Create distributed rate limiter
	limiter, err := redisrl.New(redisrl.Config{
		Client:    redisClient,
		Rate:      requestsPerMinute,
		Burst:     burstSize,
		Interval:  time.Minute,
		KeyPrefix: "example:ratelimit:",
		Logger:    logger,
		OnLimit: func(key string) {
			slog.Warn("Rate limit exceeded",
				slog.String("key", key),
				slog.Int("rate", requestsPerMinute),
				slog.Int("burst", burstSize),
			)
		},
		BucketTTL:       time.Hour,
		FallbackOnError: false, // Fail-closed on Redis errors
	})
	if err != nil {
		log.Fatalf("Failed to create rate limiter: %v", err)
	}

	// Create HTTP server
	mux := http.NewServeMux()

	// Health check endpoint (not rate limited)
	mux.HandleFunc("/health", healthHandler(redisClient))

	// API endpoint (rate limited)
	mux.HandleFunc("/api/data", rateLimitMiddleware(limiter)(apiHandler))

	// Status endpoint showing rate limit info
	mux.HandleFunc("/api/status", rateLimitMiddleware(limiter)(statusHandler(limiter)))

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		slog.Info("Starting HTTP server",
			slog.String("port", port),
			slog.Int("rate_limit", requestsPerMinute),
			slog.Int("burst", burstSize),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	slog.Info("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	// Close Redis connection
	if err := redisClient.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}

	slog.Info("Server stopped")
}

// healthHandler returns a health check endpoint.
func healthHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Check Redis connectivity
		if err := redisClient.Ping(ctx).Err(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "unhealthy",
				"reason": "redis connection failed",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
		})
	}
}

// apiHandler is a sample API endpoint that returns data.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Request successful",
		"user_id":   userID,
		"timestamp": time.Now().Unix(),
		"data": map[string]string{
			"foo": "bar",
			"baz": "qux",
		},
	})
}

// statusHandler returns rate limit status information.
func statusHandler(limiter ratelimit.RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getUserID(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"user_id": userID,
			"limits": map[string]int{
				"requests_per_minute": requestsPerMinute,
				"burst":               burstSize,
			},
			"message": "Rate limit status",
		})
	}
}

// rateLimitMiddleware wraps an HTTP handler with rate limiting.
func rateLimitMiddleware(limiter ratelimit.RateLimiter) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Extract user ID from request
			userID := getUserID(r)

			// Check rate limit
			if !limiter.Allow(r.Context(), userID) {
				// Rate limited - return 429
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", requestsPerMinute))
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("Retry-After", "60") // 1 minute
				w.WriteHeader(http.StatusTooManyRequests)

				json.NewEncoder(w).Encode(map[string]string{
					"error":   "rate_limit_exceeded",
					"message": fmt.Sprintf("Rate limit exceeded. Maximum %d requests per minute.", requestsPerMinute),
				})

				slog.Warn("Request rate limited",
					slog.String("user_id", userID),
					slog.String("path", r.URL.Path),
					slog.String("method", r.Method),
				)
				return
			}

			// Set rate limit headers
			w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", requestsPerMinute))

			// Process request
			next(w, r)
		}
	}
}

// getUserID extracts the user ID from the request.
// In a real application, this might come from JWT claims, session, etc.
func getUserID(r *http.Request) string {
	// Try X-User-ID header first
	if userID := r.Header.Get("X-User-ID"); userID != "" {
		return userID
	}

	// Fall back to IP address
	return r.RemoteAddr
}

// getEnv gets an environment variable with a default fallback.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
