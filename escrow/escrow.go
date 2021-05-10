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

type EscrowTransactions struct {
	EscrowID       int64   `json:"eid"`
	BuyerName      string  `json:"buyer"`
	SellerName     string  `json:"seller"`
	ProductID      int64   `json:"productID,omitempty"`
	ProductName    string  `json:"pname,omitempty"`
	TotalPrice     float64 `json:"total"`
	CompletionCode string  `json:"code,omitempty"`
	CreatedAt      int64
	Deadline       int64
	CompletedAt    int64
}
type EscrowOrdersReponse struct {
	Orders     []EscrowTransactions `json:"orders"`
	StatusCode int                  `json:"status"`
	Total      int64                `json:"total"`
}

func GetInvoicesTowardsSeller(sellerWalletName string) (*EscrowOrdersReponse, error) {
	seller := wallet.GetWalletByName(sellerWalletName)
	if seller == nil {
		return nil, errors.New("sadly you seem to not have a active mobile wallet")
	}
	result := new(EscrowOrdersReponse)
	rows, err := database.DB.Query(`SELECT 
		eid ,
		senderWalletName       ,		
		receiverWalletName     ,
		prodcutID,  amount   
		FROM escrowInvoices
		WHERE 
		receiverWalletName = ?;
	`, seller.WalletName())
	if err != nil {
		result.StatusCode = -500
		return result, err
	}

	for rows.Next() {
		singleOrder := EscrowTransactions{}
		if err := rows.Scan(
			&singleOrder.EscrowID,
			&singleOrder.BuyerName,
			&singleOrder.SellerName,
			&singleOrder.ProductID,
			&singleOrder.TotalPrice,
		); err != nil {
			result.StatusCode = -501
			return result, err
		}
		result.Total += int64(singleOrder.TotalPrice)
		result.Orders = append(result.Orders, singleOrder)
	}
	if err != nil {
		result.StatusCode = -502
		return result, err
	}
	if result.Orders == nil {
		result.StatusCode = -503
	}
	defer rows.Close()
	return result, nil
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

//Settle Escrow
func (escrow *EscrowInvoice) Settle() error {
	//Send Money to seller
	//Inform the buyer
	return nil
}

//Reverse Escrow
func (escrow *EscrowInvoice) Reverse() error {
	//Send Money back to the owner
	//Inform the seller that the purchase did not go through due to time
	return nil
}

//Get orders / Escrows for a particular seller
