# Используем актуальный базовый образ с Go
FROM golang:1.22.8 AS builder

WORKDIR /app
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN swag init -g ./cmd/main.go

RUN go build -o medods ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/medods ./
COPY schema/* ./schema/
COPY .env ./

ENTRYPOINT ["./medods"]
