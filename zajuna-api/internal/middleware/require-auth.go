package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequireAuth(c *gin.Context) {
	log.Info("En el Middleware")

	c.Next()
}
