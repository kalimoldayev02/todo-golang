package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Роут для 404
func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "endpoint is not found",
			})
		},
	)
}
