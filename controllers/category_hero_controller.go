package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/me/golang-simple_crud-gin_gonic/model"
	"github.com/me/golang-simple_crud-gin_gonic/handler"
	"net/http"
)

type CategoryHeroController struct {
	Db *sql.DB
	code int
}

func(chController *CategoryHeroController) GetAllCategoryHero(ctx *gin.Context){
	var (
		scanSuccess = true
		categoryHeroes []model.CategoryGetDTO
		result gin.H
	)
	chController.code = http.StatusOK
	if rows, err := chController.Db.Query("SELECT category_name FROM category_hero"); err != nil{
		result = *handler.ErrorScanQuery(&err,&result)
		chController.code = http.StatusInternalServerError
	} else {
		for rows.Next() {
			var categoryHero model.CategoryGetDTO
			err = rows.Scan(&categoryHero.Category_hero)
			if err != nil {
				scanSuccess = false
			}
			categoryHeroes = append(categoryHeroes, categoryHero)
		}
		if len(categoryHeroes) <= 0 {
			result = gin.H{
				"result": nil,
				"count":  0}
		} else {
			result = gin.H{
				"result": categoryHeroes,
				"count":  len(categoryHeroes)}
		}
		if scanSuccess == false {
			result = *handler.ErrorScanQuery(&err, &result)
		}
	}
	ctx.JSON(chController.code, result)
}