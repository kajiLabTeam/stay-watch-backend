package controller

import (

	// "strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func BackUpDB(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
    "message": "Success",
	})
}