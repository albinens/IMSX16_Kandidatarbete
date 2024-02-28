package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb2.Client

func InitTS() {
	client = influxdb2.NewClient("http://localhost:8086", "yctwnvCURF2Trj_H1e1gqsyB65dSa701xYlnAw3sl94s4i8xIR1bfi7MgBzDppuAT0k5NO-YOPKSPReiVPIVLw==")
}

func WriteTimeSeriesData(p *write.Point) {
	writeAPI := client.WriteAPI("liveinfo", "liveinfo")
	writeAPI.WritePoint(p)
}

func TSReader() api.QueryAPI {
	return client.QueryAPI("liveinfo")
}
