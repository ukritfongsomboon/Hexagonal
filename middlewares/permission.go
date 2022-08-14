package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValPer(listPer []int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO อยากรับ Argument เพิ่มอ่ะครับ
		for _, role := range listPer {
			if role == c.Locals("role") {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    fiber.StatusForbidden,
			"status":  false,
			"message": "permission denied",
			"data":    "",
		})
	}
}

// AuthReq middleware
func AuthReq(data string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(data)
		return nil
	}
}
