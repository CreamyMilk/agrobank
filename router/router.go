package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Post("/treg", TempRegistrationHandler)
	app.Post("/stkcall", StkcallHandler)
	app.Post("/login", LoginHandler)

	wallet := app.Group("/wallet")

	wallet.Post("/deposit", depositCashHandler)
	wallet.Post("/sendmoney", sendMoneyHandler)
	wallet.Post("/balance", getBalanceHandler)
	wallet.Post("/verify", verifyTransactionHandler)
	wallet.Post("/transactions", getTransactionsHandler)

	store := app.Group("/store")

	store.Post("/add", addProductHandler)
	store.Put("/update", upadateProductHandler)
	store.Post("/stock", getUserStockhandler)
	store.Get("/categories", getAllCategoriesHandler)
	store.Post("/products", getAllProductsByCategoryHandler)

	invoice := app.Group("/invoice")

	invoice.Post("/create", createPurchaseInvoiceHandler)
	invoice.Post("/orders", SellersOrdersHandler)

	api := app.Group("/api")
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
