package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"Stay_watch/model"
	"Stay_watch/service"

	"github.com/gin-gonic/gin"
)

const (
	BEACON_NAME_IPHONE          = "iPhone"
	BEACON_NAME_ANDROID         = "Android"
	BEACON_NAME_FCS1301         = "FCS1301"
	BEACON_NAME_STAYWATCHBEACON = "StayWatchBeacon"
)

func Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func createUuid(communityId int64, beaconName string, userId int64, requestUuid string) string {
	// コミュニティIDを16進数3桁に変換
	communityIdHex := fmt.Sprintf("%03x", communityId)
	newUuid := ""

	if beaconName == BEACON_NAME_IPHONE {
		// iPhoneの場合「8ebc21144abd00000000ff01000 + (userId(16進数))」
		userIdHex := fmt.Sprintf("%05x", userId)
		newUuid = "8ebc21144abd00000000ff01000" + userIdHex
	} else if beaconName == BEACON_NAME_ANDROID {
		// Androidの場合ユーザIDから取得した値を用いる
		// ユーザIDを16進数5桁に変換
		userIdHex := fmt.Sprintf("%05x", userId)
		newUuid = "8ebc21144abd" + "ba0d" + "b7c6" + "ff0a" + communityIdHex + userIdHex
	} else if beaconName == BEACON_NAME_FCS1301 {
		newUuid = "8ebc21144abd" + "ba0d" + "b7c6" + "ff0f" + communityIdHex + requestUuid
	} else {
		// その他のビーコンの場合リクエストの値をそのまま用いる(StayWatchBeaconもここに含まれる)
		newUuid = requestUuid
	}

	return newUuid
}

func isValidCreateUserRequest(request model.UserCreateRequest) bool {
	if request.PrivateKey == "" && request.Uuid == "" {
		// PrivateKeyとUUIDがどちらとも未入力はNG
		fmt.Println("privateKey and uuid are null")
		return false
	}
	if request.PrivateKey != "" && request.Uuid != "" {
		// PrivateKeyとUUIDがどちらとも入力はNG
		fmt.Println("privateKey and uuid are not-null")
		return false
	}
	if request.PrivateKey != "" && len(request.PrivateKey) != 32 {
		// privateKeyは32文字固定
		fmt.Println("privateKey is 32 words")
		return false
	}
	return true
}

func isValidUpdateUserRequest(request model.UserUpdateRequest) bool {
	if request.PrivateKey != nil && request.Uuid != nil {
		// PrivateKeyとUUIDがどちらとも入力はNG
		return false
	}
	if request.PrivateKey != nil && len(*request.PrivateKey) != 32 {
		// privateKeyは32文字固定
		return false
	}
	return true
}

func CreateUser(c *gin.Context) {
	UserCreateRequest := model.UserCreateRequest{}
	c.Bind(&UserCreateRequest)

	UserService := service.UserService{}
	BeaconService := service.BeaconService{}
	TagService := service.TagService{}

	// PrivateKeyとUUIDが正しい値出ない場合BadRequestを返す
	if !isValidCreateUserRequest(UserCreateRequest) {
		fmt.Println("privateKey and uuid are not correct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request is not correct"})
		return
	}

	// 同じPrivateKeyが既に登録済みだったら409を返す
	if UserCreateRequest.PrivateKey != "" {
		isRegisterdPrivateKey, err := UserService.IsPrivateKeyAlreadyRegistered(UserCreateRequest.PrivateKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check private key is arleady registerd"})
			return
		}
		if isRegisterdPrivateKey {
			// 同じPrivateKeyが既に登録済みの場合
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Arleady Registered Private Key"})
			return
		}
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
		PrivateKey:  UserCreateRequest.PrivateKey,
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
	newUuid := createUuid(communityId, UserCreateRequest.BeaconName, registerdUserId, UserCreateRequest.Uuid)

	// UUIDを上書き
	err = UserService.UpdateUuid(newUuid, registerdUserId)
	if err != nil {
		fmt.Printf("Cannot update uuid: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update uuid"})
	}

	// タグ名からタグを取得
	tags, err := TagService.GetTagsByTagNames(UserCreateRequest.TagNames, UserCreateRequest.CommunityId)
	if err != nil {
		fmt.Printf("Cannot get tags: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	}
	// タグマップに登録
	for _, tag := range tags {
		err = TagService.CreateTagMap(&model.TagMap{UserID: registerdUserId, TagID: int64(tag.ID)})
		if err != nil {
			fmt.Printf("Cannot register tags map: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register tags map"})
		}
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

	// PrivateKeyとUUIDが正しい値出ない場合BadRequestを返す
	if !isValidUpdateUserRequest(UserUpdateRequest) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request is not correct"})
		return
	}

	beacon, err := BeaconService.GetBeaconByBeaconName(UserUpdateRequest.BeaconName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// 更新前のユーザの情報を取得
	user, err := UserService.GetUserByUserId(UserUpdateRequest.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// PrivateKeyに変更がある場合、同じPrivateKeyが既に登録済みだったら409を返す
	if UserUpdateRequest.PrivateKey != nil && user.PrivateKey != *UserUpdateRequest.PrivateKey {
		isRegisterdPrivateKey, err := UserService.IsPrivateKeyAlreadyRegistered(*UserUpdateRequest.PrivateKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check private key is arleady registerd"})
			return
		}
		if isRegisterdPrivateKey {
			// 同じPrivateKeyが既に登録済みの場合
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Arleady Registered Private Key"})
			return
		}
	}

	// メールアドレスに変更がある場合、そのメールアドレスが他のユーザに既に使われているかをチェックする処理をプラスする
	if UserUpdateRequest.Email != nil && user.Email != *UserUpdateRequest.Email {
		isRegisterdEmail, err := UserService.IsEmailAlreadyRegistered(*UserUpdateRequest.Email)
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
	newUuid := ""
	if UserUpdateRequest.Uuid != nil && UserUpdateRequest.CommunityId != nil {
		newUuid = createUuid(*UserUpdateRequest.CommunityId, UserUpdateRequest.BeaconName, UserUpdateRequest.ID, *UserUpdateRequest.Uuid)
	}

	// 値が存在するフィールドのみ更新
	user.Name = UserUpdateRequest.Name
	user.BeaconId = int64(beacon.ID)
	// nilでない場合のみ更新
	if UserUpdateRequest.Email != nil {
		user.Email = *UserUpdateRequest.Email
	}
	if UserUpdateRequest.Role != nil {
		user.Role = *UserUpdateRequest.Role
	}
	if UserUpdateRequest.CommunityId != nil {
		user.CommunityId = *UserUpdateRequest.CommunityId
	}

	// StayWatchBeaconの場合(PrivateKeyを使う場合)UUIDは""にし、そうでない場合はPrivateKeyを""にする
	if UserUpdateRequest.BeaconName == BEACON_NAME_STAYWATCHBEACON {
		user.UUID = ""
		if UserUpdateRequest.PrivateKey != nil {
			user.PrivateKey = *UserUpdateRequest.PrivateKey
		}
	} else {
		user.PrivateKey = ""
		if UserUpdateRequest.Uuid != nil {
			user.UUID = newUuid
		}
	}

	// usersテーブルにユーザ情報を保存
	err = UserService.UpdateUser(&user, UserUpdateRequest.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// タグマップ関連
	// タグマップリクエストが空でなかったらタグマップを更新
	if UserUpdateRequest.TagNames != nil {
		// タグ名からタグを取得
		tags, err := TagService.GetTagsByTagNames(UserUpdateRequest.TagNames, user.CommunityId)
		if err != nil {
			fmt.Printf("Cannot get tags: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		// tag_mapsテーブルの変更前のマップを削除
		err = TagService.DeleteTagMap(UserUpdateRequest.ID)
		if err != nil {
			fmt.Printf("Cannot delete tagMap: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tagMap"})
			return
		}
		// タグマップに登録
		for _, tag := range tags {
			err = TagService.CreateTagMap(&model.TagMap{UserID: UserUpdateRequest.ID, TagID: int64(tag.ID)})
			if err != nil {
				fmt.Printf("Cannot register tags map: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register tags map"})
			}
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
			// タグIDからタグ名を取得する
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
			// タグIDからタグ名を取得する
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
			// タグIDからタグ名を取得する
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
			// タグIDからタグ名を取得する
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
	// 構造体定義
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
	// attendaance_tmpテーブルから全てのデータを取得する
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

	// メールアドレスが存在しない場合はUserは存在しないのでリクエスト失敗
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

	// メールアドレスが存在しない場合はUserは存在しないのでリクエスト失敗
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
