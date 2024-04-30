package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santosh/crud/controller"
)

func SetupUserRoute(app *fiber.App) {
	app.Get("/api/employee", controller.GetAllEmployees)
	app.Post("/api/employee", controller.SaveEmployee)

}
