package config

import (
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/godb" +
		//"?charset=utf8&parseTime=True&loc=Local"+
		"")
	if err != nil {
		panic("failed to connect to database")
	}
	db.LogMode(true)
	return db
}