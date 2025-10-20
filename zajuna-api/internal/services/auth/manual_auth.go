package auth

import (
	log "github.com/sirupsen/logrus"
)

type ManualAuth struct{}

func (m ManualAuth) PreventLocalPasswords() bool { return false }

func (m ManualAuth) Login(username, password string) (bool, error) {
	log.Info("Verificando usuario manual:", username)
	// Aquí podrías validar usuario y contraseña contra la BD
	return true, nil
}

func init() {
	Register("manual", ManualAuth{})
}
