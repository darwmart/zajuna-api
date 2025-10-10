package middleware

import "github.com/gin-gonic/gin"

// EnableCORS aplica las cabeceras CORS a todas las rutas
func EnableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permitir cualquier origen (menos seguro pero más flexible para desarrollo)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")  // Versión restrictiva
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
