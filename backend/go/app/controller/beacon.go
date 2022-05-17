package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//ビーコン情報を受け取る
func Beacon(c *gin.Context) {

	beaconRoom := model.BeaconRoom{}

	err := c.Bind(&beaconRoom)
	if err != nil {
		fmt.Println("err=", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	RoomService := service.RoomService{}
	UserService := service.UserService{}
	//事前にStayerテーブルのデータを取得する
	pastAllStayer, err := RoomService.GetAllStayer()

	for _, pastStayer := range pastAllStayer {

		isExist := false
		targetUserRssi := -200

		for _, currentStayer := range beaconRoom.Beacons {

			pastUUID, er := UserService.GetUserUUIDByUserID(pastStayer.UserID)
			if er != nil {
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			//1つ前のstgayerテーブルにもいた場合
			if pastUUID == currentStayer.Uuid {
				targetUserRssi = int(currentStayer.Rssi)
				isExist = true
			}

		}

		//RSSIが以前より強い場合
		if isExist && targetUserRssi > int(pastStayer.Rssi) {
			//同じ部屋にいる場合は更新
			if beaconRoom.RoomID == pastStayer.RoomID {
				fmt.Println("同じ部屋にいる場合は更新 強くなる")
				err := RoomService.UpdateStayer(&model.Stayer{
					UserID: pastStayer.UserID,
					RoomID: pastStayer.RoomID,
					Rssi:   int64(targetUserRssi),
				})
				if err != nil {
					c.String(http.StatusInternalServerError, "Server Error")
					return
				}
			} else {
				fmt.Println("別の部屋にいる場合は削除")
				//別の部屋の場合Stayerテーブルから削除する
				err := RoomService.DeleteStayer(pastStayer.UserID)
				if err != nil {
					fmt.Println(err)
				}
				//logテーブルのendAtを更新する
				err = RoomService.InsertEndAt(pastStayer.UserID)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		// //RSSIが以前より弱い場合
		// if isExist && targetUserRssi < int(pastStayer.Rssi) {
		// 	//同じ部屋にいる場合は更新
		// 	if beaconRoom.RoomID == pastStayer.RoomID {
		// 		fmt.Println("同じ部屋にいる場合はRSSIを更新 弱くなる")
		// 		err := RoomService.UpdateStayer(&model.Stayer{
		// 			UserID: pastStayer.UserID,
		// 			RoomID: pastStayer.RoomID,
		// 			Rssi:   int64(targetUserRssi),
		// 		})
		// 		if err != nil {
		// 			c.String(http.StatusInternalServerError, "Server Error")
		// 			return
		// 		}
		// 	}
		// }

		//以前いた部屋のデータに存在しない場合 {Beacons:[] ,RoomID:1}
		if !isExist && pastStayer.RoomID == beaconRoom.RoomID {
			fmt.Println("以前いた部屋のデータに存在しない場合")

			//Stayerテーブルから削除する
			err := RoomService.DeleteStayer(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			//logテーブルのendAtを更新する
			err = RoomService.InsertEndAt(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	for _, currentStayer := range beaconRoom.Beacons {

		currentUserID, err := UserService.GetUserIDByUUID(currentStayer.Uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		currentTime := time.Now()
		//もし火曜日だったら
		if currentTime.Weekday() == time.Tuesday {
			//8時から12時の時
			if currentTime.Hour() >= 8 && currentTime.Hour() < 12 {
				UserService.TemporarilySavedAttendance(currentUserID, 1)
			}
		}

		//stayerテーブルを検索して該当ユーザがいるか確認する
		err, stayerFlag := RoomService.GetStayer(&model.Stayer{UserID: currentUserID})
		if err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}
		//該当ユーザがいない場合はstayerテーブルとlogテーブルに新規に追加する
		if !stayerFlag {
			err = RoomService.SetStayer(&model.Stayer{UserID: currentUserID, RoomID: beaconRoom.RoomID, Rssi: currentStayer.Rssi})
			if err != nil {
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}
			currentTime := time.Now()

			err = RoomService.SetLog(&model.Log{RoomID: beaconRoom.RoomID, StartAt: currentTime.Format("2006-01-02 15:04:05"), EndAt: "2016-01-01 00:00:00", UserID: currentUserID, Rssi: currentStayer.Rssi})
			if err != nil {
				log.Fatal(err)
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}

		}
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}
