package controller

import (

	// "strconv"

	"Stay_watch/model"
	"Stay_watch/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommunityByUserId(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64) // string -> int64

	CommunityService := service.CommunityService{}
	UserService := service.UserService{}

	// ユーザのコミュニティIDを取得する
	communityId, err := UserService.GetCommunityIdByUserId(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get communityId"})
		return
	}

	// コミュニティIDからコミュニティ情報を取得する
	community, err := CommunityService.GetCommunityById(communityId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get communigy"})
		return
	}

	// レスポンスの型に変換する
	communityGetResponses := model.CommunityGetResponse{
		CommunityId: int64(community.ID),
		Name:        community.Name,
	}

	c.JSON(http.StatusOK, communityGetResponses)

}
