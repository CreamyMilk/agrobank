package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", homeHandler)
	app.Get("/imageupload", photoUploadHandler)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(418).JSON(&fiber.Map{
			"Message": "🍏 Route not found",
		}) // => 418 "I am a tepot"
	})
}

func homeHandler(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"Message": "Hello Handler",
	})
}

func photoUploadHandler(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"url": "https://localimagesstore/upload",
	})
}
