BUILD_DIR=bin

postgres:
	docker run --name pg-xtracker -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it pg-xtracker createdb --username=root --owner=root xtracker

dropdb:
	docker exec -it pg-xtracker dropdb xtracker

create_migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/xtracker?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/xtracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

serve:
	go run main.go

test:
	go test ./... -v -cover -count=1

build:
	go build -o ./$(BUILD_DIR)/backend.bin

run:
	./$(BUILD_DIR)/backend.bin
