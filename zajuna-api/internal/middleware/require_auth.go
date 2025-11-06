package middleware

import (
	"net/http"
	"time"
	"zajunaApi/internal/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RequireAuth es un middleware de autenticación basado en sesiones.
//
// @Summary      Middleware de autenticación por sesión
// @Description  Valida la cookie "Authorization", busca la sesión en la BD y determina si el usuario está autenticado.
// @Tags         auth
//
// @Security     CookieAuth
//
// @Failure      401   {string}  string  "Unauthorized – cookie ausente o sesión inválida"
// @Failure      500   {string}  string  "Internal Server Error – error al consultar la BD"
func RequireAuth(sessionRepo repository.SessionsRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("En el Middleware")

		//Obtener el token
		token, err := c.Cookie("Authorization")

		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Info(token)
		session, err := sessionRepo.FindBySID(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if session == nil {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "Authorization",
				Value:    "",
				Path:     "/",
				Domain:   "",
				MaxAge:   -1,
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				Secure:   true, // debe coincidir con el original
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
