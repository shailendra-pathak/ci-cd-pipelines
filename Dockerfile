# ---------- Build Stage ----------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o pipelines ./cmd/server

# ---------- Runtime Stage ----------
FROM alpine:3.23

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/pipelines .

EXPOSE 8080

CMD ["./pipelines"]