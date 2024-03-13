FROM golang:1.22-alpine as builder
LABEL authors="James"
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./config ./config
COPY ./logger ./logger
COPY ./models ./models
COPY ./src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin /app/src/main.go

FROM debian:buster-slim
LABEL authors="James"

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin /app/bin
CMD ["/app/bin"]
