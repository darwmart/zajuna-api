package services

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services/auth"

	"github.com/golang-jwt/jwt/v5"

	"github.com/amoghe/go-crypt"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

func (s *UserService) Login(username, password string) (string, error) {

	username = strings.TrimSpace(strings.ToLower(username))

	//VALIDACION SI EL USUARIO EXISTE EN BD
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("usuario no encontrado")
	}

	//VALIDACION SI EL USUARIO ESTA SUSPENDIDO
	if user.Suspended == 1 {
		return "", fmt.Errorf("%s", "Suspended Login: "+username)
	}

	plugin, ok := auth.Get(user.Auth)
	if !ok {
		return "", errors.New("método de autenticación no encontrado")
	}

	//VALIDACION DE PASSWORD CON SU HASH
	if PasswordVerify(password, user.Password) {
		log.Info("Contraseña válida")
	} else {
		log.Error("Contraseña incorrecta")
		return "", errors.New("credenciales inválidas")
	}
	success, err := plugin.Login(username, password)
	if err != nil || !success {
		return "", errors.New("credenciales inválidas")
	}
	log.Info("Autenticación exitosa con método:", user.Auth)

	// TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("Error al firmar token: " + err.Error())
	}
	return tokenString, nil
}

func PasswordVerify(password string, hash string) bool {
	// crypt.Crypt() generará un hash con la misma configuración ($6$, rounds, salt)
	generated, err := crypt.Crypt(password, hash)
	if err != nil {
		return false
	}
	// Si el hash generado es exactamente igual, la contraseña es válida
	return generated == hash
}
