package model

type Hero struct {
	Id 					int
	Name				string `gorm:"column:name"`
	Atk_type			string
	IdCategoryHero		int `gorm:"column:id_category_hero"`
}

func (Hero) TableName() string {
	return "hero"
}

type HeroDTO struct {
	Pengenal 		int
	Jeneng_hero		string
	Jenis_serangan	string
	CategoryHero	string
}