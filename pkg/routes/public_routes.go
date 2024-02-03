package routes

import (
	"todo-golang/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, controller *controllers.Controller) {
	// Создаю новую группу роутов
	route := app.Group("/api")

	// Auth
	route.Post("/sign-up", controller.SignUp)
	route.Post("/sign-in", controller.SignIn)
}
