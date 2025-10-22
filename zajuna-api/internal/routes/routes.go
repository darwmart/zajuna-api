package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"zajunaApi/internal/handlers"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")

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
	userService := services.NewUserService(userRepo, sessionRepo)
	userHandler := handlers.NewUserHandler(userService)

	// --- Rutas API ---
	api.GET("/categories", categoryHandler.GetCategories)
	api.GET("/courses", courseHandler.GetCourses)
	api.GET("/courses/:id/details", courseHandler.GetCourseDetails)
	api.GET("/users", userHandler.GetUsers)
	api.POST("/login", userHandler.Login)
	api.POST("/logout", userHandler.Logout)
}
