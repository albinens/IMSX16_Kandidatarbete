package database

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb2.Client

func InitTS() {
	if os.Getenv("INFLUXDB_TOKEN") == "" {
		panic("INFLUXDB_TOKEN is not set")
	}

	client = influxdb2.NewClient("http://localhost:8086", os.Getenv("INFLUXDB_TOKEN"))
}

func WriteTimeSeriesData(p *write.Point) {
	writeAPI := client.WriteAPI("liveinfo", "liveinfo")
	writeAPI.WritePoint(p)
}

func TSReader() api.QueryAPI {
	return client.QueryAPI("liveinfo")
}
