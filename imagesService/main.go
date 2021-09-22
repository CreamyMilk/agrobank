package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Fiber instance
	app := fiber.New()
  app.Use(logger.New())
	// Routes
	app.Post("/formstuff", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
    fmt.Println("Hello new file")
		// Save file to root directory:
		c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
    return c.JSON(&fiber.Map{
        "url":file.Filename,
    })
	})

	// Start server
	log.Fatal(app.Listen(":4000"))
}
