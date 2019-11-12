package db

import (
	"blog/config"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {
	config.DbConfig()
	db, err := gorm.Open("mysql", config.User+":"+config.Pwd+"@tcp("+config.Host+":3306)/"+config.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("DB connect err:", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
}
