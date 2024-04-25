package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Port string
var InitialAuthKey string

var InfluxDB struct {
	Token  string
	Url    string
	Org    string
	Bucket string
}

var Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func Load() {
	godotenv.Load()

	loadPostgres()
	loadInfluxDB()
	Port = withDefault("PORT", "8080")
	InitialAuthKey = withDefault("INITIAL_AUTH_KEY", "123456")
}

func loadPostgres() {
	Postgres.Username = verifyEnv("POSTGRES_USERNAME")
	Postgres.Password = verifyEnv("POSTGRES_PASSWORD")

	Postgres.Host = withDefault("POSTGRES_HOST", "localhost")
	Postgres.Port = withDefault("POSTGRES_PORT", "5432")
	Postgres.Database = withDefault("POSTGRES_DATABASE", "liveinfo")
	Postgres.SSLMode = withDefault("POSTGRES_SSLMODE", "disable")
}

func loadInfluxDB() {
	InfluxDB.Token = verifyEnv("INFLUXDB_TOKEN")
	InfluxDB.Url = verifyEnv("INFLUXDB_URL")

	InfluxDB.Org = withDefault("INFLUXDB_ORG", "liveinfo")
	InfluxDB.Bucket = withDefault("INFLUXDB_BUCKET", "liveinfo")
}

func withDefault(env, def string) string {
	if os.Getenv(env) == "" {
		return def
	}

	return os.Getenv(env)
}

func verifyEnv(env string) string {
	if os.Getenv(env) == "" {
		panic(env + " is not set. Please set it in .env file or in the environment variables.")
	}

	return os.Getenv(env)
}
