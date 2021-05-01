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
type ProductsList struct {
	Products   []Product `json:"products"`
	StatusCode int       `json:"status"`
}

type Catergory struct {
	CatergoryID    int64  `json:"categoryid"`
	CatergoryName  string `json:"categoryname"`
	CatergoryImage string `json:"image"`
}
type CategoryLists struct {
	Categories []Catergory `json:"categories"`
	StatusCode int         `json:"status"`
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
func GetProductsByOwnerID(owner_id int64) (*ProductsList, error) {
	result := new(ProductsList)
	rows, err := database.DB.Query(`
	SELECT product_id, 
  	category_id,
	product_name,
	product_image,
	product_image_large,
	descriptions,
	price,stock,
	product_packtype
	FROM products WHERE owner_id=?;
	`, owner_id)

	if err != nil {
		result.StatusCode = -500
		return result, err
	}

	for rows.Next() {
		singleProduct := Product{}
		if err := rows.Scan(
			&singleProduct.ProductID,
			&singleProduct.CategoryID,
			&singleProduct.ProductName,
			&singleProduct.ProductImage,
			&singleProduct.ProductImageLarge,
			&singleProduct.Description,
			&singleProduct.Price,
			&singleProduct.Stock,
			&singleProduct.PackingType); err != nil {
			result.StatusCode = -501
			return result, err
		}
		result.Products = append(result.Products, singleProduct)
	}
	if err != nil {
		result.StatusCode = -502
		return result, err
	}
	if result.Products == nil {
		result.StatusCode = -503
		result.Products = []Product{}
	}
	defer rows.Close()
	return result, nil
}

func GetProductsByCategoryID(category_id int64) (*ProductsList, error) {
	result := new(ProductsList)
	rows, err := database.DB.Query(`
	SELECT product_id, 
	owner_id,
  	category_id,
	product_name,
	product_image,
	product_image_large,
	descriptions,
	price,stock,
	product_packtype
	FROM products WHERE category_id=?;
	`, category_id)

	if err != nil {
		result.StatusCode = -500
		return result, err
	}

	for rows.Next() {
		singleProduct := Product{}
		if err := rows.Scan(
			&singleProduct.ProductID,
			&singleProduct.OwnerID,
			&singleProduct.CategoryID,
			&singleProduct.ProductName,
			&singleProduct.ProductImage,
			&singleProduct.ProductImageLarge,
			&singleProduct.Description,
			&singleProduct.Price,
			&singleProduct.Stock,
			&singleProduct.PackingType); err != nil {
			result.StatusCode = -501
			return result, err
		}
		result.Products = append(result.Products, singleProduct)
	}
	if err != nil {
		result.StatusCode = -502
		return result, err
	}
	if result.Products == nil {
		result.StatusCode = -503
		result.Products = []Product{}
	}
	defer rows.Close()
	return result, nil
}

func GetCategories() (*CategoryLists, error) {
	result := new(CategoryLists)
	rows, err := database.DB.Query("SELECT category_id,category_name,category_image FROM categories")
	if err != nil {
		result.StatusCode = -500
		return result, err
	}

	for rows.Next() {
		singleCategory := Catergory{}
		if err := rows.Scan(&singleCategory.CatergoryID, &singleCategory.CatergoryName, &singleCategory.CatergoryImage); err != nil {
			result.StatusCode = -501
			return result, err
		}
		result.Categories = append(result.Categories, singleCategory)
	}
	if err != nil {
		result.StatusCode = -502
		return result, err
	}
	//To avoid passing null back to the user
	if result.Categories == nil {
		result.StatusCode = -503
		result.Categories = []Catergory{}
	}
	defer rows.Close()
	return result, nil
}
