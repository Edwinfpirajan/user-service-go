package providers

import (
	"log"

	"github.com/Edwinfpirajan/user-service-go/config"
	"github.com/Edwinfpirajan/user-service-go/internal/database"
	"github.com/Edwinfpirajan/user-service-go/internal/server"
)

type AppContainer struct {
	Config *config.Config
	DB     *database.GormDB
	Server *server.Server
}

// Inicializador de contenedor de dependencias
func Initialize() (*AppContainer, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.NewGormDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)
	if err != nil {
		return nil, err
	}

	log.Println("Dependencias inicializadas correctamente")

	return &AppContainer{
		Config: cfg,
		DB:     db,
		Server: server.NewServer(cfg),
	}, nil
}
