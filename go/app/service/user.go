package service

import (
	"Stay_watch/model"
	"fmt"
	"strconv"
)

type UserService struct{}

func (UserService) RegisterSampleUser(user *model.User) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Create(user)
	if result.Error != nil {
		fmt.Printf("ユーザ登録処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

// 新しいuuidを生成する
func (UserService) NewUUID() string {

	//dbから一番最後に登録されたuuidを取得する
	user := model.User{}
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return ""
	}
	defer closer.Close()
	result := DbEngine.Last(&user)
	if result.Error != nil {
		fmt.Printf("uuid取得失敗 %v", result.Error)
		return ""
	}

	uuid := user.UUID
	fowardTarget := uuid[0:28]
	backTarget := uuid[28:]
	//10進数に変換
	targetInt, _ := strconv.ParseInt(backTarget, 16, 64)
	targetInt = targetInt + 1
	//16新数に変換
	targetHex := strconv.FormatInt(targetInt, 16)

	return fowardTarget + targetHex

}

// ユーザ登録処理new（）
func (UserService) RegisterUser(user *model.User) (int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return -1, err
	}
	defer closer.Close()
	result := DbEngine.Create(user)
	if result.Error != nil {
		fmt.Printf("ユーザ登録処理失敗 %v", result.Error)
		return -1, result.Error
	}
	return int64(user.ID), nil
}

// ユーザのアップデート
func (UserService) UpdateUser(user *model.User, userId int64) error {

	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()

	result := DbEngine.Model(&model.User{}).Where("id = ?", userId).Updates(&user)
	if result.Error != nil {
		fmt.Printf("ユーザ更新失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

// Uuidの更新
func (UserService) UpdateUuid(newUuid string, userId int64) error {

	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Model(&user).Where("id = ?", userId).Update("uuid", newUuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ユーザの削除
func (UserService) DeleteUser(userId int64) error {
	fmt.Println("ユーザ削除サービス")

	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Unscoped().Delete(&model.User{}, userId)
	if result.Error != nil {
		fmt.Printf("ユーザ削除処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

// 全てのユーザを取得する
func (UserService) GetAllUser() ([]model.User, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	users := make([]model.User, 0)
	result := DbEngine.Find(&users)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return users, nil
}

// 該当ビーコンIDのユーザを取得する
func (UserService) GetUsersByBeaconId(beaconId int64) ([]model.User, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	users := make([]model.User, 0)
	result := DbEngine.Where("beacon_id = ?", beaconId).Find(&users)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return users, nil
}

// コミュニティのユーザの編集に必要な部分を取得する
func (UserService) GetEditUsersByCommunityId(communityId int64) ([]model.User, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	users := make([]model.User, 0)
	//result := DbEngine.Find(&users)
	result := DbEngine.Where("community_id = ?", communityId).Find(&users)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return users, nil
}

// 全てのユーザネームを取得する
func (UserService) GetAllUserName() ([]string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	users := make([]model.User, 0)
	result := DbEngine.Find(&users)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return nil, result.Error
	}
	names := make([]string, 0)
	for _, user := range users {
		names = append(names, user.Name)
	}
	return names, nil
}

// IDから名前を取得する
func (UserService) GetUserNameByUserID(userID int64) (string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return "", err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("id=?", userID).Take(&user)
	if result.Error != nil {
		// 見つからない場合は削除されたユーザであるということになる
		return "削除済みユーザ", nil
	}
	return user.Name, nil
}

// IDからタグ(複数形)IDを取得する
func (UserService) GetUserTagsID(userID int64) ([]int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tags := make([]int64, 0)
	result := DbEngine.Table("tag_maps").Where("user_id=?", userID).Select("tag_id").Find(&tags)
	if result.Error != nil {
		fmt.Printf("タグID取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tags, nil
}

// タグIDからタグ名を取得する
func (UserService) GetTagName(tagID int64) (string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return "", err
	}
	defer closer.Close()

	tag := model.Tag{}
	result := DbEngine.Where("id=?", tagID).Take(&tag)
	if result.Error != nil {
		fmt.Printf("タグ名取得失敗 %v", result.Error)
		return "", result.Error
	}
	return tag.Name, nil
}

// attendanceテーブルに登録する
func (UserService) RegisterAttendance(userID int64, date string, flag bool) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	attendance := model.Attendance{
		UserID: userID,
		Date:   date,
		Flag:   flag,
	}
	result := DbEngine.Create(&attendance)
	if result.Error != nil {
		fmt.Printf("出席登録失敗 %v", result.Error)
		return result.Error
	}

	return nil
}

func (UserService) TemporarilySavedAttendance(userID int64, flag int64) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	//update
	result := DbEngine.Model(&model.AttendanceTmp{}).Where("user_id=?", userID).Update("flag", flag)
	if result.Error != nil {
		fmt.Printf("出席登録失敗 %v", result.Error)
		return result.Error
	}

	return nil
}

// attendance_tmpテーブルから登録済みのデータを全て取得する
func (UserService) GetAllAttendancesTmp() ([]model.AttendanceTmp, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	attendanceTmp := make([]model.AttendanceTmp, 0)
	result := DbEngine.Find(&attendanceTmp)
	if result.Error != nil {
		fmt.Printf("出席取得失敗 %v", result.Error)
		return nil, result.Error
	}
	return attendanceTmp, nil
}

// user_idからuuidを求める
func (UserService) GetUserUUIDByUserID(userID int64) (string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return "", err
	}
	defer closer.Close()

	user := model.User{}
	result := DbEngine.Where("id=?", userID).Take(&user)
	if result.Error != nil {
		fmt.Printf("ユーザIDからユーザ名取得失敗 %v", result.Error)
		return "", result.Error
	}

	return user.UUID, nil
}

// uuidからuser_idを求める
func (UserService) GetUserIDByUUID(uuid string) (int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return 0, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("uuid=?", uuid).Take(&user)
	if result.Error != nil {
		fmt.Printf("UUIDからユーザ名取得失敗 %v", result.Error)
		return 0, result.Error
	}

	return int64(user.ID), nil
}

func (UserService) GetUserIDByEmail(email string) (int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return 0, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("email=?", email).Take(&user)
	if result.Error != nil {
		fmt.Printf("ユーザID取得失敗 %v", result.Error)
		return 0, result.Error
	}

	return int64(user.ID), nil

}

func (UserService) GetEmailByUserId(userId int64) (string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return "", err
	}
	defer closer.Close()

	user := model.User{}
	result := DbEngine.Where("id=?", userId).Take(&user)
	if result.Error != nil {
		fmt.Printf("メールアドレス取得失敗 %v", result.Error)
		return "", result.Error
	}

	return user.Email, nil
}

func (UserService) IsEmailAlreadyRegistered(email string) (bool, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		// 接続できなかった場合もtrueとする(どちらにしてもいい)
		return true, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("email=?", email).Take(&user)
	// エラーの時はメールアドレスが見つからなかった時と同じなため
	if result.Error != nil || email == "" {
		return false, nil
	}
	return true, nil
}

func (UserService) IsPrivateKeyAlreadyRegistered(privateKey string) (bool, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		// 接続できなかった場合もtrueとする(どちらにしてもいい)
		return true, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("private_key=?", privateKey).Take(&user)
	// エラーの時はPrivateKeyが見つからなかった時と同じなため
	if result.Error != nil || privateKey == "" {
		return false, nil
	}
	return true, nil
}

//指定されたログリストと同じ時間にいたユーザを取得する
// func (UserService) GetSameTimeUser(logs []model.Log) ([]model.SimultaneousStayUserGetResponse, error) {
// 	targetLogs := make([]model.Log, 0)
// 	fmt.Println(logs)
// 	dates := make([]string, 0)
// 	for _, log := range logs {
// 		dates = append(dates, log.StartAt.Format("2006-01-02"))
// 		//時間が被るログを取得
// 		err := DbEngine.Table("log").Asc("start_at").Where("start_at >= ?", log.StartAt).And("start_at <= ?", log.EndAt).Or(
// 			"end_at >= ? and end_at <= ?", log.StartAt, log.EndAt).Or(
// 			"start_at <= ? and end_at >= ?", log.StartAt, log.EndAt).And("room_id = ?", log.RoomID).Find(&targetLogs)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		// DbEngine.Table("log").Where("start_at >= ?", log.StartAt).And("start_at <= ?", log.EndAt)
// 		// DbEngine.Table("log").Where("end_at >= ?", log.StartAt).And("end_at <= ?", log.EndAt)
// 	}

// 	simultaneousStayUserGetResponses := make([]model.SimultaneousStayUserGetResponse, 0)

// 	UserService := UserService{}

// 	utilService := util.Util{}
// 	dates = utilService.SliceUniqueString(dates)
// 	fmt.Println(dates)
// 	for _, date := range dates {

// 		userIDs := make([]int64, 0)
// 		for _, log := range targetLogs {
// 			if log.StartAt.Format("2006-01-02") == date {
// 				userIDs = append(userIDs, log.UserID)
// 			}
// 		}
// 		uniqueUserIDs := utilService.SliceUniqueNumber(userIDs)

// 		names := make([]model.Name, 0)
// 		for _, uniqueUserID := range uniqueUserIDs {
// 			userName, err := UserService.GetUserNameByUserID(uniqueUserID)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 			names = append(names, model.Name{
// 				Name: userName,
// 				ID:   uniqueUserID,
// 			})
// 		}

// 		simultaneousStayUserGetResponses = append(simultaneousStayUserGetResponses, model.SimultaneousStayUserGetResponse{
// 			Date:  date,
// 			Names: names,
// 		})
// 	}
// 	return simultaneousStayUserGetResponses, nil
// }

// emailからユーザを取得する
func (UserService) GetUserByEmail(email string) (model.User, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return model.User{}, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("email=?", email).Take(&user)
	if result.Error != nil {
		fmt.Printf("Emailからユーザ名取得失敗 %v", result.Error)
		return model.User{}, result.Error
	}
	return user, nil
}

func (UserService) GetUserByUserId(userId int64) (model.User, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return model.User{}, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("id=?", userId).Take(&user)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return model.User{}, result.Error
	}
	return user, nil
}

func (UserService) GetCommunityIdByUserId(userId int64) (int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return 0, err
	}
	defer closer.Close()
	user := model.User{}
	result := DbEngine.Where("id=?", userId).Take(&user)
	if result.Error != nil {
		fmt.Printf("ユーザ取得失敗 %v", result.Error)
		return 0, result.Error
	}
	return int64(user.CommunityId), nil
}
