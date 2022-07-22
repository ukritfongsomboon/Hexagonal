package middlewares

import (
	"hexagonal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func ValToken(c *fiber.Ctx) error {
	var access_token string
	cookie := c.Cookies("Accesstoken")

	authorizationHeader := c.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token = fields[1]
	} else {
		access_token = cookie
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"status":  false,
			"message": "unauthorized",
			"data":    "",
		})
	}

	sub, err := utils.ValidateToken(access_token, viper.GetString("app.access_token_public_key"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	c.Locals("user_id", sub.Id)
	c.Locals("role", sub.Role)

	return c.Next()
}
