package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/api/controller"
	"github.com/tonus-gps-tracker/server/internal/api/route"
	"github.com/tonus-gps-tracker/server/internal/common"
)

type HttpServer struct {
	gpsTrackerController *controller.GpsTrackerController
}

func (httpServer *HttpServer) registerRoutes(app *gin.Engine) {
	rootGroup := app.Group("/")
	apiGroup := rootGroup.Group("api")

	route.HomeRoutes(rootGroup)
	route.HealthRoutes(rootGroup)
	route.GpsTrackerRoutes(apiGroup, httpServer.gpsTrackerController)
}

func (httpServer *HttpServer) Run() {
	gin.SetMode(common.GetEnv("GIN_MODE"))
	gin.DisableConsoleColor()

	app := gin.Default()

	httpServer.registerRoutes(app)

	log.Println("[INFO][API] Server started")
	err := app.Run(":" + common.GetEnv("HTTP_SERVER_PORT"))

	if err != nil {
		log.Fatalf("[ERROR] HttpServer_Run, app.Run: %s\n", err)
	}
}

func NewHttpServer(gpsTrackerController *controller.GpsTrackerController) *HttpServer {
	return &HttpServer{
		gpsTrackerController: gpsTrackerController,
	}
}
