package api

import (
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/http/middlewares"
	"github.com/BenMeredithConsult/locagri-apps/utils/jwt"
	"github.com/gofiber/fiber/v3"
)

type apiRouter struct {
	mainRouter   fiber.Router
	StorageSrv   gateways.StorageService
	CacheSrv     gateways.CacheService
	Adapter      *database.Adapter
	RedisAdapter *database.RedisAdapter
	Adapters     map[string]*database.Adapter
	Tenant       string
	Producer     gateways.EventProducer
	jwt          *jwt.JWT
}

func NewAPIRouter(params []any) *apiRouter {
	instance := &apiRouter{
		jwt: jwt.NewJWT(),
	}
	instance.jwt.GenerateKey()
	return instance.instantiate(params)
}

func (r *apiRouter) Router(app *fiber.App) {
	api := app.Group("/api")
	r.mainRouter = api
	UnauthorizedRoutes(r.mainRouter, r)
	r.mainRouter.Use(middlewares.Authenticate(r.jwt))
	r.renderRoutes()
}

func (r *apiRouter) instantiate(params []any) *apiRouter {
	for _, param := range params {
		if adapter, ok := param.(*database.Adapter); ok {
			r.Adapter = adapter
			continue
		}
		if adapter, ok := param.(*database.RedisAdapter); ok {
			r.RedisAdapter = adapter
			continue
		}
		if eventProducer, ok := param.(gateways.EventProducer); ok {
			r.Producer = eventProducer
			continue
		}
		if storageService, ok := param.(gateways.StorageService); ok {
			r.StorageSrv = storageService
			continue
		}
		if cacheService, ok := param.(gateways.CacheService); ok {
			r.CacheSrv = cacheService
			continue
		}

	}
	return r
}

func (r *apiRouter) renderRoutes() {
	FarmerRoutes(r.mainRouter, r)
}
