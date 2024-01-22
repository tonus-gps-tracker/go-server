package repository

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tonus-gps-tracker/server/internal/common"
	"github.com/tonus-gps-tracker/server/internal/dto"
	"github.com/tonus-gps-tracker/server/internal/influxdb"
)

type LocationRepository struct {
	connection *influxdb.Connection
}

func (repository *LocationRepository) Save(t time.Time, gpsData dto.GPSData) {
	dataPoint := influxdb2.NewPointWithMeasurement(common.GetEnv("INFLUXDB_MEASUREMENT"))
	dataPoint.AddField("latitude", gpsData.Latitude)
	dataPoint.AddField("longitude", gpsData.Longitude)
	dataPoint.AddField("altitude", gpsData.Altitude)
	dataPoint.AddField("speed", gpsData.Speed)
	dataPoint.AddField("nSatellites", gpsData.NSatellites)
	dataPoint.SetTime(t)

	repository.connection.WriteAPI.WritePoint(dataPoint)
}

func NewLocationRepository(connection *influxdb.Connection) *LocationRepository {
	return &LocationRepository{
		connection: connection,
	}
}
