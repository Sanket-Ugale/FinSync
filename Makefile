# FinSync Makefile

.PHONY: help build up down logs clean dev test

# Default target
help:
	@echo "Available commands:"
	@echo "  build     - Build the Docker containers"
	@echo "  up        - Start the application (production mode)"
	@echo "  down      - Stop the application"
	@echo "  logs      - View logs"
	@echo "  clean     - Clean up containers, volumes, and images"
	@echo "  dev       - Start in development mode"
	@echo "  test      - Run tests"
	@echo "  db-shell  - Connect to PostgreSQL shell"
	@echo "  redis-cli - Connect to Redis CLI"

# Build containers
build:
	docker-compose build

# Start application (production)
up:
	docker-compose up -d

# Start in development mode
dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

# Stop application
down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Clean up everything
clean:
	docker-compose down -v --rmi all --remove-orphans
	docker system prune -f

# Connect to PostgreSQL
db-shell:
	docker-compose exec postgres psql -U postgres -d finsync

# Connect to Redis CLI
redis-cli:
	docker-compose exec redis redis-cli

# Run tests (when you add tests)
test:
	go test ./...
