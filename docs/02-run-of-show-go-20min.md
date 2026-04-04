# Run of show - Go block (20 min)

## Goal

Show one repeatable pattern for production work with Copilot:
- minimal scoped change
- mandatory quality gate (lint + tests)
- iterate until green

## Minute-by-minute plan

- 00:00-01:30: Framing
  - "Today is not about big refactors."
  - "It is about small diffs + continuous verification."

- 01:30-03:30: Guardrails first
  - Open `demo/go/AGENTS.md`.
  - Show the non-negotiable loop:
    1. smallest diff
    2. run `./scripts/quality_gate.sh`
    3. if fail, fix and repeat until pass

- 03:30-05:00: Show instruction wiring
  - Open `demo/go/.github/copilot-instructions.md`.
  - Explain this file tells Copilot to honor `AGENTS.md` before edits.

- 05:00-12:00: Prompt Copilot for a tiny change
  - Work on `demo/go/internal/probe/service.go`.
  - Ask for a minimal targeted improvement only.
  - Reject broad suggestions; keep one small accepted diff.

- 12:00-15:30: Enforce quality gate
  - Run `./scripts/quality_gate.sh` (or `make quality`).
  - If it fails, keep prompting for minimal fixes.
  - End this section only when gate passes.

- 15:30-18:30: One more micro-change
  - Do a second tiny enhancement (e.g., error message clarity or test edge case).
  - Re-run gate and confirm green again.

- 18:30-20:00: Close
  - 3 takeaways:
    - Small diffs are easier to trust.
    - AI speed must be paired with automated checks.
    - Team conventions belong in repo-level instructions (`AGENTS.md` or similar).

## Timing fallback

If you get only 15 minutes:
- Show `AGENTS.md`.
- Do one small change.
- Run quality gate and stop on green.

If you get only 10 minutes:
- Show `AGENTS.md` + one prompt + one quality gate run.
