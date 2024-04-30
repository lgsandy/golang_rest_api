package controller

import "github.com/gofiber/fiber/v2"

func GetAllEmployees(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"name": "Sandy", "age": 26, "branch": "hyderabad"})
}

func SaveEmployee(c *fiber.Ctx) error {
	return c.Send(c.Body())
}
