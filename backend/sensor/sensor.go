package sensor

import (
	"errors"
	"time"

	"example.com/m/v2/database"
)

var cache map[string]string = nil
var lastCacheUpdate time.Time = time.Time{}

func RoomFromMac(mac string) (string, error) {
	if cache == nil || time.Since(lastCacheUpdate) > 10*time.Minute {
		if err := Update(); err != nil {
			return "", err
		}
	}
	room, ok := cache[mac]
	if !ok {
		return "", errors.New("mac is not part of cache")
	}

	return room, nil
}

func Update() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT sensor, name FROM rooms")
	if err != nil {
		return err
	}
	defer rows.Close()

	cache = make(map[string]string)

	for rows.Next() {
		var sensor string
		var name string
		err := rows.Scan(&sensor, &name)
		if err != nil {
			return err
		}

		cache[sensor] = name
	}

	lastCacheUpdate = time.Now()

	return nil
}
