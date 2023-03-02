package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

func GetBuildingsEditor(c *gin.Context) {
	BuildingService := service.BuildingService{}
	buildings, err := BuildingService.GetAllBuildings()
	if err != nil {
		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	buildingsEditorGetResponse := []model.BuildingsEditorGetResponse{}
	
	for _, building := range buildings {
		buildingId := int64(building.ID)
		buildingName := building.Name
		buildingImagePath := building.MapFile

		buildingsEditorGetResponse = append(buildingsEditorGetResponse, model.BuildingsEditorGetResponse{
			BuildingID: buildingId,
			Name: buildingName,
			MapImagePath: buildingImagePath,
		})
	}

	c.JSON(200, buildingsEditorGetResponse)

}