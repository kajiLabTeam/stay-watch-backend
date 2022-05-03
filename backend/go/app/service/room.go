package service

import (
	"Stay_watch/model"
	"fmt"
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
	_, err := DbEngine.Table("log").Desc("start_at").Limit(1).Where("user_id=?", userID).Update(map[string]string{"end_at": currentTime.Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (RoomService) GetSimultaneousList(userID string) ([]model.Log, error) {
	currentTime := time.Now()
	logs := make([]model.Log, 0)
	err := DbEngine.Table("log").Where("user_id=?", userID).And("start_at >= ?", currentTime.Add(time.Hour*-24*7).Format("2006-01-02 15:04:05")).Find(&logs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	dates := make([]string, 0)
	roomIDs := make([]int64, 0)
	for _, log := range logs {
		dates = append(dates, log.StartAt[:10])
		roomIDs = append(roomIDs, log.RoomID)
	}
	uniqueDates := sliceUniqueString(dates)
	uniqueRoomIDs := sliceUniqueNumber(roomIDs)

	dateSql := ""
	for index, uniqueDate := range uniqueDates {
		dateSql += "start_at like '" + uniqueDate + "%' or "
		if index == len(uniqueDates)-1 {
			dateSql = dateSql[:len(dateSql)-4]
		}
	}

	roomSql := ""
	for index, uniqueRoomID := range uniqueRoomIDs {
		roomSql += "room_id=" + fmt.Sprintf("%d", uniqueRoomID) + " or "
		if index == len(uniqueRoomIDs)-1 {
			roomSql = roomSql[:len(roomSql)-4]
		}
	}

	sameDayAndRoomlogs := make([]model.Log, 0)
	err = DbEngine.Table("log").Where(dateSql).And(roomSql).Asc("user_id").Find(&sameDayAndRoomlogs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	// userIDs := make([]string, 0)
	// for _, sameDayAndRoomlog := range sameDayAndRoomlogs {
	// 	userIDs = append(userIDs, sameDayAndRoomlog.UserID)
	// }
	// uniqueUserIDs := sliceUniqueString(userIDs)
	UserService := UserService{}
	RoomService := RoomService{}
	userRoomTimeLogGetResponses := make([]model.UserRoomTimeLogGetResponse, 0)
	roomStayTimes := make([]model.RoomStayTime, 0)
	timeRooms := make([]model.TimeRoom, 0)
	times := make([]int, 0)

	userRoomTimeLogGetResponse := model.UserRoomTimeLogGetResponse{}
	for index, sameDayAndRoomlog := range sameDayAndRoomlogs {
		//Ascで昇順にしているため、違うuserIDになるまでループする
		if sameDayAndRoomlog.UserID == sameDayAndRoomlogs[index+1].UserID {
			roomStayTime := model.RoomStayTime{}
			if sameDayAndRoomlog.StartAt[:10] == sameDayAndRoomlogs[index+1].StartAt[:10] {
				timeRoom := model.TimeRoom{}
				if sameDayAndRoomlog.RoomID == sameDayAndRoomlogs[index+1].RoomID {
					//同じ部屋で同じ日に同じ時間がある場合
					times = append(times, 9)
				} else {
					timeRoom.ID = int(sameDayAndRoomlog.RoomID)
					roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					timeRoom.Name = roomName
					//同じ部屋最後の時間を追加
					times = append(times, 9)
					timeRoom.Times = times
					timeRooms = append(timeRooms, timeRoom)
				}
			} else {
				//同じ部屋で違う日に同じ時間がある場合
				roomStayTime.Date = sameDayAndRoomlog.StartAt[:10]
				roomStayTime.TimeRooms = timeRooms
				roomStayTimes = append(roomStayTimes, roomStayTime)
			}
		} else {
			//同じ部屋で違うuserIDがある場合
			userRoomTimeLogGetResponse.ID = sameDayAndRoomlog.UserID
			userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			userRoomTimeLogGetResponse.RoomStayTimes = roomStayTimes
			userRoomTimeLogGetResponse.Name = userName
			userRoomTimeLogGetResponses = append(userRoomTimeLogGetResponses, userRoomTimeLogGetResponse)
		}
	}

	// for _, uniqueUserID := range uniqueUserIDs {
	// 	UserService.GetUserName(uniqueUserID)
	// 	for _, sameDayAndRoomlog := range sameDayAndRoomlogs {
	// 		if sameDayAndRoomlog.UserID == uniqueUserID {
	// 			userRoomTimeLogGetResponses = append(userRoomTimeLogGetResponses, model.UserRoomTimeLogGetResponse{
	// 				UserName: UserService.UserName,
	// 				StartAt:  sameDayAndRoomlog.StartAt,
	// 				EndAt:    sameDayAndRoomlog.EndAt,
	// 			})
	// 		}
	// 	}
	// }

	return sameDayAndRoomlogs, nil
}

func sliceUniqueString(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func sliceUniqueNumber(target []int64) (unique []int64) {
	m := map[int64]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
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
