# Logistics checklist

## T-7 days

- Confirm your exact time slot and block order.
- Confirm you can demo Copilot in Agent Mode (if planned).
- Verify `AGENTS.md` and `.github/copilot-instructions.md` are visible in the repo.

## T-2 days

- Rehearse with timer (18 min target + 2 min buffer).
- Rehearse one "quality gate fails" scenario and recovery.
- Dry-run:
  - `cd demo/go`
  - `./scripts/quality_gate.sh`

## T-1 day

- Re-run full demo end to end.
- Keep fallback command ready:
  - `GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh`
- Prepare one static screenshot for each key moment:
  - AGENTS rules
  - first prompt
  - green quality gate output

## Event day

- Open these files before starting:
  - `demo/go/AGENTS.md`
  - `demo/go/.github/copilot-instructions.md`
  - `demo/go/internal/probe/service.go`
  - `demo/go/internal/probe/service_test.go`
- Keep a clock visible and preserve 60 seconds for closing takeaways.
