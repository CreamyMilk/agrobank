package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/registration"
	"github.com/gofiber/fiber/v2"
)

func TempRegistrationHandler(c *fiber.Ctx) error {
	r := new(registration.RegistrationLimbo)

	if err := c.BodyParser(r); err != nil {
		return c.JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	fmt.Printf("%+v", r)
	if err := r.TempCreate(); err != nil {
		return c.JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Registraion was successful",
	})
}
