package auth

import (
	log "github.com/sirupsen/logrus"
)

type LdapAuth struct{}

func (l LdapAuth) PreventLocalPasswords() bool {
	return true
}

func (l LdapAuth) Login(username, password string) (bool, error) {
	log.Info("Verificando usuario CAS:", username)
	// Aquí iría la lógica real de autenticación LDAP
	return true, nil
}

func init() {
	Register("cas", LdapAuth{})
}
