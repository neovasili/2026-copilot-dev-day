# Run of show - Go block (20 min)

## Goal

In 20 minutes, attendees should see a complete mini workflow:
1. Ask Copilot for a focused Go implementation.
2. Verify behavior with tests.
3. Refine code to align with idiomatic Go.

## Minute-by-minute plan

- 00:00-01:30: Framing
  - State one outcome: "we will build a small HTTP probe service with tests and a concurrency option."
  - Remind audience this is practical: short prompts, quick validation.

- 01:30-04:00: Context setup
  - Open `demo/go/internal/probe/service.go`.
  - Explain the `HTTPClient` interface and why this keeps code testable.
  - Mention success criteria: valid URL handling, timeout, preserved order in batch probes.

- 04:00-10:00: Build with Copilot
  - Use Copilot Chat to generate/refine `ProbeOnce`.
  - Focus points: wrapped errors (`%w`), `context.WithTimeout`, clear return values.
  - Run `go test ./...`.

- 10:00-15:30: Batch probing and light concurrency
  - Walk through `ProbeAll`.
  - Emphasize bounded workers and stable output order.
  - Re-run tests and show that behavior is still deterministic.

- 15:30-18:30: Idiomatic cleanup pass
  - Ask Copilot for a readability pass without changing behavior.
  - Show one accepted suggestion and one rejected suggestion.
  - Reinforce: Copilot assists, developer decides.

- 18:30-20:00: Close + takeaways
  - 3 practical takeaways:
    - Ask for constraints, not only outcomes.
    - Keep a red/green loop with tests.
    - Use Go tooling (`go test`, `go fmt`) as guardrails.

## Timing fallback (if previous block overruns)

If you get only 15 minutes:
- Keep `ProbeOnce` + tests.
- Explain `ProbeAll` conceptually without full code walk.
- Show one concurrency snippet and move to Q&A.

If you get only 10 minutes:
- Show one single-file flow (`ProbeOnce`), one test run, and one prompt pattern attendees can reuse.
