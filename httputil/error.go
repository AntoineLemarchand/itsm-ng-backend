package httputil

import "github.com/gofiber/fiber/v2"

func NewError(c *fiber.Ctx, code int, message string) error {
    return c.Status(code).JSON(fiber.Map{"error": message})
}

type HTTPError struct {
    Code    int
    Message string
}
