package services

import (
	"errors"
	"fmt"
	"net"
	"net/http"
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
	repo        *repository.UserRepository
	sessionRepo *repository.SessionsRepository
	courseRepo  *repository.CourseRepository
}

func NewUserService(repo *repository.UserRepository, sessionRepo *repository.SessionsRepository, courseRepo *repository.CourseRepository) *UserService {
	return &UserService{repo: repo, sessionRepo: sessionRepo, courseRepo: courseRepo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

func (s *UserService) Login(r *http.Request, username, password string) (string, error) {

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

	//VALIDACION SI EL USUARIO TIENE CURSOS VINCULADOS
	countCourses, err := s.courseRepo.CountUserCourses(int(user.ID))
	if err != nil {
		return "", errors.New("Error al buscar los cursos del usuario")
	}
	if countCourses == 0 {
		return "", errors.New("El usuario no tiene cursos vinculados")
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

	// GENERAR Y FIRMAR EL TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		//"user": username,
		//"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("Error al firmar token: " + err.Error())
	}

	//GUARDAR TOKEN EN DB
	session := &models.Sessions{
		State:        0,
		SID:          tokenString,
		UserID:       user.ID,
		SessData:     nil,
		TimeCreated:  time.Now().Unix(),
		TimeModified: time.Now().Unix(),
		FirstIp:      getRemoteAddr(r),
		LastIp:       getRemoteAddr(r),
	}
	log.Info("TIEMPO EN SERVICE", time.Now().Add(time.Hour*3).Unix())
	err = s.sessionRepo.InsertSession(session)
	if err != nil {
		// Hubo un error al insertar
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) Logout(sid string) (string, error) {
	err := s.sessionRepo.DeleteSession(sid)
	if err != nil {
		return "", err
	}
	return "Sesion deleted", nil
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

// getRemoteAddr devuelve la IP del cliente remoto.
// Si no se puede determinar, devuelve "0.0.0.0".
func getRemoteAddr(r *http.Request) string {
	// Revisa la cabecera X-Forwarded-For
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		ip := strings.TrimSpace(parts[0])
		if net.ParseIP(ip) != nil {
			return normalizeLoopback(ip)
		}
	}

	// Usa la IP de la conexión TCP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && net.ParseIP(ip) != nil {
		return normalizeLoopback(ip)
	}

	return "0.0.0.0"
}

func normalizeLoopback(ip string) string {
	if ip == "::1" {
		return "127.0.0.1"
	}
	return ip
}

func (s *UserService) DeleteUsers(userIDs []int) error {
	return s.repo.DeleteUsers(userIDs)
}

func (s *UserService) UpdateUsers(users []models.User) (int64, error) {
	return s.repo.UpdateUsers(users)
}
