package hero

import (
	"../../model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDB) CreateHeroWithPostForm (context *gin.Context)  {
	hero := model.Hero{
		Name 		: context.PostForm("name_hero"),
		Atk_type 	: context.PostForm("attack_type"),
	}
	idb.DB.Save(hero)
	context.JSON(http.StatusOK, hero)
}

func (idb *InDB) CreateHero (context *gin.Context){
	//initialize db & other
	postHero := model.PostHero{}

	//binding request to Json and put to the postHero
	context.ShouldBindJSON(&postHero)
	idCategoryHero, _ := strconv.Atoi(postHero.CategoryHero)

	//mapping from postHero(DTO) to the real model in DB
	hero := model.Hero{
		Name			: postHero.Jeneng_hero,
		Atk_type		: postHero.Jenis_serangan,
		IdCategoryHero	: idCategoryHero,
	}

	//save to the DB
	idb.DB.Create(&hero)

	//returning json
	context.JSON(http.StatusOK, hero)
}
