package api

import (
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/application"
	"github.com/BenMeredithConsult/locagri-apps/app/application/auth"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/BenMeredithConsult/locagri-apps/utils/jwt"
	"github.com/gofiber/fiber/v3"
)

type authHandler struct {
	service gateways.AuthService
}

func NewAuthHandler(db *database.Adapter, producer gateways.EventProducer, jwt *jwt.JWT, cacheSrv gateways.CacheService) *authHandler {
	return &authHandler{
		service: auth.NewAuthService(auth.NewAuthRepo(db), jwt, cacheSrv, producer),
	}
}

func (h *authHandler) Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.LoginAndRestRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.Login(request.Phone)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *authHandler) Register() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.CreateUserRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.Register(request)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *authHandler) RefreshToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		data, err := h.service.RefreshToken(c.Get("X-Refresh-Token"))
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}

func (h *authHandler) ResetAccount() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.LoginAndRestRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.ResetAccount(request.Phone)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *authHandler) SendOTP() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.LoginAndRestRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.SendOTP(request.Phone)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.MessageResponse(data))
	}
}

func (h *authHandler) VerifyOTP() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.VerifyOTPRequest)
		if err := application.BodyParser(c, &request); err != nil {
			return c.Status(400).JSON(presenters.ErrorResponse(err))
		}
		data, err := h.service.VerifyOTP(request.Phone, request.Code)
		if err != nil {
			status, err := application.HandleErrors(err)
			return c.Status(status).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}
