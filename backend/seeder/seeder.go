package seeder

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"example.com/m/v2/database"
	"example.com/m/v2/flag"
	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

func SeedDevelopmentData() {
	if utils.IsProduction() {
		return
	}

	populateSQL()
	populateRoomTimeSeriesData()
}

func populateSQL() {
	if hasPopulatedSQL() {
		return
	}

	db := database.GetDB()

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

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	flag.GetFlag("sql_test_rooms").Set()
}

func hasPopulatedSQL() bool {
	return flag.GetFlag("sql_test_rooms").IsSet()
}

func handleTXError(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}

func populateRoomTimeSeriesData() {
	if hasPopulatedTimeSeries() {
		return
	}

	rooms, err := room.AllRooms()
	if err != nil {
		panic(err)
	}

	for _, r := range rooms {
		for i := 0; i < 100; i++ {
			p := room.CreateDataPoint(r.Name, rand.Int63n(7), randPreviousTime())
			database.WriteTimeSeriesData(p)
		}
	}

	flag.GetFlag("ts_test_rooms").Set()
}

func hasPopulatedTimeSeries() bool {
	return flag.GetFlag("ts_test_rooms").IsSet()
}

func randPreviousTime() time.Time {
	const week = time.Hour * 24 * 7
	offset := time.Duration(rand.Int63n(int64(week)))
	return time.Now().Add(-offset)
}
