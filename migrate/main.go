package main

import (
	"log"

	"github.com/rajnarayan1729/goCrudPostgres/model"
	"github.com/rajnarayan1729/goCrudPostgres/utils"
)

func main() {

	db, err := utils.GetDbConn()

	if err != nil {

		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Product{})

	// Create
	db.Create(&model.Product{Code: "D42", Price: 100})

	// Read
	var product *model.Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(model.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
