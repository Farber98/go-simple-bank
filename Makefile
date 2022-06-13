postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=toor -p 5432:5432 -d postgres:alpine

restartpostgres:
	docker start postgres

server:
	go run main.go

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

tests:
	go test -cover -race -v -count=1 ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go go-simple-bank/db/sqlc Store
	
.PHONY:
	postgres createdb dropdb migrateup migratedown server migrateup1 migratedown1