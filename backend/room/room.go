package room

import (
	"context"
	"fmt"
	"time"

	"example.com/m/v2/database"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/pingcap/errors"
)

type Status string

const (
	Available Status = "available"
	Occupied  Status = "occupied"
	Booked    Status = "booked"
)

type Room struct {
	Room     string `json:"room"`
	Building string `json:"building"`
	Status   Status `json:"status"`
}

type RoomDBObject struct {
	Name     string
	Sensor   string
	Building string
}

func StatusOfAllRooms() ([]Room, error) {
	rooms, err := AllRooms()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all relevant rooms from the database")
	}

	currentOccupation, err := currentRoomOccupancy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get room occupancy data to determine room status")
	}

	roomsWithOccupancy := make([]Room, 0, len(rooms))

	for _, room := range rooms {
		var status Status
		if currentOccupation[room.Name] == 0 {
			status = Available
		} else {
			status = Occupied
		}

		roomsWithOccupancy = append(roomsWithOccupancy, Room{
			Room:     room.Name,
			Building: room.Building,
			Status:   status,
		})
	}

	return roomsWithOccupancy, nil
}

func AllRooms() ([]RoomDBObject, error) {
	db := database.GetDB()

	var rooms []RoomDBObject
	err := db.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		return nil, errors.Wrap(err, "failed to query database for all rooms")
	}

	return rooms, nil
}

func RoomByName(name string) (*RoomDBObject, error) {
	db := database.GetDB()

	var room RoomDBObject
	err := db.Get(&room, "SELECT * FROM rooms WHERE name = $1", name)
	if err != nil {
		return &RoomDBObject{}, errors.Wrap(err, "failed to query database for room by name")
	}

	return &room, nil
}

func CreateDataPoint(room string, numberOfPeople int64, time time.Time) *write.Point {
	p := influxdb2.NewPointWithMeasurement("status").AddTag("room", room).
		AddField("number_of_people", numberOfPeople).
		SetTime(time)
	return p
}

func AddRoom(name, sensor, building string) error {
	db := database.GetDB()

	_, err := db.Queryx("INSERT INTO rooms (name, sensor, building) VALUES ($1, $2, $3)", name, sensor, building)
	return err
}

func DeleteRoom(name string) error {
	db := database.GetDB()

	_, err := db.Queryx("DELETE FROM rooms WHERE name = $1", name)
	return err
}

func StatusOfRoom(name string) (*Room, error) {
	room, err := RoomByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get room by name")
	}

	occupancy, err := currentRoomOccupancy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current room occupancy")
	}

	status := occupationToStatus(occupancy[room.Name])

	return &Room{
		Room:     room.Name,
		Building: room.Building,
		Status:   status,
	}, nil
}

type WeekdayAverageRoomOccupancy map[string]map[string]float32
type roomOccupancyWithEntries struct {
	Occupancy int64
	Entries   int
}

func RoomOccupancyPerDayOfWeek(from, to time.Time) (WeekdayAverageRoomOccupancy, error) {
	result, err := dataBetween(from, to)
	if err != nil {
		return nil, err
	}

	start := time.Now()

	data := make(map[string]map[string]*roomOccupancyWithEntries, 0)

	for result.Next() {
		room := result.Record().ValueByKey("room").(string)
		day := result.Record().Time().Weekday().String()

		var occupancy int64
		if result.Record().Value() == nil {
			continue
		}
		occupancy = result.Record().Value().(int64)

		if _, ok := data[room]; !ok {
			addAllWeekDays(data, room)
		}

		roomDayData := data[room][day]
		roomDayData.Occupancy += occupancy
		roomDayData.Entries++
	}

	average := toWeekdayAverage(data)

	time_since := time.Since(start)
	fmt.Println("Time to calculate room occupancy per day of week: ", time_since)

	return average, nil
}

func dataBetween(from, to time.Time) (*api.QueryTableResult, error) {
	if from.After(to) {
		return nil, errors.New("from date must be before to date")
	}

	reader := database.TimeSeriesReader()
	query := fmt.Sprintf(`
		from(bucket: "liveinfo")
		|> range(start: %d, stop: %d)
		|> filter(fn: (r) => r._measurement == "status")
		|> filter(fn: (r) => r._field == "number_of_people") 
		|> group(columns: ["room", "_time"])
	`, from.Unix(), to.Unix())

	result, err := reader.Query(context.Background(), query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query influxdb for room occupancy per day of week")
	}

	return result, nil
}

func addAllWeekDays(data map[string]map[string]*roomOccupancyWithEntries, room string) {
	data[room] = make(map[string]*roomOccupancyWithEntries)

	for i := 0; i < 7; i++ {
		data[room][time.Weekday(i).String()] = &roomOccupancyWithEntries{
			Occupancy: 0,
			Entries:   0,
		}
	}
}

func toWeekdayAverage(data map[string]map[string]*roomOccupancyWithEntries) WeekdayAverageRoomOccupancy {
	total := make(WeekdayAverageRoomOccupancy)
	for room, days := range data {
		total[room] = make(map[string]float32)
		for day, entry := range days {
			average := float32(entry.Occupancy) / float32(entry.Entries)
			twoDecimals := float32(int(average*100)) / 100
			total[room][day] = twoDecimals
		}
	}

	return total
}

func currentRoomOccupancy() (map[string]int64, error) {
	reader := database.TimeSeriesReader()
	query := `
		from(bucket: "liveinfo")
		|> range(start: -31d)
		|> filter(fn: (r) => r._measurement == "status")
		|> filter(fn: (r) => r._field == "number_of_people")
		|> last()
	`
	result, err := reader.Query(context.Background(), query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query influxdb for current room occupancy")
	}

	currentOccupation := make(map[string]int64)

	for result.Next() {
		currentOccupation[result.Record().ValueByKey("room").(string)] = result.Record().Value().(int64)
	}

	return currentOccupation, nil
}

func occupationToStatus(occupation int64) Status {
	if occupation == 0 {
		return Available
	}
	return Occupied
}
