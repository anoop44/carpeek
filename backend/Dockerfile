# Multi-stage Dockerfile for CarPeek Backend

# Build stage
FROM golang:1.24-alpine AS builder

# Install git (needed for go modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o carpeek-backend cmd/server/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and tzdata for timezone support
RUN apk --no-cache add ca-certificates tzdata

# Create a directory for the app and set it as working directory
RUN mkdir -p /app
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/carpeek-backend .
COPY --from=builder /app/migrations ./migrations

# Set environment variables
ENV MIGRATIONS_PATH=/app/migrations

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./carpeek-backend"]