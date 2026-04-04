# Live demo script (Go + Copilot)

Use this as a speaking track. Keep the pace high and practical.

## 1) Opening (60-90 sec)

Say:
- "In this block we will use Copilot on real Go code, validate with tests, and keep control with clear constraints."
- "Goal: useful patterns you can apply today, not a perfect final app."

## 2) Show the target file (2 min)

Open:
- `demo/go/internal/probe/service.go`

Say:
- "This service probes URLs with timeout and returns structured results."
- "We use an interface so tests do not depend on real networks."

## 3) Prompt Copilot for implementation/refinement (5-6 min)

Prompt A:
```text
Implement or improve ProbeOnce in idiomatic Go.
Constraints:
- validate URL and require http/https
- apply context.WithTimeout
- wrap external errors with fmt.Errorf("...: %w", err)
- return deterministic, test-friendly output
Do not add third-party dependencies.
```

Prompt B:
```text
Review ProbeAll and suggest a bounded worker-pool approach that preserves input order.
Keep code simple and readable for workshop explanation.
```

Talk track:
- Read suggestion before accepting.
- Accept chunks, not entire files at once.

## 4) Validate (3-4 min)

Run:
```bash
cd demo/go
go test ./...
```

Say:
- "Green tests are our permission to move fast."
- "When a suggestion fails, we iterate prompt + code, not blind accept."

## 5) Quality pass (2-3 min)

Prompt C:
```text
Propose a small readability refactor without changing behavior.
List risks before proposing edits.
```

Say:
- "Useful Copilot usage includes saying no to suggestions that add complexity."

## 6) Close (60 sec)

Say:
- "Three habits: constrain prompts, verify with tests, and keep human judgment in charge."
- "If you only keep one thing: ask for constraints and acceptance criteria, not just output."
