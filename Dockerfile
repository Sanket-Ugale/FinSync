# Use the official Golang image to create a build artifact.
# This build stage installs dependencies, builds the application,
# and performs any other necessary steps.
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image to run the built binary.
# This image will only contain the compiled binary and the assets required to run it.
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/.env .env

# Expose port 8080 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]
