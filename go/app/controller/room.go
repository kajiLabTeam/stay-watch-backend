package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Stayer(c *gin.Context) {

	RoomService := service.RoomService{}
	UserService := service.UserService{}

	//Stayerテーブルから全てのデータを取得する
	allStayer, err := RoomService.GetAllStayer()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	fmt.Println("ok")

	stayerGetResponse := []model.StayerGetResponse{}

	for _, stayer := range allStayer {

		userName, err := UserService.GetUserNameByUserID(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomNameByRoomID(stayer.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		tagsGetResponse := make([]model.TagGetResponse, 0)

		tagsID, err := UserService.GetUserTagsID(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			tag := model.TagGetResponse{
				ID:   tagID,
				Name: tagName,
			}
			tagsGetResponse = append(tagsGetResponse, tag)
		}

		stayerGetResponse = append(stayerGetResponse, model.StayerGetResponse{
			ID:     stayer.UserID,
			Name:   userName,
			Room:   roomName,
			RoomID: int(stayer.RoomID),
			Tags:   tagsGetResponse,
		})
	}
	c.JSON(200, stayerGetResponse)
}

func Log(c *gin.Context) {
	RoomService := service.RoomService{}
	UserService := service.UserService{}

	//ページング処理
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	//ページごとにLogテーブルからデータを取得する
	pageLog, err := RoomService.GetLogsByPage(pageInt)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	logGetResponse := []model.LogGetResponse{}

	for _, log := range pageLog {

		userName, err := UserService.GetUserNameByUserID(log.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomNameByRoomID(log.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		logGetResponse = append(logGetResponse, model.LogGetResponse{
			ID:      int64(log.ID),
			Name:    userName,
			Room:    roomName,
			StartAt: log.StartAt.Format("2006-01-02"),
			EndAt:   log.EndAt.Format("2006-01-02"),
		})
	}
	c.JSON(200, logGetResponse)
}

// func SimultaneousList(c *gin.Context) {
// 	userID := c.Param("user_id")

// 	RoomService := service.RoomService{}

// 	//userIDをint64に変換
// 	userIDInt, err := strconv.ParseInt(userID, 10, 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Bad Request")
// 		return
// 	}

// 	SimultaneousList, err := RoomService.GetSimultaneousList(userIDInt)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}

// 	c.JSON(http.StatusOK, SimultaneousList)
// }

func LogGantt(c *gin.Context) {

	RoomService := service.RoomService{}
	GanttLogs, err := RoomService.GetGanttLog()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusOK, GanttLogs)
}
