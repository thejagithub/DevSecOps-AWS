# Multi-stage build
# Stage 1: Build the binary
FROM golang:1.21-alpine AS builder

# Install ca-certificates for HTTPS support in minimal images
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy go mod files
COPY go.mod go.mod

# Copy source code
COPY main.go main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Stage 2: Create minimal runtime image
FROM scratch

# Copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary from builder
COPY --from=builder /app/app /app

# Create a non-root user (numeric ID for scratch image)
USER 1000

# Expose port
EXPOSE 8080

# Set environment variable
ENV APP_PORT=8080

# Run the application
ENTRYPOINT ["/app"]
