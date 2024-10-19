package api

import (
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/application/appconst"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/gofiber/fiber/v3"
)

type appconstHandler struct {
	service gateways.AppConstService
}

func NewAppConstHandler(db *database.Adapter, cacheSrv gateways.CacheService) *appconstHandler {
	return &appconstHandler{
		service: appconst.NewAppConstService(appconst.NewAppConstRepo(db), cacheSrv),
	}
}

func (h *appconstHandler) GetLanguages() fiber.Handler {
	return func(c fiber.Ctx) error {
		data, err := h.service.GetLanguages()
		if err != nil {
			return c.Status(500).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}

func (h *appconstHandler) GetCountries() fiber.Handler {
	return func(c fiber.Ctx) error {
		data, err := h.service.GetCountries()
		if err != nil {
			return c.Status(500).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}

func (h *appconstHandler) GetNationalities() fiber.Handler {
	return func(c fiber.Ctx) error {
		data, err := h.service.GetNationalities()
		if err != nil {
			return c.Status(500).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(200).JSON(presenters.SuccessResponse(data))
	}
}
