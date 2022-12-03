### STAGE 1: Build ###
FROM golang:1.19.3 AS builder

WORKDIR /api

COPY api ./api
COPY go.mod .
COPY go.sum .
COPY pkg ./pkg
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o main ./api/cmd/api

### STAGE 2: Run ###
FROM alpine:3.17.0

WORKDIR /api

RUN touch .env

COPY --from=builder /api/main .
RUN chmod +x main
CMD ["./main"]
