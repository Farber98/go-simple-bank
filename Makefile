# Run postgres container.
postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=toor -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:
	postgres createdb dropdb migrateup migratedown