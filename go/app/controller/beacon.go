package controller

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Stay_watch/model"
	"Stay_watch/presence"
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

// ランダム値とハッシュ値からDBのユーザテーブル内のキーでフルテーブルスキャンする
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
	return 0, errors.New("no matching user found")
}

// ラズパイ受信機から送られてきた電波情報からユーザを特定
func convertBeaconsStayers(inputBeacons []*model.BeaconSignal) []model.Stayer {
	UserService := service.UserService{}

	outStayers := []model.Stayer{}
	for _, inputBeacon := range inputBeacons {
		// MSDを用いる場合とUUIDを用いる場合で分ける
		// MSDを用いる場合（現在はPrivBeaconのみ）
		if inputBeacon.Msd != "" {
			// PrivBeaconはMSDの頭4文字が固定値かつMSDの長さが20文字より大きい
			if strings.HasPrefix(inputBeacon.Msd, PRIVBEACON_MSD_PREFIX) && len(inputBeacon.Msd) > 20 {
				// PrivBeaconの場合MSDからランダム値部分とハッシュ値部分を取り出して、DB内のPrivBeaconユーザを総当たりでハッシュ化しユーザ特定する
				hashValue := inputBeacon.Msd[4:20]
				randomValue := inputBeacon.Msd[20:]
				privBeaconUserId, err := getUserIdBySipHash(randomValue, hashValue)
				if err == nil {
					// 見つかったら滞在者リストに追加して次のforへ
					outStayers = append(outStayers, model.Stayer{
						UserID: privBeaconUserId,
						Rssi:   inputBeacon.Rssi,
					})
					continue
				}
			}
		}
		// UUIDを用いる場合（iPhoneBeaconやAndroidビーコン、市販ビーコンなど）
		if inputBeacon.Uuid != "" {
			// iPhoneビーコンは送られてくるUUID情報を含む値の長さが38文字
			if len(inputBeacon.Uuid) == 38 {
				// iPhoneビーコンの場合UUIDを取得する処理が必要(例："4c000180000021000021000021000021000021" -> "8ebc21144abd00000000ff0100000021")
				tmpUUID := getEndUUIDByManufacturer(inputBeacon.Uuid)
				uuid := "8ebc21144abd00000000ff01000" + tmpUUID
				iphoneUserId, err := UserService.GetUserIDByUUID(uuid)
				if err == nil {
					// 見つかったら滞在者リストに追加して次のforへ
					outStayers = append(outStayers, model.Stayer{
						UserID: iphoneUserId,
						Rssi:   inputBeacon.Rssi,
					})
					continue
				}
			}
			// Androidビーコンや市販ビーコンの場合は単純にUUIDからユーザを検索
			commonBeaconUserId, err := UserService.GetUserIDByUUID(inputBeacon.Uuid)
			if err == nil {
				// 見つかったら滞在者リストに追加して次のforへ
				outStayers = append(outStayers, model.Stayer{
					UserID: commonBeaconUserId,
					Rssi:   inputBeacon.Rssi,
				})
				continue
			}
		}
	}
	return outStayers
}

// 同じユーザが複数件含まれる場合（PrivBeaconのMSDが毎回変わる、複数ビーコン所持など）に
// 1ユーザ1件へまとめる。RSSIは最も強い値を採用する。
func mergeStayersByUser(stayers []model.Stayer) []model.Stayer {
	merged := []model.Stayer{}
	index := map[int64]int{}
	for _, s := range stayers {
		if i, ok := index[s.UserID]; ok {
			if s.Rssi > merged[i].Rssi {
				merged[i].Rssi = s.Rssi
			}
			continue
		}
		index[s.UserID] = len(merged)
		merged = append(merged, s)
	}
	return merged
}

// 在室中のユーザを退室させる（Stayer削除・ログのendAt更新・Slack通知）
func exitStayer(pastStayer model.Stayer) error {
	RoomService := service.RoomService{}
	UserService := service.UserService{}
	BotService := service.BotService{}

	if err := RoomService.DeleteStayer(pastStayer.UserID); err != nil {
		return fmt.Errorf("cannot delete stayer: %w", err)
	}
	if err := RoomService.UpdateEndAt(pastStayer.UserID, time.Now()); err != nil {
		return fmt.Errorf("cannot update endAt: %w", err)
	}
	userName, err := UserService.GetUserNameByUserID(pastStayer.UserID)
	if err != nil {
		return fmt.Errorf("cannot get user name: %w", err)
	}
	roomName, err := RoomService.GetRoomNameByRoomID(pastStayer.RoomID)
	if err != nil {
		return fmt.Errorf("cannot get room name: %w", err)
	}
	if err := BotService.SendMessage(fmt.Sprintf("%sさんが%sから退室しました", userName, roomName), "B03J95EL3ME/9MLCZ8VTkEFGDVwTxkqYLKyj"); err != nil {
		return fmt.Errorf("cannot send message: %w", err)
	}
	return nil
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

	requestStayers := mergeStayersByUser(convertBeaconsStayers(beaconRoom.Beacons))
	requestRoomId := beaconRoom.RoomID

	RoomService := service.RoomService{}
	UserService := service.UserService{}
	BotService := service.BotService{}
	Util := util.Util{}

	requestStayerByUser := map[int64]model.Stayer{}
	for _, s := range requestStayers {
		requestStayerByUser[s.UserID] = s
	}

	// 事前にStayerテーブルのデータを全て取得する
	pastAllStayer, err := RoomService.GetAllStayer()
	if err != nil {
		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stayer"})
		return
	}

	now := time.Now()
	for _, pastStayer := range pastAllStayer {
		currentStayer, isExist := requestStayerByUser[pastStayer.UserID]

		action := presence.Decide(
			presence.Current{RoomID: pastStayer.RoomID, Rssi: pastStayer.Rssi, LastSeen: pastStayer.UpdatedAt},
			presence.Observation{RoomID: requestRoomId, Seen: isExist, Rssi: currentStayer.Rssi},
			now,
		)

		switch action {
		case presence.Update:
			// 同じ部屋で観測できた場合はRSSIと最終観測時刻(updated_at)を更新
			err := RoomService.UpdateStayer(&model.Stayer{
				UserID: pastStayer.UserID,
				RoomID: pastStayer.RoomID,
				Rssi:   currentStayer.Rssi,
			})
			if err != nil {
				fmt.Printf("failed: Cannnot update stayer %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stayer"})
				return
			}
		case presence.Exit:
			// 退室（別部屋の報告による移動の場合は、この後の処理でその部屋に入室する）
			fmt.Printf("exit: user=%d room=%d rssi=%d -> reportRoom=%d rssi=%d seen=%v\n",
				pastStayer.UserID, pastStayer.RoomID, pastStayer.Rssi, requestRoomId, currentStayer.Rssi, isExist)
			if err := exitStayer(pastStayer); err != nil {
				fmt.Println("failed:", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to exit stayer"})
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
