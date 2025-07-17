package controller

import (
	"net/http"
	"strconv"

	"Stay_watch/constant"
	"Stay_watch/model"
	"Stay_watch/service"

	// "strconv"

	"github.com/gin-gonic/gin"
)

func GetTagsByCommunityIdHandler(c *gin.Context) {
	communityId, err := strconv.ParseInt(c.Param("communityId"), 10, 64) // string -> int64
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Type is not number")
	}

	TagService := service.TagService{}

	// DBからどこのコミュニティにも該当するタグを持ってくる
	publicTags, err := TagService.GetTagsByCommunityId(constant.PublicTagID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public tags"})
		return
	}
	// DBからコミュニティのタグネームを持ってくる
	communityTags, err := TagService.GetTagsByCommunityId(communityId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get private tags"})
		return
	}
	tags := append(publicTags, communityTags...)

	tagsResponse := []model.TagsGetResponse{}

	for _, tag := range tags {
		tagsResponse = append(tagsResponse, model.TagsGetResponse{
			Id:   int64(tag.ID),
			Name: tag.Name,
		})
	}

	c.JSON(http.StatusOK, tagsResponse)
}

func GetTagNamesByCommunityId(c *gin.Context) {
	communityId, err := strconv.ParseInt(c.Param("communityId"), 10, 64) // string -> int64
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Type is not number")
	}

	TagService := service.TagService{}

	// DBからどこのコミュニティにも該当するタグネームを持ってくる
	publicTagNames, err := TagService.GetTagNamesByCommunityId(constant.PublicTagID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public tag name"})
		return
	}
	// DBからコミュニティのタグネームを持ってくる
	communityTagNames, err := TagService.GetTagNamesByCommunityId(communityId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get private tag name"})
		return
	}
	tagNames := append(publicTagNames, communityTagNames...)

	tagNamesResponse := []model.TagsNamesGetResponse{}

	for _, tagName := range tagNames {
		tagNamesResponse = append(tagNamesResponse, model.TagsNamesGetResponse{
			Name: tagName,
		})
	}

	c.JSON(http.StatusOK, tagNamesResponse)
}
