include .env
export

all: run

run:
	go run ./api/cmd/api

build:
	go build ./api/cmd/api

sqlc:
	sqlc generate

lint:
	golangci-lint run ./...

migrateup:
	migrate -database ${MAINDB_DSN} -path pkg/db/migrations up ${step}

migratedown:
	migrate -database ${MAINDB_DSN} -path pkg/db/migrations down ${step}

migrateforce:
	migrate -database ${MAINDB_DSN} -path pkg/db/migrations force ${version}

.PHONY: run build sqlc lint migrateup migratedown