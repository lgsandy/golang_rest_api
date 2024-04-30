package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santosh/crud/routes"
)

func main() {
	fmt.Println("Hello")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": 200, "msz": "Hwllo from gofiber"})
	})

	//setup routes
	routes.SetupUserRoute(app)
	routes.SetupAuthRouts(app)

	app.Listen(":3000")
}
