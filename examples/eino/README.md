# Eino + Fortify

This example shows how to wrap an [Eino](https://github.com/cloudwego/eino)
chat-model component with the Fortify `middleware.LLMCall` preset to
get a cost-aware, retry-bounded resilience layer for free.

Eino owns prompts, tools, planning, memory. Fortify owns the
resilience layer around each model call. They compose; they do not
compete.

## What this example demonstrates

- `middleware.LLMCall` capping the dollar cost of a retry storm during
  a provider incident, via `budget.Cost{USDMicros, Tokens}`.
- Token-aware `Charge` reading the response Usage so the budget
  reflects actual provider billing, not a guess.
- Non-idempotent retry semantics by default: uncategorised errors do
  not trigger retries because partial completions already cost
  tokens.
- Structured handling of `budget.ErrBudgetExceeded` to short-circuit
  remaining work cleanly.

## What this example is not

It does **not** import `github.com/cloudwego/eino` directly so the
example builds standalone in CI. A minimal local `ChatModel` interface
mirrors the shape of Eino's `model.ChatModel`. To wire it to a real
Eino model:

```go
import (
    "github.com/cloudwego/eino/components/model"
    "github.com/cloudwego/eino-ext/components/model/openai"
)

m, _ := openai.NewChatModel(ctx, &openai.ChatModelConfig{Model: "gpt-5"})
// m satisfies model.ChatModel; adapt it to this example's ChatModel
// interface (one-method shim) and pass it where fakeChatModel is used.
```

## Run

```bash
cd examples/eino
go run .
```

The fake model is configured to fail twice, then succeed; the budget
ceiling will eventually fire so you can watch the breach handler do
its job without a real LLM bill.

## Sensitive payloads

The example never logs prompt or completion content. The structured
errors emitted by Fortify implement `slog.LogValuer`, so logging via
`slog.Any("err", err)` produces structured fields without ever
inspecting payload contents. See `docs/PRODUCTION.md` ("Observability
and sensitive payloads") for the full posture.
