package database

import (
	"context"
	"database/sql"

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
	populateSQL()
}

func GetDB() *sqlx.DB {
	return db
}

func populateSQL() {
	if !utils.IsProduction() {
		if hasPopulatedSQL() {
			return
		}

		tx, err := db.BeginTx(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		_, err = tx.Exec("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", "EG-2515", "d7:6c:09:bb:f0:c4", "NC")
		handleTXError(tx, err)
		_, err = tx.Exec("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", "EG-2516", "65:1f:01:8a:26:b6", "NC")
		handleTXError(tx, err)
		_, err = tx.Exec("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", "M1203A", "c0:fc:7b:de:85:7d", "Maskinhuset")
		handleTXError(tx, err)
		_, err = tx.Exec("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", "Vasa-G14", "a9:af:39:d2:f1:cb", "Vasa")
		handleTXError(tx, err)

		_, err = tx.Exec("INSERT INTO flags (name, value) VALUES ($1, $2)", "sql_test_rooms", "true")

		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}
}

func hasPopulatedSQL() bool {
	res, err := db.Queryx("SELECT value FROM flags WHERE name = 'sql_test_rooms'")
	if err != nil {
		panic(err)
	}

	var value string
	for res.Next() {
		err = res.Scan(&value)
		if err != nil {
			panic(err)
		}
	}

	if &value == nil {
		return false
	}

	return value == "true"
}

func handleTXError(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback()
		panic(err)
	}
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
