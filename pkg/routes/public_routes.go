package routes

import (
	"todo-golang/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	// Создаю новую группу роутов
	route := app.Group("/api")

	// Роуты users
	route.Get("/sign-up", controllers.SignUp) // TODO: auth_controller
}
