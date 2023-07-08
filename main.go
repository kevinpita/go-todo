package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinpita/go-todo/database"
	"github.com/kevinpita/go-todo/routes"
	"log"
)

func main() {
	const addr = ":3000"

	err := setupApp(addr)
	if err != nil {
		log.Fatalf("App could not be started on %v\n", addr)
	}
}

func setupApp(addr string) error {
	app := fiber.New()
	app.Use(setupLogger())

	routes.RegisterRoutes(app, database.CreateDatabase())

	return app.Listen(addr)
}

func setupLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})
}
