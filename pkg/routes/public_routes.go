package routes

import "github.com/gofiber/fiber/v2"

func PublicRoutes(app *fiber.App) {
	// Создаю новую группу роутов
	route := app.Group("/api")

	// Роуты users
	route.Get("/users") // TODO: add controller
	route.Post("/users")
}
