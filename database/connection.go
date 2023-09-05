package database

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/tonus-gps-tracker/server/utils"
)

var InfluxDB InfluxDBConnection

type InfluxDBConnection struct {
	Client   influxdb2.Client
	WriteAPI api.WriteAPI
	QueryAPI api.QueryAPI
}

func (conn *InfluxDBConnection) Setup() {
	dbHost := fmt.Sprintf("%s:%s", utils.GetEnv("INFLUXDB_HOST"), utils.GetEnv("INFLUXDB_PORT"))

	log.Printf("[INFO][InfluxDB] Connecting to %s\n", dbHost)

	client := influxdb2.NewClient(dbHost, utils.GetEnv("INFLUXDB_TOKEN"))

	_, err := client.Health(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] InfluxDBConnection_Setup, influxdb2.NewClient: %s\n", err)
	}

	InfluxDB.Client = client
	InfluxDB.WriteAPI = client.WriteAPI(utils.GetEnv("INFLUXDB_ORG"), utils.GetEnv("INFLUXDB_BUCKET"))
	InfluxDB.QueryAPI = client.QueryAPI(utils.GetEnv("INFLUXDB_ORG"))

	log.Println("[INFO][InfluxDB] Connected")
}

func NewInfluxDBConnection() {
	conn := new(InfluxDBConnection)
	conn.Setup()
}
