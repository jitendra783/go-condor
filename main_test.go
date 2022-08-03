package db

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fetal("cannot load config", err)

	}
	testDB, err := sql.open(config.dbDriver, config.dbSource)
	if err != nil {
		log.fetal("cannot connect to db", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
