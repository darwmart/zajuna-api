package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	DSN     string
}

// LoadConfig detecta el entorno y carga el .env correcto
func LoadConfig() *Config {
	env := getEnv("APP_ENV", "development")

	envFile := fmt.Sprintf("internal/config/.env.%s", env)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("⚠️ No se encontró %s, usando variables del sistema", envFile)
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
		AppEnv:  env,
		AppPort: getEnv("APP_PORT", "8080"),
		DSN:     dsn,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
