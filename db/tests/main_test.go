package db_test

import (
	db "Fin-Remittance/db/sqlc"
	"Fin-Remittance/utils"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cant load env config", err)
	}
	connection, err := sql.Open(config.DBdriver, config.DB_source)
	if err != nil {
		log.Fatal("could not connect to db", err)
	}
	testQuery = db.New(connection)
	//exit the process when u r done running test
	os.Exit(m.Run())
}
