export

DATABASE_URL = postgres://postgres:1234@localhost:5433/postgres?sslmode=disable
MIGRATE_SOURCE = file://migration/postgres/auth

# Compose run
compose:
	docker compose run --rm migrator
	docker compose build --no-cache app
	docker compose up -d notify
	docker compose up -d app


# Local run
run:
	go run cmd/app/main.go

db:
	docker run --name no_api-db -e POSTGRES_PASSWORD=1234 -p 5433:5432 -d postgres

kafka:
	docker run -d \
  --name kafka \
  -p 9092:9092 \
  -p 9094:9094 \
  -e KAFKA_CFG_NODE_ID=0 \
  -e KAFKA_CFG_PROCESS_ROLES=controller,broker \
  -e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS='0@localhost:9093' \
  -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=false \
  -e KAFKA_CFG_LISTENERS='PLAINTEXT://:9092,CONTROLLER://:9093,EXTERNAL://:9094' \
  -e KAFKA_CFG_ADVERTISED_LISTENERS='PLAINTEXT://localhost:9092,EXTERNAL://localhost:9094' \
  -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP='CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT,PLAINTEXT:PLAINTEXT' \
  -e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER \
  bitnami/kafka:3.9

# kafka-init:
#   docker exec -it kafka kafka-topics.sh --create --topic auth-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4

migrate-up:
	migrate -database "$(DATABASE_URL)" -source "$(MIGRATE_SOURCE)" up

migrate-down:
	migrate -database "$(DATABASE_URL)" -source "$(MIGRATE_SOURCE)" down -all

migrate-drop:
	migrate -database "$(DATABASE_URL)" -source "$(MIGRATE_SOURCE)" drop -f