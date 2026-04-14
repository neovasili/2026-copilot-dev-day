# Copilot Dev Days Caceres - Session Prep Repo

Prep materials for Juan Manuel Ruiz Fernandez's Go block in **Session 03** of GitHub Copilot Developer Days (Caceres edition).

## Confirmed context from event website

Source pages:

- <https://copilot-dev-caceres.vercel.app/index.html>
- <https://copilot-dev-caceres.vercel.app/pages/agenda.html>
- <https://copilot-dev-caceres.vercel.app/pages/session-03.html>

Confirmed details:

- Event date: **Friday, April 17, 2026**
- Event schedule: **17:00-20:00** (local time)
- Session 03: **18:40-19:40** (60 min)
- Session 03 format: practical workshop with three blocks of about 20 minutes each
- Go block speaker listed as: **Juan Manuel Ruiz (Juanma)**

## Session strategy in this repo

The content is optimized for:

- Go fundamentals in practice: packages/modules, interfaces, and idiomatic error handling
- lightweight concurrency when it adds value (workers + context cancellation)
- Copilot in real editor workflows: complete tests, refactor small functions, explore stdlib usage
- clear code and language conventions first, with mandatory verification loop
- time reserved for tooling Q&A (`gofmt`, tests, modules, common ecosystem practices)

## Repo structure

- `docs/live-demo-script.md`: single demo playbook (run of show + talk track + instruction-file explainer)
- `docs/micro-change-tasks.md`: concrete easy/medium/hard tasks for live selection
- `demo/go/AGENTS.md`: prewritten guardrails for Copilot behavior
- `demo/go/.github/copilot-instructions.md`: Copilot-compatible instruction hook
- `demo/go/scripts/quality_gate.sh`: format/lint/test quality gate
- `demo/go`: runnable mini-project for the workshop

## Requirements

- GitHub account with GitHub Copilot access
- IDE with Copilot enabled (VS Code or JetBrains recommended)
- Go installed (1.26+)
- `gofmt` and `go vet` available in `PATH` (included with Go)
- `make` (optional, only for `make quality`)

## Quick start

```bash
cd demo/go
./scripts/quality_gate.sh
```

Alternative:

```bash
cd demo/go
make quality
```

If your machine has a toolchain mismatch, use:

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```

## Suggested prep sequence

1. Rehearse `docs/live-demo-script.md` exactly once with timer.
2. Start by showing `go.mod`, package layout, `demo/go/AGENTS.md`, and `demo/go/.github/copilot-instructions.md`.
3. Pick one task from `docs/micro-change-tasks.md` focused on tests, refactor, or concurrency.
4. Run the quality gate after each change and iterate to green.
5. Reserve closing minutes for tooling questions (`gofmt`, tests, modules) and takeaways.
