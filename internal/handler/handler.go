package handler

import (
	"log"
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
		log.Fatalf(err.Error())
	}

	return c.Status(http.StatusOK).JSON(data)
}
