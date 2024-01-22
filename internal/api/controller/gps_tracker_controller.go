package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/api/service"
)

type GpsTrackerController struct {
	gpsTrackerService *service.GpsTrackerService
}

func (controller *GpsTrackerController) Post(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[ERROR] GpsTrackerController_Post, io.ReadAll: %s\n", err)
	}

	err = controller.gpsTrackerService.Save(string(body))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "OK")
}

func NewGpsTrackerController(gpsTrackerService *service.GpsTrackerService) *GpsTrackerController {
	return &GpsTrackerController{
		gpsTrackerService: gpsTrackerService,
	}
}
