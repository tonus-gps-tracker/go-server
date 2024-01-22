package influxdb

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/tonus-gps-tracker/server/internal/common"
)

type Connection struct {
	Client   influxdb2.Client
	WriteAPI api.WriteAPI
	QueryAPI api.QueryAPI
}

func (conn *Connection) setup() {
	dbHost := fmt.Sprintf("%s:%s", common.GetEnv("INFLUXDB_HOST"), common.GetEnv("INFLUXDB_PORT"))

	log.Printf("[INFO][InfluxDB] Connecting to %s\n", dbHost)

	client := influxdb2.NewClient(dbHost, common.GetEnv("INFLUXDB_TOKEN"))

	_, err := client.Health(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] Connection_Setup, influxdb2.NewClient: %s\n", err)
	}

	conn.Client = client
	conn.WriteAPI = client.WriteAPI(common.GetEnv("INFLUXDB_ORG"), common.GetEnv("INFLUXDB_BUCKET"))
	conn.QueryAPI = client.QueryAPI(common.GetEnv("INFLUXDB_ORG"))

	log.Println("[INFO][InfluxDB] Connected")
}

func NewConnection() *Connection {
	conn := new(Connection)
	conn.setup()

	return conn
}
