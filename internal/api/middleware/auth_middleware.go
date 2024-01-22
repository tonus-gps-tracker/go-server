package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonus-gps-tracker/server/internal/common"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("x-api-secret") != common.GetEnv("HTTP_SERVER_SECRET") {
			c.String(http.StatusUnauthorized, "UNAUTHORIZED")
			c.Abort()
		}

		c.Next()
	}
}
