package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Edwinfpirajan/user-service-go/config"
)

type DBConnection struct {
	*gorm.DB
}

// Instancia de la base de datos
func NewGormDB(cfg *config.Config) (*DBConnection, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Bogota",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
	return &DBConnection{db}, nil
}
