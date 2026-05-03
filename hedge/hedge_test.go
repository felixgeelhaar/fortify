package hedge

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestHedge_PrimarySucceedsImmediately_NoHedgeFired(t *testing.T) {
	h := New[int](Config{MaxAttempts: 3, HedgeDelay: 100 * time.Millisecond})

	var calls atomic.Int32
	got, err := h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		calls.Add(1)
		return 42, nil
	})

	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if got != 42 {
		t.Fatalf("value = %d, want 42", got)
	}
	if c := calls.Load(); c != 1 {
		t.Fatalf("calls = %d, want 1 (no hedges)", c)
	}
}

func TestHedge_FiresHedgeWhenPrimarySlow(t *testing.T) {
	h := New[int](Config{MaxAttempts: 2, HedgeDelay: 10 * time.Millisecond})

	var calls atomic.Int32
	got, err := h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		n := calls.Add(1)
		if n == 1 {
			// Primary is slow but still wins-or-loses. Block until either
			// the hedge fires (so calls.Load() advances) or ctx cancels.
			for ctx.Err() == nil && calls.Load() < 2 {
				time.Sleep(time.Millisecond)
			}
			return 1, ctx.Err()
		}
		// Hedge attempt: succeed quickly.
		return 2, nil
	})

	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if got != 2 {
		t.Fatalf("value = %d, want 2 (hedge winner)", got)
	}
	if c := calls.Load(); c != 2 {
		t.Fatalf("calls = %d, want 2 (primary + 1 hedge)", c)
	}
}

func TestHedge_AllAttemptsFail_ReturnsFirstError(t *testing.T) {
	wantErr := errors.New("attempt failed")
	h := New[int](Config{MaxAttempts: 3, HedgeDelay: 5 * time.Millisecond})

	got, err := h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, wantErr
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if got != 0 {
		t.Fatalf("value = %d, want 0", got)
	}
}

func TestHedge_ContextCancelStopsHedging(t *testing.T) {
	h := New[int](Config{MaxAttempts: 5, HedgeDelay: 1 * time.Millisecond})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := h.Execute(ctx, func(ctx context.Context) (int, error) {
		<-ctx.Done()
		return 0, ctx.Err()
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want DeadlineExceeded", err)
	}
}

func TestHedge_CapsMaxAttempts(t *testing.T) {
	h := New[int](Config{MaxAttempts: 1000, HedgeDelay: time.Microsecond})

	concrete, ok := h.(*hedge[int])
	if !ok {
		t.Fatal("type assertion failed")
	}
	if concrete.config.MaxAttempts != maxAttemptsCap {
		t.Fatalf("MaxAttempts = %d, want capped at %d", concrete.config.MaxAttempts, maxAttemptsCap)
	}
}

func TestHedge_OnHedgeCallback(t *testing.T) {
	var hedgeCount atomic.Int32
	var calls atomic.Int32
	h := New[int](Config{
		MaxAttempts: 2,
		HedgeDelay:  5 * time.Millisecond,
		OnHedge: func(attempt int) {
			hedgeCount.Add(1)
		},
	})

	_, err := h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		n := calls.Add(1)
		if n == 1 {
			// Primary: wait until hedge fires, then return error.
			for ctx.Err() == nil && calls.Load() < 2 {
				time.Sleep(time.Millisecond)
			}
			return 0, errors.New("primary lost")
		}
		return 7, nil
	})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if c := hedgeCount.Load(); c == 0 {
		t.Fatalf("OnHedge was never called")
	}
}

func TestHedge_OnHedgeCallbackPanicRecovered(t *testing.T) {
	h := New[int](Config{
		MaxAttempts: 2,
		HedgeDelay:  1 * time.Millisecond,
		OnHedge: func(attempt int) {
			panic("boom")
		},
	})

	// Should not crash the test process.
	got, err := h.Execute(context.Background(), func(ctx context.Context) (int, error) {
		time.Sleep(5 * time.Millisecond)
		return 7, nil
	})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if got != 7 {
		t.Fatalf("value = %d, want 7", got)
	}
}

func TestHedge_DefaultsApplied(t *testing.T) {
	h := New[int](Config{})
	concrete, ok := h.(*hedge[int])
	if !ok {
		t.Fatal("type assertion failed")
	}
	if concrete.config.MaxAttempts != defaultMaxAttempts {
		t.Errorf("MaxAttempts default = %d, want %d", concrete.config.MaxAttempts, defaultMaxAttempts)
	}
	if concrete.config.HedgeDelay != defaultHedgeDelay {
		t.Errorf("HedgeDelay default = %v, want %v", concrete.config.HedgeDelay, defaultHedgeDelay)
	}
}

func TestHedge_PreCancelledContext(t *testing.T) {
	h := New[int](Config{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var calls atomic.Int32
	_, err := h.Execute(ctx, func(ctx context.Context) (int, error) {
		calls.Add(1)
		return 0, nil
	})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want Canceled", err)
	}
	if c := calls.Load(); c != 0 {
		t.Fatalf("calls = %d, want 0 (ctx pre-cancelled)", c)
	}
}
