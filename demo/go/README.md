# Go HTTP Probe Demo

This project implements a small HTTP probing service and a runnable example.
It is designed to be simple, testable, and idiomatic Go.

## What the project does

- probes one URL and reports status code, duration, and error
- probes multiple URLs with bounded concurrency
- preserves input order in batch results
- validates URLs before sending requests
- supports cancellation and per-request timeout via context

## How it is organized

- `cmd/workshop/main.go`: runnable example that starts a local test server and probes several URLs
- `internal/probe/service.go`: core `Service` implementation (`ProbeOnce`, `ProbeAll`, URL validation)
- `internal/probe/result.go`: `Result` model returned per URL
- `internal/probe/service_test.go`: unit and behavior tests
- `scripts/quality_gate.sh`: format, vet, and test checks

## How it works

### `ProbeOnce`

1. Validates the input URL.
2. Creates a request-scoped timeout context.
3. Builds a `GET` request and executes it through the injected `HTTPClient` interface.
4. Returns a `Result` with URL, status code, duration, and any wrapped error.

### `ProbeAll`

1. Normalizes worker count.
2. Sends indexed jobs to a worker pool through a channel.
3. Each worker calls `ProbeOnce`.
4. Stores each result at the original index to keep deterministic order.
5. Returns all results and a batch error when failures occur.

## Run the example

```bash
cd demo/go
go run ./cmd/workshop
```

The example prints a batch error plus one line per URL including status, duration, and per-item error.

## Run tests and checks

```bash
cd demo/go
go test ./...
```

```bash
cd demo/go
./scripts/quality_gate.sh
```

Alternative:

```bash
cd demo/go
make quality
```

If your machine reports a toolchain mismatch, pin one toolchain explicitly:

```bash
cd demo/go
GOTOOLCHAIN=go1.26.0 ./scripts/quality_gate.sh
```
