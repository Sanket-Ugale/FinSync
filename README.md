# Neosurge Portfolio Management Service

This is a simplified portfolio management service.

## Features

- User management (registration, authentication)
- Portfolio management (create, update, delete portfolios)
- Asset tracking (add, update, remove assets)
- Basic analytics (total portfolio value, average return)

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Redis

## Setup

1. Clone the repository:

git clone https://github.com/sanket-ugale/FinSync.git

cd FinSync

2. Install dependencies:
go mod tidy

3. Set up the environment variables:
- Copy the `.env.example` file to `.env`
- Fill in the required information in the `.env` file

4. Run the application:
go run main.go

The server will start on `http://localhost:80`.

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
