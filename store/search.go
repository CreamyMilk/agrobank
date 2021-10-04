package store

import "github.com/CreamyMilk/agrobank/database"

func SearchProducts(queryString string) (*ProductsList, error) {
	list := new(ProductsList)
	res := database.DB.Find(&list.Products, "product_name LIKE ?", "%"+queryString+"%")
	if res.Error != nil {
		return list, errCouldNotGetProducts
	}
	return list, nil
}
