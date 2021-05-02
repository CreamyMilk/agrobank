package escrow

import (
	"errors"

	"github.com/CreamyMilk/agrobank/store"
	"github.com/CreamyMilk/agrobank/wallet"
)

//This package is responisible for the creation of secure payment channles between buyers and sellers
type EscrowInvoice struct {
	EscrowID    int64
	Buyer       *wallet.Wallet
	Seller      *wallet.Wallet
	Product     *store.Product
	TotalPrice  float64
	CreatedAt   int64
	Deadline    int64
	CompletedAt int64
}

func CreateEscrowInvoice(buyerWalletName string, productID int64, quantity int64) (*EscrowInvoice, error) {
	buyer := wallet.GetWalletByName(buyerWalletName)
	if buyer == nil {
		return nil, errors.New("sadly you seem to not have a mobile wallet")
	}
	product := store.GetProductByProductID(productID)

	if product == nil {
		return nil, errors.New("the product you wish to purchase has been removed by the seller")
	}
	seller := product.GetWalletOfProductOwner()
	if seller == nil {
		return nil, errors.New("the seller is currently unavailable")
	}

	totalPrice := product.Price * float64(quantity)
	currentBuyerBalance := buyer.GetBalance()
	if currentBuyerBalance <= int64(totalPrice) {
		return nil, errors.New("the buyer has insuffiecnt funds to complete trasaction")
	}
	err := product.DeceremtStockBy(quantity)
	if err != nil {
		return nil, err
	}
	//buyer.Payescrow()

	tempInvoice := new(EscrowInvoice)
	tempInvoice.Buyer = buyer
	tempInvoice.Seller = seller
	tempInvoice.Product = product
	tempInvoice.TotalPrice = totalPrice

	return tempInvoice, nil
}
