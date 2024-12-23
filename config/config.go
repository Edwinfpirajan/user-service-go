package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ConfigDB estructura de configuración de la base de datos.
type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// ConfigApp estructura de configuración de la aplicación.
type ConfigApp struct {
	Port int
}

// Config estructura principal de configuración de la aplicación.
type Config struct {
	App *ConfigApp
	DB  *ConfigDB
}

func NewConfig() (*Config, error) {
	if err := loadEnvFile(); err != nil {
		logWarning(err)
	}

	// Convertir SERVER_PORT a entero (si quieres manejarlo como int)
	appPortStr := os.Getenv("SERVER_PORT")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		logWarning(fmt.Errorf("no se pudo convertir SERVER_PORT a entero, usando valor por defecto 8080. Error: %v", err))
	}

	// Convertir DB_PORT a entero
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		logWarning(fmt.Errorf("no se pudo convertir DB_PORT a entero, usando valor por defecto 5432. Error: %v", err))
	}

	logSuccess("Configuración cargada correctamente")

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

func loadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("no se encontró un archivo .env, usando variables de entorno existentes")
	}
	logSuccess("Archivo .env cargado correctamente")
	return nil
}

// logSuccess imprime mensajes de éxito en verde
func logSuccess(message string) {
	fmt.Printf("\033[32m[SUCCESS]\033[0m %s\n", message)
}

// logWarning imprime mensajes de advertencia en amarillo
func logWarning(err error) {
	fmt.Printf("\033[33m[WARNING]\033[0m %s\n", err)
}
