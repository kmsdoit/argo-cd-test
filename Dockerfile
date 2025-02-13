FROM golang:1.23-alpine as builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED=1

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Fiber app
RUN go build -o go-test

# Use a minimal ARM64 Docker image to run the binary
#FROM --platform=linux/arm64 alpine:3.18
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/go-test .

# Expose the port that the app will run on
EXPOSE 3009

# Command to run the executable
CMD ["./go-test"]