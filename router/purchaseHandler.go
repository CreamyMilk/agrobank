package router

import (
	"github.com/CreamyMilk/agrobank/escrow"
	"github.com/gofiber/fiber/v2"
)

type purchaseRequest struct {
	Walletname   string `json:"walletname"`
	ProductID    int64  `json:"productid"`
	Quantity     int64  `json:"quantity"`
	PasswordHash string `json:"passwordHash"`
	Delivery     int    `json:"acceptDelivery"`
}

func createPurchaseInvoiceHandler(c *fiber.Ctx) error {
	req := new(purchaseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})

	}
	_, err := escrow.CreateEscrowTransaction(req.Walletname, req.ProductID, req.Quantity)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  1,
			"message": err.Error(),
		})
	}
	//invoice.InformSeller()
	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Order Placed Sucesfully",
	})
}
