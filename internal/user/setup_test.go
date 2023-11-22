package user

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB *sql.DB
)

var (
	driver = "postgres"
	postgresDNS = "postgres://root:123456@localhost:5433/go_tdd?sslmode=disable"
)

func TestMain(m *testing.M) {
	dbCon, err := sql.Open(driver, postgresDNS)
	if err != nil {
		log.Fatal(err)
	}

	err = dbCon.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testDB = dbCon

	code := m.Run()

	err = testDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
