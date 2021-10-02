package store

import (
	"errors"
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
)

var (
	errCouldNotCreateCategory   = errors.New("could not create category")
	errCouldNotGetCategories    = errors.New("could not fetch categories")
	errCouldNotDeleteCategories = errors.New("could not delete the category")
)

type Catergory struct {
	CatergoryID    int64  `json:"categoryid"`
	CatergoryName  string `json:"categoryname"`
	CatergoryImage string `json:"image"`
}
type CategoryLists struct {
	Categories []models.Category `json:"categories"`
	StatusCode int               `json:"status"`
}

func (c *Catergory) AddCategory() error {
	newCategory := models.Category{
		CatergoryName:  c.CatergoryName,
		CatergoryImage: c.CatergoryImage,
	}

	res := database.DB.Create(&newCategory)
	if res.Error != nil {
		log.Println(res.Error)
		return errCouldNotCreateCategory
	}

	return nil
}

func (c *Catergory) UpdateCategory() error {
	var singleCategory models.Category
	r := database.DB.First(&singleCategory).Where(c.CatergoryID)
	if r.Error != nil {
		log.Println(r)
		return errCouldNotUpdateProduct
	}
	singleCategory.CatergoryImage = c.CatergoryImage
	singleCategory.CatergoryName = c.CatergoryName

	r = database.DB.Save(&singleCategory)
	count := r.RowsAffected
	if count == 0 {
		return errNoProductWasUpdated
	}
	return nil
}

func (c *Catergory) DeleteCategory() error {
	var singleCategory models.Category
	r := database.DB.Delete(&singleCategory, c.CatergoryID)
	if r.Error != nil {
		log.Println(r)
		return errCouldNotDeleteCategories
	}

	count := r.RowsAffected
	if count == 0 {
		return errCouldNotDeleteCategories
	}

	return nil
}
func GetCategories() (*CategoryLists, error) {
	list := new(CategoryLists)
	res := database.DB.Find(&list.Categories)
	if res.Error != nil {
		return list, errCouldNotGetCategories
	}
	return list, nil
}
