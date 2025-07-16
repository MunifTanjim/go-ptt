.PHONY: all fmt test build build-proto-gen

all: fmt vet test

fmt:
	go fmt ./...

lint:
	go vet

test:
	go test -v ./...

build: build-proto-gen
	go build -o bin/ptt cmd/ptt/main.go

build-proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		cmd/ptt/server/proto/ptt.proto
