# ============
# VARIABLES
# ============
DB_URL=postgres://sunil:password123@postgres:5432/pillnet?sslmode=disable
MIGRATIONS_DIR=internal/db/pg/migrations

# Docker compose override
DC=docker compose

# ============
# MIGRATIONS
# ============

# Create a migration file
migration:
	@if [ -z "$(name)" ]; then \
		echo "❌ Please provide a migration name: make migration name=create_users"; \
		exit 1; \
	fi
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Run migrations UP
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Run migrations DOWN (dangerous)
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

# Print migration status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status


# ============
# DOCKER COMMANDS
# ============

# Build all docker services (server + worker)
docker-build:
	$(DC) build

# Start all containers in background
docker-up:
	$(DC) up 

# Start and view logs live
docker-up-logs:
	$(DC) up

# Stop containers
docker-down:
	$(DC) down

# Restart (down + up)
docker-restart:
	$(DC) down
	$(DC) up -d

# ============
# DEV / UTIL
# ============

# Format Go code
fmt:
	go fmt ./...

# Run server locally
run-server:
	go run cmd/app/main.go

# Run worker locally
run-worker:
	go run cmd/worker/main.go

#Run all services
docker-up-all:
	docker compose up --build


