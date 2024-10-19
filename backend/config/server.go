package config

import (
	"log"
	"os"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

type server struct {
	HTTP *fiber.App
	WG   *sync.WaitGroup
}

func NewServer() *server {
	var wg = sync.WaitGroup{}
	return &server{
		HTTP: fiber.New(
			fiber.Config{
				CaseSensitive:     Server().CaseSensitive,
				StrictRouting:     Server().StrictRouting,
				ServerHeader:      Server().ServerHeader,
				AppName:           App().Name,
				BodyLimit:         10485760,
				StreamRequestBody: true,
				JSONEncoder:       json.Marshal,
				JSONDecoder:       json.Unmarshal,
				Views:             html.New("./app/framework/http/templates", ".html"),
			},
		),
		WG: &wg,
	}

}

func (http *server) Run() {
	if os.Getenv("APP_ENV") == "production" {
		port := os.Getenv("PORT")
		if port == "" {
			port = App().PORT
		}
		log.Fatal(http.HTTP.Listen(":" + port))
	} else {
		log.Fatal(http.HTTP.Listen(":" + App().PORT))
	}
}
