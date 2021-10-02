package models

import "gorm.io/gorm"

type Registation struct {
	gorm.Model
	Fullname     string
	StoreName    string
	IDNumber     string
	DOB          string
	PhoneNumber  string
	Email        string
	IDPhoto      string
	ProfilePhoto string
	PasswordHash string
	Role         string
	CheckoutID   string
}
