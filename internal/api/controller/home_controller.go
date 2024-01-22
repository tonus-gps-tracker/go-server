package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/common"
)

type HomeController struct{}

func (*HomeController) Get(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://%s:%s", c.Request.Host, common.GetEnv("GRAFANA_PORT")))
}
