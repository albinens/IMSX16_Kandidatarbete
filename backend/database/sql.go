package database

import (
	"fmt"

	"example.com/m/v2/env"
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

CREATE TABLE IF NOT EXISTS api_keys (
	key TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS gateway_users (
	username TEXT PRIMARY KEY,
	password TEXT NOT NULL
);
`

const devSchemaAdditions = `
	CREATE TABLE IF NOT EXISTS flags (
		name TEXT PRIMARY KEY,
		value TEXT
	);
`

var db *sqlx.DB

func InitSQL() error {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.Postgres.Username,
		env.Postgres.Password,
		env.Postgres.Host,
		env.Postgres.Port,
		env.Postgres.Database,
		env.Postgres.SSLMode,
	)

	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return err
	}
	db = dbConn

	return pushSQLSchema()
}

func GetDB() *sqlx.DB {
	return db
}

func pushSQLSchema() error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	if !utils.IsProduction() {
		db.Exec(devSchemaAdditions)
	}

	return nil
}
