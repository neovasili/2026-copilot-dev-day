# Live demo script (Go + Copilot)

Use this as a practical speaking track focused on minimal changes and feedback loops.

## 1) Opening (60 sec)

Say:

- "In this block we optimize for trust: tiny changes, fast checks, repeat until green."
- "Copilot writes suggestions; quality gates decide if we keep them."

## 2) Show guardrails before coding (2-3 min)

Open:

- `docs/09-instructions-files-explainer.md`
- `demo/go/AGENTS.md`
- `demo/go/.github/copilot-instructions.md`

Say:

- "This is how we encode style, testability, and scope in the repo itself."
- "`AGENTS.md` defines the rules; `copilot-instructions.md` makes Copilot apply them automatically."
- "Our rule: no finished answer without passing the quality gate."

## 3) First micro-task prompt (5-6 min)

Open:

- `docs/08-micro-change-tasks.md`
- `demo/go/internal/probe/service.go`

Prompt A:

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff in service.go to improve readability of ProbeOnce error handling without changing behavior.
After editing, run ./scripts/quality_gate.sh.
If anything fails, keep iterating with minimal fixes until the quality gate passes.
Then summarize changes in 4 bullets.
```

Talk track:

- Pick Easy/Medium/Hard based on remaining time.
- Accept only minimal in-scope edits.
- Reject over-broad refactors.

## 4) Run and show gate (3 min)

Run:

```bash
cd demo/go
make quality
```

Fallback (toolchain mismatch):

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```

Say:

- "Green is the finish line, not generated text."

## 5) Second micro-task prompt (4-5 min)

Prompt B:

```text
Add one small table-driven test that improves edge-case coverage for service.go.
Keep the change minimal.
Run ./scripts/quality_gate.sh and iterate until PASS.
```

Say:

- "Same loop again. Repeatability is the win."

## 6) Close (60 sec)

Say:

- "Put standards in files, not only in people’s heads."
- "Copilot + quality loop gives speed with accountability."
