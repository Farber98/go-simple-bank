package db_test

import (
	"database/sql"
	db "go-simple-bank/db/sqlc"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:toor@localhost:5432/simple-bank?sslmode=disable"

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
