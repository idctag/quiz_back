#!/bin/sh
set -e

# Wait for PostgreSQL using the full URL
echo "Waiting for PostgreSQL..."
until psql "$DATABASE_URL" -c '\l' &> /dev/null; do
  echo "PostgreSQL not ready yet..."
  sleep 2
done

echo "Running migrations..."
migrate -path /app/db/migrations -database "$DATABASE_URL" up

echo "Starting application..."
exec ./app
