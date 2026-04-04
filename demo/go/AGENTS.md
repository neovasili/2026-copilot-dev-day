# AGENTS.md - Copilot Guardrails for This Go Workshop

## Mission

Make the smallest safe change that solves the requested task, then prove quality with tooling before stopping.

## Non-negotiable workflow

For every code change request:

1. Restate the task in one sentence.
2. Propose the smallest possible diff (avoid unrelated refactors).
3. Add tests to cover new code added.
4. Apply the change.
5. Run the quality gate:
   - `./scripts/quality_gate.sh`
6. If the quality gate fails, keep iterating automatically:
   - read the failing output
   - apply minimal corrective changes
   - re-run `./scripts/quality_gate.sh`
7. Stop only when the quality gate passes.

## Go style and design constraints

- Prefer idiomatic Go and standard library.
- Avoid single letter variable names, use meaningful variable names (exception: 't' for testing).
- Keep functions focused and readable.
- Keep code testable.
- Handle errors explicitly and wrap with `%w` when propagating.
- Keep public APIs stable unless asked otherwise.
- Preserve deterministic behavior in tests.
- Prefer table-driven tests.

## Scope discipline

- Do not rename files, move packages, or re-architect unless requested.
- Do not introduce dependencies unless requested.
- Do not change behavior outside the requested scope.

## Output expectations

When done, return:

- what changed (short)
- what quality commands were run
- final status (`PASS`/`FAIL`)
- any remaining risk or assumption
