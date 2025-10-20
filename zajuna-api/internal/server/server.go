package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"zajunaApi/internal/config"
	"zajunaApi/internal/middleware"
	"zajunaApi/internal/routes"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
	cfg    *config.Config
}

func New() *Server {
	cfg := config.LoadConfig()

	// Conectar BD
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la BD: %v", err)
	}
	log.Println("Conexión a la base de datos exitosa")

	// Inicializar Gin
	router := gin.Default()
	router.Use(middleware.EnableCORS())

	// Registrar rutas
	routes.RegisterRoutes(router, db)

	return &Server{
		router: router,
		db:     db,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.cfg.AppPort)
	log.Printf("Servidor corriendo en http://localhost%s", addr)
	return s.router.Run(addr)
}
