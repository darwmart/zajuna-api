package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthMiddleware valida la sesión de Moodle y setea el userID en el contexto
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sid string

		// Intentar obtener token desde Header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			// Formato esperado: "Bearer <sid>"
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				sid = parts[1]
				log.Printf("AuthMiddleware: Token desde Authorization header - SID=%s", sid[:min(10, len(sid))])
			}
		}

		// Si no hay token en el header, intentar obtenerlo de la cookie MoodleSession
		if sid == "" {
			cookieSID, err := c.Cookie("MoodleSession")
			if err == nil && cookieSID != "" {
				sid = cookieSID
				log.Printf("AuthMiddleware: Token desde cookie MoodleSession - SID=%s", sid[:min(10, len(sid))])
			}
		}

		// Si no hay token ni en header ni en cookie, retornar error
		if sid == "" {
			log.Printf("AuthMiddleware ERROR: No Authorization header ni cookie MoodleSession - Path: %s", c.Request.URL.Path)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			c.Abort()
			return
		}

		// Validar sesión en la BD de Moodle
		var session struct {
			UserID       int   `gorm:"column:userid"`
			TimeModified int64 `gorm:"column:timemodified"`
			State        int   `gorm:"column:state"`
		}

		err := db.Table("mdl_sessions").
			Select("userid, timemodified, state").
			Where("sid = ? AND state = 0", sid).
			First(&session).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("AuthMiddleware ERROR: Session NOT FOUND in DB - SID=%s", sid[:min(10, len(sid))])
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
			} else {
				log.Printf("AuthMiddleware ERROR: DB Error - %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating session"})
			}
			c.Abort()
			return
		}

		log.Printf("AuthMiddleware: Session FOUND - UserID=%d, SID=%s", session.UserID, sid[:min(10, len(sid))])

		// Verificar expiración de sesión (2 horas = 7200 segundos)
		sessionTimeout := int64(7200)
		expiration := session.TimeModified + sessionTimeout
		currentTime := time.Now().Unix()

		if currentTime > expiration {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}

		// CRÍTICO: Setear userID en el contexto para que los middlewares de permisos lo usen
		c.Set("userID", session.UserID)

		// Continuar con la siguiente función del middleware
		c.Next()
	}
}

// OptionalAuthMiddleware es similar a AuthMiddleware pero no bloquea si no hay token
// Útil para rutas que pueden ser accedidas con o sin autenticación
func OptionalAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// No hay token, continuar sin setear userID
			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Token inválido, continuar sin setear userID
			c.Next()
			return
		}
		sid := parts[1]

		var session struct {
			UserID       int   `gorm:"column:userid"`
			TimeModified int64 `gorm:"column:timemodified"`
			State        int   `gorm:"column:state"`
		}

		err := db.Table("mdl_sessions").
			Select("userid, timemodified, state").
			Where("sid = ? AND state = 0", sid).
			First(&session).Error

		if err == nil {
			// Verificar expiración
			sessionTimeout := int64(7200)
			expiration := session.TimeModified + sessionTimeout
			currentTime := time.Now().Unix()

			if currentTime <= expiration {
				// Sesión válida, setear userID
				c.Set("userID", session.UserID)
			}
		}

		c.Next()
	}
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
