package escrow

import (
	"errors"
	"fmt"
	"time"

	"github.com/CreamyMilk/agrobank/database"
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

func CreateEscrowTransaction(buyerWalletName string, productID int64, quantity int64) (*EscrowInvoice, error) {
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

	if buyer.WalletName() == seller.WalletName() {
		return nil, errors.New("as a seller you cannot purchase your own goods")
	}
	//Checks that the user has the purchasing power to buy the requested product
	totalPrice := product.Price * float64(quantity)
	currentBuyerBalance := buyer.GetBalance()
	if currentBuyerBalance <= int64(totalPrice) {
		return nil, errors.New("the buyer has insuffiecnt funds to complete trasaction")
	}
	//this ensures that an order has benn placed an the new stock has been updated
	tx, err := database.DB.Begin()
	if err != nil {
		fmt.Printf("Was unabke ti naje database trabsation %v", err)
	}

	//Create invoice
	//This number is required for merchent reconciliation and payment money to be formalized
	reconciliationCode := ReconciliationCodeGen()
	_, err = tx.Exec(`INSERT INTO escrowInvoices(
		reconciliationcode,
		senderWalletName,
		receiverWalletName,
		prodcutID,
		amount,
		CreatedAt,
		Deadline) 
		VALUES (?,?,?,?,?,?,?)`,
		reconciliationCode,
		buyer.WalletName(),
		seller.WalletName(),
		product.ProductID,
		totalPrice,
		time.Now(),
		time.Now().Add(time.Hour*STANDARD_DELIVERY_DURATION),
	)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return nil, errors.New("could not create escrow invoice between you and the seller kindly try again later")
	}

	err = product.DeceremtStockBy(tx, quantity)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = buyer.PayEscrow(tx, product.GetProductShortName(), reconciliationCode, int64(totalPrice))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tempInvoice := new(EscrowInvoice)
	tempInvoice.Buyer = buyer
	tempInvoice.Seller = seller
	tempInvoice.Product = product
	tempInvoice.TotalPrice = totalPrice

	tx.Commit()
	return tempInvoice, nil
}
