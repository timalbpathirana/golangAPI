package models

import (
	"github.com/jinzhu/gorm"
	"github.com/timalbpathirana/golangAPI/pkg/config"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	ProductID           int
	ProductName         string
	ProductSize         string
	ProductColour       string
	ProductPrice        int
	ProductLocation     string
	ProductAvailability string
}

//https://v1.gorm.io/docs/query.html

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	//https://v1.gorm.io/docs/create.html
	if b := db.NewRecord(p); b != false {
		panic("Error Creating the Product")
	}
	return p
}

func GetProduct() []Product {
	var product []Product
	db.Find(&product)
	return product
}

func GetProductById(i int64) (*Product, *gorm.DB) {
	var product Product
	// Get all matched records
	// SELECT * FROM users WHERE ProductID = '10';
	db.Where("product_id = ?", i).Find(&product)
	return &product, db
}

func DeleteProductById(id int64) Product {
	var product Product
	db.Where("product_id = ?", id).Delete(product)
	return product
}
