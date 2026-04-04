# Logistics checklist

## T-7 days

- Confirm your block order (Python -> TypeScript -> Go) with organizers.
- Confirm exact speaking window in the room clock.
- Confirm internet availability and fallback plan.

## T-2 days

- Rehearse once with a timer (target 18 minutes + 2 minutes buffer).
- Rehearse once with a "delay scenario" (only 15 minutes).
- Verify all commands run from clean terminal.

## T-1 day

- Update local repo and run:
  - `cd demo/go`
  - `go test ./...`
  - `go run ./cmd/workshop`
- Prepare one backup screenshot for each critical step.
- Prepare local notes with the 3 key takeaways.

## Event day (before session)

- Close unrelated apps/tabs.
- Set terminal/editor font size for room visibility.
- Open the files you will use in advance:
  - `demo/go/internal/probe/service.go`
  - `demo/go/internal/probe/service_test.go`
  - `docs/04-live-demo-script.md`

## Event day (during session)

- Keep one clock visible.
- If over time, skip optional concurrency deep dive.
- End with reusable prompt pattern and test-first loop.
