package router

import (
	"github.com/CreamyMilk/agrobank/machinary"
	"github.com/gofiber/fiber/v2"
)

type getOwnersMachinarysRequest struct {
	OwnerID int64 `json:"ownerid"`
}

type GetMachinarysByCategoryIDRequest struct {
	CategoryID int64 `json:"categoryid"`
}

func addMachinaryHandler(c *fiber.Ctx) error {
	tempMachinary := new(machinary.Machinary)

	if err := c.BodyParser(tempMachinary); err != nil {
		//fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	if err := tempMachinary.AddMachinary(); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Added Machinary succesfully",
	})
}

func upadateMachinaryHandler(c *fiber.Ctx) error {
	tempMachinary := new(machinary.Machinary)

	if err := c.BodyParser(tempMachinary); err != nil {
		//fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	if err := tempMachinary.UpdateMachinary(); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Updated Machinary succesfully",
	})
}

func getUserMachineshandler(c *fiber.Ctx) error {
	req := new(getOwnersMachinarysRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Malformed request",
		})
	}
	Machinarys, err := machinary.GetMachinarysByOwnerID(req.OwnerID)

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(Machinarys)
}

func getAllMachineCategoriesHandler(c *fiber.Ctx) error {
	categories, err := machinary.GetCategories()

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(categories)
}

func getAllMachinarysByCategoryHandler(c *fiber.Ctx) error {
	req := new(GetMachinarysByCategoryIDRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Malformed request",
		})
	}
	Machinarys, err := machinary.GetMachinarysByCategoryID(req.CategoryID)

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(Machinarys)
}
