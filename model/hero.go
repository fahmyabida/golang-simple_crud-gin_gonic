package model

type Hero struct {
	Id             int		`json:"id"`
	Name           string	`json:"name"`
	Atk_type       string	`json:"attack_type"`
	IdCategoryHero int		`json:"id_category_hero"`
}

type HeroDelete struct {
	Id				string 	`json:"id_hero"`
}

type HeroDTO struct {
	Id             	int				`json:"id"`
	Name           	string			`json:"name"`
	Atk_type       	string			`json:"attack_type"`
	CategoryHero 	CategoryHero	`json:"category_hero"`
}
