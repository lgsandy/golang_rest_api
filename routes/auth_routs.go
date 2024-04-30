package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santosh/crud/controller"
)

func SetupAuthRouts(app *fiber.App) {
	app.Post("/api/login", controller.Login)
}
