package api

import (
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/handlers/api"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/middlewares"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/requests"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func UnauthorizedRoutes(r fiber.Router, router *apiRouter) {

	//APP CONST
	appHandler := api.NewAppConstHandler(router.Adapter, router.CacheSrv)
	app := r.Group("/config")
	app.Get("/languages", appHandler.GetLanguages())
	app.Get("/countries", appHandler.GetCountries())
	app.Get("/nationalities", appHandler.GetNationalities())

	//AUTH
	authHandler := api.NewAuthHandler(router.Adapter, router.Producer, router.jwt, router.CacheSrv)
	auth := r.Group("/auth")
	auth.Post("/login", authHandler.Login(), adaptor.HTTPMiddleware(requests.ValidateLogin))
	auth.Post("/register", authHandler.Register(), adaptor.HTTPMiddleware(requests.ValidateCreateUser))
	auth.Post("/resend-otp", authHandler.SendOTP(), adaptor.HTTPMiddleware(requests.ValidateLogin))
	auth.Post("/verify-otp", authHandler.VerifyOTP(), adaptor.HTTPMiddleware(requests.ValidateVerifyOTP))
	auth.Post("/refresh-token", authHandler.RefreshToken(), middlewares.ValidateRefreshToken())
	auth.Post("/reset-account", authHandler.ResetAccount(), adaptor.HTTPMiddleware(requests.ValidateLogin))

}
