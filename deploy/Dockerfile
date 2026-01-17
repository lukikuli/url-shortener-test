FROM golang:1.23-alpine AS builder

# Security & reproducibility
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o app ./cmd/api



FROM gcr.io/distroless/base-debian12 AS runtime

# Non-root user (distroless provides this)
USER nonroot:nonroot

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app /app/app

# Expose HTTP port
EXPOSE 8080

# Run
ENTRYPOINT ["/app/app"]
