FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /task-management-api cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /task-management-api /task-management-api

CMD ["/task-management-api"]
