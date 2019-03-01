package model

type Hero struct {
	Id             	int		`json:"id"`
	Name           	string	`json:"name"`
	Atk_type       	string	`json:"attack_type"`
	IdCategoryHero 	int		`json:"id_category_hero"`
}

type HeroPostDTO struct {
	Name           	string	`json:"name"`
	Atk_type       	string	`json:"attack_type"`
	IdCategoryHero 	string		`json:"id_category_hero"`
}

type HeroDeleteDTO struct {
	Id				string 	`json:"id_hero"`
}

type HeroGetDTO struct {
	Id             	int				`json:"id"`
	Name           	string			`json:"name"`
	Atk_type       	string			`json:"attack_type"`
	CategoryHero 	CategoryHero	`json:"category_hero"`
}

func (h *Hero) SetName(name string) {
	h.Name = name
}

func (h *Hero) SetAtkType(attack_type string) {
	h.Atk_type = attack_type
}

func (h *Hero) SetIDCategoryHero(id_category_hero int) {
	h.IdCategoryHero = id_category_hero
}