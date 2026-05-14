FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ # No go.sum yet as no external dependencies
RUN go mod download

COPY . .

RUN go build -o stress-test ./cmd/stress-test/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stress-test .

ENTRYPOINT ["./stress-test"]
