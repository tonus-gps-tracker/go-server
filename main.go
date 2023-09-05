package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tonus-gps-tracker/server/database"
	"github.com/tonus-gps-tracker/server/utils"
)

func main() {
	godotenv.Load()

	database.NewInfluxDBConnection()

	httpServer := new(HttpServer)
	go httpServer.Run()

	utils.WaitInterruption()
	onExit()
}

func onExit() {
	database.InfluxDB.WriteAPI.Flush()
	database.InfluxDB.Client.Close()

	log.Println("[INFO][InfluxDB] Connection closed")
	log.Println("[INFO][API] Server stopped")
}
