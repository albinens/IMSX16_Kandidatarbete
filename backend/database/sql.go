package database

import (
	"example.com/m/v2/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS rooms (
	name TEXT PRIMARY KEY,
	sensor TEXT UNIQUE NOT NULL,
	building TEXT NOT NULL
);
`

const devSchemaAdditions = `
	CREATE TABLE IF NOT EXISTS flags (
		name TEXT PRIMARY KEY,
		value TEXT
	);
`

var db *sqlx.DB

func InitSQL() {
	dbConn, err := sqlx.Connect("postgres", "postgres://liveinfo:liveinfo@localhost:5432/liveinfo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = dbConn

	pushSQLSchema()
}

func GetDB() *sqlx.DB {
	return db
}

func pushSQLSchema() {
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}

	if !utils.IsProduction() {
		db.Exec(devSchemaAdditions)
	}
}
