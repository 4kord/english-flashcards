all: run

run:
	go run ./api/cmd/api

build:
	go build ./api/cmd/api

sqlc:
	sqlc generate

lint:
	golangci-lint run ./...

.PHONY: run