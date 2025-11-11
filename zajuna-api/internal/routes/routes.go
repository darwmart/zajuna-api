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

	// --- Usuarios ---
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// --- Rutas API ---
	api.GET("/categories", categoryHandler.GetCategories)
	// api.GET("/courses/search", courseHandler.SearchCourses) // TODO: Implementar SearchCourses
	api.GET("/courses/:idnumber/details", courseHandler.GetCourseDetails)
	api.GET("/courses", courseHandler.GetCourses)
	api.DELETE("/courses", courseHandler.DeleteCourses)
	api.PUT("/courses", courseHandler.UpdateCourses)
	api.GET("/enrollments/course/:courseid", userHandler.GetEnrolledUsers)
	api.GET("/users", userHandler.GetUsers)
	api.DELETE("/users", userHandler.DeleteUsers)
	api.PUT("/users/update", userHandler.UpdateUsers)
	api.PUT("/users/:id/toggle-status", userHandler.ToggleUserStatus)
}
