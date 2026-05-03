package ratelimit

import (
	"context"
	"testing"
	"time"

	"pgregory.net/rapid"
)

// TestProperty_TokensWithinBucket asserts that across any sequence of
// Allow/Take calls, the bucket's stored tokens stay in [0, Burst].
func TestProperty_TokensWithinBucket(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		burst := rapid.IntRange(1, 100).Draw(t, "burst")
		rl := New(Config{
			Rate:                burst,
			Burst:               burst,
			Interval:            time.Second,
			MaxTokensPerRequest: burst,
		})
		defer func() { _ = rl.Close() }()

		ctx := context.Background()
		key := rapid.StringMatching(`[a-z]{1,8}`).Draw(t, "key")
		ops := rapid.SliceOfN(rapid.IntRange(0, 4), 0, 30).Draw(t, "ops")
		for _, op := range ops {
			switch op {
			case 0:
				rl.Allow(ctx, key)
			case 1:
				rl.Take(ctx, key, 1)
			case 2:
				rl.Take(ctx, key, rapid.IntRange(1, burst).Draw(t, "n"))
			case 3:
				rl.Allow(ctx, key)
			case 4:
				rl.Take(ctx, key, burst+1) // beyond MaxTokensPerRequest; must be denied
			}
		}

		// Sample bucket directly via Store.Get (using internal store accessor).
		concrete := rl.(*rateLimiter)
		state, err := concrete.config.Store.Get(ctx, key)
		if err != nil {
			t.Fatalf("Store.Get err = %v", err)
		}
		if state == nil {
			return
		}
		if state.Tokens < 0 {
			t.Fatalf("tokens went negative: %f", state.Tokens)
		}
		if state.Tokens > float64(burst)+0.01 {
			t.Fatalf("tokens %f > burst %d", state.Tokens, burst)
		}
	})
}

// TestProperty_BucketCountBoundedByMaxKeys asserts that the in-memory store
// never holds more buckets than its WithMaxKeys cap.
func TestProperty_BucketCountBoundedByMaxKeys(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		maxKeys := rapid.IntRange(1, 50).Draw(t, "maxKeys")
		store := NewMemoryStoreWithOptions(WithMaxKeys(maxKeys))
		defer func() { _ = store.Close() }()

		rl := New(Config{
			Rate:     10,
			Burst:    10,
			Interval: time.Second,
			Store:    store,
			FailOpen: false,
		})
		defer func() { _ = rl.Close() }()

		ctx := context.Background()
		nKeys := rapid.IntRange(0, maxKeys*3).Draw(t, "nKeys")
		for i := 0; i < nKeys; i++ {
			key := rapid.StringMatching(`[a-z]{1,16}`).Draw(t, "key")
			rl.Allow(ctx, key)
		}
		if got := store.BucketCount(); got > int64(maxKeys) {
			t.Fatalf("BucketCount %d > MaxKeys %d", got, maxKeys)
		}
	})
}

// TestProperty_AllowMatchesTakeOne asserts Allow(key) and Take(key, 1)
// produce the same admit/deny decision for the same bucket state.
func TestProperty_AllowMatchesTakeOne(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Use rate=0 (bumped to default 100 by setDefaults? no, defaults
		// to DefaultRate=100). Use Burst=1 to make decisions deterministic.
		rl := New(Config{
			Rate:                1,
			Burst:               1,
			Interval:            time.Hour, // refill effectively never
			MaxTokensPerRequest: 1,
		})
		defer func() { _ = rl.Close() }()

		ctx := context.Background()
		k1 := "allowKey"
		k2 := "takeKey"

		decisions := rapid.SliceOfN(rapid.Bool(), 1, 5).Draw(t, "calls")
		_ = decisions // use to vary call count
		for range decisions {
			a := rl.Allow(ctx, k1)
			b := rl.Take(ctx, k2, 1)
			if a != b {
				t.Fatalf("Allow=%v vs Take(.,1)=%v diverged", a, b)
			}
		}
	})
}

// TestProperty_ResetEmptiesStore asserts Reset removes all buckets.
func TestProperty_ResetEmptiesStore(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		store := NewMemoryStore()
		defer func() { _ = store.Close() }()

		rl := New(Config{
			Rate:     10,
			Burst:    10,
			Interval: time.Second,
			Store:    store,
		})
		defer func() { _ = rl.Close() }()

		ctx := context.Background()
		nKeys := rapid.IntRange(0, 30).Draw(t, "nKeys")
		for i := 0; i < nKeys; i++ {
			rl.Allow(ctx, rapid.StringMatching(`[a-z]{1,8}`).Draw(t, "k"))
		}

		if err := rl.Reset(ctx); err != nil {
			t.Fatalf("Reset err = %v", err)
		}
		if got := store.BucketCount(); got != 0 {
			t.Fatalf("BucketCount after Reset = %d, want 0", got)
		}
	})
}
