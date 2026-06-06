package streamtimeout

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.klarlabs.de/fortify/ferrors"
)

func TestNew_RejectsAllZero(t *testing.T) {
	if _, err := New[int](Config{}); err == nil {
		t.Fatal("expected error for zero config")
	}
}

func TestExecute_FirstByteTimeoutFires(t *testing.T) {
	st, _ := New[int](Config{FirstByteTimeout: 30 * time.Millisecond})
	_, err := st.Execute(context.Background(), func(ctx context.Context, _ Mark) (int, error) {
		<-ctx.Done()
		return 0, ctx.Err()
	})

	var ste *StreamTimeoutError
	if !errors.As(err, &ste) {
		t.Fatalf("expected *StreamTimeoutError, got %v", err)
	}
	if ste.Stage != StageFirstByte {
		t.Errorf("Stage = %s, want %s", ste.Stage, StageFirstByte)
	}
	if !errors.Is(err, ferrors.ErrTimeout) {
		t.Error("err should wrap ferrors.ErrTimeout")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("err should wrap context.DeadlineExceeded")
	}
}

func TestExecute_IdleTimeoutFires(t *testing.T) {
	st, _ := New[int](Config{
		FirstByteTimeout: time.Second,
		IdleTimeout:      40 * time.Millisecond,
	})

	_, err := st.Execute(context.Background(), func(ctx context.Context, mark Mark) (int, error) {
		mark()
		// Now we stop chunking; idle should fire.
		<-ctx.Done()
		return 0, ctx.Err()
	})

	var ste *StreamTimeoutError
	if !errors.As(err, &ste) {
		t.Fatalf("expected *StreamTimeoutError, got %v", err)
	}
	if ste.Stage != StageIdle {
		t.Errorf("Stage = %s, want %s", ste.Stage, StageIdle)
	}
}

func TestExecute_IdleResetsOnMark(t *testing.T) {
	st, _ := New[int](Config{
		IdleTimeout:  60 * time.Millisecond,
		TotalTimeout: 500 * time.Millisecond,
	})

	got, err := st.Execute(context.Background(), func(ctx context.Context, mark Mark) (int, error) {
		// 5 chunks, 30ms apart → never idle for 60ms in a row.
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				return 0, ctx.Err()
			case <-time.After(30 * time.Millisecond):
				mark()
			}
		}
		return 42, nil
	})

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != 42 {
		t.Errorf("result = %d, want 42", got)
	}
}

func TestExecute_TotalTimeoutFires(t *testing.T) {
	st, _ := New[int](Config{
		IdleTimeout:  500 * time.Millisecond,
		TotalTimeout: 60 * time.Millisecond,
	})

	_, err := st.Execute(context.Background(), func(ctx context.Context, mark Mark) (int, error) {
		// Mark frequently so idle never fires; total should still fire.
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return 0, ctx.Err()
			case <-ticker.C:
				mark()
			}
		}
	})

	var ste *StreamTimeoutError
	if !errors.As(err, &ste) {
		t.Fatalf("expected *StreamTimeoutError, got %v", err)
	}
	if ste.Stage != StageTotal {
		t.Errorf("Stage = %s, want %s", ste.Stage, StageTotal)
	}
}

func TestExecute_HappyPath(t *testing.T) {
	st, _ := New[string](Config{
		FirstByteTimeout: time.Second,
		IdleTimeout:      time.Second,
		TotalTimeout:     2 * time.Second,
	})

	got, err := st.Execute(context.Background(), func(_ context.Context, mark Mark) (string, error) {
		mark()
		mark()
		return "ok", nil
	})

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != "ok" {
		t.Errorf("result = %q, want %q", got, "ok")
	}
}

func TestExecute_ParentCancelPropagates(t *testing.T) {
	st, _ := New[int](Config{TotalTimeout: time.Second})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := st.Execute(ctx, func(ctx context.Context, _ Mark) (int, error) {
		<-ctx.Done()
		return 0, ctx.Err()
	})

	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestStreamTimeoutError_LogValue(t *testing.T) {
	e := &StreamTimeoutError{
		Stage:            StageIdle,
		FirstByteTimeout: time.Second,
		IdleTimeout:      100 * time.Millisecond,
		TotalTimeout:     5 * time.Second,
		Elapsed:          150 * time.Millisecond,
	}
	got := map[string]any{}
	for _, a := range e.LogValue().Group() {
		got[a.Key] = a.Value.Any()
	}
	if got["stage"] != "idle" {
		t.Errorf("stage = %v, want idle", got["stage"])
	}
	if got["error"] != "stream_timeout" {
		t.Errorf("error = %v, want stream_timeout", got["error"])
	}
}
