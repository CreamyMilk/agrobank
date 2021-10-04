package escrow

import (
	"errors"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
)

type PurchasesList struct {
	Status    int                    `json:"status"`
	Purchases []models.EscrowInvoice `json:"purchases"`
}

var (
	errCouldNotGetPurchases = errors.New("could not find any purchases")
)

func GetBuyerInvoices(wallAddr string) (*PurchasesList, error) {
	list := new(PurchasesList)
	res := database.DB.Find(&list.Purchases, "buyer_addr=?", wallAddr)
	if res.Error != nil {
		return list, errCouldNotGetPurchases
	}
	return list, nil
}
