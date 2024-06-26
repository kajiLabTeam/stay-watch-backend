package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func createUuid(communityId int64, uuidEditable bool, beaconName string, userId int64, requestUuid string) string {
	// コミュニティIDを16進数3桁に変換
	communityIdHex := fmt.Sprintf("%03x", communityId)
	newUuid := ""

	if beaconName == "iPhone" {
		// iPhoneの場合「8ebc21144abd00000000ff01000 + (userId(16進数))」
		userIdHex := fmt.Sprintf("%05x", userId)
		newUuid = "8ebc21144abd00000000ff01000" + userIdHex
	} else if uuidEditable {
		// 編集可能（物理）の場合ユーザがフォームで入力した値を用いる
		//newUuid = "e7d61ea3f8dd49c88f2ff24f" + communityIdHex + requestUuid
		newUuid = "8ebc21144abd" + "ba0d" + "b7c6" + "ff0f" + communityIdHex + requestUuid
	} else {
		// 編集不可（Android）の場合ユーザIDから取得した値を用いる
		// ユーザIDを16進数5桁に変換
		userIdHex := fmt.Sprintf("%05x", userId)
		newUuid = "8ebc21144abd" + "ba0d" + "b7c6" + "ff0a" + communityIdHex + userIdHex
	}

	return newUuid
}

func CreateUser(c *gin.Context) {
	UserCreateRequest := model.UserCreateRequest{}
	c.Bind(&UserCreateRequest)

	UserService := service.UserService{}
	BeaconService := service.BeaconService{}
	TagService := service.TagService{}

	beacon, err := BeaconService.GetBeaconByBeaconName(UserCreateRequest.BeaconName)
	// もしbeaconTypeが取得できたらerrがnilになる
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get beacon by name"})
		return
	}

	user := model.User{
		Name:        UserCreateRequest.Name,
		Email:       UserCreateRequest.Email,
		Role:        UserCreateRequest.Role,
		UUID:        "",
		BeaconId:    int64(beacon.ID),
		CommunityId: UserCreateRequest.CommunityId,
	}

	// 同じメールアドレスのユーザが既に登録済みだったら409を返す
	isRegisterd, err := UserService.IsEmailAlreadyRegistered(UserCreateRequest.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email is arleady registerd"})
		return
	}
	if isRegisterd {
		// 同じメールアドレスが既に登録済みの場合
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Arleady Registered Email"})
		return
	}

	// usersテーブルにユーザ情報を保存
	registerdUserId, err := UserService.RegisterUser(&user)
	if err != nil {
		fmt.Printf("Cannnot register user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// UUIDの作成
	// コミュニティID取得
	communityId := UserCreateRequest.CommunityId
	newUuid := createUuid(communityId, beacon.UuidEditable, UserCreateRequest.BeaconName, registerdUserId, UserCreateRequest.Uuid)

	// UUIDを上書き
	err = UserService.UpdateUuid(newUuid, registerdUserId)
	if err != nil {
		fmt.Printf("Cannot update uuid: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update uuid"})
	}

	// tag_mapsテーブルにタグのマップを保存
	for _, tagId := range UserCreateRequest.TagIds {
		tag := model.TagMap{
			UserID: int64(registerdUserId),
			TagID:  int64(tagId),
		}
		err = TagService.CreateTagMap(&tag)
		if err != nil {
			fmt.Printf("Cannot register tagMap: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag map"})
			return
		}
	}

	if !strings.HasSuffix(os.Args[0], ".test") {
		mailService := service.MailService{}
		mailService.SendMail("滞在ウォッチユーザ登録の完了のお知らせ", "ユーザ登録が完了したので滞在ウォッチを閲覧することが可能になりました\n一度プロジェクトをリセットしたので再度ログインお願いします。\nアプリドメイン\nhttps://stay-watch-go.kajilab.tk/", UserCreateRequest.Email)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func PastCreateUser(c *gin.Context) {
	RegistrationUserForm := model.RegistrationUserForm{}
	c.Bind(&RegistrationUserForm)

	UserService := service.UserService{}
	//userIDがないなら新規登録
	if RegistrationUserForm.ID == 0 {
		user := model.User{
			Name:  RegistrationUserForm.Name,
			Email: RegistrationUserForm.Email,
			Role:  RegistrationUserForm.Role,
			UUID:  UserService.NewUUID(),
		}

		err := UserService.PastRegisterUser(&user)
		if err != nil {
			fmt.Printf("Cannnot register user: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
	}

	//userIDがあるなら更新
	if RegistrationUserForm.ID != 0 {
		//userNameが空なので、userIDからuserNameを取得する
		err := UserService.PastUpdateUser(
			int(RegistrationUserForm.ID),
			RegistrationUserForm.Email,
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to"})
			return
		}
	}

	if !strings.HasSuffix(os.Args[0], ".test") {
		mailService := service.MailService{}
		mailService.SendMail("滞在ウォッチユーザ登録の完了のお知らせ", "ユーザ登録が完了したので滞在ウォッチを閲覧することが可能になりました\n一度プロジェクトをリセットしたので再度ログインお願いします。\nアプリドメイン\nhttps://stay-watch-go.kajilab.tk/", RegistrationUserForm.Email)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func DeleteUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64) // string -> int64
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Type is not number")
	}

	fmt.Print("userId: ")
	fmt.Println(userId)

	UserService := service.UserService{}
	DeletedUserService := service.DeletedUserService{}

	// 削除するユーザの情報の取得
	user, err := UserService.GetUserByUserId(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// ユーザ情報をDeletedUserの型に格納
	deletedUser := model.DeletedUser{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		UUID:        user.UUID,
		BeaconId:    user.BeaconId,
		CommunityId: user.CommunityId,
		UserId:      userId,
	}

	// deletedUserテーブルに登録
	err = DeletedUserService.CreateDeletedUser(&deletedUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deleted user"})
		return
	}

	// usersテーブルから削除
	err = UserService.DeleteUser(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func UpdateUser(c *gin.Context) {
	UserUpdateRequest := model.UserUpdateRequest{}
	c.Bind(&UserUpdateRequest)

	UserService := service.UserService{}
	BeaconService := service.BeaconService{}
	TagService := service.TagService{}

	beacon, err := BeaconService.GetBeaconByBeaconName(UserUpdateRequest.BeaconName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// 更新前のユーザの情報を取得
	currentUser, err := UserService.GetUserByUserId(UserUpdateRequest.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// メールアドレスに変更がある場合、そのメールアドレスが他のユーザに既に使われているかをチェックする処理をプラスする
	if currentUser.Email != UserUpdateRequest.Email {
		isRegisterdEmail, err := UserService.IsEmailAlreadyRegistered(UserUpdateRequest.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email is arleady registerd"})
			return
		} else if isRegisterdEmail {
			// 同じメールアドレスが既に登録済みの場合
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Arleady Registered Email"})
			return
		}
	}

	// UUIDの作成
	newUuid := createUuid(UserUpdateRequest.CommunityId, beacon.UuidEditable, UserUpdateRequest.BeaconName, UserUpdateRequest.ID, UserUpdateRequest.Uuid)

	user := model.User{
		Name:        UserUpdateRequest.Name,
		Email:       UserUpdateRequest.Email,
		Role:        UserUpdateRequest.Role,
		UUID:        newUuid,
		BeaconId:    int64(beacon.ID),
		CommunityId: UserUpdateRequest.CommunityId,
	}

	// usersテーブルにユーザ情報を保存
	err = UserService.UpdateUser(&user, UserUpdateRequest.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	/* タグマップ関連 */
	// タグマップIDを取得
	tagMapIds, err := TagService.GetTagMapIdsByUserId(UserUpdateRequest.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tagMap"})
		return
	}

	// tag_mapsテーブルの変更前のマップを削除
	for _, tagMapId := range tagMapIds {
		err = TagService.DeleteTagMap(tagMapId)
		if err != nil {
			fmt.Printf("Cannot delete tagMap: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tagMap"})
			return
		}
	}

	// tag_mapsテーブルに新しいタグのマップを保存
	for _, tagId := range UserUpdateRequest.TagIds {
		tagMap := model.TagMap{
			UserID: int64(UserUpdateRequest.ID),
			TagID:  int64(tagId),
		}
		err = TagService.CreateTagMap(&tagMap)
		if err != nil {
			fmt.Printf("Cannot register tagMap: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to registered tagMap"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func PastUserList(c *gin.Context) {

	UserService := service.UserService{}

	users, err := UserService.GetAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	userInformationGetResponse := []model.UserInformationGetResponse{}

	for _, user := range users {

		tags := make([]model.TagGetResponse, 0)
		tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user tags"})
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tag name"})
				return
			}
			tag := model.TagGetResponse{
				ID:   tagID,
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
	c.JSON(http.StatusOK, userInformationGetResponse)
}

func UserList(c *gin.Context) {

	UserService := service.UserService{}
	communityId, err := strconv.ParseInt(c.Param("communityId"), 10, 64) // string -> int64
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Type is not number"})
	}

	users, err := UserService.GetEditUsersByCommunityId(communityId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot Get Users")
	}

	userInformationGetResponse := []model.UserInformationGetResponse{}

	for _, user := range users {

		tags := make([]model.TagGetResponse, 0)
		tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user tags"})
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tag"})
				return
			}
			tag := model.TagGetResponse{
				ID:   tagID,
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
	c.JSON(http.StatusOK, userInformationGetResponse)
}

func AdminUserList(c *gin.Context) {
	UserService := service.UserService{}
	BeaconService := service.BeaconService{}
	communityId, err := strconv.ParseInt(c.Param("communityId"), 10, 64) // string -> int64
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Type is not number")
	}

	edit_users, err := UserService.GetEditUsersByCommunityId(communityId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
	}

	userEditorResponse := []model.UserEditorResponse{}

	for _, user := range edit_users {

		tags := make([]model.TagGetResponse, 0)
		tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user tags"})
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tag"})
				return
			}
			tag := model.TagGetResponse{
				ID:   tagID,
				Name: tagName,
			}
			tags = append(tags, tag)
		}

		beacon, err := BeaconService.GetBeaconByBeaconId(user.BeaconId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get beacon"})
			return
		}

		userEditorResponse = append(userEditorResponse, model.UserEditorResponse{
			ID:                 int64(user.ID),
			Name:               user.Name,
			Uuid:               user.UUID,
			Email:              user.Email,
			Role:               user.Role,
			BeaconUuidEditable: beacon.UuidEditable,
			BeaconName:         beacon.Type,
			Tags:               tags,
		})
	}
	c.JSON(http.StatusOK, userEditorResponse)
}

func ExtendedUserList(c *gin.Context) {

	UserService := service.UserService{}
	users, err := UserService.GetAllUser()
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot GetAllUser")
		return
	}

	extendedUserInformationGetResponses := []model.ExtendedUserInformationGetResponse{}
	// userInformationGetResponse := []model.UserInformationGetResponse{}

	for _, user := range users {

		tags := make([]model.TagGetResponse, 0)
		tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
		if err != nil {
			c.String(http.StatusInternalServerError, "Cannot GetUserTagsID")
			return
		}

		for _, tagID := range tagsID {
			//タグIDからタグ名を取得する
			tagName, err := UserService.GetTagName(tagID)
			if err != nil {
				c.String(http.StatusInternalServerError, "Cannot GetTagName")
				return
			}
			tag := model.TagGetResponse{
				ID:   tagID,
				Name: tagName,
			}
			tags = append(tags, tag)
		}

		extendedUserInformationGetResponses = append(extendedUserInformationGetResponses, model.ExtendedUserInformationGetResponse{
			ID:   int64(user.ID),
			Name: user.Name,
			Tags: tags,
			Role: user.Role,
			Uuid: user.UUID,
		})
	}

	c.JSON(http.StatusOK, extendedUserInformationGetResponses)
}

// for _, user := range users {

// 	tags := make([]model.TagGetResponse, 0)
// 	tagsID, err := UserService.GetUserTagsID(int64(user.Model.ID))
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to "})
// 		return
// 	}

// 	for _, tagID := range tagsID {
// 		//タグIDからタグ名を取得する
// 		tagName, err := UserService.GetTagName(tagID)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to "})
// 			return
// 		}
// 		tag := model.TagGetResponse{
// 			ID:   tagID,
// 			Name: tagName,
// 		}
// 		tags = append(tags, tag)
// 	}

// 	userInformationGetResponse = append(userInformationGetResponse, model.UserInformationGetResponse{
// 		ID:   int64(user.ID),
// 		Name: user.Name,
// 		Tags: tags,
// 	})
// }

// c.JSON(http.StatusOK, userInformationGetResponse)

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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attendance"})
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

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// func SimultaneousStayUserList(c *gin.Context) {
// 	userID := c.Param("user_id")
// 	//int64に変換
// 	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to "})
// 		return
// 	}

// 	UserService := service.UserService{}
// 	RoomService := service.RoomService{}

// 	logs, err := RoomService.GetLogByUserAndDate(userIDInt64, 14)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to "})
// 		return
// 	}
// 	simultaneousStayUserGetResponses, err := UserService.GetSameTimeUser(logs)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to "})
// 		return
// 	}

// 	c.JSON(http.StatusOK, simultaneousStayUserGetResponses)
// }

func Check(c *gin.Context) {

	firebaseUserInfo, err := verifyCheck(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "invalid token",
		})
		return
	}

	UserService := service.UserService{}
	user, err := UserService.GetUserByEmail(firebaseUserInfo["Email"])
	if err != nil {
		fmt.Printf("Cannnot find user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	fmt.Println(user)

	CommunityService := service.CommunityService{}
	community, err := CommunityService.GetCommunityById(user.CommunityId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get community"})
		return
	}

	//メールアドレスが存在しない場合はUserは存在しないのでリクエスト失敗
	if (user == model.User{}) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "権限がありません 管理者にユーザ追加を依頼してください",
		})
		return
	}

	userRole := model.UserRoleCommunityGetResponse{
		ID:            int64(user.ID),
		Role:          user.Role,
		UUID:          user.UUID,
		Name:          user.Name,
		CommunityId:   int64(community.ID),
		CommunityName: community.Name,
	}

	c.JSON(http.StatusOK,
		userRole,
	)
}

func SignUp(c *gin.Context) {

	firebaseUserInfo, err := verifyCheck(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "invalid token",
		})
		return
	}

	UserService := service.UserService{}
	user, err := UserService.GetUserByEmail(firebaseUserInfo["Email"])
	if err != nil {
		fmt.Printf("Cannnot find user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	fmt.Println(user)

	//メールアドレスが存在しない場合はUserは存在しないのでリクエスト失敗
	if (user == model.User{}) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "権限がありません 管理者にユーザ追加を依頼してください",
		})
		return
	}

	// userRole := model.UserRoleGetResponse{
	// 	ID:   int64(user.ID),
	// 	Role: user.Role,
	// }

	c.JSON(http.StatusOK,
		user,
	)
}
