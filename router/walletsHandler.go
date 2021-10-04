package router

import (
	"log"

	"github.com/CreamyMilk/agrobank/auth/registration"
	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/mpesa"
	"github.com/CreamyMilk/agrobank/wallet"
	"github.com/gofiber/fiber/v2"
)

type depositRequest struct {
	WalletAddress string `json:"walletname"`
	Phone         string `json:"phonenumber"`
	FmcToken      string `json:"fcmtoken"`
	Amount        string `json:"amount"`
}

type sendMoneyrequest struct {
	SenderWalletName     string `json:"from"`
	ReceipientWalletName string `json:"to"`
	Amount               int64  `json:"amount"`
	Key                  string `json:"key"`
}

type verifyTransactionRequest struct {
	Phone  string `json:"phonenumber"`
	Amount int64  `json:"amount"`
}

//number,amount} =>  get Rates, calculate NewBalance, USERNAME
type verificationResponse struct {
	Rates         int64  `json:"rates"`
	StatusCode    int    `json:"status"`
	Username      string `json:"username"`
	WalletAddress string `json:"wallet_address"`
	Message       string `json:"message"`
}

type getBalanceRequest struct {
	WalletAddress string `json:"walletname"`
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
	userwallet := wallet.GetWalletByAddress(req.WalletAddress)

	if userwallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "error retriving your wallet info",
		})
	}

	check, err := mpesa.SendSTK(req.Phone, req.Amount, "Deposit", req.FmcToken, mpesa.DepositTypeSTK)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": err.Error(),
		})
	}

	atmpr := &models.DepositAttempt{
		CheckID:       check,
		Amount:        req.Amount,
		WalletAddress: userwallet.WalletAddress,
	}
	res := database.DB.Save(atmpr)

	if res.Error != nil {
		log.Println(res.Error)
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

	identity := registration.GetUserDetailsByMobile(req.Phone)
	if identity == nil {
		res.StatusCode = -19
		res.Message = "Seems the User is not registerd to AgroCRM"
		return c.JSON(res)
	}
	wall := wallet.GetWalletByPhone(req.Phone)
	if wall == nil {
		res.StatusCode = -19
		res.Message = "Seems the User has no wallet to AgroCRM"
		return c.JSON(res)
	}
	res.WalletAddress = wall.WalletAddress
	res.StatusCode = 0
	res.Username = identity.Fullname
	res.Message = "Successful Registration"
	return c.JSON(res)
}

//Todo Add Password Auth
func sendMoneyHandler(c *fiber.Ctx) error {
	req := new(sendMoneyrequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	//fmt.Printf("%+v", req)
	fromWallet := wallet.GetWalletByAddress(req.SenderWalletName)
	if fromWallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "invalid sender wallet address",
		})
	}
	toWallet := wallet.GetWalletByAddress(req.ReceipientWalletName)
	if toWallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "invalid sender wallet address",
		})
	}

	sucessful, err := wallet.SendMoney(req.Amount, fromWallet.WalletAddress, toWallet.WalletAddress)
	if !sucessful {
		return c.JSON(&fiber.Map{
			"status":  -10,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":     0,
		"newbalance": wallet.GetWalletBalance(fromWallet.WalletAddress),
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

	w := wallet.GetWalletByAddress(req.WalletAddress)
	if w == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Error retriving account balance.Try again later.",
		})
	}

	return c.JSON(&fiber.Map{
		"status":  0,
		"balance": w.Balance,
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
	w := wallet.GetWalletByAddress(req.WalletName)
	if w == nil {
		return c.JSON(&fiber.Map{
			"status":  -19,
			"message": "Sadly De wallet fake",
		})
	}
	transacations, err := wallet.GetWalletTransactions(req.WalletName)

	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.JSON(transacations)
}
