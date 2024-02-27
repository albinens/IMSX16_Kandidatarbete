package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS rooms (
	name TEXT PRIMARY KEY,
	sensor TEXT,
	building TEXT
);
`

var db *sqlx.DB

func InitSQL() {
	dbPool, err := sqlx.Connect("postgres", "postgres://liveinfo:liveinfo@localhost:5432/liveinfo?sslmode=disable")
	if err != nil {
		panic(err)
	}

	_, err = dbPool.Exec(schema)

	if err != nil {
		panic(err)
	}

	db = dbPool
}

func GetDB() *sqlx.DB {
	return db
}
