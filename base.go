package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := config.DBInit()
	hero := &controllers.HeroController{Db:dbConn}
	router := gin.Default()
	router.GET("/heroes", hero.GetAllHero)
	router.GET("/v2/heroes", hero.GetAllHeroWithDTO)
	router.DELETE("/hero", hero.DeleteHero)
	router.POST("/hero", hero.CreateHeroWithPostForm)
	router.POST("/v2/hero", hero.CreateHeroWithBodyRequest)
	router.Run(":8080")
}
