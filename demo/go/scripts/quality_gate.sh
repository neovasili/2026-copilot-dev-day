#!/usr/bin/env bash
set -euo pipefail

if ! command -v go >/dev/null 2>&1; then
  echo "ERROR: go is required but not found in PATH"
  exit 1
fi

if ! command -v gofmt >/dev/null 2>&1; then
  echo "ERROR: gofmt is required but not found in PATH"
  exit 1
fi

echo "==> format check (gofmt -l .)"
UNFORMATTED="$(gofmt -l .)"
if [[ -n "$UNFORMATTED" ]]; then
  echo "Formatting issues detected in:"
  echo "$UNFORMATTED"
  echo "Run: gofmt -w <files>"
  exit 1
fi

echo "==> lint (go vet ./...)"
go vet ./...

echo "==> tests (go test ./...)"
go test ./...

echo "QUALITY GATE: PASS"
