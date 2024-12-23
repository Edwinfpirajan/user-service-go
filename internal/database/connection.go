package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	*gorm.DB
}

// Instancia de la base de datos
func NewGormDB(host string, port int, user, password, dbName string) (*GormDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Bogota",
		host, port, user, password, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")

	return &GormDB{db}, nil
}
