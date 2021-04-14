package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/callbacks"
	"github.com/CreamyMilk/agrobank/registration"
	"github.com/gofiber/fiber/v2"
)

func StkcallHandler(c *fiber.Ctx) error {
	r := new(callbacks.StkPushCallBack)

	if err := c.BodyParser(r); err != nil {
		return c.JSON(&fiber.Map{
			"ResponseCode": "00000000",
			"ResponseDesc": "success",
		})
	}

	if r.Body.StkCallback.ResultCode == 0 {
		p := registration.GetTempByID(r.Body.StkCallback.CheckoutRequestID)
		err := p.InsertPermanent()
		if err != nil {
			//Someone paid via Mpesa but didnot fill registration forms
			fmt.Print(err)
		}
	}

	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}
