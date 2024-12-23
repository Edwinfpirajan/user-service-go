package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type ConfigApp struct {
	Port int
}

type Config struct {
	App *ConfigApp
	DB  *ConfigDB
}

// Instancia de la configuración
func NewConfig() (*Config, error) {
	if err := loadEnvFile(); err != nil {
		log.Printf("Advertencia: %v", err)
	}

	appPortStr := os.Getenv("SERVER_PORT")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		log.Printf("No se pudo convertir SERVER_PORT a entero, usando valor por defecto 8080. Error: %v", err)
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Printf("No se pudo convertir DB_PORT a entero, usando valor por defecto 5432. Error: %v", err)
		dbPort = 5432
	}

	return &Config{
		App: &ConfigApp{
			Port: appPort,
		},
		DB: &ConfigDB{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}

// Cargar archivo .env
func loadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("no se encontró un archivo .env, usando variables de entorno existentes")
	}
	return nil
}
