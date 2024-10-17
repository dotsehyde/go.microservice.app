package api

import (
	"sample/app/adapters/gateways"
	"sample/app/adapters/presenters"
	"sample/app/application"
	"sample/app/application/auth"
	"sample/app/domain/requestdto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service gateways.AuthService
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		service: auth.NewAuthService(auth.NewAuthRepo()),
	}
}

func (h *authHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {

		request := new(requestdto.RegisterRequest)
		if err := application.BodyParser(c, request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		data, err := h.service.Register(*request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(data)
	}
}

func (h *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requestdto.LoginRequest)
		if err := application.BodyParser(c, request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		data, err := h.service.Login(*request)
		if err != nil {
			if strings.Contains(err.Error(), "user not found") {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(data)
	}
}

func (h *authHandler) GetAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := h.service.AllUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse{
				Message: err.Error(),
				Status:  500,
			})
		}
		return c.JSON(data)

	}
}
