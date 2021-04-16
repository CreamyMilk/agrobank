package router

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/mpesa"
	"github.com/CreamyMilk/agrobank/wallet"
	"github.com/gofiber/fiber/v2"
)

type depositRequest struct {
	WalletName string `json:"walletname"`
	Phone      string `json:"phonenumber"`
	FmcToken   string `json:"fcmtoken"`
	Amount     string `json:"amount"`
}

type sendMoneyrequest struct {
	SenderWalletName     string `json:"from"`
	ReceipientWalletName string `json:"to"`
	Amount               int64  `json:"amount"`
}

type getBalanceRequest struct {
	WalletName string `json:"walletname"`
}

func depositCashHandler(c *fiber.Ctx) error {

	req := new(depositRequest)

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("%+v", err)
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	userwallet := wallet.GetWalletByName(req.WalletName)

	if userwallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "error retriving your wallet info",
		})
	}

	why, err := mpesa.SendSTK(req.Phone, req.Amount, req.WalletName, req.FmcToken)

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": why,
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Deposit request was successful please wait",
	})
}

func sendMoneyHandler(c *fiber.Ctx) error {
	req := new(sendMoneyrequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	fromWallet := wallet.GetWalletByName(req.SenderWalletName)
	if fromWallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "invalid sender wallet address",
		})
	}
	toWallet := wallet.GetWalletByName(req.ReceipientWalletName)
	if toWallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "invalid sender wallet address",
		})
	}

	errorMessage, sucessful := fromWallet.SendMoney(req.Amount, *toWallet)
	if !sucessful {
		return c.JSON(&fiber.Map{
			"status":  -10,
			"message": errorMessage,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 0,
		"messge": "Succesfully SentMoney",
	})
}

func getBalanceHandler(c *fiber.Ctx) error {
	req := new(getBalanceRequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	walletAddress := wallet.GetWalletByName(req.WalletName)
	if walletAddress == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Error retriving account balance.Try again later.",
		})
	}
	currentBalance := walletAddress.GetBalance()

	return c.JSON(&fiber.Map{
		"status":  0,
		"balance": currentBalance,
	})
}
