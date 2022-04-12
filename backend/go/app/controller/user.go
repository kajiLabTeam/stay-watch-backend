package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"net/http"

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

func List(c *gin.Context) {
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
			Team: user.Team,
			Tags: tags,
		})
	}

	c.JSON(200, userInformationGetResponse)
}
