# Stage 1: Build the application
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy all project files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/consumer ./cmd/consumer/consumer.go

# Stage 2: Create the lightweight image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/consumer /app/consumer
COPY --from=builder /app/.env /app/.env

# Ensure the binary is executable
RUN chmod +x /app/consumer

# Set the command to run the binary
ENTRYPOINT  ["./consumer"]