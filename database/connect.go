package database

import (
	"fmt"
	"os"

	"github.com/CreamyMilk/agrobank/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB holds global database object
var DB *gorm.DB

// Connect to db
func Connect() error {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_STORAGENAME")
	)
	var err error
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SetupModels(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

func SeedTransactionCosts() {
	dtbl := []struct {
		upperLimit int64
		charge     int64
	}{
		{100, 1},
		{1000, 5},
		{10000, 10},
		{100000, 15},
		{1000000, 20},
		{10000000, 25},
		{100000000, 100},
		{999999999999, 200},
	}
	for _, t := range dtbl {
		singleCost := models.TransactionCost{
			Upper_limit: t.upperLimit,
			Charge:      t.charge,
		}
		DB.Create(&singleCost)
	}
}

func SeedCategories() {
	dtbl := []struct {
		name string
		img  string
	}{
		{"Groceries", "https://res.cloudinary.com/agrocr/image/upload/v1633492513/fresh-vegetables-on-grey-background-picture-id1251268582_tjgwa4.jpg"},
		{"Cash Crops", "https://res.cloudinary.com/agrocr/image/upload/v1633492648/african-woman-holding-tea-leaves-rwanda-picture-id1138313836_nqkdwq.jpg"},
		{"Farm Tools", "https://res.cloudinary.com/agrocr/image/upload/v1633492748/gardening-tools-seeds-and-soil-on-wooden-table-picture-id673641858_akwihu.jpg"},
		{"Fruits", "https://res.cloudinary.com/agrocr/image/upload/v1633492934/photo-1610832958506-aa56368176cf_mnbh20.jpg"},
	}
	for _, t := range dtbl {
		singleCost := models.Category{
			CatergoryName:  t.name,
			CatergoryImage: t.img,
		}
		DB.Create(&singleCost)
	}
}
