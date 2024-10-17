package api

import (
	"github.com/gofiber/fiber/v2"
)

type apiRouter struct {
	mainRouter fiber.Router
	// jwt        *jwt.JWT
}

func NewAPIRouter(params []any) *apiRouter {
	apiRouter := &apiRouter{
		// jwt: jwt.NewJWT(),
	}
	// apiRouter.jwt.GenerateKey()
	return apiRouter.instantiate(params)
}

func (r *apiRouter) instantiate(params []any) *apiRouter {
	// for _, param := range params {
	// 	switch p := param.(type) {
	// 	case *database.RedisAdapter:
	// 		r.RedisAdapter = p
	// 		continue
	// 	}
	// }
	return r
}

func (r *apiRouter) Router(app *fiber.App) {
	api := app.Group("/api")
	r.mainRouter = api.Group("")

	r.renderRoutes()
}

func (r *apiRouter) renderRoutes() {
	AuthRoutes(r.mainRouter, r)
}
