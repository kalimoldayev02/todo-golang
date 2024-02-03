package controllers

import (
	"net/http"
	"todo-golang/app/dto"
	"todo-golang/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) SignUp(ctx *fiber.Ctx) error {
	signUpDto := dto.SignUpDto{}

	if err := ctx.BodyParser(&signUpDto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(signUpDto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": utils.ValidatorErrors(err),
		})
	}

	userId, err := c.Service.SignUp(signUpDto)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"user_id": userId,
	})
}

func (c *Controller) SignIn(ctx *fiber.Ctx) error {
	signInDto := dto.SignInDto{}

	if err := ctx.BodyParser(&signInDto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(signInDto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": utils.ValidatorErrors(err),
		})
	}

	// TODO: create JWT token
	token, err := c.Service.SignIn(signInDto)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
