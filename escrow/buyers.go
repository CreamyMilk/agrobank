package escrow

import (
	"errors"
	"log"

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
	res := database.DB.Order("created_at desc").Find(&list.Purchases, "buyer_addr=?", wallAddr)
	if res.Error != nil {
		return list, errCouldNotGetPurchases
	}
	return list, nil
}

func GetBuyerInvoicesSearch(wallAddr string, query string) (*PurchasesList, error) {
	list := new(PurchasesList)
	res := database.DB.Order("created_at desc").Find(&list.Purchases, "buyer_addr=? AND product_name LIKE ?", wallAddr, "%"+query+"%")
	if res.Error != nil {
		log.Println(res.Error)
		list.Status = -1
		return list, errCouldNotGetPurchases
	}
	return list, nil
}
