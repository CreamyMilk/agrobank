package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/auth/login"
	"github.com/gofiber/fiber/v2"
)

func loginHandler(c *fiber.Ctx) error {

	req := new(login.LoginReq)

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	res, err := req.AttemptLogin()
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(res)
}
