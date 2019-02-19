package main

import (
	"./config"
	"./controllers"
	"./controllers/hero"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDBHero := &hero.InDB{DB: db}
	router := gin.Default()
	router.GET("/hello", controllers.HelloWorld)
	router.GET("/heroes", inDBHero.GetAllHero)
	router.GET("/heroes/mapping", inDBHero.GetAllHeroWithMapping)
	router.GET("/customheroes", inDBHero.CustomGetAllHero)
	router.POST("/hero", inDBHero.CreateHero)
	router.POST("/hero2", inDBHero.CreateHeroWithPostForm)
	router.Run(":8080")
}
