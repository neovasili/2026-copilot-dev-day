# Micro-change tasks for live demo (Easy / Medium / Hard)

Use this to pick one task live based on remaining time.

## How to pick quickly

- 3-5 minutes left: choose **Easy**.
- 6-9 minutes left: choose **Medium**.
- 10-14 minutes left: choose **Hard**.

All tasks use the same loop:

1. Ask for the smallest possible diff.
2. Run `./scripts/quality_gate.sh`.
3. If failing, iterate until `PASS`.

---

## Easy - explicit empty URL validation

### Task 1 - Objective

Improve error clarity by returning a specific error for empty or whitespace-only URLs.

### Task 1 - Expected files touched

- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

### Task 1 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff to make ProbeOnce return a clear "url is empty" error when the input URL is empty or whitespace-only.
Keep existing behavior unchanged for all other URL cases.
Add/adjust table-driven tests for this new behavior.
Run ./scripts/quality_gate.sh and iterate with minimal fixes until PASS.
```

### Task 1 - Done criteria

- Empty URL and whitespace URL return an error containing `url is empty`.
- Existing tests still pass.
- Quality gate ends in `QUALITY GATE: PASS`.

### Task 1 - Why this is a good opener

- Very small diff.
- Easy to explain business value: better diagnostics.

---

## Medium - treat HTTP 5xx as probe errors

### Task 2 - Objective

Keep status code information, but mark 5xx responses as errors in `ProbeOnce`.

### Task 2 - Expected files touched

- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

### Task 2 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff so ProbeOnce treats HTTP 5xx responses as errors.
Requirements:
- preserve StatusCode and Duration in Result
- keep existing behavior for 2xx/4xx responses
- wrap the returned error with useful context
Add one focused test for a 503 response.
Run ./scripts/quality_gate.sh and iterate with minimal fixes until PASS.
```

### Task 2 - Done criteria

- A 503 response returns `err != nil` and `result.StatusCode == 503`.
- Non-5xx behavior stays unchanged.
- Quality gate ends in `QUALITY GATE: PASS`.

### Task 2 - Why this is good in mid-slot

- Still small scope.
- Demonstrates behavior change + test updates.

---

## Hard - return joined batch errors instead of only first error

### Task 3 - Objective

Improve `ProbeAll` so it returns a joined error containing all per-item probe errors (not just the first one), while preserving deterministic result order.

### Task 3 - Expected files touched

- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

### Task 3 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff in ProbeAll so batch errors include all per-item probe errors (use errors.Join), instead of only the first error.
Constraints:
- keep output order deterministic
- keep per-item Result.Err behavior
- avoid broad refactors
Add one focused test that proves multiple invalid URLs are represented in the returned batch error.
Run ./scripts/quality_gate.sh and iterate with minimal fixes until PASS.
```

### Task 3 - Done criteria

- Batch with multiple invalid URLs returns one joined error that reflects multiple failures.
- `results[i]` order still matches input URL order.
- Quality gate ends in `QUALITY GATE: PASS`.

### Task 3 - Why this is a strong closer

- Shows more advanced error semantics with minimal architecture change.
- Great example of “AI can help with tricky edits, but tests keep us safe.”

---

## Universal fallback prompt (if time is almost over)

```text
Read AGENTS.md and follow it strictly.
Do one minimal readability-only change in service.go with no behavior change.
Run ./scripts/quality_gate.sh and iterate until PASS.
```
