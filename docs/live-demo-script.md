# Live demo script (Go + Copilot)

This is the main playbook for your Go segment.
If you only open one prep file before the session, use this one.

## Session 03 timing context

- Session 03 keeps the same overall window: 18:40-19:40.
- Activities now run by parallel track.
- This script targets the hands-on track where Python runs first (30 min) and Go runs second (30 min).

## Block objective

Show a repeatable workflow for using Copilot in Go while staying idiomatic:

- clear code aligned with Go conventions
- small, verifiable changes
- mandatory quality gate after each change

## Topics to cover

- packages and modules (quick read of `go.mod` and folder structure)
- interfaces and idiomatic error handling (`%w`, `errors.Is`, useful context)
- lightweight concurrency when it adds value (workers, `context`, cancellation)
- Copilot usage inside the editor for:
  - completing tests
  - refactoring small functions
  - exploring stdlib options without leaving the editor
- final time for tooling questions (`gofmt`, tests, modules)

## Files to keep open

- `demo/go/AGENTS.md`
- `demo/go/.github/copilot-instructions.md`
- `docs/micro-change-tasks.md`
- `demo/go/go.mod`
- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

## 30-minute run of show

- 00:00-02:00: Opening
  - "Today we prioritize idiomatic Go and clear code."
  - "Copilot can accelerate delivery, but quality is decided by the gate."

- 02:00-05:00: Go context
  - Show `go.mod` and package structure.
  - Explain in 30 seconds how this connects to imports, tests, and modules.

- 05:00-08:00: Guardrails + working contract
  - Open `demo/go/AGENTS.md` and `demo/go/.github/copilot-instructions.md`.
  - Repeat the mandatory loop:
    1. minimal change
    2. run `./scripts/quality_gate.sh`
    3. if it fails, fix and repeat

- 08:00-16:00: Micro-change 1 (tests or error handling)
  - Choose an Easy or Medium task from `docs/micro-change-tasks.md`.
  - Emphasize interfaces and idiomatic error handling while Copilot proposes changes.
  - Accept only small, readable diffs.

- 16:00-20:00: Quality and verification
  - Run the quality gate.
  - If it fails, iterate live with a single focused correction.

- 20:00-27:00: Optional micro-change 2 (light concurrency or refactor)
  - Choose a Hard task only if time allows.
  - Reinforce that concurrency should be used where it adds value, not for complexity.

- 27:00-30:00: Closing + tooling Q&A
  - Leave room for questions: `gofmt`, tests, modules, package organization.
  - Suggested closing:
    - "Speed without quality does not scale."
    - "Copilot + conventions + checks gives teams a reusable workflow."

## Prompt starters

Prompt A:

```text
Read AGENTS.md and follow it strictly.
Complete or improve table-driven tests in service_test.go for one edge case in URL validation or error handling.
Keep changes minimal and idiomatic.
After editing, run ./scripts/quality_gate.sh and iterate until PASS.
```

Prompt B:

```text
Read AGENTS.md and follow it strictly.
Refactor one small function in service.go to improve readability while preserving behavior.
Keep error wrapping with %w and avoid broad refactors.
Run ./scripts/quality_gate.sh and iterate until PASS.
```

Prompt C:

```text
Read AGENTS.md and follow it strictly.
From the current code in service.go, explain which stdlib APIs are central (context, net/http, errors, sync) and propose one minimal improvement.
Apply only that improvement and run ./scripts/quality_gate.sh until PASS.
```

## Commands to run live

```bash
cd demo/go
make quality
```

Fallback (toolchain mismatch):

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```

## Timing fallback

If you get only 20 minutes:

- show modules/packages in 30 seconds
- do one micro-change (tests or errors)
- finish on a green quality gate

If you get only 10 minutes:

- show `AGENTS.md`
- run one short prompt
- run the quality gate once
