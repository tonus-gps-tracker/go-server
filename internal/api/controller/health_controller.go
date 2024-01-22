package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (*HealthController) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
