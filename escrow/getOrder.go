package escrow

import (
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/wallet"
)

type DueOrdersResponse struct {
	Status     int               `json:"status"`
	Orders     []DueOrder        `json:"orders"`
	TotalValue int64             `json:"total"`
	Mappers    map[string]string `json:"mappings"`
}
type DueOrder struct {
	BuyerAddr   string `json:"buyer"`
	SellerID    int64  `json:"seller"`
	ProductID   uint   `json:"productID,omitempty"`
	ProductName string `json:"pname,omitempty"`
	Quantity    int64  `json:"quantity"`
	TotalPrice  int64  `json:"total"`
	//CompletionCode string `json:"code,omitempty"`
	Deadline    int64
	CompletedAt int64
	Status      string
}

func GetOrdersTowardsWalletAddr(wallAddr string) (*DueOrdersResponse, error) {
	resp := new(DueOrdersResponse)

	sellwall := wallet.GetWalletByAddress(wallAddr)
	if sellwall == nil {
		resp.Status = -1
		resp.Orders = []DueOrder{}
		return resp, nil
	}

	//&models.EscrowInvoice{}.Limit(10).Find(&DueOrder{})
	err := database.DB.Model(&models.EscrowInvoice{}).Where("seller_id=? AND status=?", sellwall.UserID, "Placed").Order("created_at desc").Find(&resp.Orders).Error
	if err != nil {
		log.Println(err)
		resp.Status = -2
		return resp, nil
	}

	//fakeList := make([]DueOrder, len(resp.Orders), len(resp.Orders))
	walls := make(map[string]string)
	for _, order := range resp.Orders {
		resp.TotalValue += order.TotalPrice
		_, found := walls[order.BuyerAddr]
		if !found {
			fetchedWall := wallet.GetWalletByAddress(order.BuyerAddr)
			walls[order.BuyerAddr] = fetchedWall.PhoneNumber
		}
	}
	resp.Mappers = walls

	return resp, nil
}
