// Example: wrapping an Eino-shaped chat model with the Fortify
// middleware.LLMCall preset.
//
// Eino (https://github.com/cloudwego/eino) is the prevailing Go LLM
// application framework and exposes ChatModel-style components as the
// composable unit of work. Fortify is complementary: it does not own
// prompts, tools, planning, or memory. It owns the resilience layer
// around each model call.
//
// This example uses a minimal local interface that mirrors the shape
// of Eino's model.ChatModel so the example builds standalone in CI
// without pinning to a specific Eino version. To wire it up for real:
//
//	import "github.com/cloudwego/eino/components/model"
//
// then satisfy this example's ChatModel interface with an Eino model
// (or a *eino-ext provider implementation), and the LLMCall chain
// construction below stays the same.
//
// What this example demonstrates:
//
//   - middleware.LLMCall capping the dollar cost of an LLM retry storm
//     during a provider incident via budget.Cost{USDMicros}.
//   - Token-aware Charge based on response Usage so the budget reflects
//     actual provider billing, not a guess.
//   - Non-idempotent retry semantics by default: we don't retry uncategorised
//     errors because partial completions already cost tokens.
//   - Structured error handling: budget.ErrBudgetExceeded short-circuits
//     remaining attempts; the caller can either degrade gracefully or
//     surface the error.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"go.klarlabs.de/fortify/budget"
	"go.klarlabs.de/fortify/middleware"
)

// Message mirrors Eino's schema.Message minimally. Real callers should
// substitute the upstream type.
type Message struct {
	Role    string
	Content string
}

// Usage carries token counts as reported by the provider. Real Eino
// callers can read these from the underlying provider response and
// promote them onto a struct of this shape.
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Response is the example shape. Real callers will already have a
// response type; the only requirement is that token usage is
// readable from the result so the budget Charge callback can read it.
type Response struct {
	Message Message
	Usage   Usage
}

// ChatModel is the Eino-shaped seam. A real Eino model satisfies this
// trivially because the Generate method already exists with a closely
// matching signature.
type ChatModel interface {
	Generate(ctx context.Context, in []Message) (Response, error)
}

// fakeChatModel is a placeholder so this example builds without a
// real provider. Replace with an Eino model in your application.
type fakeChatModel struct {
	calls int
	fail  int // first N calls fail with a transient error
}

func (m *fakeChatModel) Generate(ctx context.Context, _ []Message) (Response, error) {
	m.calls++
	if m.calls <= m.fail {
		return Response{}, fmt.Errorf("provider 503: transient")
	}
	return Response{
		Message: Message{Role: "assistant", Content: "ok"},
		Usage:   Usage{PromptTokens: 100, CompletionTokens: 50, TotalTokens: 150},
	}, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// In a real app, replace this with an Eino chat model bound to a
	// provider (OpenAI, Anthropic, Mistral, …):
	//
	//	model, _ := openai.NewChatModel(ctx, &openai.ChatModelConfig{Model: "gpt-5"})
	//
	model := &fakeChatModel{fail: 2}

	// Cost-aware budget. USDMicros is denominated as 1_000_000 = $1.
	// At ~$0.005 per 1k tokens, $0.05 caps roughly 10k tokens.
	const usdPer1KTokensMicros = 5_000

	chain, err := middleware.LLMCall[Response](middleware.LLMCallConfig[Response]{
		Provider:    "openai",
		Model:       "gpt-5",
		CallTimeout: 30 * time.Second,
		MaxRetries:  3,
		Budget: middleware.BudgetConfig[Response]{
			Max: budget.Cost{
				Tokens:    50_000,
				USDMicros: 250_000, // $0.25
			},
			Charge: func(_ context.Context, r Response, err error) budget.Cost {
				if err != nil {
					// Failed attempts may have consumed input tokens. If
					// your provider exposes them, charge here. We default
					// to zero for unknown failures.
					return budget.Cost{}
				}
				return budget.Cost{
					Tokens:    int64(r.Usage.TotalTokens),
					USDMicros: int64(r.Usage.TotalTokens) * usdPer1KTokensMicros / 1000,
				}
			},
			OnExceeded: func(c budget.Cost) {
				logger.Warn("budget breached",
					slog.Int64("tokens", c.Tokens),
					slog.Int64("usd_micros", c.USDMicros),
					slog.Int64("calls", c.Calls),
				)
			},
		},
		// LLM calls are not idempotent by default. Set true if your
		// caller-side logic deduplicates via an idempotency key.
		AssumeIdempotent: false,
	})
	if err != nil {
		logger.Error("LLMCall constructor failed", slog.Any("err", err))
		os.Exit(1)
	}

	// Drive the chain a few times so the budget fires.
	for i := 0; i < 5; i++ {
		out, err := chain.Execute(context.Background(), func(ctx context.Context) (Response, error) {
			return model.Generate(ctx, []Message{{Role: "user", Content: "hi"}})
		})

		switch {
		case err == nil:
			logger.Info("ok",
				slog.Int("call", i),
				slog.String("content", out.Message.Content),
				slog.Int("tokens", out.Usage.TotalTokens),
			)
		case errors.Is(err, budget.ErrBudgetExceeded):
			logger.Warn("budget exceeded; stopping", slog.Any("err", err))
			return
		default:
			logger.Error("call failed", slog.Int("call", i), slog.Any("err", err))
		}
	}
}
