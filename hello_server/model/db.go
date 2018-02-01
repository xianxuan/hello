package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Init() {
	db, _ := gorm.Open("mysql", "root:root@/hello?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&Hello{})
}
func DB() *gorm.DB {
	return db
}

func Session() *gorm.DB {
	return db.Begin()
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
