package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"zajunaApi/internal/handlers"
	"zajunaApi/internal/middleware"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")

	// --- Config ---
	configRepo := repository.NewConfigRepository(db)

	// --- RoleCapability ---
	roleCapabilityRepo := repository.NewRoleCapabilityRepository(db)

	// --- Categorías ---
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// --- Cursos ---
	courseRepo := repository.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	// --- Sessions ---
	sessionRepo := repository.NewSessionsRepository(db)

	// --- Usuarios ---
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, sessionRepo, courseRepo)
	userHandler := handlers.NewUserHandler(userService)

	authMiddleware := middleware.RequireAuth(sessionRepo)

	// --- Rutas API ---
	api.GET("/categories", authMiddleware, middleware.HasCapability(configRepo, sessionRepo, roleCapabilityRepo, "moodle/site:readallmessages"), categoryHandler.GetCategories)
	api.GET("/courses", authMiddleware, courseHandler.GetCourses)
	api.GET("/courses/:id/details", authMiddleware, courseHandler.GetCourseDetails)
	api.DELETE("/courses", authMiddleware, courseHandler.DeleteCourses)
	api.GET("/users", authMiddleware, userHandler.GetUsers)
	api.DELETE("/users", authMiddleware, userHandler.DeleteUsers)
	api.PUT("/users/update", authMiddleware, userHandler.UpdateUsers)
	api.POST("/login", userHandler.Login)
	api.POST("/logout", userHandler.Logout)
}
