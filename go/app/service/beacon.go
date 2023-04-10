package service

import (
	"Stay_watch/model"
	"fmt"
)

type BeaconService struct{}

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
