package models

import "gorm.io/gorm"

type Verifieduser struct {
	gorm.Model
	Fullname     string
	StoreName    string
	IDNumber     string
	DOB          string
	PhoneNumber  string `gorm:"unique"`
	Email        string
	IDPhoto      string
	ProfilePhoto string
	PasswordHash string
	Role         string
}
