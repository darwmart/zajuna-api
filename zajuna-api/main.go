package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"zajunaApi/internal/handlers"
	"zajunaApi/internal/middleware"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services"
)

func main() {
	// --- Configuración de conexión a PostgreSQL ---
	dsn := "host=localhost user=postgres password=12345 dbname=zajuna port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
	}

	// --- Dependencias de Categorías ---
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// --- Dependencias de Cursos ---
	courseRepo := repository.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	// --- Dependencias de Usuarios ---
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// --- Inicializar servidor Gin ---
	router := gin.Default()

	// --- Middlewares globales ---
	router.Use(middleware.EnableCORS()) // Comentado para permitir acceso desde cualquier origen
	// router.Use(middleware.LoggingMiddleware())  // Comentado para deshabilitar logging

	// --- Rutas API ---
	api := router.Group("/api")
	{
		api.GET("/categories", categoryHandler.GetCategories)
		api.GET("/courses", courseHandler.GetCourses)
		//api.GET("/courses/:id/roles", courseHandler.GetCourseRoles)
		api.GET("/courses/:id/details", courseHandler.GetCourseDetails)
		api.GET("/users", userHandler.GetUsers)
	}

	// --- Inicio del servidor ---
	log.Println("servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error iniciando el servidor:", err)
	}
}
