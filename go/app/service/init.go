package service

import (
	"Stay_watch/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db := connect()
	closer, err := db.DB()
	if err != nil {
		return
	}
	defer closer.Close()

	db.AutoMigrate(&model.User{}, &model.Log{}, &model.Room{}, &model.Stayer{}, &model.Tag{}, &model.TagMap{})
}

func connect() *gorm.DB {
	// connect to the database
	// flag.Parse()
	// dsn := "docker:docker@tcp(localhost:33063)/app-db?charset=utf8mb4&parseTime=True&loc=Local"

	// if flag.Arg(0) == "production" {
	// 	dsn = "docker:docker@tcp(db:3306)/app-db?charset=utf8mb4&parseTime=True&loc=Local"
	// }
	dsn := "root:root@tcp(vol_mysql:3306)/app?charset=utf8mb4&parseTime=true"

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database")
	return gormDB
}
