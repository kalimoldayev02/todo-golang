package handler

import (
	"net/http"
	"todo-golang/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	data, err := h.Service.GetUsers()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(data)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {

}
