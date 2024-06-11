# Use the official Golang image as the base image for building the application
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY resource_monitor/go.mod resource_monitor/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY resource_monitor/. .

# Build the Go application
RUN go build -o main .

# Use a minimal image for running the application
FROM alpine:latest

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main /app/main

# Set the entry point to the binary
ENTRYPOINT ["/app/main"]