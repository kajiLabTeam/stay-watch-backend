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
