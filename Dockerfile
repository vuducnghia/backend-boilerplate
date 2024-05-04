FROM golang:1.22-alpine as builder
LABEL authors="James Vu"
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./config ./config
COPY ./logger ./logger
COPY ./models ./models
COPY ./src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin /app/src/main.go

FROM debian:buster-slim as app
LABEL authors="James Vu"

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
RUN apt update && apt install -y curl

COPY --from=builder /app/bin /app/bin
CMD ["/app/bin"]
