#!/bin/bash

# FinSync Health Check Script
echo "🔍 Checking FinSync Application Health..."

# Check if containers are running
echo "📦 Checking Docker containers..."
if docker-compose ps | grep -q "Up"; then
    echo "✅ Docker containers are running"
else
    echo "❌ Docker containers are not running"
    echo "Run 'make up' or 'docker-compose up -d' to start"
    exit 1
fi

# Check application health
echo "🌐 Checking application health..."
if curl -f -s http://localhost:8080/ > /dev/null; then
    echo "✅ Application is responding"
else
    echo "❌ Application is not responding"
    echo "Check logs with 'make logs'"
fi

# Check database connection
echo "🗄️  Checking database connection..."
if docker-compose exec -T postgres pg_isready -U postgres > /dev/null 2>&1; then
    echo "✅ PostgreSQL is ready"
else
    echo "❌ PostgreSQL is not ready"
fi

# Check Redis connection
echo "🔴 Checking Redis connection..."
if docker-compose exec -T redis redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis is responding"
else
    echo "❌ Redis is not responding"
fi

echo "🎉 Health check completed!"
