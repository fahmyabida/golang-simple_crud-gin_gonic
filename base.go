package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := config.DBInit()
	hero := &controllers.HeroController{Db: dbConn}
	router := gin.Default()
	version1 := router.Group("/v1")
	{
		version1.GET("/heroes", hero.GetAllHero)
		version1.DELETE("/hero", hero.DeleteHero)
		version1.POST("/hero", hero.CreateHeroWithPostForm)
	}

	version2 := router.Group("/v2")
	{
		version2.GET("/v2/heroes", hero.GetAllHeroWithDTO)
		version2.POST("/v2/hero", hero.CreateHeroWithBodyRequest)
		router.Run(":8080")
	}
}
