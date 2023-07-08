package routes

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func createErrorMap(err error) fiber.Map {
	return fiber.Map{
		"success": false,
		"error":   err.Error(),
	}
}

func (h handler) GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    h.DB.FetchAll(),
	})
}

func (h handler) GetTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, errConvert := strconv.Atoi(idParam)
	if errConvert != nil {
		err := fmt.Errorf("id %v is not from a valid type", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(createErrorMap(err))
	}

	todoText, err := h.DB.FetchTodo(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(createErrorMap(err))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"id":      id,
		"todo":    todoText,
	})
}

func (h handler) CreateTodo(c *fiber.Ctx) error {
	todoText := c.FormValue("todo", "")

	if todoText == "" {
		err := errors.New("Empty TODO cannot be created, 'todo' key is required")
		return c.Status(fiber.StatusBadRequest).JSON(createErrorMap(err))
	}

	id := h.DB.CreateTodo(todoText)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"id":      id,
		"todo":    todoText,
	})
}
