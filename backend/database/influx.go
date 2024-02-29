package database

import (
	"example.com/m/v2/env"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb2.Client

func InitTS() {
	client = influxdb2.NewClient(env.InfluxDB.Url, env.InfluxDB.Token)
}

func WriteTimeSeriesData(p *write.Point) {
	writeAPI := client.WriteAPI(env.InfluxDB.Org, env.InfluxDB.Bucket)
	writeAPI.WritePoint(p)
}

func TSReader() api.QueryAPI {
	return client.QueryAPI("liveinfo")
}
