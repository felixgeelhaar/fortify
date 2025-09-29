package fallback_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/felixgeelhaar/fortify/fallback"
)

// Example demonstrates basic fallback pattern usage.
func Example() {
	// Create a fallback that returns cached data when API fails
	fb := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			// Return cached or default data
			return "cached data", nil
		},
	})

	// Execute operation with fallback
	result, err := fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		// Simulate API call failure
		return "", errors.New("API unavailable")
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: cached data
}

// Example_cache demonstrates using fallback with a cache.
func Example_cache() {
	// Simulate a cache
	cache := map[string]string{
		"user:123": `{"id": "123", "name": "Alice"}`,
	}

	fb := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			// Try to get data from cache
			if data, ok := cache["user:123"]; ok {
				return data, nil
			}
			return "", errors.New("not found in cache")
		},
	})

	// Try to fetch from database, fall back to cache on failure
	result, err := fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		// Simulate database failure
		return "", errors.New("database unavailable")
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User data: %s\n", result)
	// Output: User data: {"id": "123", "name": "Alice"}
}

// Example_defaultValue demonstrates returning a default value on failure.
func Example_defaultValue() {
	type UserPreferences struct {
		Theme    string
		Language string
	}

	fb := fallback.New[UserPreferences](fallback.Config[UserPreferences]{
		Fallback: func(ctx context.Context, err error) (UserPreferences, error) {
			// Return default preferences
			return UserPreferences{
				Theme:    "light",
				Language: "en",
			}, nil
		},
	})

	// Try to load user preferences, use defaults on failure
	prefs, err := fb.Execute(context.Background(), func(ctx context.Context) (UserPreferences, error) {
		// Simulate loading failure
		return UserPreferences{}, errors.New("user preferences not found")
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Theme: %s, Language: %s\n", prefs.Theme, prefs.Language)
	// Output: Theme: light, Language: en
}

// Example_selective demonstrates conditional fallback execution.
func Example_selective() {
	var temporaryError = errors.New("temporary error")
	var permanentError = errors.New("permanent error")

	fb := fallback.New[int](fallback.Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 100, nil
		},
		ShouldFallback: func(err error) bool {
			// Only fallback for temporary errors
			return err == temporaryError
		},
	})

	// Temporary error - triggers fallback
	result, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, temporaryError
	})
	fmt.Printf("Temporary error - Result: %d, Error: %v\n", result, err)

	// Permanent error - does not trigger fallback
	result, err = fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, permanentError
	})
	fmt.Printf("Permanent error - Result: %d, Error: %v\n", result, err)
	// Output:
	// Temporary error - Result: 100, Error: <nil>
	// Permanent error - Result: 0, Error: permanent error
}

// Example_callbacks demonstrates using callbacks for monitoring.
func Example_callbacks() {
	var fallbackCount int

	fb := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			return "fallback response", nil
		},
		OnFallback: func(err error) {
			fallbackCount++
			fmt.Printf("Fallback triggered: %v\n", err)
		},
		OnSuccess: func() {
			fmt.Println("Primary succeeded")
		},
	})

	// Successful execution
	_, _ = fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "success", nil
	})

	// Failed execution with fallback
	_, _ = fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", errors.New("operation failed")
	})

	fmt.Printf("Total fallbacks: %d\n", fallbackCount)
	// Output:
	// Primary succeeded
	// Fallback triggered: operation failed
	// Total fallbacks: 1
}

// Example_http demonstrates HTTP client with fallback to stale cache.
func Example_http() {
	// Simulate a stale cache
	staleCache := map[string]string{
		"/api/users": `[{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]`,
	}

	fb := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			// Return stale cache data
			if data, ok := staleCache["/api/users"]; ok {
				return data + " (stale)", nil
			}
			return "", errors.New("no cached data")
		},
	})

	// Simulate HTTP request
	fetchUsers := func(ctx context.Context) (string, error) {
		// Simulate HTTP failure
		return "", errors.New("HTTP 503 Service Unavailable")
	}

	result, err := fb.Execute(context.Background(), fetchUsers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Users: %s\n", result)
	// Output: Users: [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}] (stale)
}

// Example_multiTier demonstrates multi-tier fallback strategy.
func Example_multiTier() {
	// Primary: Database
	// Fallback 1: Redis cache
	// Fallback 2: Default value

	redisCache := map[string]string{
		"config:feature_flags": `{"feature_x": true}`,
	}

	// Second-level fallback (default value)
	secondFallback := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			return `{"feature_x": false}`, nil // Default config
		},
	})

	// First-level fallback (Redis)
	firstFallback := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			// Try Redis cache
			if data, ok := redisCache["config:feature_flags"]; ok {
				return data, nil
			}

			// Redis also failed, use second-level fallback
			return secondFallback.Execute(ctx, func(ctx context.Context) (string, error) {
				return "", errors.New("redis unavailable")
			})
		},
	})

	// Try database first
	config, err := firstFallback.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", errors.New("database unavailable")
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Config: %s\n", config)
	// Output: Config: {"feature_x": true}
}

// Example_timeout demonstrates fallback with timeout handling.
func Example_timeout() {
	fb := fallback.New[string](fallback.Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			return "fast fallback response", nil
		},
		ShouldFallback: func(err error) bool {
			// Fallback on timeout errors
			return errors.Is(err, context.DeadlineExceeded)
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	result, err := fb.Execute(ctx, func(ctx context.Context) (string, error) {
		// Simulate slow operation
		select {
		case <-time.After(200 * time.Millisecond):
			return "slow response", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: fast fallback response
}

// Example_externalService demonstrates graceful degradation with external services.
func Example_externalService() {
	type WeatherData struct {
		Temperature float64
		Condition   string
	}

	// Mock weather service that sometimes fails
	var serviceUnavailable bool = true

	fb := fallback.New[WeatherData](fallback.Config[WeatherData]{
		Fallback: func(ctx context.Context, err error) (WeatherData, error) {
			// Return last known good data or generic message
			return WeatherData{
				Temperature: 20.0,
				Condition:   "Data temporarily unavailable",
			}, nil
		},
	})

	weather, err := fb.Execute(context.Background(), func(ctx context.Context) (WeatherData, error) {
		if serviceUnavailable {
			return WeatherData{}, errors.New("weather service unavailable")
		}

		// Simulate API call
		resp, err := http.Get("https://api.weather.example/current")
		if err != nil {
			return WeatherData{}, err
		}
		defer resp.Body.Close()

		// Parse response...
		return WeatherData{Temperature: 22.5, Condition: "Sunny"}, nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Temperature: %.1f°C, Condition: %s\n", weather.Temperature, weather.Condition)
	// Output: Temperature: 20.0°C, Condition: Data temporarily unavailable
}