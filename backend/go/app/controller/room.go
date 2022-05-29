package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
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

	stayerGetResponse := []model.StayerGetResponse{}

	for _, stayer := range allStayer {

		userName, err := UserService.GetUserName(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomName(stayer.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		tags := make([]model.Tag, 0)

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
			tag := model.Tag{
				ID:   tagID,
				Name: tagName,
			}
			tags = append(tags, tag)
		}

		stayerGetResponse = append(stayerGetResponse, model.StayerGetResponse{
			ID:     stayer.UserID,
			Name:   userName,
			Room:   roomName,
			RoomID: int(stayer.RoomID),
			Tags:   tags,
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

		userName, err := UserService.GetUserName(log.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomName(log.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		logGetResponse = append(logGetResponse, model.LogGetResponse{
			ID:      log.ID,
			Name:    userName,
			Room:    roomName,
			StartAt: log.StartAt,
			EndAt:   log.EndAt,
		})
	}
	c.JSON(200, logGetResponse)
}

func SimultaneousList(c *gin.Context) {
	userID := c.Param("user_id")

	RoomService := service.RoomService{}

	//userIDをint64に変換
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	SimultaneousList, err := RoomService.GetSimultaneousList(userIDInt)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, SimultaneousList)
}
