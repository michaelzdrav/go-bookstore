package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//gorm db connection: "dbType",  "username:password/nameOfTable?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:ANSKk08aPEDbFjDO@tcp(localhost:3306)/testing?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
