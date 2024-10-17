package api

import (
	"sample/app/framework/http/handlers/api"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewAuthHandler()
	r.Post("/register", handler.Register())
	r.Post("/login", handler.Login())
	r.Get("/users", handler.GetAllUsers())

}
