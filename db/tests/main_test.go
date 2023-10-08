package db_test

import (
	db "Fin-Remittance/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/finTech_postgres_db?sslmode=disable")

	if err != nil {
		log.Fatal("could not connect to db", err)
	}
	testQuery = db.New(connection)
	//exit the process when u r done running test
	os.Exit(m.Run())
}
