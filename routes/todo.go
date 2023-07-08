package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/narvikd/errorskit"
	"github.com/narvikd/fiberparser"
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
		errorString := fmt.Sprintf("%v is not a valid id", idParam)
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
	var body CreateRequestsBody

	err := fiberparser.ParseAndValidate(c, &body)
	if err != nil {
		errorskit.LogWrap(err, "CreateTodo invalid parameter")
		return FailResponse("Invalid request body", c, fiber.StatusBadRequest)
	}

	todoText := body.Todo
	if todoText == "" {
		return FailResponse("todo field is empty", c, fiber.StatusBadRequest)
	}

	data := CreateResponseMap(h.DB.CreateTodo(todoText), todoText)
	return SuccessResponse(data, c)
}
