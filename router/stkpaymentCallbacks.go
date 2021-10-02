package router

import (
	"log"

	"github.com/CreamyMilk/agrobank/callback"
	"github.com/gofiber/fiber/v2"
)

func registrationStkCallHandler(c *fiber.Ctx) error {
	r := new(callback.StkpushCallbackResponse)

	if err := c.BodyParser(r); err != nil {
		log.Println("So they changed the structure of callbacks")
		return c.JSON(&fiber.Map{
			"ResponseCode": "00000000",
			"ResponseDesc": "success",
		})
	}

	go r.ParseRegistrations()

	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}

func depositStkCallHandler(c *fiber.Ctx) error {
	r := new(callback.StkpushCallbackResponse)

	if err := c.BodyParser(r); err != nil {
		log.Println("So they changed the structure of callbacks")
		return c.JSON(&fiber.Map{
			"ResponseCode": "00000000",
			"ResponseDesc": "success",
		})
	}

	go r.ParseDeposits()

	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}
