package service

import (
	"Stay_watch/model"
	"log"
	"time"
)

type RoomService struct{}

func (RoomService) SetLog(Log *model.Log) error {

	_, err := DbEngine.Table("log").Insert(Log)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

//該当ユーザが存在するか確認
func (RoomService) GetStayer(stayer *model.Stayer) (error, bool) {
	affected, err := DbEngine.Get(stayer)
	// fmt.Println("affected=", affected)
	// fmt.Printf("%T\n", affected) // int

	if err != nil {
		log.Fatal(err.Error())
		return err, false
	}

	if affected == true {
		return nil, true
	}
	return nil, false
}

//滞在者全体を取得する
func (RoomService) GetAllStayer() ([]model.Stayer, error) {
	stayers := make([]model.Stayer, 0)
	err := DbEngine.Find(&stayers)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return stayers, nil
}

//滞在者の一部を取得する
func (RoomService) GetStayerByRoomID(roomID int64) ([]model.Stayer, error) {
	stayers := make([]model.Stayer, 0)
	err := DbEngine.Table("stayer").Where("room_id=?", roomID).Find(&stayers)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return stayers, nil
}

func (RoomService) SetStayer(stayer *model.Stayer) error {

	_, err := DbEngine.Table("stayer").Insert(stayer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (RoomService) UpdateStayer(stayer *model.Stayer) error {

	_, err := DbEngine.Table("stayer").Where("user_id=?", stayer.UserID).Update(stayer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (RoomService) DeleteStayer(userID string) error {

	_, err := DbEngine.Table("stayer").Where("user_id=?", userID).Delete(model.Stayer{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (RoomService) InsertEndAt(userID string) error {

	currentTime := time.Now()
	_, err := DbEngine.Table("log").Where("user_id=?", userID).Update(map[string]string{"end_at": currentTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

//ルームIDからルームの名前を取得する
func (RoomService) GetRoomName(roomID int64) (string, error) {
	room := model.Room{}
	_, err := DbEngine.Table("room").Where("id=?", roomID).Get(&room)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return room.Name, nil
}

//IDから名前を取得する
func (RoomService) GetUserName(userID string) (string, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.Name, nil
}

//IDからチーム名を取得する
func (RoomService) GetUserTeam(userID string) (string, error) {
	user := model.User{}
	_, err := DbEngine.Table("user").Where("id=?", userID).Get(&user)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return user.Team, nil
}

//全てのログを取得する
func (RoomService) GetAllLog() ([]model.Log, error) {
	logs := make([]model.Log, 0)
	err := DbEngine.Find(&logs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return logs, nil
}
