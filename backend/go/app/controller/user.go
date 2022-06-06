package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func Register(c *gin.Context) {

	user := model.User{}
	c.BindJSON(&user)

	userService := service.UserService{}

	err := userService.RegisterUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func UserList(c *gin.Context) {

	UserService := service.UserService{}
	users, err := UserService.GetAllUser()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	userInformationGetResponse := []model.UserInformationGetResponse{}

	for _, user := range users {

		tags := make([]model.Tag, 0)

		tagsID, err := UserService.GetUserTagsID(user.ID)
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

		userInformationGetResponse = append(userInformationGetResponse, model.UserInformationGetResponse{
			ID:   user.ID,
			Name: user.Name,
			Tags: tags,
		})
	}
	c.JSON(200, userInformationGetResponse)
}

func Attendance(c *gin.Context) {

	//構造体定義
	type Meeting struct {
		ID string `json:"meetingID"`
	}
	var meeting Meeting
	c.Bind(&meeting)

	fmt.Println(meeting.ID)
	UserService := service.UserService{}
	//attendaance_tmpテーブルから全てのデータを取得する
	allAttendancesTmp, err := UserService.GetAllAttendancesTmp()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	ExcelService := service.ExcelService{}

	ExcelService.WriteExcel(allAttendancesTmp, meeting.ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func SimultaneousStayUserList(c *gin.Context) {
	userID := c.Param("user_id")
	//int64に変換
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	UserService := service.UserService{}
	RoomService := service.RoomService{}

	logs, err := RoomService.GetLogByUserAndDate(userIDInt64, 14)
	simultaneousStayUserGetResponses, err := UserService.GetSameTimeUser(logs)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(200, simultaneousStayUserGetResponses)
}
