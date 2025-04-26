.PHONY: all fmt test lint build

build:
	go build -ldflags "-s -w" -trimpath ./cmd/noctx/

fmt:
	go tool golangci-lint fmt ./...

lint:
	go tool golangci-lint run ./...

test:
	go test -race ./...

test_coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
