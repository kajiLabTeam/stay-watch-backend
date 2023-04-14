package service

import (
	"Stay_watch/model"
	"fmt"
)

type DeletedUserService struct{}

// 新規登録
func (DeletedUserService) CreateDeletedUser(user *model.DeletedUser) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Create(user)
	if result.Error != nil {
		fmt.Printf("削除されるユーザ登録処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}
