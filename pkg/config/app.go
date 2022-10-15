package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//https://v1.gorm.io/
	connectionString := "yourSqlUsername:yourSqlPassword@tcp(host:port)/ikeaDB?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Error:", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
