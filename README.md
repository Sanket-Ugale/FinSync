# FinSync - Portfolio Management Service

This is a comprehensive portfolio management service built with Go, PostgreSQL, and Redis.

## Features

- User management (registration, authentication with OTP)
- Portfolio management (create, update, delete portfolios)
- Asset tracking (add, update, remove assets)
- Basic analytics (total portfolio value, average return)
- JWT-based authentication
- Email notifications with OTP verification
- Docker support for easy deployment

## Prerequisites

- Docker and Docker Compose
- Go 1.20+ (for local development)

## Quick Start with Docker

1. **Clone the repository:**
   ```bash
   git clone https://github.com/sanket-ugale/FinSync.git
   cd FinSync
   ```

2. **Set up environment variables:**
   ```bash
   cp .env.example .env
   ```
   Update the `.env` file with your configuration (especially SMTP settings).

3. **Start the application:**
   ```bash
   # Using make (recommended)
   make up
   
   # Or using docker-compose directly
   docker-compose up -d
   ```

4. **View logs:**
   ```bash
   make logs
   # Or: docker-compose logs -f
   ```

The server will start on `http://localhost:8080`.

## Development Setup

For development with hot reload:

```bash
# Start in development mode
make dev

# Or using docker-compose
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up
```

## Database Management

```bash
# Connect to PostgreSQL shell
make db-shell

# Connect to Redis CLI
make redis-cli
```

## Local Development (without Docker)

If you prefer to run locally:

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Set up PostgreSQL and Redis locally**

3. **Update .env file** with local database URLs:
   ```
   DATABASE_URL=postgresql://username:password@localhost:5432/finsync?sslmode=disable
   REDIS_URL=localhost:6379
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## Docker Services

The application includes the following services:

- **app**: Main Go application (port 8080)
- **postgres**: PostgreSQL 15 database (port 5432)
- **redis**: Redis cache (port 6379)

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Application port | 8080 |
| `DATABASE_URL` | PostgreSQL connection string | - |
| `REDIS_URL` | Redis connection string | - |
| `JWT_SECRET` | JWT signing secret | - |
| `SMTP_*` | Email configuration | - |

## Available Make Commands

- `make build` - Build Docker containers
- `make up` - Start application (production)
- `make dev` - Start in development mode
- `make down` - Stop application
- `make logs` - View logs
- `make clean` - Clean up containers and volumes
- `make db-shell` - Connect to PostgreSQL
- `make redis-cli` - Connect to Redis CLI

## API Documentation

### Authentication

#### Register a new user
- `POST /auth/register`
- Body: `{ "email": "user@example.com", "password": "password123" }`

#### Login
- `POST /auth/login`
- Body: `{ "email": "user@example.com", "password": "password123" }`

#### Verify OTP
- `POST /auth/verify-otp`
- Body: `{ "email": "user@example.com", "otp": "123456" }`

### User Management

#### Get user details
- `GET /api/user`
- Header: `Authorization: Bearer <token>`

#### Update user
- `PUT /api/user`
- Header: `Authorization: Bearer <token>`
- Body: `{ "name": "John Doe" }`

### Portfolio Management

#### Create portfolio
- `POST /api/portfolio`
- Header: `Authorization: Bearer <token>`
- Body: `{ "name": "My Portfolio" }`

#### Get all portfolios
- `GET /api/portfolio`
- Header: `Authorization: Bearer <token>`

#### Get single portfolio
- `GET /api/portfolio/:id`
- Header: `Authorization: Bearer <token>`

#### Update portfolio
- `PUT /api/portfolio/:id`
- Header: `Authorization: Bearer <token>`
- Body: `{ "name": "Updated Portfolio Name" }`

#### Delete portfolio
- `DELETE /api/portfolio/:id`
- Header: `Authorization: Bearer <token>`

### Asset Management

#### Add asset to portfolio
- `POST /api/portfolio/:id/asset`
- Header: `Authorization: Bearer <token>`
- Body: `{ "name": "AAPL", "type": "stock", "quantity": 10, "value": 150.50 }`

#### Update asset
- `PUT /api/portfolio/:id/asset/:assetId`
- Header: `Authorization: Bearer <token>`
- Body: `{ "name": "AAPL", "type": "stock", "quantity": 15, "value": 155.75 }`

#### Delete asset
- `DELETE /api/portfolio/:id/asset/:assetId`
- Header: `Authorization: Bearer <token>`

#### Get asset details
- `GET /api/portfolio/:id/asset/:assetId`
- Header: `Authorization: Bearer <token>`

### Analytics

#### Get portfolio value
- `GET /api/portfolio/:id/value`
- Header: `Authorization: Bearer <token>`

#### Get portfolio return
- `GET /api/portfolio/:id/return?start_date=2023-01-01&end_date=2023-12-31`
- Header: `Authorization: Bearer <token>`
