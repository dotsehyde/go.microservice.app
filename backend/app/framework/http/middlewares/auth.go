package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/BenMeredithConsult/locagri-apps/utils/jwt"
)

func ValidateRefreshToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		if c.Get("X-Refresh-Token") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "X-Refresh-Token header missing or empty value"})
		}
		return c.Next()
	}
}

// func ValidateOAuthToken() fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		if c.Get("X-OAuth-Token") == "" {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "X-OAuth-Token header missing or empty value"})
// 		}
// 		return c.Next()
// 	}
// }

func Authenticate(jwt *jwt.JWT) fiber.Handler {
	return func(c fiber.Ctx) error {
		token, err := bearerToken(c.Get("Authorization"))
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": false, "message": "Forbidden"})
		}
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
		}
		c.Locals("user", claims["session"])
		return c.Next()
	}
}

// bearerToken extracts the content from the header, striping the Bearer prefix
func bearerToken(rawToken string) (string, error) {
	pieces := strings.SplitN(rawToken, " ", 2)
	if len(pieces) < 2 {
		return "", fmt.Errorf("token with incorrect bearer format")
	}
	token := strings.TrimSpace(pieces[1])
	return token, nil
}
