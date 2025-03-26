package controller

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Stay_watch/model"
	"Stay_watch/service"
	"Stay_watch/util"

	"github.com/dchest/siphash"
	"github.com/gin-gonic/gin"
)

const (
	BEACON_ID_STAYWATCHBEACON = 4
	LENGTH_PRIVATE_KEY        = 32
	PRIVBEACON_MSD_PREFIX     = "ffff"
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

func getUserIdBySipHash(randomValue string, hashValue string) (int64, error) {
	UserService := service.UserService{}
	users, err := UserService.GetUsersByBeaconId(BEACON_ID_STAYWATCHBEACON)
	if err != nil {
		fmt.Println("failed to get users by db: ", err)
		return 0, err
	}

	for _, user := range users {
		if len(user.PrivateKey) == LENGTH_PRIVATE_KEY {
			// 例：01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 00
			// 		privateKey := "0102030405060708090a0b0c0d0e0f00"
			privateKey := user.PrivateKey

			// 文字列をバイトスライスに変換
			bytes, err := hex.DecodeString(privateKey)
			if err != nil {
				fmt.Println("failed to decode hex string: ", err)
				return 0, err
			}

			// バイトスライスをuint64に変換(8バイトずつ分割して変換)
			var key1 uint64
			var key2 uint64
			for i := 0; i < 8; i++ {
				key1 |= uint64(bytes[i]) << (8 * i)
				key2 |= uint64(bytes[8+i]) << (8 * i)
			}

			// ハッシュ化したいデータ
			msg := []byte(randomValue) // ランダム値

			// SipHashを計算
			hash := siphash.Hash(key1, key2, msg) // ここで第1, 第2引数にカスタム値を指定可能

			// 結果を出力
			// 結果が一緒になればそのユーザの秘密キーがビーコン固有の秘密キー
			if hashValue == fmt.Sprintf("%016x", hash) {
				return int64(user.ID), nil
			}
		}
	}

	// 見つからなかった場合エラーを返す
	return 0, err
}

func convertBeaconsStayers(inputBeacons []*model.BeaconSignal) []model.Stayer {
	UserService := service.UserService{}

	outStayers := []model.Stayer{}
	for _, inputBeacon := range inputBeacons {

		userId := int64(0)
		if strings.HasPrefix(inputBeacon.Msd, PRIVBEACON_MSD_PREFIX) {
			if len(inputBeacon.Msd) < 20 {
				continue
			}
			// PrivBeaconの場合
			hashValue := inputBeacon.Msd[4:20]
			randomValue := inputBeacon.Msd[20:]
			tmpUserId, err := getUserIdBySipHash(randomValue, hashValue)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			userId = tmpUserId
		} else if len(inputBeacon.Uuid) == 38 {
			// iPhoneビーコンの場合UUIDを取得する処理が必要(例："4c000180000021000021000021000021000021" -> "8ebc21144abd00000000ff0100000021")
			tmpUUID := getEndUUIDByManufacturer(inputBeacon.Uuid)
			uuid := "8ebc21144abd00000000ff01000" + tmpUUID
			tmpUserId, err := UserService.GetUserIDByUUID(uuid)
			if err != nil {
				fmt.Println("Cannot get userid by uuid:", err)
				continue
			}
			userId = tmpUserId
		} else {
			tmpUserId, err := UserService.GetUserIDByUUID(inputBeacon.Uuid)
			userId = tmpUserId
			if err != nil {
				// もし見つからなかった場合基本的に旧滞在ウォッチビーコンである
				randomValue := inputBeacon.Uuid[:16]
				hashValue := inputBeacon.Uuid[16:]
				tmpUserId, err := getUserIdBySipHash(randomValue, hashValue)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				userId = tmpUserId
			}
		}

		// 0でない(ユーザが見つかった場合)のみ追加
		if userId != 0 {
			tmpStayer := model.Stayer{
				UserID: userId,
				Rssi:   inputBeacon.Rssi,
			}
			outStayers = append(outStayers, tmpStayer)
		}
	}

	return outStayers
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

	// requestBeacons := convertBeacons(beaconRoom.Beacons)
	requestStayers := convertBeaconsStayers(beaconRoom.Beacons)
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
		for _, currentStayer := range requestStayers {
			// 現在いる部屋に以前からいた場合
			if pastStayer.UserID == currentStayer.UserID {
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
				err = RoomService.UpdateEndAt(pastStayer.UserID, time.Now())
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

		// 以前いた部屋のデータに現在はいない場合退室処理
		if !isExist && pastStayer.RoomID == requestRoomId {
			fmt.Println("以前いた部屋のデータに存在しない場合")

			// Stayerテーブルから削除する
			err := RoomService.DeleteStayer(pastStayer.UserID)
			if err != nil {
				fmt.Println(err)
			}

			// logテーブルのendAtを更新する
			err = RoomService.UpdateEndAt(pastStayer.UserID, time.Now())
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

	for _, currentStayer := range requestStayers {

		currentUserID := currentStayer.UserID
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
