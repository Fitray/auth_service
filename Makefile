include .env
export
export PROJECT_ROOT := $(shell pwd)
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable

gen:
	@protoc -I api/proto api/proto/*.proto --go_out=./api/gen/go --go_opt=paths=source_relative --go-grpc_out=./api/gen/go/ --go-grpc_opt=paths=source_relative

run:
	@go run cmd/server/main.go