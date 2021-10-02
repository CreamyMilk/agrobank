package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/firenotifier"
	"github.com/CreamyMilk/agrobank/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	firenotifier.Init()
	router.SetupRoutes(app)
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	database.SetupModels(
		&models.Registation{},
		&models.Verifieduser{},
		&models.Category{},
		&models.Product{},
		&models.Wallet{},
		&models.Transaction{},
		&models.TransactionCost{},
		&models.DepositAttempt{},
		&models.EscrowInvoice{},
	)
	if os.Getenv("SEED_RATES") != "" {
		database.SeedTransactionCosts()
	}

	if os.Getenv("SEED_CATEGORIES") != "" {
		database.SeedCategories()
	}

	log.Fatal(app.Listen(":3000"))
}
