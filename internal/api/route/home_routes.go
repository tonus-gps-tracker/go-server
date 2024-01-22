package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/api/controller"
)

func HomeRoutes(rGroup *gin.RouterGroup) {
	homeController := new(controller.HomeController)
	rGroup.GET("/", homeController.Get)
	rGroup.GET("/home", homeController.Get)
}
