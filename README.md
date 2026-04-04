# Copilot Dev Days Caceres - Session Prep Repo

Prep materials for Juan Manuel Ruiz Fernandez's Go block in **Session 03** of GitHub Copilot Developer Days (Caceres edition).

## Confirmed context from event website

Source pages:
- https://copilot-dev-caceres.vercel.app/index.html
- https://copilot-dev-caceres.vercel.app/pages/agenda.html
- https://copilot-dev-caceres.vercel.app/pages/session-03.html

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
- `demo/go/AGENTS.md`: prewritten guardrails for Copilot behavior
- `demo/go/.github/copilot-instructions.md`: Copilot-compatible instruction hook
- `demo/go/scripts/quality_gate.sh`: format/lint/test quality gate
- `demo/go`: runnable mini-project for the workshop

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
2. Start the demo by showing `demo/go/AGENTS.md`.
3. Perform one micro-change, run quality gate, iterate to green.
4. Repeat once and close with the 3 takeaways.
