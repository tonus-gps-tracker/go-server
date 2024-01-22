package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/api/controller"
)

func GpsTrackerRoutes(rGroup *gin.RouterGroup, gpsTrackerController *controller.GpsTrackerController) {
	rGroup.POST("/gps-tracker", gpsTrackerController.Post)
}
