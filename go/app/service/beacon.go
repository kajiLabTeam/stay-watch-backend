package service

import (
	"Stay_watch/model"
	"fmt"
)

type BeaconService struct{}

// 全てのビーコンを取得する
func (BeaconService) GetAllBeacon() ([]model.Beacon, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	beacons := make([]model.Beacon, 0)
	result := DbEngine.Find(&beacons)
	if result.Error != nil {
		fmt.Printf("ビーコン取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return beacons, nil
}

// ビーコンIDからビーコン情報を取得する
func (BeaconService) GetBeaconByBeaconId(beaconId int64) (model.Beacon, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return model.Beacon{}, err
	}
	defer closer.Close()
	beacon := model.Beacon{}
	result := DbEngine.Where("id=?", beaconId).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return model.Beacon{}, result.Error
	}

	return beacon, nil
}

func (BeaconService) GetBeaconByBeaconName(beaconName string) (model.Beacon, error) {
	DbEngine := connect()
	close, err := DbEngine.DB()
	if err != nil {
		return model.Beacon{}, err
	}
	defer close.Close()
	beacon := model.Beacon{}
	result := DbEngine.Where("type=?", beaconName).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ビーコン情報取得失敗 %v", result.Error)
		return model.Beacon{}, result.Error
	}

	return beacon, nil
}

func (BeaconService) GetBeaconIdByBeaconName(beaconName string) (int64, error) {
	DbEngine := connect()
	close, err := DbEngine.DB()

	if err != nil {
		return 0, err
	}
	defer close.Close()
	beacon := model.Beacon{}
	result := DbEngine.Where("name=?", beaconName).Take(&beacon)
	if result.Error != nil {
		fmt.Printf("ユーザ名取得失敗 %v", result.Error)
		return 0, result.Error
	}

	return int64(beacon.ID), nil
}
