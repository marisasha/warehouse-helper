ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DATABASE_URL ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=./migrations

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-up-one:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up 1

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(version)

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

print-db-url:
	@echo "postgres://$(DB_USERNAME):****@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)"

run:
	go run ./cmd/app