package main

import (
	"database/sql"
	"log"
	"net/http"
	"zajunaApi/internal/handlers"
	"zajunaApi/internal/middleware"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=zajuna sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
	}
	defer db.Close()

	// --- Dependencias de categorías ---
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// --- Dependencias de cursos ---
	courseRepo := repository.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	// --- Dependencias de usuarios ---
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// --- Rutas ---
	mux := http.NewServeMux()
	mux.HandleFunc("/api/categories", categoryHandler.GetCategories)
	mux.HandleFunc("/api/courses", courseHandler.GetCourses)
	mux.HandleFunc("/api/users", userHandler.GetUsers)

	// --- Middleware global (CORS + Logging) ---
	handler := middleware.LoggingMiddleware(middleware.EnableCORS(mux))

	log.Println("Backend corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
