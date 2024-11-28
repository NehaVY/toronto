
# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o app .

# Stage 2: Create a minimal image for running the application
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./app"]
