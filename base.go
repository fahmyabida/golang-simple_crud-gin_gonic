package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	router.GET("/hello", controllers.Hello)
	router.GET("/heroes", inDB.GetAllHero)
	router.GET("/heroes/mapping", inDB.GetAllHeroWithMapping)
	router.GET("/customheroes", inDB.CustomGetAllHero)
	router.POST("/hero", inDB.CreateHeroWithPostForm)
	router.Run(":8080")
}
