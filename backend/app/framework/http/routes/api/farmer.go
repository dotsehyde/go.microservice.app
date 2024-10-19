package api

import (
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/handlers/api"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/requests"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func FarmerRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewFarmerHandler(router.Adapter, router.Producer, router.StorageSrv, router.CacheSrv)
	farmer := r.Group("/farmer")
	farmer.Get("/", handler.GetProfile())
	farmer.Put("/documents", handler.UpdateIDInfo(), adaptor.HTTPMiddleware(requests.ValidateIDUpload))
	farmer.Put("/profile-photo", handler.UpdateProfilePhoto())
	farmer.Put("/update", handler.UpdateUserInfo(), adaptor.HTTPMiddleware(requests.ValidateFarmerUpdateInfo))
}
