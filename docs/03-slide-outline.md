# Slide outline for your Go segment

Keep it to 4 slides max. This block is workshop-first.

## Slide 1 - What we will do in 20 minutes

- Build a small Go probe service with Copilot.
- Validate with tests.
- Improve readability without losing behavior.

Speaker note:
- Set expectation: "You can copy this pattern to any Go service in your stack."

## Slide 2 - Prompting pattern that works

- Give context: file + role + constraints.
- Ask for idiomatic Go and explicit error handling.
- Request tests or acceptance criteria.

Example prompt:
- "Implement `ProbeOnce` in Go. Keep context timeout, wrap errors with `%w`, and return deterministic results suitable for table-driven tests."

## Slide 3 - Human-in-the-loop quality loop

- Generate suggestion.
- Run tests.
- Accept/reject/refine.
- Keep project conventions above model convenience.

Speaker note:
- Mention one rejected suggestion live to model critical review.

## Slide 4 - Reusable takeaways

- Constrain prompts, do not over-specify every line.
- Use tests as safety rails.
- Keep changes small and reviewable.
- Copilot increases speed, not accountability transfer.
