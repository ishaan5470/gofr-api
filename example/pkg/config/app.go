package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// the whole point of this file is to return a variable called db which will help other files to interact with db

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "ishaan:ishaan5470@/tablee?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
