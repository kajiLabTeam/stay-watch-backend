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

func Stayer(c *gin.Context) {

	RoomService := service.RoomService{}
	//Stayerテーブルから全てのデータを取得する
	allStayer, err := RoomService.GetAllStayer()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	fmt.Println("allStayer=", allStayer)
	fmt.Println("allStayer=", allStayer[0].UserID)

	stayerGetResponse := []model.StayerGetResponse{}

	for _, stayer := range allStayer {

		userName, err := RoomService.GetUserName(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomName(stayer.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		teamName, err := RoomService.GetUserTeam(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		stayerGetResponse = append(stayerGetResponse, model.StayerGetResponse{
			ID:   stayer.UserID,
			Name: userName,
			Team: teamName,
			Room: roomName,
		})
	}
	c.JSON(200, stayerGetResponse)
}

func Log(c *gin.Context) {
	RoomService := service.RoomService{}
	//Logテーブルから全てのデータを取得する
	allLog, err := RoomService.GetAllLog()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	logGetResponse := []model.LogGetResponse{}

	for _, log := range allLog {

		userName, err := RoomService.GetUserName(log.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomName(log.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		teamName, err := RoomService.GetUserTeam(log.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		logGetResponse = append(logGetResponse, model.LogGetResponse{
			ID:      log.UserID,
			Name:    userName,
			Room:    roomName,
			StartAt: log.StartAt,
			EndAt:   log.EndAt,
			Team:    teamName,
		})
	}
	c.JSON(200, logGetResponse)
}

//ビーコン情報を受け取る
func Beacon(c *gin.Context) {

	beaconRoom := model.BeaconRoom{}
	// fmt.Println("c=", c.Bind(&beaconRoom))
	err := c.Bind(&beaconRoom)
	if err != nil {
		fmt.Println("err=", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	RoomService := service.RoomService{}
	//事前にStayerテーブルのデータを取得する
	pastAllStayer, err := RoomService.GetAllStayer()

	for _, pastStayer := range pastAllStayer {

		isExist := false
		targetUserRssi := -200

		//前のStayerが現在も同じ部屋にいるか線形探索で確認
		for _, currentStayer := range beaconRoom.Beacons {

			// if pastStayer.cu == currentStayer.Rssi
			//現在も同じ部屋にいる場合
			if pastStayer.UserID == currentStayer.Uuid {
				targetUserRssi = int(currentStayer.Rssi)
				isExist = true
			}

		}

		fmt.Println(targetUserRssi, int(pastStayer.Rssi))

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

		//RSSIが以前より弱い場合
		if isExist && targetUserRssi < int(pastStayer.Rssi) {
			//同じ部屋にいる場合は更新
			if beaconRoom.RoomID == pastStayer.RoomID {
				fmt.Println("同じ部屋にいる場合はRSSIを更新 弱くなる")
				err := RoomService.UpdateStayer(&model.Stayer{
					UserID: pastStayer.UserID,
					RoomID: pastStayer.RoomID,
					Rssi:   int64(targetUserRssi),
				})
				if err != nil {
					c.String(http.StatusInternalServerError, "Server Error")
					return
				}
			}
		}

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

		//stayerテーブルを検索して該当ユーザがいるか確認する
		err, stayerFlag := RoomService.GetStayer(&model.Stayer{UserID: currentStayer.Uuid})
		if err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}
		//該当ユーザがいない場合はstayerテーブルとlogテーブルに新規に追加する
		if !stayerFlag {
			err = RoomService.SetStayer(&model.Stayer{UserID: currentStayer.Uuid, RoomID: beaconRoom.RoomID, Rssi: currentStayer.Rssi})
			if err != nil {
				c.String(http.StatusBadRequest, "Bad Request")
				return
			}
			currentTime := time.Now()
			err = RoomService.SetLog(&model.Log{RoomID: beaconRoom.RoomID, StartAt: currentTime.Format("2006-01-02 15:04:05"), EndAt: "2016-01-01 00:00:00", UserID: currentStayer.Uuid, Rssi: currentStayer.Rssi})
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
