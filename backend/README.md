# Backend

This is the backend of the project.
It is a RESTful API that is built using Go.
It is responsible for handling all the requests from the frontend and interacting with the database.

## Starting the backend

### Prerequisites

- Go 1.22 or higher
- Docker
- Docker Compose

### Configuration

Configuring the backend is done using environment variables. These can either be configured in the runtime (such as docker-compose) or in a `.env` file in the root of the backend directory. Some of these options _must_ be configured in order for the backend to work properly. These are marked as `(required)` in the table below.

The following environment variables are available:

| Environment variable | Description                                                                                                                                          | Default value | Example value            |
| :------------------: | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ------------- | ------------------------ |
|   `INFLUXDB_TOKEN`   | Token used to authenticate with the InfluxDB instance.                                                                                               | none          | YOUR_TOKEN               |
|    `INFLUXDB_URL`    | The URL of the InfluxDB instance. (required)                                                                                                         | none          | http://localhost:8086    |
|    `INFLUXDB_ORG`    | The org to use for InfluxDB.                                                                                                                         | liveinfo      | liveinfo                 |
|  `INFLUXDB_BUCKET`   | The InfluxDB to store data in.                                                                                                                       | liveinfo      | liveinfo                 |
|    `POSTGRES_DB`     | The internal database used by the Postgres instance.                                                                                                 | liveinfo      | liveinfo                 |
|   `POSTGRES_HOST`    | This is the host machine of the Postgres instance.                                                                                                   | localhost     | localhost                |
| `POSTGRES_USERNAME`  | The username used to authenticate with the Postgres database. (required)                                                                             | none          | liveinfo                 |
| `POSTGRES_PASSWORD`  | The passowrd used to authenticate with the Postgres database. (required)                                                                             | none          | SUPER_SECURE_PASSWORD123 |
|   `POSTGRES_PORT`    | The port used by the Postgres instance.                                                                                                              | 5432          | 5432                     |
|  `POSTGRES_SSLMODE`  | Whether or not SSL should be used when connecting to the database. Available options [here](https://www.postgresql.org/docs/current/libpq-ssl.html). | disable       | prefer                   |
|        `PORT`        | The port used by the server to listen for requests.                                                                                                  | 8080          | 3000                     |

### Running the backend

1. Clone the repository
2. Navigate to the backend directory
3. Start the databases using `docker-compose up -d`. This will start a PostgreSQL and a InfluxDB instance.
4. Copy the `.env.example` file to `.env` and configure the environment variables. The default values are fine for local development.
5. Run the backend using `go run .`. This will start the backend on port 8080.
