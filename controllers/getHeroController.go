package controllers

import (
	"../model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (idb *InDB) GetAllHero (context *gin.Context){
	var (
		heroes []model.Hero
		result gin.H
		)
	idb.DB.Table("hero").Find(&heroes)
	if len(heroes) <= 0 {
		result = gin.H{
			"result" : nil,
			"count" : 0 }
	} else{
		result = gin.H{
			"result": heroes,
			"count":  len(heroes) }
	}
	context.JSON(http.StatusOK, result)
}

func (idb *InDB) GetAllHeroWithMapping (c *gin.Context)  {
	var (
		heroes []model.Hero
		heroesDto []model.HeroDTO
		result gin.H
	)
	idb.DB.Table("hero").Find(&heroes)
	if len(heroes) <= 0 {
		result = gin.H{
			"result" : nil,
			"count" : 0 }
	} else{
		//for item := 0; item <= len(heroes); item++ {
		for _, item := range heroes  {
			var categoryHero model.CategoryHero
			//simple definition forEach in GO --> https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go
			idb.DB.Table("category_hero").Find(&categoryHero, "id = "+strconv.Itoa(item.IdCategoryHero))
			heroesDto = append(heroesDto, model.HeroDTO{
				Pengenal:       item.Id,
				Jeneng_hero:    item.Name,
				Jenis_serangan: item.Atk_type,
				CategoryHero:   categoryHero.Category_name})
				//Pengenal:       heroes[item].Id,
				//Jeneng_hero:    heroes[item].Name,
				//Jenis_serangan: heroes[item].Atk_type,
				//CategoryHero:   categoryHero.Category_name})
			//a++
		}
			result = gin.H{
				"result": heroesDto,
				"count":  len(heroes)}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CustomGetAllHero(context *gin.Context){
	var (
		heroes []model.Hero
		result  gin.H
	)
	idb.DB.Table("hero").
		Select("hero.name, hero.atk_type, category_hero.category_name").
		Joins("INNER JOIN category_hero ON hero.id_category_hero=category_hero.id").
		Find(&heroes)
	log.Println("halo hero :", heroes)
	if len(heroes) <= 0 { //check length of heroes
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": heroes,
			"count":  len(heroes),
		}
	}
	context.JSON(http.StatusOK, result)
}
