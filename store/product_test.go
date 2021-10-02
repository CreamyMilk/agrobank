package store

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/CreamyMilk/agrobank/database"
// )

// func TestAddProduct(t *testing.T) {
// 	if err := database.Connect(); err != nil {
// 		fmt.Printf("DB ERROR %v", err)
// 	}
// 	defer database.DB.Close()

// 	tt := []struct {
// 		categoryID        int64
// 		ownerID           int64
// 		productName       string
// 		productImage      string
// 		productImageLarge string
// 		description       string
// 		packingType       string
// 		stock             int
// 		price             float64
// 		possible          bool
// 	}{
// 		{1, 1, "Carrots", "http://peodIMage", "largeImage", "descrtiption", "box", 60, 12.2, true},
// 		{1, 99, "Unkown Owner", "http://peodIMage", "largeImage", "descrtiption", "box", 60, 12.2, false},
// 		{1, 1, "Beans", "http://peodIMage", "largeImage", "descrtiption", "box", 60, 12.2, true},
// 		{2, 1, "Unknown Category", "http://peodIMage", "largeImage", "descrtiption", "box", 60, 12.2, false},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.productName, func(t *testing.T) {
// 			w := Product{CategoryID: tc.categoryID,
// 				ProductName:       tc.productName,
// 				ProductImage:      tc.productImage,
// 				ProductImageLarge: tc.productImageLarge,
// 				Description:       tc.description,
// 				PackingType:       tc.packingType,
// 				Stock:             tc.stock,
// 				Price:             tc.price}
// 			e := w.AddProduct()
// 			if e != nil && tc.possible {
// 				t.Errorf("Cannot create account because %v", e)
// 			}
// 			if err := w.DeleteProduct(); err != nil {
// 				t.Errorf("Cannot delete wallet %s because : %v", w.ProductName, err)
// 			}
// 		})
// 	}
// }
