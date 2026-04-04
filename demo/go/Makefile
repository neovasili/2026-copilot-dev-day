.PHONY: quality test lint fmt-check run

quality:
	./scripts/quality_gate.sh

test:
	go test ./...

lint:
	go vet ./...

fmt-check:
	@unformatted="$$(gofmt -l .)"; \
	test -z "$$unformatted" || (echo "Run gofmt -w on listed files"; echo "$$unformatted"; exit 1)

run:
	go run ./cmd/workshop
