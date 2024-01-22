package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tonus-gps-tracker/server/internal/api/controller"
	"github.com/tonus-gps-tracker/server/internal/api/service"
	"github.com/tonus-gps-tracker/server/internal/common"
	"github.com/tonus-gps-tracker/server/internal/influxdb"
	"github.com/tonus-gps-tracker/server/internal/influxdb/repository"
	"github.com/tonus-gps-tracker/server/internal/server"
)

var influxDbConnection *influxdb.Connection

func main() {
	godotenv.Load()

	influxDbConnection = influxdb.NewConnection()

	startServer()

	common.WaitInterruption()
	onExit()
}

func startServer() {
	locationRepository := repository.NewLocationRepository(influxDbConnection)
	gpsTrackerService := service.NewGpsTrackerService(locationRepository)
	gpsTrackController := controller.NewGpsTrackerController(gpsTrackerService)

	httpServer := server.NewHttpServer(gpsTrackController)

	go httpServer.Run()
}

func onExit() {
	influxDbConnection.WriteAPI.Flush()
	influxDbConnection.Client.Close()

	log.Println("[INFO][InfluxDB] Connection closed")
	log.Println("[INFO][API] Server stopped")
}
