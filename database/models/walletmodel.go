package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	WalletAddress string `json:"address"`
	UserID        uint   `json:"userid"`
	Balance       int64  `json:"balance"`
	WalletHash    string
	PhoneNumber   string
}

type DepositAttempt struct {
	gorm.Model
	CheckID       string
	Amount        string
	WalletAddress string
	Proccessed    bool `gorm:"default:0"`
}
