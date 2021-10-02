package store

import (
	"errors"
	"fmt"
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
)

var (
	errCouldNotCreate        = errors.New("could not create product")
	errCouldNotUpdateProduct = errors.New("could not update product")
	errNoProductWasUpdated   = errors.New("the product you wish to update was not affected")
	errCouldNotDeleteProduct = errors.New("could not delete the product")
	errCouldNotGetProducts   = errors.New("could not fetch products")
)

type Product struct {
	ProductID         int64  `json:"productID"`
	CategoryID        int64  `json:"categoryID"`
	OwnerID           int64  `json:"ownerID"`
	ProductName       string `json:"productname"`
	ProductImage      string `json:"image"`
	ProductImageLarge string `json:"imagelarge"`
	Description       string `json:"description"`
	PackingType       string `json:"packingtype"`
	Stock             int64  `json:"stock"`
	Price             int64  `json:"price"`
	PriceString       string `json:"priceString"`
}

type ProductsList struct {
	Products   []models.Product `json:"products"`
	StatusCode int              `json:"status"`
}

func (p *Product) AddProduct() error {
	newProduct := &models.Product{
		CategoryID:   p.CategoryID,
		OwnerID:      p.OwnerID,
		ProductName:  p.ProductName,
		ProductImage: p.ProductImage,
		Description:  p.Description,
		PackingType:  p.PackingType,
		Stock:        p.Stock,
		Price:        p.Price,
	}

	res := database.DB.Create(newProduct)
	if res.Error != nil {
		log.Println(res.Error)
		return errCouldNotCreate
	}
	return nil
}

func (p *Product) UpdateProduct() error {
	var singleProduct models.Product
	r := database.DB.First(&singleProduct, p.ProductID)
	if r.Error != nil {
		log.Println(r)
		return errCouldNotUpdateProduct
	}
	fmt.Printf("Req :%+v\n", p)
	fmt.Printf("DB :%+v\n", singleProduct)
	singleProduct.CategoryID = p.CategoryID
	singleProduct.ProductName = p.ProductName
	singleProduct.ProductImage = p.ProductImage
	singleProduct.Description = p.Description
	singleProduct.PackingType = p.PackingType
	singleProduct.Stock = p.Stock
	singleProduct.Price = p.Price

	r = database.DB.Save(&singleProduct)
	if r.Error != nil {
		log.Println(r.Error)
		return errNoProductWasUpdated
	}
	count := r.RowsAffected
	if count == 0 {
		return errNoProductWasUpdated
	}

	return nil
}

func DeleteProduct(OwnerID string, ProductID string) error {
	//Todo : See if the product is refrencedn in any orders first
	r := database.DB.Where("id=? AND ownerid=?", OwnerID, ProductID).Delete(&models.Product{})
	if r.Error != nil {
		return errCouldNotDeleteProduct
	}
	return nil
}

func GetCurrentStock(ProductId int) int {
	type StockReq struct {
		Stock int
	}
	var strq StockReq
	database.DB.Model(&models.Product{}).Limit(1).Find(strq)
	return strq.Stock
}

//Database transactions mehtod that is dependat if the invoice was placed succefully
func GetProductsByOwnerID(owner_id int64) (*ProductsList, error) {
	list := new(ProductsList)
	res := database.DB.Find(&list.Products, "owner_id=?", owner_id)
	if res.Error != nil {
		return list, errCouldNotGetProducts
	}
	return list, nil
}

func GetProductsByCategoryID(category_id int64) (*ProductsList, error) {
	list := new(ProductsList)
	res := database.DB.Find(&list.Products, "category_id=?", category_id)
	if res.Error != nil {
		return list, errCouldNotGetProducts
	}
	return list, nil
}

func GetProductByProductID(productID int64) *models.Product {
	var singleProduct models.Product
	res := database.DB.First(&singleProduct, productID)
	if res.Error != nil {
		return nil
	}
	return &singleProduct
}
