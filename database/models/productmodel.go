package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID   int64  `json:"categoryID"`
	OwnerID      int64  `json:"ownerID"`
	ProductName  string `json:"productname"`
	ProductImage string `json:"image"`
	Description  string `json:"description"`
	PackingType  string `json:"packingtype"`
	Stock        int64  `json:"stock"`
	Price        int64  `json:"price"`
}
