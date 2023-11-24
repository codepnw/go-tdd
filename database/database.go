package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func ConnectDB() *Database {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DNS"))
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}

	return &Database{db: db}
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
