# Live demo script (Go + Copilot)

This is the single demo-first playbook for your Go block.
If you only open one prep doc, use this one.

## Demo objective

Show one repeatable pattern for production work with Copilot:

- minimal scoped change
- mandatory quality gate (format + lint + tests)
- iterate until green

## Files to keep open during the demo

- `demo/go/AGENTS.md`
- `demo/go/.github/copilot-instructions.md`
- `docs/micro-change-tasks.md`
- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

## 20-minute run of show

- 00:00-01:30: Opening
	- "In this block we optimize for trust: tiny changes, fast checks, repeat until green."
	- "Copilot writes suggestions; quality gates decide if we keep them."

- 01:30-03:30: Guardrails first
	- Open `demo/go/AGENTS.md`.
	- Open `demo/go/.github/copilot-instructions.md`.
	- Explain the loop:
		1. smallest diff
		2. run `./scripts/quality_gate.sh`
		3. if fail, fix and repeat until pass

- 03:30-05:00: 30-second instruction file explainer
	- `AGENTS.md` = behavior contract (scope, style, tests, stop condition).
	- `copilot-instructions.md` = integration point that makes Copilot load repo instructions by default.
	- One-line summary: "`AGENTS.md` defines the rules; `copilot-instructions.md` makes Copilot apply them automatically."

- 05:00-12:00: First micro-change
	- Pick Easy/Medium/Hard from `docs/micro-change-tasks.md` based on remaining time.
	- Accept only minimal in-scope edits.
	- Reject broad refactors.

- 12:00-15:30: Enforce quality gate
	- Run quality gate and iterate to green.
	- Say: "Green is the finish line, not generated text."

- 15:30-18:30: Optional second micro-change
	- Do one more tiny enhancement and re-run the gate.

- 18:30-20:00: Close
	- "Put standards in files, not only in people’s heads."
	- "Copilot + quality loop gives speed with accountability."

## Prompt starters

Prompt A:

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff in service.go to improve readability of ProbeOnce error handling without changing behavior.
After editing, run ./scripts/quality_gate.sh.
If anything fails, keep iterating with minimal fixes until the quality gate passes.
Then summarize changes in 4 bullets.
```

Prompt B:

```text
Add one small table-driven test that improves edge-case coverage for service.go.
Keep the change minimal.
Run ./scripts/quality_gate.sh and iterate until PASS.
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

If you get only 15 minutes:

- show guardrails
- do one micro-change
- finish on green quality gate

If you get only 10 minutes:

- show `AGENTS.md`
- run one prompt
- run one quality gate
