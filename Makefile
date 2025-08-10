export

DATABASE_URL = postgres://postgres:1234@localhost:5433/postgres?sslmode=disable
MIGRATE_SOURCE = file://migration/postgres/auth

# Compose run
compose:
	docker compose run --rm migrator
	docker compose build --no-cache app
	docker compose up -d app


# Local run
run:
	go run cmd/app/main.go

db:
	docker run --name no_api-db -e POSTGRES_PASSWORD=1234 -p 5433:5432 -d postgres

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4

migrate-up:
	migrate -database "$(DATABASE_URL)" -source "$(MIGRATE_SOURCE)" up

migrate-down:
	migrate -database "$(DATABASE_URL)" -source "$(MIGRATE_SOURCE)" down -all