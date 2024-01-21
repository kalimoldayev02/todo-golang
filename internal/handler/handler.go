package handler

import (
	"todo-golang/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.Service
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes(api *fiber.App) {
	api.Group("/api")
	api.Get("/users")
}
