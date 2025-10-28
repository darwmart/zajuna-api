package middleware

import (
	"net/http"
	"time"
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
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "Authorization",
				Value:    "",
				Path:     "/", // o el path con el que se creó
				Domain:   "",
				MaxAge:   -1,
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				Secure:   true, // debe coincidir con el original
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
