package database

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client

func InitTS() {
	newClient := influxdb2.NewClient("http://localhost:8086", "6040q69W8P-oiQFKpo_ZMs1ghiK0uRpZl-lgPcQxRDISepfugOwezAdmenxT-oxkv81ZC1ATFk0ESf7TNyUeew==")

	writer := newClient.WriteAPIBlocking("liveinfo", "liveinfo")
	writer.WritePoint(context.Background(),
		influxdb2.NewPointWithMeasurement("status").
			AddTag("room", "4058").AddField("number_of_people", 4).
			SetTime(time.Now()))

	writer.WritePoint(context.Background(),
		influxdb2.NewPointWithMeasurement("status").
			AddTag("room", "4059").AddField("number_of_people", 0).
			SetTime(time.Now()))

	client = newClient
}

func TSClient() influxdb2.Client {
	return client
}
