package model

type CategoryHero struct {
	Id 				int 	`json:"id"`
	Category_name 	string 	`json:"category_name"`
}

type CategoryGetDTO struct {
	Category_hero	string	`json:"category_hero"`
}
