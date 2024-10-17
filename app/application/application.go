package application

import (
	"sample/utils/bodyparser"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func BodyParser(c *fiber.Ctx, v any) error {
	httpReq, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}
	return bodyparser.Parse(httpReq, v)
}
