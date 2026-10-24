# Go Project: `llmroute` — a small LLM router with per-team cost tracking

**Why this project:** the Figma Internal AI role (and a growing number of platform roles) asks for experience with model routing, cost/usage instrumentation, and audit trails for AI workloads. You haven't built any of that, and using AI tools isn't the same thing. This project gives you a small, real, truthful example, and it builds on the Go you already know from `dnsprobe`: `context` timeouts, the `(result, error)` idiom, structs, and turning a CLI into a service.

**What you're really building:** a program that accepts a prompt and a team name, picks a model by a rule you define, calls that model's API, and records how many tokens it used and what that cost, tagged by team. That's it. Everything else is small steps toward that.

Six milestones. Each has a goal, new concepts, a skeleton (signatures only, you write the bodies), a checkpoint, and a hint. No full solutions on purpose. Do the "think first" questions before you write code. They're the part you'll be asked about in an interview.

---

## Before you start (30 minutes)

1. **Get an API key** from an LLM provider (any one is fine). Read their "getting started" docs and find three things: the URL you POST to, the headers it needs, and the shape of the JSON request and response.
2. **Set a monthly spending cap** in the provider's console before you make your first call. Use the cheapest model they offer. With short prompts, testing this whole project should cost pennies to a few dollars.
3. **Never put the key in your code or git.** Read it from an environment variable (`os.Getenv`). Add a `.gitignore` entry for any file that might hold it.
4. Create the module: `go mod init llmroute`.

**Think first:** why would a company route requests across models at all, instead of using the best one for everything? Write two sentences. (Hint: cost and latency are not the same as capability.)

---

## M1 — One request, one model

**Goal:** send a fixed prompt to one model and print the reply text.

**New concepts:** `net/http` (build a request, set headers, send it), `encoding/json` (marshal a struct into a request body, unmarshal the response), reading a secret from the environment.

**Skeleton:**
```go
type Request struct { /* fields from the provider's docs */ }
type Response struct { /* only the fields you need: reply text, token counts */ }

func callModel(apiKey, model, prompt string) (*Response, error) {
    // TODO
}
```

**Checkpoint:** `go run .` prints a sensible reply. With a deliberately wrong key it prints a clean error that includes the HTTP status. It must not panic and must not print the key.

**Think first:** your `Response` struct only needs a few fields. Why decode only those instead of everything the API returns?

**Hint if stuck:** `json.Marshal`, `http.NewRequest`, and `http.Client{}.Do`. Check the status code before you try to decode the body.

---

## M2 — Two models behind one function

**Goal:** make "which model?" a parameter, so the rest of your program doesn't care which one it's talking to.

**New concepts:** interfaces. Define what a "model" can do, then let two different implementations satisfy it. This is the core abstraction in any gateway.

**Skeleton:**
```go
type Model interface {
    Name() string
    Complete(ctx context.Context, prompt string) (Result, error)
}

type Result struct {
    Text         string
    InputTokens  int
    OutputTokens int
}
```
Start with two models from the same provider, one cheap and one stronger. A second provider is optional later. Providers name the token fields differently, so converting them into your own `Result` is the point.

**Checkpoint:** the same prompt runs through both models by changing one value. Your `main` has no provider-specific code.

**Think first:** what belongs in the interface and what doesn't? For example, should the price per token live in the model, or somewhere else?

---

## M3 — A routing rule

**Goal:** pick the model automatically.

**New concepts:** keeping policy separate from mechanism. The rule can be crude, such as prompt length, an explicit "complexity" tag from the caller, or a keyword list. What matters is that the rule is one small function you can change and test.

**Skeleton:**
```go
func chooseModel(prompt string, team string, models []Model) Model {
    // TODO — your rule
}
```

**Checkpoint:** short prompts go to the cheap model and long or flagged ones go to the stronger one, and you can show it. Print which model was chosen and why.

**Think first:** a routing decision has to be explainable later to a security or finance person. What would you record so that someone could ask "why did this request go to that model?"

---

## M4 — Usage and cost log, tagged by team

**Goal:** after every call, append one structured record: timestamp, team, model, input tokens, output tokens, estimated cost, and the reason the model was chosen.

**New concepts:** `log/slog` (structured logging), a small price table, and writing to a file as JSON lines.

**Skeleton:**
```go
type UsageRecord struct { /* fields above */ }

func estimateCost(model string, in, out int) float64 {
    // TODO — look up per-token prices, multiply
}

func recordUsage(r UsageRecord) error {
    // TODO — append one JSON line
}
```

**Checkpoint:** run twenty prompts across two teams. Then write a short second program, or a command-line flag on this one, that reads the log and prints total cost per team. Make it sensible against hand arithmetic on three of the rows.

**Think first:** prices change. Where should the price table live so you don't have to recompile to update it? What happens to old records when a price changes?

---

## M5 — Turn it into an HTTP service

**Goal:** other programs can POST a prompt and a team to your service and get the reply.

**New concepts:** `net/http` as a server (handlers, status codes), request validation, reading the team from a header, and a timeout on every outbound call using `context.WithTimeout`, the same pattern as `dnsprobe` M3.

**Skeleton:**
```go
func handleComplete(w http.ResponseWriter, r *http.Request) {
    // TODO — validate input, pick model, call it with a timeout, log usage, reply
}
```

**Checkpoint:** `curl` against your service works from a second terminal. A request with no team header gets a clean 400. A model that takes too long returns an error to the caller instead of hanging.

**Think first:** this is a shared service, so what stops one team from running up the whole bill? Don't build the answer yet, just write down two options.

---

## M6 — Tests and a failure path

**Goal:** prove the routing and the cost math work, and handle a failing model.

**New concepts:** `testing`, table-driven tests, and `net/http/httptest` to fake a model API so your tests don't spend money.

**Skeleton:**
```go
func TestChooseModel(t *testing.T) { /* table of prompt/team/expected model */ }
func TestEstimateCost(t *testing.T) { /* table of tokens/model/expected cost */ }
```

**Checkpoint:** `go test ./...` passes with no network access. When the cheap model returns an error, your service retries once on the stronger model and records that it did.

**Think first:** is a fallback to a more expensive model always the right behavior? When would you rather fail?

---

## Optional extensions (only if you want them)

- A per-team daily budget that refuses requests once exceeded.
- A second provider behind the same interface.
- A sanitized audit field. Record who asked and what was routed, but decide deliberately whether to store prompt text at all.
- Run it in a container and deploy it the way you deploy the DNSSEC platform.

---

## How to talk about it honestly

**You can say:** "I built a small LLM router in Go to learn how gateways work. It routes by a rule I wrote, logs token usage and estimated cost per team, and has tests that run without calling the API."

**Don't say:** that you ran an AI platform, that it served real users, or that it's production-grade. It's a learning project, and that's a respectable thing to say it is. When asked what you'd do differently at company scale, your answers to the "think first" questions are the material. Authentication, rate limits, budgets, and audit storage are exactly what the real teams work on.

**Resume placement:** only after it exists. A line under a "Projects" heading would be fine. Not before.

## Time estimate
Roughly one evening for each of M1 through M4, and one to two for each of M5 and M6. Four to five weeks at a few hours a week.
