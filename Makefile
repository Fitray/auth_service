include .env
export
export PROJECT_ROOT := $(shell pwd)
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable

# Main commands

run: .env
	@go run $(PROJECT_ROOT)/cmd/app/main.go

env-up:
	@make docker-up
	@make wait-db
	@make wait-redis
	@make migrate-up

env-down:
	@make docker-down

gen:
	@protoc -I api/proto api/proto/*.proto --go_out=./api/gen/go --go_opt=paths=source_relative --go-grpc_out=./api/gen/go/ --go-grpc_opt=paths=source_relative

# Additional commands

docker-up:
	@docker compose up -d

docker-down:
	@docker compose down

docker-restart:
	@docker compose down
	@docker compose up -d

docker-logs:
	@docker compose logs -f

wait-db:
	@docker compose exec -T postgres \
		sh -c 'until pg_isready -U $$POSTGRES_USER -d $$POSTGRES_DB; do sleep 1; done'

wait-redis:
	@docker compose exec redis sh -c 'until redis-cli --user app --pass $$REDIS_PASSWORD ping; do sleep 1; done'

db-shell:
	@docker exec -it server-template-postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

redis-cli:
	@docker exec -it server-template-redis \
		redis-cli --user app --pass $(REDIS_PASSWORD)

logs-db:
	@docker logs server-template-postgres

logs-redis:
	@docker logs server-template-redis

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test ./...

build:
	@go build ./cmd/app

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=create_users"; \
		exit 1; \
	fi
	@docker compose run --rm migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq \
		$(name)
	
migrate-up:
	@docker compose run --rm migrate \
		-path /migrations \
		-database "$(DATABASE_URL)" \
		up

migrate-down:
	@docker compose run --rm migrate \
		-path /migrations \
		-database "$(DATABASE_URL)" \
		down

help:
	@echo "Main commands:"
	@echo "  env-up            Start Docker services, wait for DB, and apply migrations"
	@echo "  env-down          Stop Docker services"
	@echo "  run               Run application"
	@echo ""
	@echo "Docker:"
	@echo "  docker-up         Start Docker services"
	@echo "  docker-down       Stop Docker services"
	@echo "  docker-restart    Restart Docker services"
	@echo "  docker-logs       Show Docker logs"
	@echo ""
	@echo "Database:"
	@echo "  wait-db           Wait until PostgreSQL is ready"
	@echo "  db-shell          Open PostgreSQL shell"
	@echo "  logs-db           Show PostgreSQL logs"
	@echo ""
	@echo "Redis:"
	@echo "  wait-redis        Wait until Redis is ready"
	@echo "  redis-cli         Open Redis CLI"
	@echo "  logs-redis        Show Redis logs"
	@echo ""
	@echo "Migrations:"
	@echo "  migrate-create    Create migration (name=...)"
	@echo "  migrate-up        Apply all migrations"
	@echo "  migrate-down      Rollback last migration"
	@echo ""
	@echo "Go:"
	@echo "  fmt               Format code"
	@echo "  vet               Run go vet"
	@echo "  test              Run tests"
	@echo "  build             Build application"