package room

import (
	"context"

	"example.com/m/v2/database"
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

type RoomTimeObject struct {
	Room           string
	Time           uint64
	NumberOfPeople uint8
}

func GetStatusOfAllRooms() {
	tsClient := database.TSClient()
	db := database.GetDB()

	var rooms []RoomDBObject

	// Get all rooms from sql
	err := db.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		panic(err)
	}

	for _, room := range rooms {
		println(room.Name)
		println(room.Sensor)
	}

	// Get status of all rooms from influx
	queryClient := tsClient.QueryAPI("liveinfo")
	query := `
		from(bucket: "liveinfo")
		|> range(start: -1h)
		|> filter(fn: (r) => r._measurement == "status")
		|> filter(fn: (r) => r._field == "number_of_people")
		|> last()
	`
	result, err := queryClient.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		println(result.Record().ValueByKey("room").(string))
		println(result.Record().Value().(int64))
	}
}
