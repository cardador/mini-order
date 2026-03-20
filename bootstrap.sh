#!/bin/bash

# Exit on any error
set -e

# Use 'docker compose' (V2) instead of 'docker-compose' (V1)
DOCKER_CMD="docker compose"

echo "🚀 Starting Docker containers..."
$DOCKER_CMD up -d

echo "⏳ Waiting for PostgreSQL to be ready..."
# We target the service name 'postgres' defined in your docker-compose.yml
until docker exec $(docker ps -qf "name=postgres") pg_isready -U postgres; do
  echo "Database is starting up..."
  sleep 2
done

echo "📂 Creating 'orders' table..."
# Using the service name directly is safer
docker exec -i $(docker ps -qf "name=postgres") psql -U postgres <<EOF
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    item TEXT NOT NULL,
    amount DOUBLE PRECISION NOT NULL
);
EOF

echo "✅ Database initialized successfully!"
echo "🏃 You can now run: go run main.go"
