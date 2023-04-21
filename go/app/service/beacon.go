package service

import (
	"Stay_watch/model"
	"fmt"
)

type BeaconService struct{}

// 全てのビーコンを取得する
func (BeaconService) GetAllBeaconType() ([]model.BeaconType, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	beaconTypes := make([]model.BeaconType, 0)
	result := DbEngine.Find(&beaconTypes)
	if result.Error != nil {
		fmt.Printf("ビーコン取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return beaconTypes, nil
}

// ビーコンIDからビーコン情報を取得する
func (BeaconService) GetBeaconByBeaconId(beaconId int64) (model.BeaconType, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return model.BeaconType{}, err
	}
	defer closer.Close()
	beacon := model.BeaconType{}
	result := DbEngine.Where("id=?", beaconId).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return model.BeaconType{}, result.Error
	}

	return beacon, nil
}

func (BeaconService) GetBeaconTypeByBeaconName(beaconName string) (model.BeaconType, error) {
	DbEngine := connect()
	close, err := DbEngine.DB()
	if err != nil {
		return model.BeaconType{}, err
	}
	defer close.Close()
	beacon := model.BeaconType{}
	result := DbEngine.Where("name=?", beaconName).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ビーコン情報取得失敗 %v", result.Error)
		return model.BeaconType{}, result.Error
	}

	return beacon, nil
}

func (BeaconService) GetBeaconTypeIdByBeaconName(beaconName string) (int64, error) {
	DbEngine := connect()
	close, err := DbEngine.DB()
	if err != nil {
		return 0, err
	}
	defer close.Close()
	beacon := model.BeaconType{}
	result := DbEngine.Where("name=?", beaconName).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ユーザ名取得失敗 %v", result.Error)
		return 0, result.Error
	}

	return int64(beacon.ID), nil
}
