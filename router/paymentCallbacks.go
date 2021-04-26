package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/registration"
	"github.com/gofiber/fiber/v2"
)

type StkPushCallBack struct {
	ID   string `json:"_id"`
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

func StkcallHandler(c *fiber.Ctx) error {
	r := new(StkPushCallBack)

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
			//Someone paid via Mpesa but did not fill registration forms
			fmt.Print(err)
		}
	}

	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}
