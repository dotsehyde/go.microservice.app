package router

import (
	"sample/app/framework/http/routes/api"
	"sample/app/framework/http/routes/web"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Router(fiber *fiber.App)
}

func NewRouter(app *fiber.App, params ...any) {
	setup(app, api.NewAPIRouter(params), web.NewWebRouter(params))
}

func setup(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.Router(app)
	}
}
