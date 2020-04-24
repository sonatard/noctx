.PHONY: all imports test lint

all: imports test lint

imports:
	goimports -w ./

test:
	go test ./...

lint:
	golangci-lint run ./...

