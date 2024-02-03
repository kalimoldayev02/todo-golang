package routes

import (
	"todo-golang/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, controller *controllers.Controller) {
	// Создаю новую группу роутов
	route := app.Group("/api")

	// Роуты users
	route.Post("/sign-up", controller.SignUp) // TODO: auth_controller
}
