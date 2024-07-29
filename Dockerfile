# Start a new stage from scratch
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy necessary folders
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/.env ./.env

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]