package controllers

import (
	_ "../config"
	"../handler"
	"../model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type HeroController struct {
	Db *sql.DB
	code int
}

func(heroController *HeroController) GetAllHero(ctx *gin.Context) {
	var (
		heroes      []model.Hero
		result      gin.H
		scanSuccess= true
	)
	rows, err := heroController.Db.Query("SELECT * FROM hero")
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

func(heroController *HeroController) GetAllHeroWithDTO(ctx *gin.Context) {
	var (
		heroDTOs    []model.HeroGetDTO
		result      gin.H
		scanSuccess= true
	)
	rows, err := heroController.Db.Query("SELECT * FROM hero")
	if err != nil {
		result = *handler.ErrorQuery(&err, &result)
		ctx.JSON(http.StatusOK, result)
		return
	}
	for rows.Next() {
		var (
			hero model.Hero
			categoryHero model.CategoryHero
			heroDTO model.HeroGetDTO
		)
		err = rows.Scan(&hero.Id, &hero.Name, &hero.Atk_type, &hero.IdCategoryHero)
		if err != nil {
			scanSuccess = false
		}
		row := heroController.Db.QueryRow("SELECT * FROM category_hero WHERE id=?", hero.IdCategoryHero)
		if err2 := row.Scan(&categoryHero.Id, &categoryHero.Category_name); err2 != nil {
			scanSuccess = false
		}
		heroDTO = model.HeroGetDTO{hero.Id, hero.Name, hero.Atk_type, categoryHero}
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
		heroDelete model.HeroDeleteDTO
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
		row, err := heroController.Db.Exec(sqlStatement, hero.Id)
		rowAffected,_ := row.RowsAffected()
		if err != nil || rowAffected == 0 {
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

func (heroController *HeroController) CreateHeroWithPostForm (ctx *gin.Context)  {
	result := gin.H{}
	hero := model.Hero{}
	if id_category_hero, err := strconv.Atoi(ctx.PostForm("id_category_hero")); err != nil {

	} else {
		hero.SetName(ctx.PostForm("name_hero"))
		hero.SetAtkType(ctx.PostForm("attack_type"))
		hero.SetIDCategoryHero(id_category_hero)
	}
	row, err := heroController.Db.Exec("INSERT INTO hero (name, atk_type, id_category_hero) VALUES (?,?,?)",
		&hero.Name, &hero.Atk_type, &hero.IdCategoryHero)
	lastInsertedID, _ := row.LastInsertId()
	if err != nil {

		result = gin.H{
			"error":"something wrong with insert data",
		}
	} else if rowAffected,_ :=row.RowsAffected(); rowAffected > 1 {
		result = gin.H{
			"error":"something wrong with insert data",
		}
		log.Panic("Data inserted more than 1 row")
	} else {
		result = gin.H{
			"message":"data succesfully inserted with id "+strconv.FormatInt(lastInsertedID,10),
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func(heroController *HeroController) CreateHeroWithBodyRequest (ctx *gin.Context){
	postHero := model.HeroPostDTO{}
	result := gin.H{}
	err := ctx.ShouldBindJSON(&postHero)
	if (err != nil){
		heroController.code = http.StatusOK
		result = gin.H{
			"error":"body request not valid",
		}
	} else {
		//save to the DB
		row, err := heroController.Db.Exec("INSERT INTO hero(name, atk_type, id_category_hero) VALUES (?, ?, ?)",
			&postHero.Name, &postHero.Atk_type, &postHero.IdCategoryHero)
		lastInsertedID, _ := row.LastInsertId()
		if err != nil {
			result = gin.H{
				"error":"something wrong with insert data",
			}
		} else if rowAffected,_ :=row.RowsAffected(); rowAffected > 1 {
			result = gin.H{
				"error":"something wrong with insert data",
			}
			log.Panic("Data inserted more than 1 row")
		} else {
			heroController.code = http.StatusOK
			result = gin.H{
				"message":"data succesfully inserted with id "+strconv.FormatInt(lastInsertedID,10),
				"data":postHero,
			}
		}
	}
	ctx.JSON(heroController.code, result)
}