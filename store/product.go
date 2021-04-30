package store

import (
	"errors"

	"github.com/CreamyMilk/agrobank/database"
)

type Product struct {
	ProductID         int64   `json:"productID"`
	CategoryID        int64   `json:"categoryID"`
	OwnerID           int64   `json:"ownerID"`
	ProductName       string  `json:"productname"`
	ProductImage      string  `json:"image"`
	ProductImageLarge string  `json:"imagelarge"`
	Description       string  `json:"description"`
	PackingType       string  `json:"packingtype"`
	Stock             int     `json:"stock"`
	Price             float64 `json:"price"`
}

func (p *Product) AddProduct() error {
	res, err := database.DB.Exec(`
	INSERT INTO products 
	(category_id,owner_id,product_name,product_image,product_image_large,descriptions,price,stock,product_packtype)
	VALUES (?,?,?,?,?,?,?,?,?)`, p.CategoryID, p.OwnerID, p.ProductName, p.ProductImage,
		p.ProductImageLarge, p.Description, p.Price, p.Stock, p.PackingType)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return errors.New("could not get the latest id")
	}
	p.ProductID = id
	return nil
}
func (p *Product) UpdateProduct() error {
	res, err := database.DB.Exec(`UPDATE products SET 
	category_id=?,
	owner_id=?,
    product_name=?,
    product_image=?,
    product_image_large=?,
    descriptions=?,
    price=?,
    stock=?,
    product_packtype=?
	WHERE product_id=?;`, p.CategoryID, p.OwnerID, p.ProductName, p.ProductImage,
		p.ProductImageLarge, p.Description, p.Price, p.Stock, p.PackingType, p.ProductID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("the product you wish to update does not exist")
	}
	return nil
}
func (p *Product) DeleteProduct() error {
	_, err := database.DB.Exec("DELETE FROM products WHERE product_id=?", p.ProductID)
	if err != nil {
		return err
	}
	return nil
}
