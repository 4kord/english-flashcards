FROM golang:1.19.0-alpine3.16

WORKDIR /api

COPY api/ ./
RUN go build -o main ./cmd/api
CMD ["./main"]
