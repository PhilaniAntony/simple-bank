# Load variables from app.env
include .env
export $(shell sed 's/=.*//' .env)

create-migrations:
	migrate create -ext sql -dir db/migration seq init_schema

postgres:
	docker run --name postgres17 -p 5432:5432 \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-d postgres:17.5-alpine

createdb:
	docker exec -it postgres17 createdb --username=$(DB_USER) --owner=$(DB_USER) $(shell echo $(DB_SOURCE) | sed -E 's/.*\/([^?]+).*/\1/')

dropdb:
	docker exec -it postgres17 dropdb $(shell echo $(DB_SOURCE) | sed -E 's/.*\/([^?]+).*/\1/')

migrate-up:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrate-up1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migrate-down:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migrate-down1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: create-migrations postgres createdb dropdb migrate-up migrate-up1 migrate-down migrate-down1 sqlc test server