package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
type Blog_article struct {
	 gorm.Model

	 Id int8
	 title string
	 UserId int
	 Description string
	 Photo string
	 Public int
	 Content string
	 CreateAt time.Time
	 UpdateAt time.Time
}
