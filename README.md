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

Assumption used in this repo:

- You (Juan Manuel Ruiz Fernandez) correspond to the Go speaker listed as Juan Manuel Ruiz.
- Based on the website order (Python -> TypeScript -> Go), your likely slot is around **19:20-19:40**.

## Session strategy in this repo

The content is optimized for your latest direction:

- keep code changes minimal
- enforce a strict prompt -> lint/test -> fix loop
- show repo-level AI guardrails before coding (`AGENTS.md`)

## Repo structure

- `docs/01-session-brief.md`: website extract summarized for your slot
- `docs/02-run-of-show-go-20min.md`: minute-by-minute facilitation plan
- `docs/03-slide-outline.md`: slide story (guardrails + quality loop)
- `docs/04-live-demo-script.md`: live talking track + prompt templates
- `docs/05-attendee-handout.md`: participant handout
- `docs/06-logistics-checklist.md`: pre-flight checklist
- `docs/07-qa-bank.md`: likely Q&A with concise answers
- `docs/08-micro-change-tasks.md`: concrete easy/medium/hard tasks for live selection
- `docs/09-instructions-files-explainer.md`: `AGENTS.md` vs `copilot-instructions.md` purpose and tradeoffs
- `demo/go/AGENTS.md`: prewritten guardrails for Copilot behavior
- `demo/go/.github/copilot-instructions.md`: Copilot-compatible instruction hook
- `demo/go/scripts/quality_gate.sh`: format/lint/test quality gate
- `demo/go`: runnable mini-project for the workshop

## Requirements

- GitHub account with GitHub Copilot access
- IDE with Copilot enabled (VS Code or JetBrains recommended)
- Go installed (1.22+; tested in this repo with 1.26.0)
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

1. Rehearse `docs/04-live-demo-script.md` exactly once with timer.
2. Start the demo by showing `docs/09-instructions-files-explainer.md` and `demo/go/AGENTS.md`.
3. Pick one task from `docs/08-micro-change-tasks.md` (easy/medium/hard).
4. Perform one micro-change, run quality gate, iterate to green.
5. If time allows, repeat with a second task and close with takeaways.
