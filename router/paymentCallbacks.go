package router

import (
	"errors"
	"fmt"

	"github.com/CreamyMilk/agrobank/deposit"
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
		inv := deposit.GetInvoiceByID(r.Body.StkCallback.CheckoutRequestID)
		if inv == nil {
			fmt.Print(errors.New("invoice not found"))
			return c.JSON(&fiber.Map{
				"ResponseCode": "00000000",
				"ResponseDesc": "success",
			})
		}
		mpesaReceiptNumber := r.Body.StkCallback.CallbackMetadata.Item[1].Value.(string)
		err := inv.PayOut(mpesaReceiptNumber)
		if err != nil {
			p := registration.GetTempByID(r.Body.StkCallback.CheckoutRequestID)
			err := p.InsertPermanent()
			fmt.Print(err)
		}
	}

	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}
