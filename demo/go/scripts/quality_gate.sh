#!/usr/bin/env bash

set -euo pipefail

# Colorize output
BLACK='\e[0;30m'
RED='\e[0;31m'
GREEN='\e[0;32m'
YELLOW='\e[0;33m'
BLUE='\e[0;34m'
MAGENTA='\e[0;35m'
CYAN='\e[0;36m'
WHITE='\e[0;37m'
GREY='\e[0;90m'
RESET='\e[0m'

gate_failed() {
  printf "\n${GREY}QUALITY GATE: ${RED}❌ FAIL${RESET}\n"
}
trap 'gate_failed' ERR

printf "\n${MAGENTA}Running Quality Gate...${RESET}\n"

if ! command -v go >/dev/null 2>&1; then
  printf "${RED}ERROR: go is required but not found in PATH${RESET}\n" >&2
  exit 1
fi

if ! command -v gofmt >/dev/null 2>&1; then
  printf "${RED}ERROR: gofmt is required but not found in PATH${RESET}\n" >&2
  exit 1
fi

printf "\n${MAGENTA}Format check (gofmt -l .)${RESET}\n"

UNFORMATTED="$(gofmt -l .)"
if [[ -n "$UNFORMATTED" ]]; then
  printf "${RED}Formatting issues detected in:${RESET}\n" >&2
  printf "${GREY}%s${RESET}\n" "$UNFORMATTED"
  printf "${YELLOW}Please run 'gofmt -w <files>' to fix formatting issues.${RESET}\n" >&2
  exit 1
fi

printf "\n${MAGENTA}Linting (go vet ./...)${RESET}\n"
go vet ./...

printf "\n${MAGENTA}Running tests (go test ./...)${RESET}\n"
grc go test ./... -v || go test ./... -v # Fallback to non-colored output if grc is missing

printf "\n${GREY}QUALITY GATE: ${GREEN}✅ PASS${RESET}\n"
