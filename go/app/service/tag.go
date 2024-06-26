package service

import (
	"Stay_watch/model"
	"fmt"
)

type TagService struct{}

func (TagService) CreateTagMap(tagMap *model.TagMap) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Create(tagMap)
	if result.Error != nil {
		fmt.Printf("タグ登録処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

func (TagService) DeleteTagMap(userId int64) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Unscoped().Delete(&model.TagMap{}, userId)
	if result.Error != nil {
		fmt.Printf("ユーザ削除処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

func (TagService) GetTagMapIdsByUserId(userId int64) ([]int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tagMapIds := make([]int64, 0)
	result := DbEngine.Table("tag_maps").Where("user_id=?", userId).Select("id").Find(&tagMapIds)
	if result.Error != nil {
		fmt.Printf("タグID取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tagMapIds, nil
}

func (TagService) GetTagsByCommunityId(communityId int64) ([]model.Tag, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tags := make([]model.Tag, 0)

	// tagsテーブルからcommunity_idカラムがcommunityIdのnameの値をtagNamesに格納
	result := DbEngine.Where("community_id = ?", communityId).Find(&tags)
	if result.Error != nil {
		fmt.Printf("タグ一覧取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tags, nil
}

func (TagService) GetTagNamesByCommunityId(communityId int64) ([]string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tagNames := make([]string, 0)

	// tagsテーブルからcommunity_idカラムがcommunityIdのnameの値をtagNamesに格納
	result := DbEngine.Table("tags").Where("community_id=?", communityId).Select("name").Find(&tagNames)
	if result.Error != nil {
		fmt.Printf("タグ名一覧取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tagNames, nil
}
