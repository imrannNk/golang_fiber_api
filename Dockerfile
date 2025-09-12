# ===== Build stage =====
FROM golang:1.25-alpine AS builder

# Set up environment
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o app

# ===== Run stage (minimal image) =====
FROM alpine:latest

# Set working dir
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port
EXPOSE 8080

# Start the app
CMD ["./app"]
