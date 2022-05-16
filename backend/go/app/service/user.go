package service

import (
	"Stay_watch/model"
	"fmt"
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
func (UserService) GetUserName(userID int64) (string, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.Name, nil
}

//IDからタグ(複数形)IDを取得する
func (UserService) GetUserTagsID(userID int64) ([]int64, error) {

	fmt.Println("タグID取得")
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
	fmt.Println("タグ名取得")
	tag := model.Tag{}
	_, err := DbEngine.Table("tag").Where("id=?", tagID).Get(&tag)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return tag.Name, nil
}

//attendanceテーブルに登録する
func (UserService) RegisterAttendance(userID int64, date string, flag bool) error {
	attendance := model.Attendance{
		UserID: userID,
		Date:   date,
		Flag:   flag,
	}
	_, err := DbEngine.Table("attendance").Insert(&attendance)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (UserService) TemporarilySavedAttendance(userID int64, flag int64) error {

	//update
	_, err := DbEngine.Table("attendance_tmp").Where("user_id=?", userID).Update(map[string]interface{}{"flag": flag})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

//attendance_tmpテーブルから登録済みのデータを全て取得する
func (UserService) GetAllAttendancesTmp() ([]model.AttendanceTmp, error) {
	attendanceTmp := make([]model.AttendanceTmp, 0)
	err := DbEngine.Table("attendance_tmp").Find(&attendanceTmp)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return attendanceTmp, nil
}

// user_idからuuidを求める
func (UserService) GetUserUUIDByUserID(userID int64) (string, error) {

	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.UUID, nil
}

//uuidからuser_idを求める
func (UserService) GetUserIDByUUID(uuid string) (int64, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("uid=?", uuid).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	return user.ID, nil
}

//
