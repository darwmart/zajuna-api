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

func HasCapability(
	configRepo *repository.ConfigRepository,
	sessionRepo *repository.SessionsRepository,
	roleCapabilityRepository *repository.RoleCapabilityRepository,
	capability string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roles := []string{}
		allowed := false

		//Obtener ids por defecto desde configs
		defaultUserRoleID, err := configRepo.FindByName("defaultuserroleid")

		if err != nil {
			c.AbortWithError(405, err)
		}
		if defaultUserRoleID != nil {
			roles = append(roles, defaultUserRoleID.Value)
		}

		defaultFrontPageRoleID, err := configRepo.FindByName("defaultfrontpageroleid")
		log.Info(defaultUserRoleID.Value)

		if err != nil {
			c.AbortWithError(405, err)
		}
		if defaultFrontPageRoleID != nil {
			roles = append(roles, defaultFrontPageRoleID.Value)
		}

		//Obtener el token de la cookie para tomar el userID de la BD
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
		//log.Info(session)

		capabilities, err := roleCapabilityRepository.FindByUserID(int64(session.UserID), roles, capability)

		if err != nil {
			c.AbortWithError(405, err)
		}

		log.Info(capabilities)

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

// HasCapability revisa si un usuario tiene una capability específica
// basándose solo en los roles del usuario (sin contextos ni config global).
