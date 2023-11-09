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

test:
	go test -v -cover ./tests/...

server:
	go run main.go


.PHONY: postgres createdb dropdb migrate_up migrate_down sqlc test server create_migrate