package main

import (
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tonus-gps-tracker/server/database"
	"github.com/tonus-gps-tracker/server/utils"
)

type GPSData struct {
	latitude    float32
	longitude   float32
	altitude    float32
	speed       float32
	nSatellites int
}

type Tracker struct {
	delimiter string
}

func (tracker *Tracker) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func (tracker *Tracker) Save(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Tracker_Save, io.ReadAll: %s\n", err)
	}

	data := strings.Split(string(body), tracker.delimiter)

	if len(data) != reflect.TypeOf(GPSData{}).NumField() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gpsData := GPSData{
		latitude:    utils.StringToFloat32(data[0]),
		longitude:   utils.StringToFloat32(data[1]),
		altitude:    utils.StringToFloat32(data[2]),
		speed:       utils.StringToFloat32(data[3]),
		nSatellites: utils.StringToInt(data[4]),
	}

	tracker.save(gpsData)

	w.WriteHeader(http.StatusOK)
}

func (tracker *Tracker) save(gpsData GPSData) {
	dataPoint := influxdb2.NewPointWithMeasurement(utils.GetEnv("INFLUXDB_MEASUREMENT"))
	dataPoint.AddField("latitude", gpsData.latitude)
	dataPoint.AddField("longitude", gpsData.longitude)
	dataPoint.AddField("altitude", gpsData.altitude)
	dataPoint.AddField("speed", gpsData.speed)
	dataPoint.AddField("nSatellites", gpsData.nSatellites)
	dataPoint.SetTime(time.Now())

	database.InfluxDB.WriteAPI.WritePoint(dataPoint)
}

func NewTracker() *Tracker {
	tracker := new(Tracker)
	tracker.delimiter = ","

	return tracker
}
