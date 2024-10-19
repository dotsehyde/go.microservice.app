package web

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
)

type webRouter struct {
	Adapter       *database.Adapter
	CacheSrv      gateways.CacheService
	EventProducer gateways.EventProducer
}

func NewWebRouter(params []any) *webRouter {
	instance := &webRouter{}
	return instance.instantiate(params)
}

func (r *webRouter) Router(app *fiber.App) {

	webGroup := app.Group("")
	r.index(webGroup)
	r.monitor(webGroup)
	// Custom config
	webGroup.Static(
		"/", "./storage/public", fiber.Static{
			Compress:      true,
			Browse:        false,
			CacheDuration: 10 * time.Second,
			MaxAge:        3600,
		},
	)
	// 404 Handler
	app.Use(
		func(c fiber.Ctx) error {
			return c.SendStatus(404) // => 404 "Not Found"
		},
	)
}

func (r *webRouter) index(router fiber.Router) {
	router.Get(
		"/", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		},
	)
}

func (r *webRouter) monitor(router fiber.Router) {
	return
	// router.Get("/monitor", monitor.New())
}

func (r *webRouter) instantiate(params []any) *webRouter {
	for _, param := range params {
		if adapter, ok := param.(*database.Adapter); ok {
			r.Adapter = adapter
			continue
		}
		if cacheService, ok := param.(gateways.CacheService); ok {
			r.CacheSrv = cacheService
			continue
		}
		if eventProducer, ok := param.(gateways.EventProducer); ok {
			r.EventProducer = eventProducer
			continue
		}
	}
	return r
}
