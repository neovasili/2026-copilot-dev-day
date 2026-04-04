# Go workshop demo

This mini-project is tuned for a 20-minute Copilot practical block.

## Learning goals

- Use Copilot with constraints in a real Go file.
- Keep behavior safe with tests.
- Apply idiomatic error handling and bounded concurrency.

## Files to use live

- `internal/probe/service.go`
- `internal/probe/service_test.go`
- `cmd/workshop/main.go`

## Demo commands

```bash
go test ./...
go run ./cmd/workshop
```

If your machine reports a toolchain mismatch (for example `compile version ... does not match go tool version ...`), pin one matching toolchain explicitly:

```bash
GOTOOLCHAIN=go1.26.0 go test ./...
GOTOOLCHAIN=go1.26.0 go run ./cmd/workshop
```

## Suggested live flow

1. Open `service.go` and explain the problem.
2. Prompt Copilot to refine `ProbeOnce` and `ProbeAll`.
3. Run tests.
4. Do a short readability pass and recap.

## Prompt starters

```text
Implement/refine ProbeOnce with these constraints:
- validate URL (http/https)
- timeout with context.WithTimeout
- wrap errors with %w
- keep output deterministic and test-friendly
```

```text
Suggest a simple worker-pool implementation for ProbeAll.
Requirements:
- preserve input order
- bounded workers
- clear code for workshop explanation
```
