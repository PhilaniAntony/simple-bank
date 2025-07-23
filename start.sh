#!/bin/sh

set -e

echo "Waiting for Postgres..."
/app/wait-for.sh postgres:5432 -- echo "Postgres is up"

echo "Running DB migrations..."
source /app/app.env
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "Starting API server..."
exec /app/main