package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santosh/crud/routes"
)

func InItServer() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
	//setup routes
	routes.SetupUserRoute(app)
	routes.SetupAuthRouts(app)
	app.Listen(":3000")
}
