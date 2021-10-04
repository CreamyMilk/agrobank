package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/escrow"
	"github.com/gofiber/fiber/v2"
)

type sellerOrdersRequest struct {
	WalletName string `json:"walletname"`
}

type buyerOrdersRequest struct {
	WalletName string `json:"walletname"`
}

func SellersOrdersHandler(c *fiber.Ctx) error {

	req := new(sellerOrdersRequest)

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	response, err := escrow.GetOrdersTowardsWalletAddr(req.WalletName)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(response)
}

func buyerPlacedOrdersHandler(c *fiber.Ctx) error {
	req := new(buyerOrdersRequest)

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	response, err := escrow.GetBuyerInvoices(req.WalletName)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(response)
}
