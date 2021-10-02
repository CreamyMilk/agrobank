package models

import (
	"gorm.io/gorm"
)

type EscrowInvoice struct {
	gorm.Model
	BuyerAddr      string `json:"buyer"`
	SellerID       int64  `json:"seller"`
	ProductID      uint   `json:"productID,omitempty"`
	ProductName    string `json:"pname,omitempty"`
	Quantity       int64
	TotalPrice     int64  `json:"total"`
	CompletionCode string `json:"code,omitempty"`
	Deadline       int64
	CompletedAt    int64
	Status         string
}
