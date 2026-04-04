# Go workshop demo

This mini-project is tuned for a 20-minute Copilot practical block focused on trustable speed.

## Core method

1. Read repo guardrails (`AGENTS.md`).
2. Ask Copilot for the smallest possible change.
3. Run the quality gate.
4. If it fails, iterate until pass.

For live selection by difficulty, use:

- [08-micro-change-tasks.md](../../docs/08-micro-change-tasks.md)

For the short explanation of instruction files, use:

- [09-instructions-files-explainer.md](../../docs/09-instructions-files-explainer.md)

## Files to use live

- `AGENTS.md`
- `.github/copilot-instructions.md`
- `internal/probe/service.go`
- `internal/probe/service_test.go`
- `scripts/quality_gate.sh`

## Demo commands

```bash
cd demo/go
./scripts/quality_gate.sh
```

Alternative:

```bash
cd demo/go
make quality
```

If your machine reports a toolchain mismatch (for example `compile version ... does not match go tool version ...`), pin one toolchain explicitly:

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```

## Prompt starter

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff for <task> in <file>.
Keep behavior stable unless explicitly requested.
Run ./scripts/quality_gate.sh after editing.
If checks fail, iterate with minimal fixes until PASS.
Return only when quality gate is green.
```
