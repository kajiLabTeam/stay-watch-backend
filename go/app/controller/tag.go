package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
)

func GetTagNamesByCommunityId(c *gin.Context) {
	communityId, _ := strconv.ParseInt(c.Param("communityId"), 10, 64) // string -> int64
	TagService := service.TagService{}

	// DBからどこのコミュニティにも該当するタグネームを持ってくる
	publicTagNames, err := TagService.GetTagNamesByCommunityId(-1)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	// DBからコミュニティのタグネームを持ってくる
	communityTagNames, err := TagService.GetTagNamesByCommunityId(communityId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	tagNames := append(publicTagNames, communityTagNames...)

	tagNamesResponse := []model.TagsNamesGetResponse{}

	for _, tagName := range tagNames {

		tagNamesResponse = append(tagNamesResponse, model.TagsNamesGetResponse{
			Name: tagName,
		})
	}

	c.JSON(200, tagNamesResponse)
}
