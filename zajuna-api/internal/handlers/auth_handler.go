package handlers

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha512_crypt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo    *repository.UserRepository
	permService *services.PermissionService
}

func NewAuthHandler(userRepo *repository.UserRepository, permService *services.PermissionService) *AuthHandler {
	return &AuthHandler{
		userRepo:    userRepo,
		permService: permService,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Success            bool          `json:"success"`
	Token              string        `json:"token,omitempty"`
	User               *UserResponse `json:"user,omitempty"`
	Error              string        `json:"error,omitempty"`
	IsAdmin            bool          `json:"isAdmin"`
	CanAccessDashboard bool          `json:"canAccessDashboard"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	IDNumber  string `json:"idnumber"`
}

// Logout invalida la sesión actual del usuario
// Elimina la sesión de la BD y la cookie del navegador
func (h *AuthHandler) Logout(c *gin.Context) {
	var sid string

	// Obtener SID desde Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			sid = parts[1]
		}
	}

	// Si no hay en header, obtener de cookie
	if sid == "" {
		cookieSID, err := c.Cookie("MoodleSession")
		if err == nil && cookieSID != "" {
			sid = cookieSID
		}
	}

	// Si hay SID, invalidar la sesión en la BD
	if sid != "" {
		// Eliminar la sesión de la tabla mdl_sessions
		err := h.userRepo.DB.Exec("DELETE FROM mdl_sessions WHERE sid = ?", sid).Error
		if err != nil {
			log.Printf("Error al eliminar sesión %s: %v", sid[:min(10, len(sid))], err)
		} else {
			log.Printf("Sesión %s eliminada correctamente", sid[:min(10, len(sid))])
		}
	}

	// Eliminar cookie MoodleSession del navegador
	c.SetCookie(
		"MoodleSession",
		"",
		-1,          // MaxAge negativo elimina la cookie
		"/",
		"",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout exitoso",
	})
}

// Login maneja la autenticación de usuarios
// Valida credenciales y verifica permisos para acceso al dashboard
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, LoginResponse{
			Success: false,
			Error:   "Datos de entrada inválidos",
		})
		return
	}

	// Buscar usuario por username o idnumber
	var user models.User
	err := h.userRepo.DB.Where("username = ? OR idnumber = ?", req.Username, req.Username).
		Where("deleted = 0").
		Where("suspended = 0").
		Where("confirmed = 1").
		First(&user).Error

	if err != nil {
		// Log para debug
		log.Printf("Usuario no encontrado o inactivo: %s (error: %v)", req.Username, err)
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Success: false,
			Error:   "Credenciales inválidas",
		})
		return
	}

	log.Printf("Usuario encontrado: ID=%d, Username=%s, Email=%s", user.ID, user.Username, user.Email)

	// Verificar contraseña
	// Moodle almacena passwords hasheados con bcrypt en el campo 'password'
	var passwordHash string
	err = h.userRepo.DB.Table("mdl_user").
		Select("password").
		Where("id = ?", user.ID).
		Scan(&passwordHash).Error

	if err != nil {
		log.Printf("Error al obtener hash de password para user ID %d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, LoginResponse{
			Success: false,
			Error:   "Error al verificar credenciales",
		})
		return
	}

	log.Printf("Hash de password obtenido: %s", passwordHash[:min(50, len(passwordHash))])

	// Verificar password (soporta bcrypt y SHA-512 de Moodle)
	err = verifyPassword(req.Password, passwordHash)
	if err != nil {
		log.Printf("Password incorrecta para user %s: %v", user.Username, err)
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Success: false,
			Error:   "Credenciales inválidas",
		})
		return
	}

	log.Printf("Password correcta para user %s", user.Username)

	// Verificar permisos para acceso al dashboard
	// Solo usuarios con rol de administrador o manager pueden acceder
	isAdmin, err := h.permService.IsSiteAdmin(int(user.ID))
	if err != nil {
		// Si hay error al verificar permisos, aún permitir login pero sin acceso al dashboard
		isAdmin = false
	}

	// Verificar si tiene capacidades de gestión
	canManageCategories := false
	canManageCourses := false
	canManageUsers := false

	systemCtx, err := h.permService.GetSystemContext()
	if err == nil {
		// Verificar capacidades administrativas
		canManageCategories, _ = h.permService.HasCapability(int(user.ID), "moodle/category:manage", systemCtx.ID)
		canManageCourses, _ = h.permService.HasCapability(int(user.ID), "moodle/course:update", systemCtx.ID)
		canManageUsers, _ = h.permService.HasCapability(int(user.ID), "moodle/user:update", systemCtx.ID)
	}

	// Un usuario puede acceder al dashboard si:
	// 1. Es administrador del sitio (site admin)
	// 2. Puede gestionar categorías
	// 3. Puede gestionar cursos
	// 4. Puede gestionar usuarios
	canAccessDashboard := isAdmin || canManageCategories || canManageCourses || canManageUsers

	// Crear sesión en Moodle (tabla mdl_sessions)
	token, err := h.createMoodleSession(user.ID, c.ClientIP())
	if err != nil {
		log.Printf("Error al crear sesión: %v", err)
		c.JSON(http.StatusInternalServerError, LoginResponse{
			Success: false,
			Error:   "Error al generar token de sesión",
		})
		return
	}

	log.Printf("Sesión creada exitosamente: SID=%s", token[:min(10, len(token))])

	// Establecer cookie MoodleSession (compatible con Moodle)
	// IMPORTANTE: Para localhost HTTP usamos SameSite=Lax
	// SameSite=None requiere Secure=true (solo funciona con HTTPS)
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"MoodleSession", // nombre de la cookie
		token,           // valor (SID)
		7200,            // maxAge en segundos (2 horas)
		"/",             // path
		"",              // domain (vacío para permitir subdominios en localhost)
		false,           // secure (false para localhost HTTP, true en producción con HTTPS)
		false,           // httpOnly (false para permitir acceso desde JS durante desarrollo)
	)

	log.Printf("Cookie MoodleSession establecida con SameSite=Lax: %s", token[:min(10, len(token))])

	// Respuesta exitosa
	c.JSON(http.StatusOK, LoginResponse{
		Success: true,
		Token:   token,
		User: &UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			IDNumber:  user.IDNumber,
		},
		IsAdmin:            isAdmin,
		CanAccessDashboard: canAccessDashboard,
	})
}

// createMoodleSession crea una sesión en la tabla mdl_sessions de Moodle y retorna el SID
// Este SID es compatible con el middleware AuthMiddleware que valida sesiones
func (h *AuthHandler) createMoodleSession(userID uint, clientIP string) (string, error) {
	// Generar un SID único de 26 caracteres (formato de Moodle)
	sid := generateRandomSID()

	// Obtener timestamp actual
	currentTime := time.Now().Unix()

	// Crear la sesión en mdl_sessions
	// state=0 significa sesión activa
	err := h.userRepo.DB.Exec(`
		INSERT INTO mdl_sessions (state, sid, userid, sessdata, timecreated, timemodified, firstip, lastip)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, 0, sid, userID, "", currentTime, currentTime, clientIP, clientIP).Error

	if err != nil {
		return "", fmt.Errorf("error al crear sesión en BD: %w", err)
	}

	return sid, nil
}

// generateRandomSID genera un Session ID aleatorio de 26 caracteres
// Compatible con el formato de sesiones de Moodle (igual que PHP session_id())
func generateRandomSID() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	const length = 26

	// Generar bytes aleatorios criptográficamente seguros
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Si falla crypto/rand, usar timestamp como último recurso
		log.Printf("Warning: crypto/rand failed, using timestamp fallback")
		timestamp := time.Now().UnixNano()
		for i := range randomBytes {
			randomBytes[i] = byte((timestamp + int64(i*137)) % 256)
		}
	}

	// Convertir bytes aleatorios a caracteres del charset
	sid := make([]byte, length)
	for i := 0; i < length; i++ {
		sid[i] = charset[int(randomBytes[i])%len(charset)]
	}

	return string(sid)
}

// Helper function para min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// verifyPassword verifica la contraseña contra el hash
// Soporta tanto bcrypt como SHA-512
func verifyPassword(password, hash string) error {
	// Determinar el tipo de hash por su prefijo
	if strings.HasPrefix(hash, "$2a$") || strings.HasPrefix(hash, "$2b$") || strings.HasPrefix(hash, "$2y$") {
		// Formato bcrypt
		return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	} else if strings.HasPrefix(hash, "$6$") {
		// Formato SHA-512 de Moodle: $6$rounds=N$salt$hash
		// Usar la librería crypt que implementa correctamente el algoritmo Unix crypt()
		crypter := crypt.SHA512.New()
		err := crypter.Verify(hash, []byte(password))
		if err != nil {
			return fmt.Errorf("contraseña incorrecta")
		}
		return nil
	}
	return fmt.Errorf("formato de hash no soportado")
}
