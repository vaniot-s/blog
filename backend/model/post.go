package model

type Post struct {
	Model
	Name      string `gorm:"column:name" json:"name"`
	Description   string `gorm:"column:description" json:"description"`
	content   string `gorm:"column:comment" json:"comment"`
}
