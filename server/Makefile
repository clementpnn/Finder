include .env
export

RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
NC := \033[0m

MIGRATIONS_PATH := database/migrations

APP = main.go

build:
	go build

dev: run-local db-up
	air

dev-np: run-local db-up
	go run ${APP}

migrate:
	@echo "${YELLOW}===> Enter the name of the table to create : ${NC}"; \
	read table; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$table

db-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

db-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

run-local:
	docker-compose up -d

run:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
	db-up

add:
	@echo "${YELLOW}===> Enter the name of the package to add : ${NC}"; \
	read packageName; \
	go get $$packageName

clean:
	go mod tidy

generate:
	sqlc generate