# Start a new stage from scratch
FROM golang:1.18

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/assets ./assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/.env ./.env

EXPOSE 80

CMD ["./main"]