.DEFAULT_GOAL := build

BINARY := $(shell go list -m)

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)

test: build
	go test ./...

coverage: test
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out