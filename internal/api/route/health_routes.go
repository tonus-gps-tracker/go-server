package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/api/controller"
)

func HealthRoutes(rGroup *gin.RouterGroup) {
	healthController := new(controller.HealthController)

	rGroup.GET("/health", healthController.GetHealth)
}
