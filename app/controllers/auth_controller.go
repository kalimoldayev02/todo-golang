package controllers

import (
	"net/http"
	"todo-golang/app/models"
	"todo-golang/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := c.BodyParser(signUp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(signUp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": utils.ValidatorErrors(err),
		})
	}

	return nil
}
