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
- Because the page order is Python -> TypeScript -> Go, your likely live slot is approximately **19:20-19:40** (inferred from the 18:40-19:40 session window).

## Repo structure

- `docs/01-session-brief.md`: website extract summarized for your slot
- `docs/02-run-of-show-go-20min.md`: minute-by-minute facilitation plan
- `docs/03-slide-outline.md`: minimal slide story for a practical workshop
- `docs/04-live-demo-script.md`: live talking track + Copilot prompt script
- `docs/05-attendee-handout.md`: participant instructions
- `docs/06-logistics-checklist.md`: pre-flight checklist
- `docs/07-qa-bank.md`: high-probability questions + strong answers
- `demo/go`: runnable Go mini-project for your live block

## Quick start (demo project)

```bash
cd demo/go
go test ./...
go run ./cmd/workshop
```

If your local Go setup has toolchain mismatch errors, use:

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 go test ./...
GOTOOLCHAIN=go1.26.0 go run ./cmd/workshop
```

## Suggested prep sequence

1. Review `docs/01-session-brief.md` and `docs/02-run-of-show-go-20min.md`.
2. Rehearse once with `docs/04-live-demo-script.md` and `demo/go`.
3. Run checklist in `docs/06-logistics-checklist.md` the day before.
