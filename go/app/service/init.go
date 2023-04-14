package service

import (
	"Stay_watch/model"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	db.AutoMigrate(&model.User{}, &model.Log{}, &model.Room{}, &model.Stayer{}, &model.Tag{}, &model.TagMap{}, &model.Building{}, &model.BeaconType{}, &model.DeletedUser{})

	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		//複数のユーザーを作成
		users := []model.User{
			{
				Name:         "test",
				Email:        "toge7113+test-stay-watch@gmail.com",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ab00",
				BeaconTypeId: 1,
				CommunityId:  1,
			},
			{
				Name:         "kaji",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac00",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ogane",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac01",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "miyagawa",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac02",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ayato",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac03",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ken",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac04",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "suzaki",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac05",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "akito",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac06",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "fueta",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac07",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "kameda",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac08",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "maruyama",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac09",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ohashi",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0a",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "rui",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0b",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "shamo",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0c",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "terada",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0d",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "toyama",
				Email:        "tatu2425@gmail.com",
				Role:         2,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0e",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ukai",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac0f",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "isiguro",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac10",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ao",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac11",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "fuma",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac12",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "ueji",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac13",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "oiwa",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac14",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "togawa",
				Email:        "toge7113@gmail.com",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac15",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "yada",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac16",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "yokoyama",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac17",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "kazuo",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac18",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "sakai",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac19",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "iwaguti",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac1a",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
			{
				Name:         "makino",
				Email:        "",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07ac1b",
				BeaconTypeId: 1,
				CommunityId:  2,
			},
		}
		db.Create(&users)
	}

	db.Model(&model.DeletedUser{}).Count(&count)
	if count == 0 {
		deletedUsers := []model.DeletedUser{
			{
				Name:         "deleted-test",
				Email:        "deleted-test-staywatch@gmail.com",
				Role:         1,
				UUID:         "e7d61ea3f8dd49c88f2ff2484c07deleted-test",
				BeaconTypeId: 1,
				CommunityId:  1,
				UserId:       0,
			},
		}
		db.Create(&deletedUsers)
	}

	db.Model(&model.Building{}).Count(&count)
	if count == 0 {
		buildings := []model.Building{
			{
				Name:    "4号館",
				MapFile: "/4g-honkan-bekkan.jpg",
			},
			{
				Name:    "4号館別館",
				MapFile: "/4goubekkan.jpg",
			},
		}
		db.Create(&buildings)
	}

	db.Model(&model.Room{}).Count(&count)
	if count == 0 {
		rooms := []model.Room{
			{
				Name:        "梶研-学生部屋",
				BuildingID:  1,
				CommunityID: 2,
				Polygon:     "0,0-0,0",
			},
			{
				Name:        "梶研-スマートルーム",
				BuildingID:  1,
				CommunityID: 2,
				Polygon:     "0,0-0,0",
			},
			{
				Name:        "梶研-院生部屋",
				BuildingID:  1,
				CommunityID: 2,
				Polygon:     "0,0-0,0",
			},
			{
				Name:        "梶研-FA部屋",
				BuildingID:  1,
				CommunityID: 2,
				Polygon:     "0,0-0,0",
			},
			{
				Name:        "梶研-先生部屋",
				BuildingID:  1,
				CommunityID: 2,
				Polygon:     "0,0-0,0",
			},
		}
		db.Create(&rooms)
	}

	db.Model(&model.Tag{}).Count(&count)
	if count == 0 {
		tags := []model.Tag{
			{
				Name: "テストタグ",
			},
			{
				Name: "梶研",
			},
			{
				Name: "ロケーション",
			},
			{
				Name: "インタラクション",
			},
			{
				Name: "センシング",
			},
			{
				Name: "B1",
			},
			{
				Name: "B2",
			},
			{
				Name: "B3",
			},
			{
				Name: "B4",
			},
			{
				Name: "M1",
			},
			{
				Name: "M2",
			},
			{
				Name: "Professor",
			},
		}
		db.Create(&tags)
	}

	db.Model(&model.TagMap{}).Count(&count)
	if count == 0 {
		tagMaps := []model.TagMap{
			{
				UserID: 1,
				TagID:  1,
			},
			{
				UserID: 2,
				TagID:  2,
			},
			{
				UserID: 2,
				TagID:  12,
			},
			{
				UserID: 3,
				TagID:  2,
			},
			{
				UserID: 3,
				TagID:  3,
			},
			{
				UserID: 4,
				TagID:  12,
			},
			{
				UserID: 5,
				TagID:  3,
			},
			{
				UserID: 5,
				TagID:  4,
			},
			{
				UserID: 5,
				TagID:  12,
			},
			{
				UserID: 5,
				TagID:  2,
			},
			{
				UserID: 5,
				TagID:  5,
			},
			{
				UserID: 5,
				TagID:  10,
			},
			{
				UserID: 6,
				TagID:  2,
			},
			{
				UserID: 6,
				TagID:  5,
			},
			{
				UserID: 6,
				TagID:  10,
			},
			{
				UserID: 7,
				TagID:  2,
			},
			{
				UserID: 7,
				TagID:  3,
			},
			{
				UserID: 7,
				TagID:  10,
			},
			{
				UserID: 8,
				TagID:  2,
			},
			{
				UserID: 8,
				TagID:  3,
			},
			{
				UserID: 8,
				TagID:  9,
			},
			{
				UserID: 9,
				TagID:  2,
			},
			{
				UserID: 9,
				TagID:  4,
			},
			{
				UserID: 9,
				TagID:  9,
			},
			{
				UserID: 10,
				TagID:  2,
			},
			{
				UserID: 10,
				TagID:  3,
			},
			{
				UserID: 10,
				TagID:  9,
			},
			{
				UserID: 11,
				TagID:  2,
			},
			{
				UserID: 11,
				TagID:  5,
			},
			{
				UserID: 11,
				TagID:  9,
			},
			{
				UserID: 12,
				TagID:  2,
			},
			{
				UserID: 12,
				TagID:  3,
			},
			{
				UserID: 12,
				TagID:  9,
			},
			{
				UserID: 13,
				TagID:  2,
			},
			{
				UserID: 13,
				TagID:  5,
			},
			{
				UserID: 13,
				TagID:  9,
			},
			{
				UserID: 14,
				TagID:  2,
			},
			{
				UserID: 14,
				TagID:  4,
			},
			{
				UserID: 14,
				TagID:  9,
			},
			{
				UserID: 15,
				TagID:  2,
			},
			{
				UserID: 15,
				TagID:  3,
			},
			{
				UserID: 16,
				TagID:  2,
			},
			{
				UserID: 16,
				TagID:  9,
			},
			{
				UserID: 17,
				TagID:  2,
			},
			{
				UserID: 17,
				TagID:  3,
			},
			{
				UserID: 17,
				TagID:  9,
			},
			{
				UserID: 18,
				TagID:  2,
			},
			{
				UserID: 18,
				TagID:  8,
			},
			{
				UserID: 19,
				TagID:  2,
			},
			{
				UserID: 19,
				TagID:  8,
			},
			{
				UserID: 20,
				TagID:  2,
			},
			{
				UserID: 20,
				TagID:  8,
			},
			{
				UserID: 21,
				TagID:  2,
			},
			{
				UserID: 21,
				TagID:  8,
			},
			{
				UserID: 22,
				TagID:  2,
			},
			{
				UserID: 22,
				TagID:  8,
			},
			{
				UserID: 23,
				TagID:  2,
			},
			{
				UserID: 23,
				TagID:  8,
			},
			{
				UserID: 24,
				TagID:  2,
			},
			{
				UserID: 24,
				TagID:  8,
			},
			{
				UserID: 25,
				TagID:  2,
			},
			{
				UserID: 25,
				TagID:  8,
			},
			{
				UserID: 26,
				TagID:  2,
			},
			{
				UserID: 26,
				TagID:  8,
			},
			{
				UserID: 27,
				TagID:  2,
			},
			{
				UserID: 27,
				TagID:  8,
			},
			{
				UserID: 28,
				TagID:  2,
			},
			{
				UserID: 28,
				TagID:  8,
			},
			{
				UserID: 29,
				TagID:  2,
			},
			{
				UserID: 29,
				TagID:  8,
			},
		}
		db.Create(&tagMaps)
	}

	db.Model(&model.BeaconType{}).Count(&count)
	if count == 0 {
		tags := []model.BeaconType{
			{
				Name:         "FCS1301",
				UuidEditable: true,
			},
			{
				Name:         "Android",
				UuidEditable: false,
			},
			{
				Name:         "iPhone",
				UuidEditable: false,
			},
		}
		db.Create(&tags)
	}

	// db.Model(&model.User{}).Count()

}

func connect() *gorm.DB {

	_, envFlag := os.LookupEnv("ENVIRONMENT")
	//環境変数がセットされていない場合は.envを読み込む（dockerを使用しない開発時)
	if !envFlag {
		fmt.Println("ENVIRONMENT is not set")

		envPath := "../../.env"
		//test実行時はenvのディレクトリが変わる
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

	//本番環境の場合はコンテナ名で接続する
	if os.Getenv("ENVIRONMENT") == "production" {
		dsn = "root:root@tcp(vol_mysql:3306)/app?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("DB connected")
	return gormDB
}
