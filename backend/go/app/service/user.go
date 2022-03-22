package service

import (
	"Stay_watch/model"
	"log"
)

type UserService struct{}

//ユーザ登録処理
func (UserService) RegisterUser(user *model.User) error {

	_, err := DbEngine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

//全てのユーザを取得する
func (UserService) GetAllUser() ([]model.User, error) {
	users := make([]model.User, 0)
	err := DbEngine.Find(&users)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return users, nil
}
