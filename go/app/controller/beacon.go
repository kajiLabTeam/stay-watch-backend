package controller

import (
	"fmt"
	"net/http"
	"time"

	"Stay_watch/model"
	"Stay_watch/service"

	"github.com/gin-gonic/gin"
)

// ビーコン情報を受け取る
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
	BotService := service.BotService{}
	// 事前にStayerテーブルのデータを取得する
	pastAllStayer, err := RoomService.GetAllStayer()
	if err != nil {
		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	for _, pastStayer := range pastAllStayer {

		isExist := false
		targetUserRssi := -200
		for _, currentStayer := range beaconRoom.Beacons {
			pastUUID, err := UserService.GetUserUUIDByUserID(pastStayer.UserID)
			if err != nil {
				fmt.Printf("failed: Cannnot get user uuid %v", err)
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			// 1つ前のstgayerテーブルにもいた場合
			if pastUUID == currentStayer.Uuid {
				targetUserRssi = int(currentStayer.Rssi)
				isExist = true
			}
		}

		// RSSIが以前より強い場合
		if isExist && targetUserRssi > int(pastStayer.Rssi) {
			// 同じ部屋にいる場合は更新
			if beaconRoom.RoomID == pastStayer.RoomID {
				fmt.Println("同じ部屋にいる場合は更新 RSSI強くなる")
				err := RoomService.UpdateStayer(&model.Stayer{
					UserID: pastStayer.UserID,
					RoomID: pastStayer.RoomID,
					Rssi:   int64(targetUserRssi),
				})
				if err != nil {
					fmt.Printf("failed: Cannnot update stayer %v", err)
					c.String(http.StatusInternalServerError, "Server Error")
					return
				}
			} else {
				fmt.Println("別の部屋にいる場合は削除")
				// 別の部屋の場合Stayerテーブルから削除する
				err := RoomService.DeleteStayer(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot delete stayer")
					c.String(http.StatusInternalServerError, "Server Error")
				}
				// logテーブルのendAtを更新する
				err = RoomService.InsertEndAt(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot update endAt")
					c.String(http.StatusInternalServerError, "Server Error")
				}

				pastStayerUserName, err := UserService.GetUserNameByUserID(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot get user name")
					c.String(http.StatusInternalServerError, "Server Error")
					return
				}
				pastRoomName, err := RoomService.GetRoomNameByRoomID(pastStayer.RoomID)
				if err != nil {
					fmt.Println("failed: Cannnot get room name")
					c.String(http.StatusInternalServerError, "Server Error")
					return
				}
				err = BotService.SendMessage(fmt.Sprintf("%sさんが%sから退室しました", pastStayerUserName, pastRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
				if err != nil {
					fmt.Println("failed: Cannnot send message")
					c.String(http.StatusInternalServerError, "Server Error")
					return
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

		// 以前いた部屋のデータに存在しない場合 {Beacons:[] ,RoomID:1}
		if !isExist && pastStayer.RoomID == beaconRoom.RoomID {
			fmt.Println("以前いた部屋のデータに存在しない場合")

			// Stayerテーブルから削除する
			err := RoomService.DeleteStayer(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			// logテーブルのendAtを更新する
			err = RoomService.InsertEndAt(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			pastStayerUserName, err := UserService.GetUserNameByUserID(pastStayer.UserID)
			if err != nil {
				fmt.Println("failed: Cannnot get user name")
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			pastRoomName, err := RoomService.GetRoomNameByRoomID(pastStayer.RoomID)
			if err != nil {
				fmt.Println("failed: Cannnot get room name")
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			err = BotService.SendMessage(fmt.Sprintf("%sさんが%sから退室しました", pastStayerUserName, pastRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
			if err != nil {
				fmt.Println("failed: Cannnot send message")
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
		}

	}

	for _, currentStayer := range beaconRoom.Beacons {

		currentUserID, err := UserService.GetUserIDByUUID(currentStayer.Uuid)
		if err != nil {
			fmt.Println("failed: Cannnot get user id")
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		currentTime := time.Now()
		// もし火曜日だったら
		if currentTime.Weekday() == time.Tuesday {
			// 8時から12時の時
			if currentTime.Hour() >= 7 && currentTime.Hour() < 12 {
				UserService.TemporarilySavedAttendance(currentUserID, 1)
			}
		}

		// stayerテーブルを検索して該当ユーザがいるか確認する
		err, stayerFlag := RoomService.GetStayer(currentUserID)
		if err != nil {
			fmt.Println("failed: Cannnot get stayer")
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}
		// 該当ユーザがいない場合はstayerテーブルとlogテーブルに新規に追加する
		if !stayerFlag {
			err = RoomService.SetStayer(&model.Stayer{UserID: currentUserID, RoomID: beaconRoom.RoomID, Rssi: currentStayer.Rssi})
			if err != nil {
				fmt.Println("failed: Cannnot set stayer")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}
			currentTime := time.Now()
			endAt, err := time.Parse("2006-01-02 15:04:05", "2016-01-01 00:00:00")
			if err != nil {
				fmt.Println("failed: Cannnot parse time")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}

			err = RoomService.SetLog(&model.Log{RoomID: beaconRoom.RoomID, StartAt: currentTime, EndAt: endAt, UserID: currentUserID, Rssi: currentStayer.Rssi})
			if err != nil {
				fmt.Println("failed: Cannnot set log")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}

			currentUserName, err := UserService.GetUserNameByUserID(currentUserID)
			if err != nil {
				fmt.Println("failed: Cannnot get user name")
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}

			currentRoomName, err := RoomService.GetRoomNameByRoomID(beaconRoom.RoomID)
			if err != nil {
				fmt.Println("failed: Cannnot get room name")
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}

			err = BotService.SendMessage(fmt.Sprintf("%sさんが%sに入室しました", currentUserName, currentRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
			if err != nil {
				fmt.Println("failed: Cannnot send message")
				c.String(http.StatusInternalServerError, "Server Error")
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
