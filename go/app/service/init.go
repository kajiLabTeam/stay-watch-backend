package service

import (
	"fmt"
	"os"
	"strings"

	"Stay_watch/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	db := connect()
	closer, err := db.DB()
	if err != nil {
		return
	}
	defer closer.Close()
	db.AutoMigrate(
		&model.User{},
		&model.Log{},
		&model.Room{},
		&model.Stayer{},
		&model.Tag{},
		&model.TagMap{},
		&model.Building{},
		&model.Beacon{},
		&model.DeletedUser{},
		&model.Community{},
	)
}

func connect() *gorm.DB {
	_, envFlag := os.LookupEnv("ENVIRONMENT")
	// 環境変数がセットされていない場合は.envを読み込む（dockerを使用しない開発時)
	if !envFlag {
		fmt.Println("ENVIRONMENT is not set")

		envPath := "../../.env"
		// test実行時はenvのディレクトリが変わる
		if strings.HasSuffix(os.Args[0], ".test") {
			envPath = "../../../.env"
		}

		err := godotenv.Load(
			envPath,
		)
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	dsn := "root:root@tcp(localhost:33066)/app?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	// 本番環境の場合はコンテナ名で接続する
	if os.Getenv("ENVIRONMENT") == "production" {
		dsn = "root:root@tcp(vol_mysql:3306)/app?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB
}
