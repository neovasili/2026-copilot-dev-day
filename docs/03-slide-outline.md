# Slide outline for your Go segment

Keep it to 4 slides max.

## Slide 1 - The promise

- Copilot can be fast without becoming risky.
- Method: smallest possible change + automatic quality gate.

Speaker note:
- "The workflow is the product, not the single code snippet."

## Slide 2 - Guardrails in the repo

- `AGENTS.md`: coding rules and mandatory loop.
- `.github/copilot-instructions.md`: tells Copilot to apply those rules.
- Team style stays explicit and reusable.

## Slide 3 - Prompt + feedback loop

- Ask for one tiny scoped change.
- Run quality gate (`./scripts/quality_gate.sh`).
- If red: fix and repeat until green.

Example prompt:
- "Apply the smallest diff that improves X in `service.go`, keep behavior stable, then run `./scripts/quality_gate.sh` and iterate until it passes."

## Slide 4 - What to reuse tomorrow

- Keep changes small.
- Put standards in repo files.
- Never finish on generated code alone; finish on green checks.
