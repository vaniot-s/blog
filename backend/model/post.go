package model

type App struct {
	Model
	//Configs   []AppConfig `gorm:"foreignkey:app_id;association_foreignkey::Id" json:"config"`
	Name      string `gorm:"column:name" json:"name"`
	Description   string `gorm:"column:version" json:"version"`
	content   string `gorm:"column:comment" json:"comment"`
}
