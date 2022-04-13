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
	d, err := gorm.Open("mysql", "root:leksyking7@/bookStoredb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error ti poju")
		panic(err)
		fmt.Println("Error ti poju")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
