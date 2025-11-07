# syntax=docker/dockerfile:1

##########
# Build Stage
##########
FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata build-base

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

##########
# Runtime Stage
##########
FROM alpine:3.20

RUN addgroup -S app && adduser -S -G app app \
    && apk add --no-cache ca-certificates tzdata wget

WORKDIR /app

COPY --from=builder /out/server /app/server

ENV DB_USER=postgres \
    DB_PASSWORD=postgres \
    DB_HOST=postgres \
    DB_PORT=5432 \
    DB_NAME=appdb \
    DB_SSLMODE=disable \
    PORT=8080 \
    GIN_MODE=release

EXPOSE 8080
USER app

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- "http://127.0.0.1:${PORT}/health" >/dev/null || exit 1

ENTRYPOINT ["/app/server"]

