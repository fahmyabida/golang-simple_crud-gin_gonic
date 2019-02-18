package model

type CategoryHero struct {
	Id				int
	Category_name	string
}

func (CategoryHero) TableName() string {
	return "category_hero"
}