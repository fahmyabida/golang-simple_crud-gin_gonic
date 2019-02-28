package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	hero := &controllers.HeroController{db}
	router := gin.Default()
	router.GET("/heroes", hero.GetAllHero)
	router.GET("/heroesv2", hero.GetAllHeroV2)
	router.DELETE("/hero", hero.DeleteHero)
	router.Run(":8080")
}
