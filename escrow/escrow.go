package escrow

import (
	"errors"
	"log"
	"time"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/firenotifier"
	"github.com/CreamyMilk/agrobank/store"
	"github.com/CreamyMilk/agrobank/wallet"
	"gorm.io/gorm"
)

type EscrowOrdersReponse struct {
	Orders     []models.EscrowInvoice `json:"orders"`
	StatusCode int                    `json:"status"`
	Total      int64                  `json:"total"`
}

var (
	errBuyerNotFound             = errors.New("could not get buyers wallet")
	errProductNotFound           = errors.New("could not get product")
	errSellerNotFound            = errors.New("the seller can't take this order currently")
	errSelfPurchase              = errors.New("as a seller you cant purchase your own goods")
	errLessStock                 = errors.New("their is no enough stock")
	errNotEnoughMoney            = errors.New("you dont have enough money")
	errCouldNotChargeForProduct  = errors.New("could not charge for the product")
	errCouldNotUpdateStock       = errors.New("could not update stock")
	errCouldNotCreateTransaction = errors.New("could not update stock")
	errInvoiceCreationFailed     = errors.New("could not create invoice")
)

func CreateEscrowTransaction(buyerWalletAddr string, productID int64, quantity int64) (*models.EscrowInvoice, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		//get buyer wallet
		buyer := wallet.GetWalletByAddress(buyerWalletAddr)
		if buyer == nil {
			return errBuyerNotFound
		}
		//get product
		product := store.GetProductByProductID(productID)
		if product == nil {
			return errProductNotFound
		}
		//get sellers Address
		seller := wallet.GetWalletByUserID(product.OwnerID)
		if seller == nil {
			return errSellerNotFound
		}
		//Own Purchase
		if seller.WalletAddress == buyerWalletAddr {
			return errSelfPurchase
		}
		//ensure stock Checkout
		if product.Stock < quantity {
			return errLessStock
		}
		//calculate total and check purchasing power
		totalPrice := product.Price * quantity
		if buyer.Balance < int64(totalPrice) {
			return errNotEnoughMoney
		}
		//charge users wallet
		buyer.Balance = buyer.Balance - int64(totalPrice)
		chargeErr := tx.Save(buyer).Error
		if chargeErr != nil {
			return errCouldNotChargeForProduct
		}
		//Update Stock
		product.Stock = product.Stock - quantity
		stockErr := tx.Save(product).Error
		if stockErr != nil {
			return errCouldNotUpdateStock
		}
		//Create reconciliation
		theFinalID := ReconciliationCodeGen()
		//create transaction
		var transs models.Transaction
		transs.Amount = int64(totalPrice)
		transs.Charge = 0
		transs.From = buyerWalletAddr
		transs.To = product.ProductName
		transs.TypeName = models.PurchaseType
		transs.TrackID = theFinalID
		persistErr := tx.Create(&transs).Error
		if persistErr != nil {
			log.Println("Transation persist failed")
			log.Println(persistErr)
			return errCouldNotCreateTransaction
		}
		//create escrow
		tempInvoice := &models.EscrowInvoice{
			BuyerAddr:      buyerWalletAddr,
			SellerID:       product.OwnerID,
			ProductID:      product.ID,
			ProductName:    product.ProductName,
			Quantity:       quantity,
			TotalPrice:     product.Price * quantity,
			CompletionCode: theFinalID,
			Deadline:       time.Now().Add(time.Hour * 96).Unix(),
			Status:         "Placed",
		}
		invErr := tx.Create(tempInvoice).Error

		if invErr != nil {
			log.Println("Could not create invoice")
			log.Println(invErr)
			return errInvoiceCreationFailed
		}

		//That id can be a zero i presume
		go firenotifier.SuccesfulPurchaseNotif(*product, seller.WalletAddress, "ORDER ID HERE")
		return nil
	})

	if err != nil {
		return nil, err
	}
	return nil, nil
}
