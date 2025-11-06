package middleware

import (
	"net/http"
	"zajunaApi/internal/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	CAP_INHERIT  = 0
	CAP_ALLOW    = 1
	CAP_PREVENT  = -1
	CAP_PROHIBIT = -1000
)

// HasCapability es un middleware de autorización basado en roles y capacidades.
//
// @Summary      Middleware de autorización
// @Description  Verifica si el usuario autenticado posee la capacidad requerida para acceder a la ruta.
// @Tags         auth
// @Security     CookieAuth
//
// @Param        capability   path      string   true   "Nombre de la capacidad requerida para acceder"
// @Failure      401          {string}  string   "Unauthorized – el usuario no tiene permisos o no está autenticado"
// @Failure      500          {string}  string   "Internal Server Error – error consultando la base de datos"
func HasCapability(
	configRepo repository.ConfigRepositoryInterface,
	sessionRepo repository.SessionsRepositoryInterface,
	roleCapabilityRepository repository.RoleCapabilityRepositoryInterface,
	capability string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roles := []string{}
		allowed := false

		//Obtener ids por defecto desde configs
		defaultUserRoleID, err := configRepo.FindByName("defaultuserroleid")

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if defaultUserRoleID != nil {
			roles = append(roles, defaultUserRoleID.Value)
		}

		defaultFrontPageRoleID, err := configRepo.FindByName("defaultfrontpageroleid")

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if defaultFrontPageRoleID != nil {
			roles = append(roles, defaultFrontPageRoleID.Value)
		}

		//Obtener el token de la cookie para tomar el userID de la BD
		token, err := c.Cookie("Authorization")

		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		session, err := sessionRepo.FindBySID(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		capabilities, err := roleCapabilityRepository.FindByUserID(int64(session.UserID), roles, capability)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		for _, capability := range *capabilities {
			switch capability.Permission {
			case CAP_PROHIBIT:
				// Si algún rol prohíbe, no hay permiso
				c.AbortWithStatus(http.StatusUnauthorized)
			case CAP_ALLOW:
				// Si al menos uno permite, lo marcamos
				allowed = true
			}
		}

		if !allowed {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
