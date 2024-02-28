package database

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var client influxdb2.Client

func InitTS() {
	client = influxdb2.NewClient("http://localhost:8086", "Z60qgr7SH8VNE22HpePF8WpL4I9P8blEF6P9mi1BXpVD3OEs4AonoeDgm-YT_GEFi32lBj0DXPRjmb_-YazDhA==")
	generateInitialTSData()
}

func TSClient() influxdb2.Client {
	return client
}

func TSWriter() api.WriteAPI {
	return client.WriteAPI("liveinfo", "liveinfo")
}

func TSReader() api.QueryAPI {
	return client.QueryAPI("liveinfo")
}

func generateInitialTSData() {
	writer := client.WriteAPIBlocking("liveinfo", "liveinfo")
	writer.WritePoint(context.Background(),
		influxdb2.NewPointWithMeasurement("status").
			AddTag("room", "4058").AddField("number_of_people", 4).
			SetTime(time.Now()))

	writer.WritePoint(context.Background(),
		influxdb2.NewPointWithMeasurement("status").
			AddTag("room", "4059").AddField("number_of_people", 0).
			SetTime(time.Now()))
}
