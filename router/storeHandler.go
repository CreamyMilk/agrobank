package router

import (
	"github.com/CreamyMilk/agrobank/store"
	"github.com/gofiber/fiber/v2"
)

type getOwnersProductsRequest struct {
	OwnerID int64 `json:"ownerid"`
}

func addProductHandler(c *fiber.Ctx) error {
	tempProduct := new(store.Product)

	if err := c.BodyParser(tempProduct); err != nil {
		//fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	if err := tempProduct.AddProduct(); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "An error occured while inserting the product",
		})
	}
	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Added Product succesfully",
	})
}

func upadateProductHandler(c *fiber.Ctx) error {
	tempProduct := new(store.Product)

	if err := c.BodyParser(tempProduct); err != nil {
		//fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	if err := tempProduct.UpdateProduct(); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Updated Product succesfully",
	})
}

func getAllProductsHandler(c *fiber.Ctx) error {
	req := new(getOwnersProductsRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Malformed request",
		})
	}
	products, err := store.GetProductsByOwnerID(req.OwnerID)

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(products)
}
