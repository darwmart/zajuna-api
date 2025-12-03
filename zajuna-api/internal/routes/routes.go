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
	// --- Inicializar servicios de repositorio ---

	// Categorías
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Cursos
	courseRepo := repository.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	// Usuarios
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Permisos (Sistema de permisos de Moodle)
	permRepo := repository.NewPermissionRepository(db)
	permService := services.NewPermissionService(permRepo)

	// Autenticación
	authHandler := handlers.NewAuthHandler(userRepo, permService)

	// --- Rutas públicas (sin autenticación) ---
	// Ruta de login - NO requiere autenticación
	router.POST("/api/login", authHandler.Login)

	// Ruta de logout - NO requiere autenticación estricta (puede fallar el middleware si la sesión ya expiró)
	router.POST("/api/logout", authHandler.Logout)

	// --- Grupo API con middleware de autenticación y permisos ---
	api := router.Group("/api")
	// Aplicar middleware de autenticación primero
	api.Use(middleware.AuthMiddleware(db))
	// Inyectar el servicio de permisos en todas las rutas autenticadas
	api.Use(middleware.PermissionMiddleware(permService))

	// --- Rutas de Categorías ---
	// Obtener categorías (requiere capacidad de ver categorías)
	api.GET("/categories",
		middleware.RequireCapability(permService, "moodle/category:viewcourselist"),
		categoryHandler.GetCategories)

	// Crear categorías (requiere capacidad de gestionar categorías)
	api.POST("/categories",
		middleware.CanManageCategories(permService),
		categoryHandler.CreateCategories)

	// Mover categoría (requiere capacidad de gestionar categorías)
	api.POST("/categories/move",
		middleware.CanManageCategories(permService),
		categoryHandler.MoveCategory)

	// --- Rutas de Cursos ---
	// Obtener detalles de un curso (requiere capacidad de ver curso)
	api.GET("/courses/:idnumber/details",
		courseHandler.GetCourseDetails) // Sin middleware por ahora, agregarlo si es necesario

	// Listar cursos (requiere capacidad de ver cursos)
	api.GET("/courses",
		courseHandler.GetCourses) // Sin middleware - permite listar cursos públicos

	// Eliminar cursos (requiere capacidad de eliminar cursos)
	api.DELETE("/courses",
		middleware.RequireCapability(permService, "moodle/course:delete"),
		courseHandler.DeleteCourses)

	// Actualizar cursos (requiere capacidad de actualizar cursos)
	api.PUT("/courses",
		middleware.RequireCapability(permService, "moodle/course:update"),
		courseHandler.UpdateCourses)

	// Mover cursos (requiere capacidad de gestionar cursos)
	api.POST("/courses/move",
		middleware.RequireCapability(permService, "moodle/course:update"),
		courseHandler.MoveCourses)

	// --- Rutas de Usuarios Inscritos ---
	// Obtener usuarios inscritos en un curso (requiere capacidad de ver participantes)
	api.GET("/enrollments/course/:courseid",
		middleware.RequireCapabilityInCourse(permService, "moodle/course:viewparticipants"),
		userHandler.GetEnrolledUsers)

	// --- Rutas de Usuarios ---
	// Listar usuarios (requiere ser administrador o tener capacidad de ver usuarios)
	api.GET("/users",
		middleware.RequireCapability(permService, "moodle/user:viewalldetails"),
		userHandler.GetUsers)

	// Eliminar usuarios (requiere ser administrador)
	api.DELETE("/users",
		middleware.RequireAdmin(permService),
		userHandler.DeleteUsers)

	// Actualizar usuarios (requiere capacidad de actualizar usuarios)
	api.PUT("/users/update",
		middleware.RequireCapability(permService, "moodle/user:update"),
		userHandler.UpdateUsers)

	// Alternar estado de usuario (requiere capacidad de gestionar usuarios)
	api.PUT("/users/:id/toggle-status",
		middleware.RequireCapability(permService, "moodle/user:update"),
		userHandler.ToggleUserStatus)

	// --- Rutas de búsqueda de cursos (comentado - implementar cuando esté listo) ---
	// api.GET("/courses/search",
	//	 middleware.RequireCapability(permService, "moodle/course:view"),
	//	 courseHandler.SearchCourses)

}
