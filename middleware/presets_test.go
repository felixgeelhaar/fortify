package middleware_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/budget"
	"github.com/felixgeelhaar/fortify/middleware"
)

func TestHTTPClient_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.HTTPClient(middleware.HTTPClientConfig{Timeout: 0})
	if err == nil {
		t.Fatal("want error for zero Timeout, got nil")
	}
}

func TestHTTPClient_ExecutesOperation(t *testing.T) {
	chain, err := middleware.HTTPClient(middleware.HTTPClientConfig{Timeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	want := &http.Response{StatusCode: 200}
	got, err := chain.Execute(context.Background(), func(ctx context.Context) (*http.Response, error) {
		return want, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestHTTPClient_RetriesOnTransientError(t *testing.T) {
	chain, err := middleware.HTTPClient(middleware.HTTPClientConfig{
		Timeout:           500 * time.Millisecond,
		MaxRetries:        3,
		RetryInitialDelay: 1 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	attempts := 0
	transient := errors.New("transient")
	got, err := chain.Execute(context.Background(), func(ctx context.Context) (*http.Response, error) {
		attempts++
		if attempts < 3 {
			return nil, transient
		}
		return &http.Response{StatusCode: 200}, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", got.StatusCode)
	}
	if attempts != 3 {
		t.Fatalf("attempts = %d, want 3", attempts)
	}
}

func TestDatabaseQuery_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.DatabaseQuery(middleware.DatabaseQueryConfig{QueryTimeout: 0})
	if err == nil {
		t.Fatal("want error for zero QueryTimeout, got nil")
	}
}

func TestDatabaseQuery_ExecutesOperation(t *testing.T) {
	chain, err := middleware.DatabaseQuery(middleware.DatabaseQueryConfig{QueryTimeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	got, err := chain.Execute(context.Background(), func(ctx context.Context) (any, error) {
		return "row-data", nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got.(string) != "row-data" {
		t.Fatalf("got %v, want row-data", got)
	}
}

func TestRPCDownstream_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.RPCDownstream[string](middleware.RPCDownstreamConfig{CallTimeout: 0})
	if err == nil {
		t.Fatal("want error for zero CallTimeout, got nil")
	}
}

func TestRPCDownstream_GenericTypePropagated(t *testing.T) {
	chain, err := middleware.RPCDownstream[int](middleware.RPCDownstreamConfig{CallTimeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	got, err := chain.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got != 42 {
		t.Fatalf("got %d, want 42", got)
	}
}

func TestLLMCall_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.LLMCall[string](middleware.LLMCallConfig[string]{
		Provider: "openai",
		Model:    "gpt-5",
		Budget: middleware.BudgetConfig[string]{
			Max: budget.Cost{Calls: 5},
		},
	})
	if err == nil {
		t.Fatal("want error for zero CallTimeout, got nil")
	}
}

func TestLLMCall_RejectsZeroBudget(t *testing.T) {
	_, err := middleware.LLMCall[string](middleware.LLMCallConfig[string]{
		Provider:    "openai",
		Model:       "gpt-5",
		CallTimeout: 100 * time.Millisecond,
		// Budget all-zero
	})
	if err == nil {
		t.Fatal("want error for zero budget, got nil")
	}
}

func TestLLMCall_HappyPath(t *testing.T) {
	chain, err := middleware.LLMCall[string](middleware.LLMCallConfig[string]{
		Provider:    "openai",
		Model:       "gpt-5",
		CallTimeout: 200 * time.Millisecond,
		Budget: middleware.BudgetConfig[string]{
			Max: budget.Cost{Calls: 5},
		},
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	got, err := chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "completion", nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got != "completion" {
		t.Errorf("got %q, want %q", got, "completion")
	}
}

func TestLLMCall_BudgetCapStopsRetryStorm(t *testing.T) {
	var attempts int
	chain, err := middleware.LLMCall[string](middleware.LLMCallConfig[string]{
		Provider:    "openai",
		Model:       "gpt-5",
		CallTimeout: 200 * time.Millisecond,
		MaxRetries:  10,
		Budget: middleware.BudgetConfig[string]{
			Max: budget.Cost{Calls: 2},
		},
		// AssumeIdempotent so transient errors trigger retries.
		AssumeIdempotent: true,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	transient := errors.New("transient downstream error")
	_, err = chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		attempts++
		return "", transient
	})

	// Should fail with budget exceeded, not transient err.
	if !errors.Is(err, budget.ErrBudgetExceeded) {
		t.Fatalf("expected ErrBudgetExceeded, got %v", err)
	}
	// Two attempts charged, third refused before fn runs.
	if attempts != 2 {
		t.Errorf("attempts = %d, want 2 (budget refused 3rd before fn)", attempts)
	}
}

func TestLLMCall_NonIdempotentDefaultDoesNotRetryArbitraryErrors(t *testing.T) {
	var attempts int
	chain, err := middleware.LLMCall[string](middleware.LLMCallConfig[string]{
		Provider:    "openai",
		Model:       "gpt-5",
		CallTimeout: 200 * time.Millisecond,
		MaxRetries:  5,
		Budget: middleware.BudgetConfig[string]{
			Max: budget.Cost{Calls: 100},
		},
		// AssumeIdempotent default = false
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	want := errors.New("uncategorised LLM error")
	_, _ = chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		attempts++
		return "", want
	})

	if attempts != 1 {
		t.Errorf("attempts = %d, want 1 (non-idempotent default should not retry uncategorised errors)", attempts)
	}
}

func TestLLMHedge_RejectsEmptyAttempts(t *testing.T) {
	_, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{})
	if err == nil {
		t.Fatal("want error for empty Attempts")
	}
}

func TestLLMHedge_RejectsCrossVendorWithoutOptIn(t *testing.T) {
	noop := func(context.Context) (string, error) { return "", nil }
	_, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts: []func(context.Context) (string, error){noop, noop},
	})
	if err == nil {
		t.Fatal("want error when multiple Attempts lack opt-in")
	}
}

func TestLLMHedge_AcceptsSameVendorMulti(t *testing.T) {
	noop := func(context.Context) (string, error) { return "", nil }
	_, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts:              []func(context.Context) (string, error){noop, noop},
		AllAttemptsSameVendor: true,
	})
	if err != nil {
		t.Fatalf("AllAttemptsSameVendor should be sufficient: %v", err)
	}
}

func TestLLMHedge_PrimaryWinsWhenFast(t *testing.T) {
	r, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts: []func(context.Context) (string, error){
			func(context.Context) (string, error) { return "primary", nil },
			func(ctx context.Context) (string, error) {
				<-ctx.Done()
				return "secondary", ctx.Err()
			},
		},
		HedgeAfter:            500 * time.Millisecond,
		AllAttemptsSameVendor: true,
	})
	if err != nil {
		t.Fatalf("constructor: %v", err)
	}

	got, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run err: %v", err)
	}
	if got != "primary" {
		t.Errorf("got %q, want primary", got)
	}
}

func TestLLMHedge_HedgeWinsWhenPrimarySlow(t *testing.T) {
	var hedgeFired int
	r, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts: []func(context.Context) (string, error){
			func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(500 * time.Millisecond):
					return "primary-late", nil
				}
			},
			func(context.Context) (string, error) {
				return "secondary-fast", nil
			},
		},
		HedgeAfter:            30 * time.Millisecond,
		AllAttemptsSameVendor: true,
		OnHedge:               func(int) { hedgeFired++ },
	})
	if err != nil {
		t.Fatalf("constructor: %v", err)
	}

	got, err := r.Run(context.Background())
	if err != nil {
		t.Fatalf("Run err: %v", err)
	}
	if got != "secondary-fast" {
		t.Errorf("got %q, want secondary-fast", got)
	}
	if hedgeFired != 1 {
		t.Errorf("OnHedge fired %d times, want 1", hedgeFired)
	}
}

func TestLLMHedge_AllErrorsReturnsLast(t *testing.T) {
	r, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts: []func(context.Context) (string, error){
			func(context.Context) (string, error) { return "", errors.New("first failed") },
			func(context.Context) (string, error) { return "", errors.New("second failed") },
		},
		HedgeAfter:            10 * time.Millisecond,
		AllAttemptsSameVendor: true,
	})
	if err != nil {
		t.Fatalf("constructor: %v", err)
	}

	_, err = r.Run(context.Background())
	if err == nil {
		t.Fatal("want error when all attempts fail")
	}
}

func TestLLMHedge_ParentCancelPropagates(t *testing.T) {
	r, err := middleware.LLMHedge[string](middleware.LLMHedgeConfig[string]{
		Attempts: []func(context.Context) (string, error){
			func(ctx context.Context) (string, error) {
				<-ctx.Done()
				return "", ctx.Err()
			},
		},
		HedgeAfter: 100 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("constructor: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	_, err = r.Run(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestLLMCall_BudgetChargesWiredFromConfig(t *testing.T) {
	chain, err := middleware.LLMCall[int](middleware.LLMCallConfig[int]{
		Provider:    "anthropic",
		Model:       "claude",
		CallTimeout: 200 * time.Millisecond,
		Budget: middleware.BudgetConfig[int]{
			Max: budget.Cost{Tokens: 50},
			Charge: func(_ context.Context, r int, _ error) budget.Cost {
				return budget.Cost{Tokens: int64(r)}
			},
		},
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	// First two calls each charge 30 tokens; second call breaches.
	_, err1 := chain.Execute(context.Background(), func(context.Context) (int, error) { return 30, nil })
	_, err2 := chain.Execute(context.Background(), func(context.Context) (int, error) { return 30, nil })

	if err1 != nil {
		t.Errorf("first call err = %v, want nil", err1)
	}
	if !errors.Is(err2, budget.ErrBudgetExceeded) {
		t.Errorf("second call err = %v, want budget exceeded", err2)
	}
}
