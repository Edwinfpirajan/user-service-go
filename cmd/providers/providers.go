package providers

import (
	"log"

	"github.com/Edwinfpirajan/user-service-go/config"
	"github.com/Edwinfpirajan/user-service-go/internal/database"
	"github.com/Edwinfpirajan/user-service-go/internal/server"
)

type AppContainer struct {
	Config *config.Config
	DB     *database.DBConnection
	Server *server.Server
}

// Inicializador de contenedor de dependencias
func Initialize() (*AppContainer, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.NewGormDB(cfg)
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	log.Println("Dependencias inicializadas correctamente")

	return &AppContainer{
		Config: cfg,
		DB:     db,
		Server: server.NewServer(cfg),
	}, nil
}
