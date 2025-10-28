package middleware

import (
	"net/http"
	"zajunaApi/internal/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequireAuth(sessionRepo *repository.SessionsRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("En el Middleware")

		//Obtener el token
		token, err := c.Cookie("Authorization")

		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		log.Info(token)
		session, err := sessionRepo.FindBySID(token)
		if err != nil {
			c.AbortWithError(405, err)
		}
		if session == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
