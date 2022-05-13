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

//IDから名前を取得する
func (UserService) GetUserName(userID string) (string, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.Name, nil
}

//IDからチーム名を取得する
func (UserService) GetUserTeam(userID string) (string, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.Team, nil
}

//IDからタグ(複数形)IDを取得する
func (UserService) GetUserTagsID(userID string) ([]int64, error) {
	tags := make([]int64, 0)
	err := DbEngine.Table("tag_map").Where("user_id=?", userID).Cols("tag_id").Find(&tags)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return tags, nil
}

//タグIDからタグ名を取得する
func (UserService) GetTagName(tagID int64) (string, error) {
	tag := model.Tag{}
	_, err := DbEngine.Table("tag").Where("id=?", tagID).Get(&tag)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return tag.Name, nil
}

//attendanceテーブルに登録する
func (UserService) RegisterAttendance(userID string, date string, exit bool) error {
	attendance := model.Attendance{
		UserID: userID,
		Date:   date,
		Exit:   exit,
	}
	_, err := DbEngine.Table("attendance").Insert(&attendance)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
