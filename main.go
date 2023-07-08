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

	app := fiber.New()
	setupApp(app)

	err := app.Listen(addr)
	if err != nil {
		log.Fatalf("App could not be started on %v\n", addr)
	}
}

func setupApp(app *fiber.App) {
	app.Use(setupLogger())
	db := database.CreateDatabase()
	routes.RegisterRoutes(app, db)

}

func setupLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})
}
