package router

import (
	"github.com/CreamyMilk/agrobank/login"
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

type verifyTransactionRequest struct {
	Phone  string `json:"phonenumber"`
	Amount int64  `json:"amount"`
}

//number,amount} =>  get Rates, calculate NewBalance, USERNAME
type verificationResponse struct {
	Rates      int64  `json:"rates"`
	StatusCode int    `json:"status"`
	Username   string `json:"username"`
	Message    string `json:"message"`
}

type getBalanceRequest struct {
	WalletName string `json:"walletname"`
}

type getTransactionsRequest struct {
	WalletName string `json:"walletname"`
}

func depositCashHandler(c *fiber.Ctx) error {

	req := new(depositRequest)

	if err := c.BodyParser(req); err != nil {
		//fmt.Printf("%+v", err)
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

//Used to verify transactions halfway
func verifyTransactionHandler(c *fiber.Ctx) error {
	req := new(verifyTransactionRequest)
	res := new(verificationResponse)
	if err := c.BodyParser(req); err != nil {
		res.StatusCode = -1
		res.Message = "request is malformed"
		return c.JSON(res)
	}
	rates, err := wallet.GetTransactionPrice(req.Amount)
	res.Rates = rates
	if err != nil {
		res.StatusCode = -2
		res.Message = "Could not retrive the rates for the stated transaction"
		return c.JSON(res)
	}
	identity, err := login.GetPersonByWalletName(req.Phone)

	if err != nil {
		res.StatusCode = -19
		res.Message = "User is not registered yet"
		return c.JSON(res)
	}
	res.StatusCode = 0
	res.Username = identity
	res.Message = "Successful"
	return c.JSON(res)
}

func sendMoneyHandler(c *fiber.Ctx) error {
	req := new(sendMoneyrequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	//fmt.Printf("%+v", req)
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
		"status":     0,
		"newbalance": fromWallet.GetBalance(),
		"messge":     "Succesfully SentMoney",
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

func getTransactionsHandler(c *fiber.Ctx) error {
	req := new(getTransactionsRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Malformed request",
		})
	}
	w := wallet.GetWalletByName(req.WalletName)
	if w == nil {
		return c.JSON(&fiber.Map{
			"status":  -19,
			"message": "Sadly De wallet fake",
		})
	}
	transacations, err := w.GetTransactions()

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(transacations)
}
