# Load variables from .env if it exists
ifneq ("$(wildcard .env)","")
  include .env
  export $(shell sed 's/=.*//' .env)
endif

# Required environment variables
REQUIRED_VARS := DB_SOURCE DB_USER DB_PASSWORD

check-env:
	@$(foreach var,$(REQUIRED_VARS), \
		if [ -z "$(${var})" ]; then \
			echo "Error: Environment variable ${var} is not set."; \
			exit 1; \
		fi; \
	)

create-migrations:
	migrate create -ext sql -dir db/migration seq init_schema

postgres:
	docker run --name postgres17 --network bank-network -p 5432:5432 \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-d postgres:17.5-alpine

createdb: check-env
	docker exec -it postgres17 createdb --username=$(DB_USER) --owner=$(DB_USER) $(shell echo $(DB_SOURCE) | sed -E 's/.*\/([^?]+).*/\1/')

dropdb: check-env
	docker exec -it postgres17 dropdb $(shell echo $(DB_SOURCE) | sed -E 's/.*\/([^?]+).*/\1/')

migrate-up: check-env
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrate-up1: check-env
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migrate-down: check-env
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migrate-down1: check-env
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

compose-up:
	docker-compose --env-file app.env up --build

compose-down:
	docker-compose --env-file app.env down

.PHONY: check-env create-migrations postgres createdb dropdb migrate-up migrate-up1 migrate-down migrate-down1 sqlc test server compose-up compose-down
