package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/narvikd/errorskit"
	"strconv"
)

func (h handler) GetAll(c *fiber.Ctx) error {
	return SuccessResponse(h.DB.FetchAll(), c)
}

func (h handler) GetTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, errConvert := strconv.Atoi(idParam)
	if errConvert != nil {
		errorskit.LogWrap(errConvert, "GetTodo invalid parameter")
		errorString := fmt.Sprintf("id %v is not from a valid type", idParam)
		return FailResponse(errorString, c, fiber.StatusBadRequest)
	}

	todoText, err := h.DB.FetchTodo(id)
	if err != nil {
		errorskit.LogWrap(err, "GetTodo id not found")
		return FailResponse(err.Error(), c, fiber.StatusNotFound)
	}

	return SuccessResponse(CreateResponseMap(id, todoText), c)
}

func (h handler) CreateTodo(c *fiber.Ctx) error {
	todoText := c.FormValue("todo", "")

	if todoText == "" {
		err := errors.New("empty TODO cannot be created, 'todo' key is required")
		return FailResponse(err.Error(), c, fiber.StatusBadRequest)
	}

	id := h.DB.CreateTodo(todoText)

	return SuccessResponse(CreateResponseMap(id, todoText), c)
}
