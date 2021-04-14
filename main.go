package main

import (
	"database/sql"
	"fmt"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//DB holds global database object
var DB *sql.DB

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()
	app.Listen(":3000")
}
