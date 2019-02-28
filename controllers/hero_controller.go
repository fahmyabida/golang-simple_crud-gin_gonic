package controllers

import (
	_ "../config"
	"../handler"
	"../model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HeroController struct {
	*sql.DB
}

func(heroController *HeroController) GetAllHero(ctx *gin.Context) {
	var (
		heroes      []model.Hero
		result      gin.H
		scanSuccess= true
	)
	rows, err := heroController.Query("SELECT * FROM hero")
	if err != nil {
		result = *handler.ErrorQuery(&err, &result)
		ctx.JSON(http.StatusOK, result)
		return
	}
	for rows.Next() {
		var hero model.Hero
		err = rows.Scan(&hero.Id, &hero.Name, &hero.Atk_type, &hero.IdCategoryHero)
		if err != nil {
			scanSuccess = false
		}
		heroes = append(heroes, hero)
	}
	if len(heroes) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0}
	} else {
		result = gin.H{
			"result": heroes,
			"count":  len(heroes)}
	}
	if scanSuccess == false {
		result = *handler.ErrorScanQuery(&err, &result)
	}
	ctx.JSON(http.StatusOK, result)
}

func(heroController *HeroController) GetAllHeroV2(ctx *gin.Context) {
	var (
		heroDTOs     []model.HeroDTO
		result      gin.H
		scanSuccess= true
	)
	rows, err := heroController.Query("SELECT * FROM hero")
	if err != nil {
		result = *handler.ErrorQuery(&err, &result)
		ctx.JSON(http.StatusOK, result)
		return
	}
	for rows.Next() {
		var (
			hero model.Hero
			categoryHero model.CategoryHero
			heroDTO model.HeroDTO
		)
		err = rows.Scan(&hero.Id, &hero.Name, &hero.Atk_type, &hero.IdCategoryHero)
		if err != nil {
			scanSuccess = false
		}
		row := heroController.QueryRow("SELECT * FROM category_hero WHERE id=?", hero.IdCategoryHero)
		if err2 := row.Scan(&categoryHero.Id, &categoryHero.Category_name); err2 != nil {
			scanSuccess = false
		}
		heroDTO = model.HeroDTO{hero.Id, hero.Name, hero.Atk_type, categoryHero}
		heroDTOs = append(heroDTOs, heroDTO)
	}
	if len(heroDTOs) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0}
	} else {
		result = gin.H{
			"result": heroDTOs,
			"count":  len(heroDTOs)}
	}
	if scanSuccess == false {
		result = *handler.ErrorScanQuery(&err, &result)
	}
	ctx.JSON(http.StatusOK, result)
}


func(heroController *HeroController) DeleteHero(ctx *gin.Context) {
	var (
		heroDelete model.HeroDelete
		result gin.H
	)
	ctx.ShouldBindJSON(&heroDelete)
	idHero, err := strconv.Atoi(heroDelete.Id)
	hero := model.Hero{Id:idHero}
	if err != nil {
		result = gin.H{
			"error"		: "Cant Delete",
			"message"	: "Paramater input is not valid",
		}
	} else{
		sqlStatement := `DELETE FROM hero WHERE id = ?;`
		_, err = heroController.Exec(sqlStatement, hero.Id)
		if err != nil {
			result = gin.H{
				"error"		: "Cant Delete",
				"message"	: "Id Hero with "+ heroDelete.Id+" not found",
			}
		} else{
			result = gin.H{
				"message": "Succesfully delete with id hero : " + heroDelete.Id,
			}
		}
	}
	ctx.JSON(http.StatusOK, result)
}

//func (idb *InDB) CreateHeroWithPostForm (context *gin.Context)  {
//	hero := model.Hero{
//		Name 		: context.PostForm("name_hero"),
//		Atk_type 	: context.PostForm("attack_type"),
//	}
//	idb.DB.Save(hero)
//	context.JSON(http.StatusOK, hero)
//}
//
//func CreateHero (context *gin.Context){
//var (
//	postHero model.PostHero
//)
//db := *config.DBInit()
////binding request to Json and put to the postHero
//context.ShouldBindJSON(&postHero)
//idCategoryHero, _ := strconv.Atoi(postHero.CategoryHero)
//
////mapping from postHero(DTO) to the real model in DB
//hero := model.Hero{
//	Name			: postHero.Jeneng_hero,
//	Atk_type		: postHero.Jenis_serangan,
//	IdCategoryHero	: idCategoryHero,
//}
//
//save to the DB
//db.Exec("INSERT INTO hero(name, atk_type, id_category_hero) VALUES ('fahmy', 'range', 3)")
//
//returning json
//context.JSON(http.StatusOK, "berhasil")
//fmt.Println(&db)
//}


		//
//import (
//	"../../handler"
//	"../../config"
//	"../../model"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func GetAllHero(context *gin.Context){
//	var (
//		heroes []model.Hero
//		result gin.H
//		scanSuccess = true
//	)
//	result = gin.H{}
//	db := *config.DBInit()
//	rows, err := db.("SELECT * FROM hero")
//	if err != nil {
//		result = *handler.ErrorQuery(&err,&result)
//		context.JSON(http.StatusOK, result)
//		return
//	}
//	for rows.Next() {
//		var hero model.Hero
//		err = rows.Scan(&hero.Id, &hero.Name, &hero.Atk_type, &hero.IdCategoryHero)
//		if err != nil {
//			scanSuccess = false
//		}
//		heroes = append(heroes, hero)
//	}
//	if len(heroes) <= 0 {
//		result = gin.H{
//			"result" : nil,
//			"count" : 0 }
//	} else{
//		result = gin.H{
//			"result": heroes,
//			"count":  len(heroes) }
//	}
//	if scanSuccess == false {
//		result = *handler.ErrorScanQuery(&err,&result)
//	}
//	context.JSON(http.StatusOK, result)
//}
//
//func (idb *InDB) GetAllHeroWithMapping (c *gin.Context)  {
//	var (
//		heroes []model.Hero
//		heroesDto []model.HeroDTO
//		result gin.H
//	)
//	idb.DB.Table("hero").Find(&heroes)
//	if len(heroes) <= 0 {
//		result = gin.H{
//			"result" : nil,
//			"count" : 0 }
//	} else{
//		//for item := 0; item <= len(heroes); item++ {
//		for _, item := range heroes  {
//			var categoryHero model.CategoryHero
//			//simple definition forEach in GO --> https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go
//			idb.DB.Table("category_hero").Find(&categoryHero, "id = "+strconv.Itoa(item.IdCategoryHero))
//			heroesDto = append(heroesDto, model.HeroDTO{
//				//Pengenal:       item.Id,
//				Jeneng_hero:    item.Name,
//				Jenis_serangan: item.Atk_type,
//				CategoryHero:   categoryHero.Category_name})
//				//Pengenal:       heroes[item].Id,
//				//Jeneng_hero:    heroes[item].Name,
//				//Jenis_serangan: heroes[item].Atk_type,
//				//CategoryHero:   categoryHero.Category_name})
//			//a++
//		}
//			result = gin.H{
//				"result": heroesDto,
//				"count":  len(heroes)}
//	}
//	c.JSON(http.StatusOK, result)
//}
//
//func (idb *InDB) CustomGetAllHero(context *gin.Context){
//	var (
//		heroes []model.Hero
//		result  gin.H
//	)
//	idb.DB.Table("hero").
//		Select("hero.name, hero.atk_type, category_hero.category_name").
//		Joins("INNER JOIN category_hero ON hero.id_category_hero=category_hero.id").
//		Find(&heroes)
//	log.Println("halo hero :", heroes)
//	if len(heroes) <= 0 { //check length of heroes
//		result = gin.H{
//			"result": nil,
//			"count":  0,
//		}
//	} else {
//		result = gin.H{
//			"result": heroes,
//			"count":  len(heroes),
//		}
//	}
//	context.JSON(http.StatusOK, result)
//}
