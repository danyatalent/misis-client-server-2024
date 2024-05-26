# Use the official Golang image for building the Go application
FROM golang:1.22.3 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /question-service ./cmd/question/main.go

# Use a minimal Docker image to run the Go app
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder image
COPY --from=builder /question-service ./

# Copy the config file
COPY config/config.yaml ./config/config.yaml


# Expose port 80 for local development and testing
EXPOSE 80

CMD ["./question-service"]