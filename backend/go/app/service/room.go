package service

import (
	"Stay_watch/model"
	"Stay_watch/util"
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

func (RoomService) GetSimultaneousList(userID string) ([]model.SimulataneousStayLogGetResponse, error) {
	currentTime := time.Now()
	logs := make([]model.Log, 0)

	//指定したuserの14日以内のログを取得
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

	util := util.Util{}

	//滞在した日付
	uniqueDates := util.SliceUniqueString(dates)
	//滞在した部屋
	uniqueRoomIDs := util.SliceUniqueNumber(roomIDs)

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
	err = DbEngine.Table("log").Where(dateSql).And(roomSql).OrderBy("date_format(start_at,'%Y-%m-%d') ,room_id ").Find(&sameDayAndRoomlogs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	UserService := UserService{}
	RoomService := RoomService{}

	simulataneousStayLogsGetResponse := make([]model.SimulataneousStayLogGetResponse, 0)
	roomsGetResponse := make([]model.RoomGetResponse, 0)
	stayTimes := make([]model.StayTime, 0)
	simulataneousStayLogGetResponse := model.SimulataneousStayLogGetResponse{}
	dateCount := 1

	for index, sameDayAndRoomlog := range sameDayAndRoomlogs {
		//Ascで昇順にしているため、違うuserIDになるまでループする
		roomGetResponse := model.RoomGetResponse{}
		stayTime := model.StayTime{}

		//最後のindexの時
		if index == len(sameDayAndRoomlogs)-1 {
			stayTime.ID = sameDayAndRoomlog.ID
			userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			stayTime.UserName = userName

			locationTime, err := RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.StartAt, "Asia/Tokyo")
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			unixMilli := timeToUnixMilli(locationTime)
			stayTime.StartAt = unixMilli

			locationTime, err = RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.EndAt, "Asia/Tokyo")
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			unixMilli = timeToUnixMilli(locationTime)
			stayTime.EndAt = unixMilli

			//検索対象者は赤色にする
			if userID == sameDayAndRoomlog.UserID {
				stayTime.Color = "red"
			} else {
				stayTime.Color = "green"
			}
			stayTimes = append(stayTimes, stayTime)

			roomGetResponse.ID = sameDayAndRoomlog.RoomID
			roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			roomGetResponse.Name = roomName
			roomGetResponse.StayTimes = stayTimes
			roomsGetResponse = append(roomsGetResponse, roomGetResponse)
			stayTimes = nil

			//後でuniqueDateのindexに置き換えるかも
			simulataneousStayLogGetResponse.ID = int64(dateCount)
			dateCount++
			simulataneousStayLogGetResponse.Date = sameDayAndRoomlog.StartAt[:10]
			simulataneousStayLogGetResponse.Rooms = roomsGetResponse
			simulataneousStayLogsGetResponse = append(simulataneousStayLogsGetResponse, simulataneousStayLogGetResponse)
			roomsGetResponse = nil
		} else {
			if sameDayAndRoomlog.StartAt[:10] == sameDayAndRoomlogs[index+1].StartAt[:10] {
				if sameDayAndRoomlog.RoomID == sameDayAndRoomlogs[index+1].RoomID {
					stayTime.ID = sameDayAndRoomlog.ID
					userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					stayTime.UserName = userName

					locationTime, err := RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.StartAt, "Asia/Tokyo")
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					unixMilli := timeToUnixMilli(locationTime)
					stayTime.StartAt = unixMilli

					locationTime, err = RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.EndAt, "Asia/Tokyo")
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					unixMilli = timeToUnixMilli(locationTime)
					stayTime.EndAt = unixMilli

					//検索対象者は赤色にする
					if userID == sameDayAndRoomlog.UserID {
						stayTime.Color = "red"
					} else {
						stayTime.Color = "green"
					}
					stayTimes = append(stayTimes, stayTime)
				} else {
					stayTime.ID = sameDayAndRoomlog.ID
					userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					stayTime.UserName = userName

					locationTime, err := RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.StartAt, "Asia/Tokyo")
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					unixMilli := timeToUnixMilli(locationTime)
					stayTime.StartAt = unixMilli

					locationTime, err = RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.EndAt, "Asia/Tokyo")
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					unixMilli = timeToUnixMilli(locationTime)
					stayTime.EndAt = unixMilli

					//検索対象者は赤色にする
					if userID == sameDayAndRoomlog.UserID {
						stayTime.Color = "red"
					} else {
						stayTime.Color = "green"
					}
					stayTimes = append(stayTimes, stayTime)

					roomGetResponse.ID = sameDayAndRoomlog.RoomID
					roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
					if err != nil {
						log.Fatal(err.Error())
						return nil, err
					}
					roomGetResponse.Name = roomName
					roomGetResponse.StayTimes = stayTimes
					roomsGetResponse = append(roomsGetResponse, roomGetResponse)
					stayTimes = nil
				}
			} else {
				stayTime.ID = sameDayAndRoomlog.ID
				userName, err := UserService.GetUserName(sameDayAndRoomlog.UserID)
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				stayTime.UserName = userName

				locationTime, err := RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.StartAt, "Asia/Tokyo")
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				unixMilli := timeToUnixMilli(locationTime)
				stayTime.StartAt = unixMilli

				locationTime, err = RoomService.convertDatetimeToLocationTime(sameDayAndRoomlog.EndAt, "Asia/Tokyo")
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				unixMilli = timeToUnixMilli(locationTime)
				stayTime.EndAt = unixMilli

				//検索対象者は赤色にする
				if userID == sameDayAndRoomlog.UserID {
					stayTime.Color = "red"
				} else {
					stayTime.Color = "green"
				}
				stayTimes = append(stayTimes, stayTime)

				roomGetResponse.ID = sameDayAndRoomlog.RoomID
				roomName, err := RoomService.GetRoomName(sameDayAndRoomlog.RoomID)
				if err != nil {
					log.Fatal(err.Error())
					return nil, err
				}
				roomGetResponse.Name = roomName
				roomGetResponse.StayTimes = stayTimes
				roomsGetResponse = append(roomsGetResponse, roomGetResponse)
				stayTimes = nil

				//後でuniqueDateのindexに置き換えるかも
				simulataneousStayLogGetResponse.ID = int64(dateCount)
				dateCount++
				simulataneousStayLogGetResponse.Date = sameDayAndRoomlog.StartAt[:10]
				simulataneousStayLogGetResponse.Rooms = roomsGetResponse
				simulataneousStayLogsGetResponse = append(simulataneousStayLogsGetResponse, simulataneousStayLogGetResponse)
				roomsGetResponse = nil
			}
		}
	}

	fmt.Println(simulataneousStayLogsGetResponse)
	fmt.Println("追加")

	return simulataneousStayLogsGetResponse, nil
}

//引数datetime文字列とタイムゾーン文字列を受け取りTime型に変換する関数
func (RoomService) convertDatetimeToLocationTime(datetime string, timezone string) (time.Time, error) {
	jst, _ := time.LoadLocation(timezone)
	locationTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, jst)
	if err != nil {
		log.Fatal(err.Error())
		return time.Time{}, err
	}
	return locationTime, nil
}

func timeToUnixMilli(t time.Time) int64 {
	return t.UnixNano() / 1000000
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

//指定した時間のログを取得する
func (RoomService) GetLogsFromStartAtAndEntAt(startAt string, endAt string) ([]model.Log, error) {
	logs := make([]model.Log, 0)
	startAtTime, err := time.Parse("2006-01-02 15:04:05", startAt)
	if err != nil {
		return nil, err
	}
	endAtTime, err := time.Parse("2006-01-02 15:04:05", endAt)
	if err != nil {
		return nil, err
	}
	err = DbEngine.Where("start_at>=? and start_at<=?", startAtTime, endAtTime).Find(&logs)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return logs, nil
}
