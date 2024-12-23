// server/server.go
package server

import (
	"fmt"
	"log"

	"github.com/Edwinfpirajan/user-service-go/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

// Iniciar el servidor
func (s *Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	log.Printf("Servidor escuchando en %s", address)
	return s.app.Listen(address)
}

// Instancia del servidor
func NewServer(cfg *config.Config) *Server {
	app := fiber.New()

	// Registrar middlewares globales
	// app.Use(func(c *fiber.Ctx) error {
	// 	log.Printf("Request: %s %s", c.Method(), c.Path())
	// 	return c.Next()
	// })

	// Registrar rutas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "UP",
			"version": cfg.App.Port,
		})
	})

	return &Server{app}
}
