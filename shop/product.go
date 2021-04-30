package shop

import (
	"errors"

	"github.com/CreamyMilk/agrobank/database"
)

type Product struct {
	productID   int64
	categoryID  int64
	productName string
	stock       int
	price       float64
}

func (p *Product) AddProduct() error {
	res, err := database.DB.Exec("INSERT INTO INTO Prdoucts_table (name,stock,price,category) VALUES (?,?,?)", p.productName, p.stock, p.price, p.categoryID)
	if err != nil {
		return errors.New("400")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return errors.New("402")
	}
	p.productID = id
	return nil
}

func (p *Product) DeleteProduct() error {
	return nil
}
