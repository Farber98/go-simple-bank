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
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = db.New(testDb)

	os.Exit(m.Run())
}
