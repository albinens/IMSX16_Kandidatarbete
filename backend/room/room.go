package room

import (
	"context"
	"time"

	"example.com/m/v2/database"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Status string

const (
	Available Status = "available"
	Occupied  Status = "occupied"
	Booked    Status = "booked"
)

type Room struct {
	Room   string `json:"room"`
	Status Status `json:"status"`
}

type RoomDBObject struct {
	Name     string
	Sensor   string
	Building string
}

func StatusOfAllRooms() []Room {
	rooms, err := AllRooms()
	if err != nil {
		panic(err)
	}

	// Get status of all rooms from influx
	reader := database.TSReader()
	query := `
		from(bucket: "liveinfo")
		|> range(start: -7d)
		|> filter(fn: (r) => r._measurement == "status")
		|> filter(fn: (r) => r._field == "number_of_people")
		|> last()
	`
	result, err := reader.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	currentOccupation := make(map[string]int64)

	for result.Next() {
		currentOccupation[result.Record().ValueByKey("room").(string)] = result.Record().Value().(int64)
	}

	convertedRooms := make([]Room, 0, len(rooms))

	for _, room := range rooms {
		var status Status
		if currentOccupation[room.Name] == 0 {
			status = Available
		} else {
			status = Occupied
		}

		convertedRooms = append(convertedRooms, Room{
			Room:   room.Name,
			Status: status,
		})
	}

	return convertedRooms
}

func AllRooms() ([]RoomDBObject, error) {
	db := database.GetDB()

	var rooms []RoomDBObject
	err := db.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func CreateDataPoint(room string, numberOfPeople int64, time time.Time) *write.Point {
	p := influxdb2.NewPointWithMeasurement("status").AddTag("room", room).
		AddField("number_of_people", numberOfPeople).
		SetTime(time)
	return p
}
