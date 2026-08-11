include .env
export
export PROJECT_ROOT := $(shell pwd)
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable

.PHONY: help gen fmt build test check run dev docker-up docker-down infra-up infra-down postgres-up postgres-down redis-up redis-down infra-logs

help:
	@echo "make dev          Start Postgres and Redis, then run the gRPC server"
	@echo "make docker-up    Start Postgres and Redis"
	@echo "make run          Run the gRPC server"
	@echo "make fmt          Format Go files"
	@echo "make check        Run tests and vet"
	@echo "make docker-down  Stop local infrastructure"

gen:
	@protoc -I api/proto api/proto/*.proto --go_out=./api/gen/go --go_opt=paths=source_relative --go-grpc_out=./api/gen/go/ --go-grpc_opt=paths=source_relative

fmt:
	@gofmt -w $$(find . -name '*.go' -not -path './api/gen/*')

build:
	@go build ./cmd/server

test:
	@go test ./...

check:
	@go test ./...
	@go vet ./...

run:
	@go run ./cmd/server

dev: infra-up run

docker-up: infra-up

docker-down: infra-down

infra-up:
	@docker compose up -d --wait postgres redis

infra-down:
	@docker compose down

postgres-up:
	@docker compose up -d --wait postgres

postgres-down:
	@docker compose stop postgres

redis-up:
	@docker compose up -d --wait redis

redis-down:
	@docker compose stop redis

infra-logs:
	@docker compose logs -f postgres redis
