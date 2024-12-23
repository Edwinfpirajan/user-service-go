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

func Initialize() (*AppContainer, error) {
	// 1. Cargar configuración
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	// 2. Inicializar base de datos
	db, err := database.NewGormDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)
	if err != nil {
		return nil, err
	}

	log.Println("Dependencias inicializadas correctamente")

	// Devolver el contenedor con todas las dependencias
	return &AppContainer{
		Config: cfg,
		DB:     db,
		Server: server.NewServer(cfg),
	}, nil
}
