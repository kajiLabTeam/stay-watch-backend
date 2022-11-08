package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func Register(c *gin.Context) {

	RegistrationUserForm := model.User{}
	c.BindJSON(&RegistrationUserForm)
	fmt.Println(RegistrationUserForm)
	UserService := service.UserService{}
	//userIDがないなら新規登録
	if RegistrationUserForm.ID == 0 {
		user := model.User{
			Name:  RegistrationUserForm.Name,
			Email: RegistrationUserForm.Email,
			Role:  RegistrationUserForm.Role,
			UUID:  UserService.NewUUID(),
		}
		err := UserService.RegisterUser(&user)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
	}
	//userIDがあるなら更新
	if RegistrationUserForm.ID != 0 {
		//userNameが空なので、userIDからuserNameを取得する
		userName, err := UserService.GetUserNameByUserID(int64(RegistrationUserForm.ID))
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		uuid, err := UserService.GetUserUUIDByUserID(int64(RegistrationUserForm.ID))
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		user := model.User{
			// ID:    RegistrationUserForm.id,
			Name:  userName,
			Email: RegistrationUserForm.Email,
			Role:  RegistrationUserForm.Role,
			UUID:  uuid,
		}
		err = UserService.UpdateUser(&user)

		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
	}

	mailService := service.MailService{}
	mailService.SendMail("滞在ウォッチユーザ登録の完了のお知らせ", "ユーザ登録が完了したので滞在ウォッチを閲覧することが可能になりました\n一度プロジェクトをリセットしたので再度ログインお願いします。\nアプリドメイン\nhttps://stay-watch-go.kajilab.tk/", RegistrationUserForm.Email)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func UserList(c *gin.Context) {

	UserService := service.UserService{}
	users, err := UserService.GetAllUser()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	userInformationGetResponse := []model.UserInformationGetResponse{}

	for _, user := range users {

		tags := make([]model.Tag, 0)
		tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.String(http.StatusInternalServerError, "Server Error")
				return
			}
			tag := model.Tag{

				Name: tagName,
			}
			tags = append(tags, tag)
		}

		userInformationGetResponse = append(userInformationGetResponse, model.UserInformationGetResponse{
			ID:   int64(user.ID),
			Name: user.Name,
			Tags: tags,
		})
	}

	c.JSON(200, userInformationGetResponse)
}

func Attendance(c *gin.Context) {

	//構造体定義
	type Meeting struct {
		ID int64 `json:"meetingID"`
	}
	var meeting Meeting
	err := c.Bind(&meeting)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(meeting.ID)
	UserService := service.UserService{}
	//attendaance_tmpテーブルから全てのデータを取得する
	allAttendancesTmp, err := UserService.GetAllAttendancesTmp()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	isExist := true
	flagCount := 0
	if meeting.ID == 2 {
		for i := 0; i < 16; i++ {
			if allAttendancesTmp[i].Flag == 0 {
				flagCount++
			}
		}
		if flagCount == 16 {
			isExist = false
		}
	}
	if meeting.ID == 1 {
		for i := 16; i < 28; i++ {
			if allAttendancesTmp[i].Flag == 0 {
				flagCount++
			}
		}
		if flagCount == 12 {
			isExist = false
		}
	}

	ExcelService := service.ExcelService{}
	if isExist {
		ExcelService.WriteExcel(allAttendancesTmp, meeting.ID)
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// func SimultaneousStayUserList(c *gin.Context) {
// 	userID := c.Param("user_id")
// 	//int64に変換
// 	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}

// 	UserService := service.UserService{}
// 	RoomService := service.RoomService{}

// 	logs, err := RoomService.GetLogByUserAndDate(userIDInt64, 14)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}
// 	simultaneousStayUserGetResponses, err := UserService.GetSameTimeUser(logs)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}

// 	c.JSON(200, simultaneousStayUserGetResponses)
// }

func Check(c *gin.Context) {
	firebaseUserInfo, err := verifyCheck(c.Request)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{
			"status": "invalid token",
		})
		return
	}

	UserService := service.UserService{}
	user, err := UserService.GetUserByEmail(firebaseUserInfo["Email"])
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	fmt.Println(user)

	//メールアドレスが存在しない場合はUserは存在しないのでリクエスト失敗
	if (user == model.User{}) {
		c.JSON(403, gin.H{
			"status": "権限がありません 管理者にユーザ追加を依頼してください",
		})
		return
	}

	c.JSON(200,
		user,
	)
}
