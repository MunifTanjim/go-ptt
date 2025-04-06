.PHONY: all fmt test

all: fmt vet test

fmt:
	go fmt ./...

lint:
	go vet

test:
	go test -v ./...
