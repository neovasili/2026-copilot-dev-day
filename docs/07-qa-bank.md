# Q&A bank for your session

## 1) "How do you keep Copilot changes small?"

Answer:
- Put scope rules in `AGENTS.md`.
- Prompt explicitly for "smallest possible diff".
- Reject broad or unrelated edits.

## 2) "How do you avoid over-trusting generated code?"

Answer:
- We never stop at code generation.
- We stop only on green quality gates (format, lint, tests).

## 3) "Why show AGENTS.md first?"

Answer:
- It makes conventions explicit and reusable.
- It aligns everyone: human reviewers and AI assistant.

## 4) "Is this only for Go?"

Answer:
- No. The pattern is language-agnostic.
- Replace quality commands with your stack equivalents.

## 5) "What if the quality gate fails repeatedly?"

Answer:
- Keep changes smaller.
- Ask Copilot to fix only the specific failing output.
- Re-run gate after each tiny fix.

## 6) "Can this fit into real team workflows?"

Answer:
- Yes. Keep `AGENTS.md` in repo, run checks in CI, and mirror the same loop locally.

## 7) "Do I need Agent Mode for this?"

Answer:
- Agent Mode helps automate the loop.
- Even without it, the same prompt discipline and manual gate checks still work.
