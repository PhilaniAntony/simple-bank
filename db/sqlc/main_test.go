package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/PhilaniAntony/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	driver := config.DBDriver
	if driver == "" {
		driver = "postgres"
	}

	testDB, err = sql.Open(driver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
