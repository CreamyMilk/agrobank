package main

import (
	"github.com/CreamyMilk/agrobank/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.Connect()
	defer database.DB.Close()
}
