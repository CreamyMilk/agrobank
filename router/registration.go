package router

import (
	"github.com/CreamyMilk/agrobank/registration"
	"github.com/gofiber/fiber/v2"
)

func TempRegistrationHandler(c *fiber.Ctx) error {
	r := new(registration.RegistrationLimbo)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	if err := r.TempCreate(); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Registraion was successful",
	})
}
