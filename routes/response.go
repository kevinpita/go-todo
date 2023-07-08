package routes

import "github.com/gofiber/fiber/v2"

func CreateResponseMap(id int, todoText string) map[int]string {
	return map[int]string{id: todoText}
}

func SuccessResponse(data map[int]string, c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    data,
	})
}

func FailResponse(errMsg string, c *fiber.Ctx, status int) error {
	return c.Status(status).JSON(&fiber.Map{
		"success": false,
		"error":   errMsg,
	})
}
