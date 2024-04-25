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

func SeedDevelopmentData() error {
	if utils.IsProduction() {
		return nil
	}

	if err := populateSQL(); err != nil {
		return err
	}
	if err := populateRoomTimeSeriesData(); err != nil {
		return err
	}

	return nil
}

func populateSQL() error {
	alreadyPopulated, err := hasPopulatedSQL()
	if err != nil {
		return err
	}

	if alreadyPopulated {
		return nil
	}

	db := database.GetDB()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	if err := addRooms(tx); err != nil {
		return err
	}

	if err := addApiKeys(tx); err != nil {
		return err
	}

	if err := addGatewayUsers(tx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	flag, err := flag.Get("sql_test_rooms")
	if err != nil {
		return err
	}

	flag.Set()

	return nil
}

func hasPopulatedSQL() (bool, error) {
	flag, err := flag.Get("sql_test_rooms")
	if err != nil {
		return false, err
	}

	return flag.IsSet(), nil
}

func addRooms(tx *sql.Tx) error {
	if err := addRoom(tx, "EG-2515", "d7:6c:09:bb:f0:c4", "NC"); err != nil {
		return err
	}

	if err := addRoom(tx, "EG-2516", "d7:6c:09:bb:ff:c4", "NC"); err != nil {
		return err
	}

	if err := addRoom(tx, "M1203A", "c0:fc:7b:de:85:7d", "Maskinhuset"); err != nil {
		return err
	}

	if err := addRoom(tx, "Vasa-G15", "a9:af:39:d2:f1:cb", "Vasa"); err != nil {
		return err
	}

	return nil
}

func addGatewayUsers(tx *sql.Tx) error {
	return addGatewayUser(tx, "test", "test")
}

func addRoom(tx *sql.Tx, name, sensor, building string) error {
	_, err := tx.Exec("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", name, sensor, building)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func addApiKeys(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO api_keys (key) VALUES ('super_secret_key')")
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func addGatewayUser(tx *sql.Tx, username, password string) error {
	_, err := tx.Exec("INSERT INTO gateway_users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func populateRoomTimeSeriesData() error {
	alreadyPopulated, err := hasPopulatedTimeSeries()
	if err != nil {
		return err
	}

	if alreadyPopulated {
		return nil
	}

	rooms, err := room.AllRooms()
	if err != nil {
		return err
	}

	for _, r := range rooms {
		for i := 0; i < 100; i++ {
			p := room.CreateDataPoint(r.Name, rand.Int63n(7), randPreviousTime())
			database.WriteTimeSeriesData(p)
		}
	}

	flag, err := flag.Get("ts_test_rooms")
	if err != nil {
		return err
	}

	flag.Set()

	return nil
}

func hasPopulatedTimeSeries() (bool, error) {
	flag, err := flag.Get("ts_test_rooms")
	if err != nil {
		return false, err
	}
	return flag.IsSet(), nil
}

func randPreviousTime() time.Time {
	const week = time.Hour * 24 * 7
	offset := time.Duration(rand.Int63n(int64(week)))
	return time.Now().Add(-offset)
}
