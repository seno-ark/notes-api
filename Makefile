#!make
include .env

go-dev:
	go run cmd/api/main.go

go-test:
	go test -v -cover ./...

go-mock:
	mockery --all --dir internal --output internal/mocks

go-swag:
	swag init --parseDependency --parseInternal --dir "cmd/api,internal/api" --output cmd/api/docs

go-build:
	go build -o bin/api cmd/api/main.go

migrate-file:
	migrate create -ext sql --dir pkg/database/migration -seq $(name)

migrate-up:
	migrate -path pkg/database/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path pkg/database/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

.PHONY: go-build go-dev go-test go-mock migrate-file migrate-up migrate-down