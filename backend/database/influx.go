package database

import (
	"context"

	"example.com/m/v2/env"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func getClient() influxdb2.Client {
	return influxdb2.NewClient(env.InfluxDB.Url, env.InfluxDB.Token)
}

func WriteTimeSeriesData(p *write.Point) {
	client := getClient()
	defer client.Close()
	writeAPI := client.WriteAPI(env.InfluxDB.Org, env.InfluxDB.Bucket)
	writeAPI.WritePoint(p)
}

func TSQuery(query string) (*api.QueryTableResult, error) {
	client := getClient()
	defer client.Close()
	return client.QueryAPI("liveinfo").Query(context.Background(), query)
}
