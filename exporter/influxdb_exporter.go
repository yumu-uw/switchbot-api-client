package exporter

import (
	"context"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var influxdb_url string
var influxdb_token string
var influxdb_org string
var influxdb_bucket string
var influxdb_client influxdb2.Client
var point_time time.Time

func InitInfluxDBExporter() {
	influxdb_url = os.Getenv("INFLUXDB_URL")
	influxdb_token = os.Getenv("INFLUXDB_TOKEN")
	influxdb_org = os.Getenv("INFLUXDB_ORG")
	influxdb_bucket = os.Getenv("INFLUXDB_BUCKET")
	influxdb_client = influxdb2.NewClient(influxdb_url, influxdb_token)
	point_time = time.Now()
}

func ExportToInfluxDB(tags map[string]string, fields map[string]interface{}) {
	writeAPI := influxdb_client.WriteAPIBlocking(influxdb_org, influxdb_bucket)
	point := influxdb2.NewPoint("measurement1", tags, fields, point_time)
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
