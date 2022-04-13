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
	d, err := gorm.Open("mysql", "Leksyking:ogundipe134@/bookStoreDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connection error")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
