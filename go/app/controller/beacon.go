package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Stay_watch/model"
	"Stay_watch/service"
	"Stay_watch/util"

	"github.com/gin-gonic/gin"
)

func getEndUUIDByManufacturer(manufacturer string) string {
	slicedManufacturers := []int{}

	// (例："4c000180000021000021000021000022000021" -> [33,33,33,34,33])
	for i := 0; i < 5; i++ {
		slicedManufacturerNum, err := strconv.ParseInt(manufacturer[(5*i)+9+i:(5*i)+14+i], 16, 64)
		if err != nil {
			fmt.Printf("failed: Cannnot convert hexStr to Int %v", err)
			continue
		}
		slicedManufacturers = append(slicedManufacturers, int(slicedManufacturerNum))
	}

	// 配列の一番小さいものがUUIDの末尾
	minNumber := slicedManufacturers[0]
	for _, slicedManufacturer := range slicedManufacturers {
		if slicedManufacturer < minNumber {
			minNumber = slicedManufacturer
		}
	}

	// 16進数整数から10進数5文字列へ(例：33 -> "00021")
	resultManufacturer := fmt.Sprintf("%05x", minNumber)
	return resultManufacturer
}

func convertBeacons(inputBeacons []*model.BeaconSignal) []model.BeaconSignal {

	//BeaconService := service.BeaconService{}

	outBeacons := []model.BeaconSignal{}
	for _, inputBeacon := range inputBeacons {

		uuid := inputBeacon.Uuid

		// iPhoneビーコンの場合UUIDを取得する処理が必要(例："4c000180000021000021000021000021000021" -> "8ebc21144abd00000000ff0100000021")
		if len(inputBeacon.Uuid) == 38 {
			// 6文字目が8以上なら発信中ってことを意味する
			// 16進数の文字列を10進数の数字に変換
			advertisingStatusHexStr := inputBeacon.Uuid[6:7]
			advertisingStatusNum, err := strconv.ParseInt(advertisingStatusHexStr, 16, 64)
			if err != nil {
				fmt.Printf("failed: Cannnot convert hexStr to Int %v", err)
				continue
			}
			if advertisingStatusNum >= 8 {
				// Manufacturerの5文字をUUID(8ebc21144abd00000000ff01000は固定値)の末尾へ追加
				tmpUUID := getEndUUIDByManufacturer(inputBeacon.Uuid)
				uuid = "8ebc21144abd00000000ff01000" + tmpUUID
			}
		}

		tmpBeacon := model.BeaconSignal{
			Uuid: uuid,
			Rssi: inputBeacon.Rssi,
		}
		outBeacons = append(outBeacons, tmpBeacon)
	}

	return outBeacons
}

// ビーコン情報を受け取る
func Beacon(c *gin.Context) {

	beaconRoom := model.BeaconRoom{}
	err := c.Bind(&beaconRoom)
	if err != nil {
		fmt.Println("err=", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	requestBeacons := convertBeacons(beaconRoom.Beacons)
	requestRoomId := beaconRoom.RoomID

	RoomService := service.RoomService{}
	UserService := service.UserService{}
	BotService := service.BotService{}
	Util := util.Util{}

	// 事前にStayerテーブルのデータを全て取得する
	pastAllStayer, err := RoomService.GetAllStayer()

	if err != nil {
		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stayer"})
		return
	}

	for _, pastStayer := range pastAllStayer {

		isExist := false
		targetUserRssi := -200

		// リクエストからの滞在者リスト(beaconRoom.Beacons)とStayerテーブルの滞在者リストを比較
		for _, currentStayer := range requestBeacons {
			pastUUID, err := UserService.GetUserUUIDByUserID(pastStayer.UserID)
			if err != nil {
				fmt.Printf("Cannnot get user uuid %v", err)
				continue
			}
			// 1つ前のstayerテーブルにもいた場合
			if pastUUID == currentStayer.Uuid {
				targetUserRssi = int(currentStayer.Rssi)
				isExist = true
			}
		}

		// RSSIが以前より強い場合
		if isExist && targetUserRssi > int(pastStayer.Rssi) {
			// 同じ部屋にいる場合は更新
			if requestRoomId == pastStayer.RoomID {
				fmt.Println("同じ部屋にいる場合は更新 RSSI強くなる")
				err := RoomService.UpdateStayer(&model.Stayer{
					UserID: pastStayer.UserID,
					RoomID: pastStayer.RoomID,
					Rssi:   int64(targetUserRssi),
				})
				if err != nil {
					fmt.Printf("failed: Cannnot update stayer %v", err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stayer"})
					return
				}
			} else {
				fmt.Println("別の部屋にいる場合は削除")
				// 別の部屋の場合Stayerテーブルから削除する
				err := RoomService.DeleteStayer(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot delete stayer")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stayer"})
				}
				// logテーブルのendAtを更新する
				err = RoomService.UpdateEndAt(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot update endAt")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update endAt"})
				}

				pastStayerUserName, err := UserService.GetUserNameByUserID(pastStayer.UserID)
				if err != nil {
					fmt.Println("failed: Cannnot get user name")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user name"})
					return
				}
				pastRoomName, err := RoomService.GetRoomNameByRoomID(pastStayer.RoomID)
				if err != nil {
					fmt.Println("failed: Cannnot get room name")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get room name"})
					return
				}
				err = BotService.SendMessage(fmt.Sprintf("%sさんが%sから退室しました", pastStayerUserName, pastRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
				if err != nil {
					fmt.Println("failed: Cannnot send message")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
					return
				}
			}
		}

		// //RSSIが以前より弱い場合
		// if isExist && targetUserRssi < int(pastStayer.Rssi) {
		// 	//同じ部屋にいる場合は更新
		// 	if requestRoomId == pastStayer.RoomID {
		// 		fmt.Println("同じ部屋にいる場合はRSSIを更新 弱くなる")
		// 		err := RoomService.UpdateStayer(&model.Stayer{
		// 			UserID: pastStayer.UserID,
		// 			RoomID: pastStayer.RoomID,
		// 			Rssi:   int64(targetUserRssi),
		// 		})
		// 		if err != nil {
		// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to"})
		// 			return
		// 		}
		// 	}
		// }

		// 以前いた部屋のデータに存在しない場合 {Beacons:[] ,RoomID:1}
		if !isExist && pastStayer.RoomID == requestRoomId {
			fmt.Println("以前いた部屋のデータに存在しない場合")

			// Stayerテーブルから削除する
			err := RoomService.DeleteStayer(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			// logテーブルのendAtを更新する
			err = RoomService.UpdateEndAt(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			pastStayerUserName, err := UserService.GetUserNameByUserID(pastStayer.UserID)
			if err != nil {
				fmt.Println("failed: Cannnot get user name")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user name"})
				return
			}
			pastRoomName, err := RoomService.GetRoomNameByRoomID(pastStayer.RoomID)
			if err != nil {
				fmt.Println("failed: Cannnot get room name")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get room name"})
				return
			}
			err = BotService.SendMessage(fmt.Sprintf("%sさんが%sから退室しました", pastStayerUserName, pastRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
			if err != nil {
				fmt.Println("failed: Cannnot send message")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
				return
			}
		}

	}

	for _, currentStayer := range requestBeacons {

		// APIのリクエストのUUIDからuserIdを取得する
		currentUserID, err := UserService.GetUserIDByUUID(currentStayer.Uuid)
		if err != nil {
			fmt.Println("failed: Cannnot get user id(UUID : " + currentStayer.Uuid + ")")
			continue
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
			err = RoomService.CreateStayer(&model.Stayer{UserID: currentUserID, RoomID: requestRoomId, Rssi: currentStayer.Rssi})
			if err != nil {
				fmt.Println("failed: Cannnot set stayer")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}
			currentTime := time.Now()
			endAt, err := Util.ConvertDatetimeToLocationTime("2016-01-01 00:00:00", "Asia/Tokyo")
			fmt.Println(endAt)
			if err != nil {
				fmt.Println("failed: Cannnot parse time")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}

			err = RoomService.CreateLog(&model.Log{RoomID: requestRoomId, StartAt: currentTime, EndAt: endAt, UserID: currentUserID, Rssi: currentStayer.Rssi})
			if err != nil {
				fmt.Println("failed: Cannnot set log")
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}

			currentUserName, err := UserService.GetUserNameByUserID(currentUserID)
			if err != nil {
				fmt.Println("failed: Cannnot get user name")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user name"})
				return
			}

			currentRoomName, err := RoomService.GetRoomNameByRoomID(requestRoomId)
			if err != nil {
				fmt.Println("failed: Cannnot get room name")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get room name"})
				return
			}

			err = BotService.SendMessage(fmt.Sprintf("%sさんが%sに入室しました", currentUserName, currentRoomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj")
			if err != nil {
				fmt.Println("failed: Cannnot send message")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
				return
			}
		}
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}

func GetBeacon(c *gin.Context) {

	BeaconService := service.BeaconService{}
	beacons, err := BeaconService.GetAllBeacon()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get beacons"})
		return
	}

	beaconGetResponses := []model.BeaconGetResponse{}

	for _, beacon := range beacons {

		beaconGetResponses = append(beaconGetResponses, model.BeaconGetResponse{
			BeaconId:     int64(beacon.ID),
			BeaconName:   beacon.Type,
			UuidEditable: beacon.UuidEditable,
		})
	}

	c.JSON(http.StatusOK, beaconGetResponses)
}
