package model

type Tag struct {
	Model
	Name      string `gorm:"column:name" json:"name"`
}
