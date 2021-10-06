package escrow

import (
	"errors"
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/firenotifier"
	"gorm.io/gorm"
)

type SettlmentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	errCodeIsInvalid         = errors.New("the issued code has expired or is invalid")
	errWrongSeller           = errors.New("the order is towards a diffrent seller")
	errCouldNotSettleInvoice = errors.New("could not settle the invoice")
	errSellerWalletNotFound  = errors.New("the sellers wallet could not be found")
	errDepositFailed         = errors.New("could not deposit the funds to your wallet")
)

func SettleOrderUsingQR(settlmentCode string, sellerID int64) (*SettlmentResponse, error) {
	resp := new(SettlmentResponse)
	var invoice models.EscrowInvoice
	var sellerWall models.Wallet

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&models.EscrowInvoice{}).Where("completion_code=? AND status=?", settlmentCode, "Placed").Find(&invoice).Error
		if err != nil {
			log.Println(err)
			return errCodeIsInvalid
		}
		if invoice.SellerID != sellerID {
			return errWrongSeller
		}
		invoice.Status = "Completed"
		invres := tx.Save(&invoice)
		if invres.Error != nil {
			log.Println(invres.Error)
			return errCouldNotSettleInvoice
		}

		res := tx.First(&sellerWall, "user_id=?", sellerID)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return errSellerWalletNotFound
			}
			return res.Error
		}

		if invoice.TotalPrice < 0 {
			return errDepositFailed
		}

		//Deposit the funds to the sellers wallet
		sellerWall.Balance = sellerWall.Balance + invoice.TotalPrice
		wallErr := tx.Save(sellerWall)
		if wallErr.Error != nil {
			log.Println(wallErr.Error)
			return errDepositFailed
		}

		//Notify the buyer the order has been executed succesfully
		return nil
	})

	if err != nil {
		resp.Status = -1
		resp.Message = err.Error()
		return resp, err
	}

	go firenotifier.BuyerPurchaseComplete(invoice.BuyerAddr)
	go firenotifier.SellerInvoiceComplete(sellerWall.WalletAddress)

	resp.Status = 0
	resp.Message = "Order has been settled successfully. You will receive your fund in a few minutes."
	return resp, err
}
