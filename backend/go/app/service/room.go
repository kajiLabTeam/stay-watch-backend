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

func (RoomService) GetSimultaneousList(userID string) ([]model.UserRoomTimeLogGetResponse, error) {
	currentTime := time.Now()
	logs := make([]model.Log, 0)
	err := DbEngine.Table("log").Where("user_id=?", userID).And("start_at >= ?", currentTime.Add(time.Hour*-24*14).Format("2006-01-02 15:04:05")).Find(&logs)
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
	err = DbEngine.Table("log").Where(dateSql).And(roomSql).Asc("user_id").Asc("start_at").Asc("room_id").Find(&sameDayAndRoomlogs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	UserService := UserService{}
	RoomService := RoomService{}
	userRoomTimeLogGetResponses := make([]model.UserRoomTimeLogGetResponse, 0)
	roomStayTimes := make([]model.RoomStayTime, 0)
	timeRooms := make([]model.TimeRoom, 0)
	times := make([]int, 0)

	userRoomTimeLogGetResponse := model.UserRoomTimeLogGetResponse{}
	for index, sameDayAndRoomlog := range sameDayAndRoomlogs {
		//Ascで昇順にしているため、違うuserIDになるまでループする
		roomStayTime := model.RoomStayTime{}
		timeRoom := model.TimeRoom{}

		//最後のindexの時
		if index == len(sameDayAndRoomlogs)-1 {
			times = append(times, 9)

			timeRoom.ID = int(sameDayAndRoomlog.RoomID)
			roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			timeRoom.Name = roomName
			//同じ部屋最後の時間を追加
			timeRoom.Times = times
			timeRooms = append(timeRooms, timeRoom)
			times = nil

			//同じ部屋で違う日に同じ時間がある場合
			roomStayTime.Date = sameDayAndRoomlog.StartAt[:10]
			roomStayTime.TimeRooms = timeRooms
			roomStayTimes = append(roomStayTimes, roomStayTime)
			timeRooms = nil

			//同じ部屋で違うuserIDがある場合
			userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			userRoomTimeLogGetResponse.ID = sameDayAndRoomlog.UserID
			userRoomTimeLogGetResponse.RoomStayTimes = roomStayTimes
			userRoomTimeLogGetResponse.Name = userName
			userRoomTimeLogGetResponses = append(userRoomTimeLogGetResponses, userRoomTimeLogGetResponse)
			roomStayTimes = nil

		} else {
			if sameDayAndRoomlog.UserID == sameDayAndRoomlogs[index+1].UserID {

				if sameDayAndRoomlog.StartAt[:10] == sameDayAndRoomlogs[index+1].StartAt[:10] {
					if sameDayAndRoomlog.RoomID == sameDayAndRoomlogs[index+1].RoomID {
						//同じ部屋で同じ日に同じ時間がある場合
						times = append(times, 9)
					} else {
						times = append(times, 9)

						timeRoom.ID = int(sameDayAndRoomlog.RoomID)
						roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
						if err != nil {
							log.Fatal(err.Error())
							return nil, err
						}
						timeRoom.Name = roomName
						//同じ部屋最後の時間を追加
						timeRoom.Times = times
						timeRooms = append(timeRooms, timeRoom)
						times = nil
					}
				} else {
					times = append(times, 9)

					timeRoom.ID = int(sameDayAndRoomlog.RoomID)
					roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					timeRoom.Name = roomName
					//同じ部屋最後の時間を追加
					timeRoom.Times = times
					timeRooms = append(timeRooms, timeRoom)
					times = nil

					//同じ部屋で違う日に同じ時間がある場合
					roomStayTime.Date = sameDayAndRoomlog.StartAt[:10]
					roomStayTime.TimeRooms = timeRooms
					roomStayTimes = append(roomStayTimes, roomStayTime)
					timeRooms = nil
				}
			} else {
				times = append(times, 9)

				timeRoom.ID = int(sameDayAndRoomlog.RoomID)
				roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				timeRoom.Name = roomName
				//同じ部屋最後の時間を追加
				timeRoom.Times = times
				timeRooms = append(timeRooms, timeRoom)
				times = nil

				//同じ部屋で違う日に同じ時間がある場合
				roomStayTime.Date = sameDayAndRoomlog.StartAt[:10]
				roomStayTime.TimeRooms = timeRooms
				roomStayTimes = append(roomStayTimes, roomStayTime)
				timeRooms = nil

				//同じ部屋で違うuserIDがある場合
				userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				userRoomTimeLogGetResponse.ID = sameDayAndRoomlog.UserID
				userRoomTimeLogGetResponse.RoomStayTimes = roomStayTimes
				userRoomTimeLogGetResponse.Name = userName
				userRoomTimeLogGetResponses = append(userRoomTimeLogGetResponses, userRoomTimeLogGetResponse)
				roomStayTimes = nil
			}
		}
	}
	fmt.Println(userRoomTimeLogGetResponses)


	return userRoomTimeLogGetResponses, nil
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

func (RoomService) GetTimesFromStartAtAndEntAt(startAt string, endAt string) ([]string, error) {
	times := []string{}
	startAtTime, err := time.Parse("2006-01-02 15:04:05", startAt)
	if err != nil {
		return nil, err
	}
	endAtTime, err := time.Parse("2006-01-02 15:04:05", endAt)
	if err != nil {
		return nil, err
	}
	for startAtTime.Before(endAtTime) {
		times = append(times, startAtTime.Format("15:04"))
		startAtTime = startAtTime.Add(time.Minute * 15)
	}
	return times, nil
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
