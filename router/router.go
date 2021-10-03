package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.Post("/register", tempRegistrationHandler)
	auth.Post("/login", loginHandler)

	callbacks := v1.Group("/callbacks")
	callbacks.Post("/registrationstkpush", registrationStkCallHandler)
	callbacks.Post("/depositstkpushendpoint", depositStkCallHandler)

	wallet := v1.Group("/wallet")
	wallet.Post("/deposit", depositCashHandler)
	wallet.Post("/sendmoney", sendMoneyHandler)
	wallet.Post("/balance", getBalanceHandler)
	wallet.Post("/verify", verifyTransactionHandler)
	wallet.Post("/transactions", getTransactionsHandler)

	store := v1.Group("/store")
	store.Post("/add", addProductHandler)
	store.Put("/update", upadateProductHandler)
	store.Post("/stock", getUserStockhandler)
	store.Put("/categories", updateCategoryHandler)
	store.Get("/categories", getAllCategoriesHandler)
	store.Post("/categories", addCategoriesHandler)
	store.Delete("/categories", deleteCategoryHandler)
	store.Post("/products", getAllProductsByCategoryHandler)

	invoice := v1.Group("/invoice")
	invoice.Post("/create", createPurchaseInvoiceHandler)
	invoice.Post("/due", SellersOrdersHandler)
	// invoice.Post("/all", SellersOrdersHandler)
	// invoice.Post("/settle", SellersOrdersHandler)
	// invoice.Post("/cancel", SellersOrdersHandler)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(418).JSON(&fiber.Map{
			"Message": "🍏 Route not found",
		})
	})
}
