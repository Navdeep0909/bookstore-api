# syntax=docker/dockerfile:1
FROM golang:1.24.1-alpine

# Install CA certificates (needed if calling HTTPS APIs)
RUN apk update && apk add --no-cache ca-certificates

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first, then download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server

# Run the binary
CMD ["./main"]
