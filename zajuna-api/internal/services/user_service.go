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
	"golang.org/x/crypto/bcrypt"

	"github.com/amoghe/go-crypt"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	repo        repository.UserRepositoryInterface
	sessionRepo repository.SessionsRepositoryInterface
	courseRepo  repository.CourseRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface, sessionRepo repository.SessionsRepositoryInterface, courseRepo repository.CourseRepositoryInterface) *UserService {
	return &UserService{repo: repo, sessionRepo: sessionRepo, courseRepo: courseRepo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

// Login autentica a un usuario utilizando su username y password.
func (s *UserService) Login(r *http.Request, username, password string) (string, error) {

	username = strings.TrimSpace(strings.ToLower(username))

	//VALIDACION SI EL USUARIO EXISTE EN BD
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrUserNotFound
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
		return "", errors.New("error al buscar los cursos del usuario")
	}
	if countCourses == 0 {
		return "", errors.New("el usuario no tiene cursos vinculados")
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
		return "", ErrInvalidCredentials
	}
	log.Info("Autenticación exitosa con método:", user.Auth)

	// GENERAR Y FIRMAR EL TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		//"user": username,
		//"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := signToken(token, os.Getenv("SECRET"))

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

// Logout elimina una sesión basada en su token (SID).
func (s *UserService) Logout(sid string) (string, error) {
	err := s.sessionRepo.DeleteSession(sid)
	if err != nil {
		return "", err
	}
	return "Sesion deleted", nil
}

var cryptFunc = crypt.Crypt

// PasswordVerify compara una contraseña en texto plano con un hash almacenado.
// Internamente usa la función cryptFunc, que apunta a crypt.Crypt, lo cual permite
// simular errores durante los tests
func PasswordVerify(password string, hash string) bool {
	generated, err := cryptFunc(password, hash)
	if err != nil {
		return false
	}
	return generated == hash
}

// HashPassword genera un hash seguro usando bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
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

// normalizeLoopback convierte direcciones loopback IPv6 a IPv4.
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

var signToken = func(token *jwt.Token, secret string) (string, error) {
	return token.SignedString([]byte(secret))
}

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
