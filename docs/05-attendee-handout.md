# Attendee handout (Go block)

## Before the session

- Bring laptop with GitHub Copilot enabled.
- Use VS Code or JetBrains with Go tooling installed.
- Have Go 1.22+ available.

## During the block

We will do three things:
1. Inspect a small Go service.
2. Use Copilot to implement/refine behavior.
3. Validate with tests and keep only useful suggestions.

## Commands used

```bash
cd demo/go
go test ./...
go run ./cmd/workshop
```

## Prompt pattern you can reuse

```text
Task: <what to build>
Context: <file/module>
Constraints:
- idiomatic Go
- explicit error handling
- no external deps
Acceptance criteria:
- tests pass
- output deterministic
- code is readable for team review
```

## After the event

Try this pattern on one real task next week:
- Pick a small production-adjacent task.
- Use the constraint template.
- Validate with tests and code review.
- Keep a short note of what prompt style worked best.
