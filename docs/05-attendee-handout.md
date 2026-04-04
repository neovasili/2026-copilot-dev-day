# Attendee handout (Go block)

## Session focus

This block is intentionally narrow:
- tiny code changes
- strict quality gate after each change
- iterate until checks are green

## Files we will show

- `demo/go/AGENTS.md`
- `demo/go/.github/copilot-instructions.md`
- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

## Commands used live

```bash
cd demo/go
./scripts/quality_gate.sh
```

Fallback (if toolchain mismatch appears):

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```

## Prompt pattern to reuse

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff for <task> in <file>.
Keep behavior stable unless explicitly requested.
Run ./scripts/quality_gate.sh after editing.
If checks fail, iterate with minimal fixes until PASS.
Return only when quality gate is green.
```

## What to copy into your real projects

- Add an `AGENTS.md` (or equivalent) with style and workflow rules.
- Keep quality commands explicit and easy to run.
- Finish only when lint + tests pass.
