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
	userService := service.UserService{}
	users, err := userService.GetAllUser()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(200, users)
}
