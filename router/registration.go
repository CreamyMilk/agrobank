package router

import (
	"github.com/CreamyMilk/agrobank/auth/registration"
	"github.com/gofiber/fiber/v2"
)

func tempRegistrationHandler(c *fiber.Ctx) error {
	r := new(registration.TempRegistrationReq)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	topic, err := r.TempCreate()
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"topic":   topic,
		"message": "Registraion was successful",
	})
}
