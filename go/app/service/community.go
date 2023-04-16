package service

import (
	"Stay_watch/model"
	// "Stay_watch/util"
	"fmt"
	// "log"
	// "time"
)

type CommunityService struct{}

func (CommunityService) GetCommunityById(communityId int64) (model.Community, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return model.Community{}, err
	}
	defer closer.Close()
	community := model.Community{}
	// communitiesテーブルからcommunityIdと合致する情報を取得する
	result := DbEngine.Where("id=?", communityId).Take(&community)
	if result.Error != nil {
		fmt.Printf("コミュニティ取得失敗 %v", result.Error)
		return model.Community{}, result.Error
	}
	return community, nil
}
