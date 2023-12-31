package routes

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/narvikd/errorskit"
	"github.com/narvikd/fiberparser"
)

func userIdFromParam(c *fiber.Ctx) (int, error) {
	idParam := c.Params("id")
	id, errConvert := strconv.Atoi(idParam)
	if errConvert != nil {
		errorskit.LogWrap(errConvert, "useridfromparam invalid parameter")
		return 0, fmt.Errorf("%v is not a valid id", idParam)
	}

	return id, nil
}

func todoTextFromBody(c *fiber.Ctx) (string, error) {
	var body TodoRequestBody

	err := fiberparser.ParseAndValidate(c, &body)
	if err != nil {
		errorskit.LogWrap(err, "todotextfrombody invalid parameter")
		return "", errors.New("invalid request body")
	}

	todoText := body.Todo
	if todoText == "" {
		log.Println("todotextfrombody empty todo field")
		return "", errors.New("todo field is empty")
	}

	return todoText, nil
}

func (h handler) GetAll(c *fiber.Ctx) error {
	todoList := h.DB.FetchAll()

	if len(todoList) == 0 {
		return FailResponse("there are no TODOs to be returned", c, fiber.StatusNotFound)
	}

	return SuccessResponse(todoList, c)
}

func (h handler) GetTodo(c *fiber.Ctx) error {
	id, errReadParam := userIdFromParam(c)
	if errReadParam != nil {
		errorskit.LogWrap(errReadParam, "gettodo")
		return FailResponse(errReadParam.Error(), c, fiber.StatusBadRequest)
	}

	todoText, err := h.DB.FetchTodo(id)
	if err != nil {
		errorskit.LogWrap(err, "gettodo id not found")
		return FailResponse(err.Error(), c, fiber.StatusNotFound)
	}

	return SuccessResponse(CreateResponseMap(id, todoText), c)
}

func (h handler) CreateTodo(c *fiber.Ctx) error {
	todoText, errParseBody := todoTextFromBody(c)
	if errParseBody != nil {
		errorskit.LogWrap(errParseBody, "createtodo")
		return FailResponse(errParseBody.Error(), c, fiber.StatusBadRequest)
	}

	data := CreateResponseMap(h.DB.CreateTodo(todoText), todoText)
	return SuccessResponse(data, c)
}

func (h handler) UpdateTodo(c *fiber.Ctx) error {
	id, errReadParam := userIdFromParam(c)
	if errReadParam != nil {
		errorskit.LogWrap(errReadParam, "udpatetodo")
		return FailResponse(errReadParam.Error(), c, fiber.StatusBadRequest)
	}

	todoText, errParseBody := todoTextFromBody(c)
	if errParseBody != nil {
		errorskit.LogWrap(errParseBody, "updatetodo")
		return FailResponse(errParseBody.Error(), c, fiber.StatusBadRequest)
	}

	err := h.DB.UpdateTodo(id, todoText)
	if err != nil {
		errorskit.LogWrap(err, "updatetodo id not found")
		return FailResponse(err.Error(), c, fiber.StatusNotFound)
	}

	return SuccessResponse(CreateResponseMap(id, todoText), c)
}

func (h handler) DeleteTodo(c *fiber.Ctx) error {
	id, errReadParam := userIdFromParam(c)
	if errReadParam != nil {
		errorskit.LogWrap(errReadParam, "deletetodo")
		return FailResponse(errReadParam.Error(), c, fiber.StatusBadRequest)
	}

	err := h.DB.DeleteTodo(id)
	if err != nil {
		errorskit.LogWrap(err, "deletetodo id not found")
		return FailResponse(err.Error(), c, fiber.StatusNotFound)
	}

	// returns a map of id:"" to follow how other methods are used
	return SuccessResponse(CreateResponseMap(id, ""), c)
}
