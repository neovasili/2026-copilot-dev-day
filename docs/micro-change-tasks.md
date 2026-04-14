# Micro-change tasks for live demo (Easy / Medium / Hard)

Use this list to quickly pick a task based on remaining time.

## How to pick quickly

- 3-5 minutes: **Easy** (tests and conventions).
- 6-9 minutes: **Medium** (interfaces + idiomatic error handling).
- 10-14 minutes: **Hard** (light concurrency + aggregated errors).

All tasks use the same loop:

1. ask for the smallest possible diff
2. run `./scripts/quality_gate.sh`
3. if it fails, iterate until `PASS`

---

## Easy - complete table-driven tests

### Task 1 - Objective

Improve URL validation coverage with one small table-driven test.

### Task 1 - Expected files touched

- `demo/go/internal/probe/service_test.go`

### Task 1 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
Add one small table-driven test in service_test.go for validateURL edge cases.
Keep test names clear and assertions idiomatic.
Do not refactor production code unless needed.
Run ./scripts/quality_gate.sh and iterate until PASS.
```

### Task 1 - Done criteria

- The test covers at least two cases (for example invalid scheme and empty host).
- The diff is small and readable.
- Quality gate ends with `QUALITY GATE: PASS`.

### Task 1 - Why this is a good opener

- Demonstrates Copilot for test completion directly in the editor.
- Reinforces idiomatic Go testing style.

---

## Medium - interfaces + error wrapping

### Task 2 - Objective

Use a fake `HTTPClient` to verify wrapped errors in `ProbeOnce` with `errors.Is`.

### Task 2 - Expected files touched

- `demo/go/internal/probe/service_test.go`

### Task 2 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
In service_test.go, add a focused unit test that uses a fake HTTPClient implementation returning a sentinel error.
Verify ProbeOnce wraps the error and errors.Is works.
Keep changes minimal and idiomatic.
Run ./scripts/quality_gate.sh and iterate until PASS.
```

### Task 2 - Done criteria

- The test uses an interface-based `HTTPClient` fake defined in the test.
- Error wrapping is verified idiomatically with `errors.Is`.
- Quality gate ends with `QUALITY GATE: PASS`.

### Task 2 - Why this is good in mid-slot

- Connects interfaces, testability, and error handling in one small change.

---

## Hard - light concurrency with aggregated errors

### Task 3 - Objective

Improve `ProbeAll` to return an aggregated error with `errors.Join` when multiple URLs fail, while preserving deterministic order.

### Task 3 - Expected files touched

- `demo/go/internal/probe/service.go`
- `demo/go/internal/probe/service_test.go`

### Task 3 - Prompt to paste

```text
Read AGENTS.md and follow it strictly.
Apply the smallest possible diff in ProbeAll so batch errors include all per-item probe errors with errors.Join, not only the first error.
Constraints:
- keep deterministic result order
- keep per-item Result.Err behavior
- avoid broad refactors
Add one focused test proving that multiple invalid URLs are represented in the returned batch error.
Run ./scripts/quality_gate.sh and iterate until PASS.
```

### Task 3 - Done criteria

- The batch error reflects multiple failures.
- `results[i]` keeps input order.
- Quality gate ends with `QUALITY GATE: PASS`.

### Task 3 - Why this is a strong closer

- Shows light concurrency and idiomatic error handling in a scoped change.

---

## Bonus prompt - explore stdlib without leaving the editor

```text
Read AGENTS.md and follow it strictly.
From service.go, identify the key stdlib APIs used (context, net/http, errors, sync, net/url).
Suggest one minimal improvement grounded in those APIs, apply it, and run ./scripts/quality_gate.sh until PASS.
```
