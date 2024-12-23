package main

import (
	"log"

	"github.com/Edwinfpirajan/user-service-go/cmd/providers"
)

func main() {
	// Inicializar el contenedor de dependencias
	container, err := providers.Initialize()
	if err != nil {
		log.Fatalf("Error al inicializar dependencias: %v", err)
	}

	// Iniciar el servidor
	container.Server.Start(container.Config.App.Port)

}
