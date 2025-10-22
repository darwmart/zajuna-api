package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DSN     string
}

// LoadConfig carga las variables desde .env
func LoadConfig() *Config {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Println("⚠️ No se encontró el archivo .env, usando variables del sistema " + err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode,
	)

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DSN:     dsn,
	}
}

// getEnv devuelve una variable o su valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
