package middleware

import (
	"net/http"
	"strconv"

	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware es el middleware base que inyecta el servicio de permisos en el contexto
func PermissionMiddleware(permService *services.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Inyectar el servicio de permisos en el contexto de Gin
		c.Set("permissionService", permService)
		c.Next()
	}
}

// RequireCapability middleware que requiere una capacidad específica en el contexto del sistema
func RequireCapability(permService *services.PermissionService, capability string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener userID del contexto (debe ser seteado por un middleware de autenticación previo)
		userIDInterface, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
			c.Abort()
			return
		}

		userID, ok := userIDInterface.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de usuario"})
			c.Abort()
			return
		}

		// Obtener contexto del sistema
		systemCtx, err := permService.GetSystemContext()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener contexto del sistema"})
			c.Abort()
			return
		}

		// Verificar la capacidad
		err = permService.RequireCapability(userID, capability, systemCtx.ID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Permisos insuficientes",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireCapabilityInCourse middleware que requiere una capacidad en un curso específico
// El courseID debe venir como query param "courseId" o path param "courseid"
func RequireCapabilityInCourse(permService *services.PermissionService, capability string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInterface, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
			c.Abort()
			return
		}

		userID, ok := userIDInterface.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de usuario"})
			c.Abort()
			return
		}

		// Intentar obtener courseID desde diferentes fuentes
		var courseID int
		var err error

		// 1. Intentar desde query param "courseId"
		courseIDStr := c.Query("courseId")
		if courseIDStr == "" {
			// 2. Intentar desde path param ":courseid"
			courseIDStr = c.Param("courseid")
		}
		if courseIDStr == "" {
			// 3. Intentar desde path param ":id" (si la ruta es /courses/:id)
			courseIDStr = c.Param("id")
		}

		if courseIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó courseId"})
			c.Abort()
			return
		}

		courseID, err = strconv.Atoi(courseIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "courseId inválido"})
			c.Abort()
			return
		}

		// Verificar la capacidad en el curso
		err = permService.RequireCapabilityByCourse(userID, capability, courseID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Permisos insuficientes",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		// Guardar courseID en el contexto para uso posterior
		c.Set("courseID", courseID)
		c.Next()
	}
}

// RequireAdmin middleware que requiere que el usuario sea administrador del sitio
func RequireAdmin(permService *services.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInterface, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
			c.Abort()
			return
		}

		userID, ok := userIDInterface.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de usuario"})
			c.Abort()
			return
		}

		// Verificar si es admin
		isAdmin, err := permService.IsSiteAdmin(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar permisos de administrador"})
			c.Abort()
			return
		}

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Se requieren permisos de administrador"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAnyCapability middleware que requiere al menos una de las capacidades listadas
func RequireAnyCapability(permService *services.PermissionService, capabilities []string, contextLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDInterface, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
			c.Abort()
			return
		}

		userID, ok := userIDInterface.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de usuario"})
			c.Abort()
			return
		}

		// Obtener contexto según el nivel
		var contextID int
		var err error

		switch contextLevel {
		case models.ContextSystem:
			systemCtx, err := permService.GetSystemContext()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener contexto"})
				c.Abort()
				return
			}
			contextID = systemCtx.ID

		case models.ContextCourse:
			courseIDStr := c.Query("courseId")
			if courseIDStr == "" {
				courseIDStr = c.Param("courseid")
			}
			if courseIDStr == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó courseId"})
				c.Abort()
				return
			}

			courseID, err := strconv.Atoi(courseIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "courseId inválido"})
				c.Abort()
				return
			}

			ctx, err := permService.GetContextByLevelAndInstance(models.ContextCourse, courseID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener contexto del curso"})
				c.Abort()
				return
			}
			contextID = ctx.ID

		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nivel de contexto no soportado"})
			c.Abort()
			return
		}

		// Verificar si tiene alguna de las capacidades
		hasAny, err := permService.HasAnyCapability(userID, capabilities, contextID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar permisos"})
			c.Abort()
			return
		}

		if !hasAny {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permisos insuficientes: requiere una de las capacidades especificadas"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CanViewCourse middleware específico para verificar permiso de ver curso
func CanViewCourse(permService *services.PermissionService) gin.HandlerFunc {
	return RequireCapabilityInCourse(permService, "moodle/course:view")
}

// CanEditCourse middleware específico para verificar permiso de editar curso
func CanEditCourse(permService *services.PermissionService) gin.HandlerFunc {
	return RequireCapabilityInCourse(permService, "moodle/course:update")
}

// CanDeleteCourse middleware específico para verificar permiso de eliminar curso
func CanDeleteCourse(permService *services.PermissionService) gin.HandlerFunc {
	return RequireCapabilityInCourse(permService, "moodle/course:delete")
}

// CanManageCategories middleware específico para verificar permiso de gestionar categorías
func CanManageCategories(permService *services.PermissionService) gin.HandlerFunc {
	return RequireCapability(permService, "moodle/category:manage")
}

// CanEnrolUsers middleware específico para verificar permiso de inscribir usuarios
func CanEnrolUsers(permService *services.PermissionService) gin.HandlerFunc {
	return RequireCapabilityInCourse(permService, "enrol/manual:enrol")
}

// GetUserID obtiene el userID del contexto de Gin
func GetUserID(c *gin.Context) (int, bool) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	userID, ok := userIDInterface.(int)
	return userID, ok
}

// GetPermissionService obtiene el servicio de permisos del contexto de Gin
func GetPermissionService(c *gin.Context) (*services.PermissionService, bool) {
	permServiceInterface, exists := c.Get("permissionService")
	if !exists {
		return nil, false
	}

	permService, ok := permServiceInterface.(*services.PermissionService)
	return permService, ok
}
