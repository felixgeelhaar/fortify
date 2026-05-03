package adaptive

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

func TestNew_DefaultsApplied(t *testing.T) {
	l := New[int](Config{}).(*limiter[int])
	if l.config.InitialLimit != defaultInitialLimit {
		t.Errorf("InitialLimit = %d, want %d", l.config.InitialLimit, defaultInitialLimit)
	}
	if l.config.MinLimit != defaultMinLimit {
		t.Errorf("MinLimit = %d, want %d", l.config.MinLimit, defaultMinLimit)
	}
	if l.config.MaxLimit != defaultMaxLimit {
		t.Errorf("MaxLimit = %d, want %d", l.config.MaxLimit, defaultMaxLimit)
	}
	if l.Limit() != defaultInitialLimit {
		t.Errorf("Limit() = %d, want %d", l.Limit(), defaultInitialLimit)
	}
}

func TestExecute_AdmitsUnderLimit(t *testing.T) {
	l := New[int](Config{InitialLimit: 5, MinLimit: 1, MaxLimit: 10})

	got, err := l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if got != 42 {
		t.Fatalf("value = %d, want 42", got)
	}
	if l.InFlight() != 0 {
		t.Fatalf("InFlight = %d, want 0", l.InFlight())
	}
}

func TestExecute_RejectsAtLimit(t *testing.T) {
	l := New[int](Config{InitialLimit: 1, MinLimit: 1, MaxLimit: 10})

	// Block one slot.
	gate := make(chan struct{})
	released := make(chan struct{})
	go func() {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			close(gate)
			<-released
			return 0, nil
		})
	}()
	<-gate

	// Second call should reject with ErrLimitExceeded.
	_, err := l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		t.Fatal("should not run")
		return 0, nil
	})
	if !errors.Is(err, ErrLimitExceeded) {
		t.Fatalf("err = %v, want ErrLimitExceeded", err)
	}

	close(released)
}

func TestExecute_FailureHalvesLimit(t *testing.T) {
	l := New[int](Config{InitialLimit: 10, MinLimit: 1, MaxLimit: 100})

	wantErr := errors.New("boom")
	_, err := l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if l.Limit() != 5 {
		t.Fatalf("Limit after failure = %d, want 5 (10/2)", l.Limit())
	}
}

func TestExecute_FailureFloorsAtMinLimit(t *testing.T) {
	l := New[int](Config{InitialLimit: 4, MinLimit: 3, MaxLimit: 10})

	for i := 0; i < 5; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("fail")
		})
	}
	if l.Limit() != 3 {
		t.Fatalf("Limit = %d, want floored at MinLimit=3", l.Limit())
	}
}

func TestExecute_AdditiveIncreaseOnSuccessThreshold(t *testing.T) {
	l := New[int](Config{
		InitialLimit:     5,
		MinLimit:         1,
		MaxLimit:         100,
		SuccessThreshold: 3,
	})

	for i := 0; i < 3; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	if l.Limit() != 6 {
		t.Fatalf("Limit after 3 successes (threshold=3) = %d, want 6", l.Limit())
	}

	// 2 more successes shouldn't yet trip another increase.
	for i := 0; i < 2; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	if l.Limit() != 6 {
		t.Fatalf("Limit after 5 successes = %d, want 6 (next at 6 successes)", l.Limit())
	}

	// One more triggers the next increase.
	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, nil
	})
	if l.Limit() != 7 {
		t.Fatalf("Limit after 6 successes = %d, want 7", l.Limit())
	}
}

func TestExecute_IncreaseCappedAtMaxLimit(t *testing.T) {
	l := New[int](Config{
		InitialLimit:     5,
		MinLimit:         1,
		MaxLimit:         5,
		SuccessThreshold: 1,
	})

	for i := 0; i < 50; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	if l.Limit() != 5 {
		t.Fatalf("Limit = %d, want capped at 5", l.Limit())
	}
}

func TestExecute_FailureResetsSuccessCounter(t *testing.T) {
	l := New[int](Config{
		InitialLimit:     5,
		MinLimit:         1,
		MaxLimit:         100,
		SuccessThreshold: 5,
	})

	// 4 successes
	for i := 0; i < 4; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	// 1 failure resets counter
	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, errors.New("fail")
	})
	// Now would-be-5th success should NOT trigger increase yet (counter reset)
	for i := 0; i < 4; i++ {
		_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
	// Limit should be at 2 (was 5, halved on failure, no successes have raised it)
	if l.Limit() != 2 {
		t.Fatalf("Limit = %d, want 2 (success counter reset by failure)", l.Limit())
	}
}

func TestExecute_OnLimitChangeFires(t *testing.T) {
	var calls atomic.Int32
	l := New[int](Config{
		InitialLimit:     10,
		MinLimit:         1,
		MaxLimit:         100,
		SuccessThreshold: 1,
		OnLimitChange: func(old, new int) {
			calls.Add(1)
		},
	})

	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, nil
	})
	if calls.Load() != 1 {
		t.Fatalf("OnLimitChange calls = %d, want 1 (success raised limit)", calls.Load())
	}

	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, errors.New("fail")
	})
	if calls.Load() != 2 {
		t.Fatalf("OnLimitChange calls = %d, want 2 (failure halved limit)", calls.Load())
	}
}

func TestExecute_OnLimitChangePanicRecovered(t *testing.T) {
	l := New[int](Config{
		InitialLimit:     5,
		MinLimit:         1,
		MaxLimit:         100,
		SuccessThreshold: 1,
		OnLimitChange:    func(old, new int) { panic("boom") },
	})

	// Should not crash
	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, nil
	})
}

func TestExecute_PreCancelledContext(t *testing.T) {
	l := New[int](Config{InitialLimit: 5, MinLimit: 1, MaxLimit: 10})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var ran atomic.Bool
	_, err := l.Execute(ctx, func(ctx context.Context) (int, error) {
		ran.Store(true)
		return 0, nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want Canceled", err)
	}
	if ran.Load() {
		t.Fatal("operation ran despite cancelled ctx")
	}
}

func TestExecute_ConcurrentRespectsLimit(t *testing.T) {
	const limit = 5
	l := New[int](Config{InitialLimit: limit, MinLimit: 1, MaxLimit: limit, SuccessThreshold: 1000})

	var concurrent atomic.Int32
	var maxObserved atomic.Int32
	gate := make(chan struct{})

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
				cur := concurrent.Add(1)
				for {
					m := maxObserved.Load()
					if cur <= m || maxObserved.CompareAndSwap(m, cur) {
						break
					}
				}
				<-gate
				concurrent.Add(-1)
				return 0, nil
			})
		}()
	}

	// Let some operations get in-flight, then release.
	close(gate)
	wg.Wait()

	if m := maxObserved.Load(); m > int32(limit) {
		t.Fatalf("concurrent in-flight = %d, exceeds limit %d", m, limit)
	}
}
