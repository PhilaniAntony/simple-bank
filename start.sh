#!/bin/sh

set -e

echo "Waiting for Postgres..."
/app/wait-for.sh postgres:5432

echo "Running DB migrations..."
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "Starting API server..."
exec /app/main