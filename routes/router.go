package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinpita/go-todo/database"
)

type handler struct {
	DB *database.Database
}

func RegisterRoutes(app *fiber.App, db *database.Database) {
	todoApi := app.Group("/api/v1/todo")
	h := handler{DB: db}

	todoApi.Get("/", h.GetAll)
	todoApi.Get("/:id", h.GetTodo)

	todoApi.Post("/", h.CreateTodo)
}
