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
	//defer db.Close()
	//db.Model(&model.Hero{}).Related(&model.CategoryHero{}, "fk_id_category_hero")
	// #AutoCreate Table in DB
	//db.AutoMigrate(&model.Car{}, &model.Person{})
	db.LogMode(true)
	return db
}