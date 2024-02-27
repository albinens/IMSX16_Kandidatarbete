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

func GetStatusOfAllRooms() []Room {
	tsClient := database.TSClient()
	db := database.GetDB()

	var rooms []RoomDBObject

	// Get all rooms from sql
	err := db.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		panic(err)
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

	currentOccupation := make(map[string]int64)

	for result.Next() {
		currentOccupation[result.Record().ValueByKey("room").(string)] = result.Record().Value().(int64)
	}

	convertedRooms := make([]Room, len(rooms)-1)

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
