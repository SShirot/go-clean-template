package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS returns a CORS middleware with default configuration
func CORS() func(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins:     "*", // In production, specify exact origins
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
		MaxAge:           86400, // 24 hours
	})
}
