package service

import (
	"Stay_watch/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbEngine *gorm.DB

func init() {
	dsn := "root:root@tcp(vol_mysql:3306)/app?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{}, &model.Log{}, &model.Room{}, &model.Stayer{})
}
