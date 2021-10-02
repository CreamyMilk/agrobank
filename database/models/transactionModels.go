package models

import (
	"gorm.io/gorm"
)

type TransactionType string

const (
	DepositType   TransactionType = "DEPOSIT"
	SendMoneyType TransactionType = "SENDMONEY"
	PurchaseType  TransactionType = "PURCHASE"
)

type Transaction struct {
	gorm.Model
	TrackID  string          `json:"trackid"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Amount   int64           `json:"amount"`
	Charge   int64           `json:"charge"`
	TypeName TransactionType `json:"typename"`
}

type TransactionCost struct {
	gorm.Model
	Upper_limit int64
	Charge      int64
}
