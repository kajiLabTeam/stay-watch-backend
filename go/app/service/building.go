package service

import (
	"Stay_watch/model"
	// "Stay_watch/util"
	"fmt"
	// "log"
	// "time"
)

type BuildingService struct{}

func (BuildingService) GetAllBuildings() ([]model.Building, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	buildings := make([]model.Building, 0)
	result := DbEngine.Table("buildings").Find(&buildings)
	if result.Error != nil {
		return nil, fmt.Errorf(" failed to get all stayer: %w", result.Error)
	}

	return buildings, nil
}