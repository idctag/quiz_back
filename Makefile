-include .env

postgres:
	docker run --name quiz_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it quiz_postgres createdb --username=root --owner=root quiz

dropdb:
	docker exec -it quiz_postgres dropdb quiz

migrateup:
	migrate -path db/migrations -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DATABASE_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

