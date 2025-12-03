package middleware

import "github.com/gin-gonic/gin"

// EnableCORS aplica las cabeceras CORS a todas las rutas
func EnableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Permitir los orígenes locales durante desarrollo
		// - Landing Page: localhost:5173 (Vite)
		// - Dashboard: localhost:3000 (React)
		// - API: localhost:8080
		allowedOrigin := ""
		if origin == "http://localhost:5173" ||
		   origin == "http://localhost:3000" ||
		   origin == "http://localhost:8080" {
			allowedOrigin = origin
		}

		// IMPORTANTE: Configurar headers ANTES de cualquier return
		if allowedOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-API-Key")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200") // 12 horas

		// Manejar preflight requests (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
