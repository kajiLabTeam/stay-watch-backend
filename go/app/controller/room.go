package controller

import (
	"Stay_watch/model"
	"Stay_watch/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoreRoom(c *gin.Context) {

	// type Room struct {
	// 	Name string `json:"room_name"`
	// }
	// var room Room
	// err := c.Bind(&room)

	RoomForm := model.RoomEditorForm{}
	err := c.Bind(&RoomForm)

	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("ControllerのStoreRoom関数まで来たよ")
	fmt.Println(c)
	//fmt.Println(RoomEditorForm.Name)
	fmt.Println(RoomForm.Name)
	fmt.Println(RoomForm.Points)




	RoomService := service.RoomService{}
	sample_room := model.TmpRoom{
		// Name: "sample",
		Name:  RoomForm.Name,
		BuildingID: 2,
		CommunityID: 2,
		Polygon: RoomForm.Points,
	}
	RoomService.CreateSampleRoom(&sample_room)

	// //RoomService.CreatgeLog()
	// sampledata, sampleerr := RoomService.CreateSampleRoom()
	// if sampleerr != nil{
	// 	fmt.Printf("failed: Cannnot get stayer %v", sampleerr)
	// 	c.String(http.StatusInternalServerError, "Server Error")
	// 	return
	// }


	// allStayer, err := RoomService.GetAllStayer()
	// if err != nil {

	// 	fmt.Printf("failed: Cannnot get stayer %v", err)
	// 	c.String(http.StatusInternalServerError, "Server Error")
	// 	return
	// }

	// fmt.Println(allStayer)
	// fmt.Println(sampledata)

	//UserService := service.UserService{}

	
	// Tmp_roomService := service.Tmp_roomService{}
	// allStayer, err := Tmp_roomService.GetAllStayer()
	// if err != nil {

	// 	fmt.Printf("failed: Cannnot get stayer %v", err)
	// 	c.String(http.StatusInternalServerError, "Server Error")
	// 	return
	// }
//--------見本------------
	// RegistrationUserForm := model.RegistrationUserForm{}
	// c.Bind(&RegistrationUserForm)

	

	// UserService := service.UserService{}
	// //userIDがないなら新規登録
	// if RegistrationUserForm.ID == 0 {
	// 	user := model.User{
	// 		// Name:  RegistrationUserForm.Name,
	// 		Name: "usergo",
	// 		Email: RegistrationUserForm.Email,
	// 		Role:  RegistrationUserForm.Role,
	// 		UUID:  UserService.NewUUID(),
	// 	}

	// 	err := UserService.RegisterUser(&user)
	// 	if err != nil {
	// 		fmt.Printf("Cannnot register user: %v", err)
	// 		c.String(http.StatusInternalServerError, "Server Error")
	// 		return
	// 	}
	// }

	c.JSON(http.StatusCreated, gin.H{
		"status": "okb",
	})
}

func GetRoomsByCommunityID(c *gin.Context) {
	fmt.Printf("aaa")
	fmt.Println("コミュニティID")
	communityID, _ := strconv.ParseInt(c.Param("communityID"), 10, 64)	// string -> int64

	RoomService := service.RoomService{}

	rooms, err := RoomService.GetAllRooms()	// この関数と上の関数は別物
	if err != nil {
		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	fmt.Println(rooms)
	// fmt.Println(rooms.RoomID)

	roomsGetResponse := []model.RoomsGetResponse{}	// json.goのやつ

	for _, room := range rooms {
		// db.goの形式からjson.goの形式へ
		// db = TmpRoom[Name, BuildingID, CommunityID, Polygon]
		// json RoomsGetResponse[Names]

		fmt.Println("部屋の名前")
		fmt.Print(room.Name)	// db.goのやつ
		fmt.Println(room.CommunityID)

		if(room.CommunityID == communityID){
			// roomName, err := room.CommunityID
			// if err != nil {
			// 	c.String(http.StatusInternalServerError, "Server Error")
			// 	return
			// }
			roomName := room.Name
			roomsGetResponse = append(roomsGetResponse, model.RoomsGetResponse{
				Names: roomName,
			})
		}
	}
	c.JSON(200, roomsGetResponse)

	fmt.Println("フロントへ渡す内容")
	fmt.Println(roomsGetResponse)
	// fmt.Printf(rooms)
}


func Stayer(c *gin.Context) {

	RoomService := service.RoomService{}
	UserService := service.UserService{}

	//Stayerテーブルから全てのデータを取得する
	allStayer, err := RoomService.GetAllStayer()
	if err != nil {

		fmt.Printf("failed: Cannnot get stayer %v", err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	stayerGetResponse := []model.StayerGetResponse{}

	for _, stayer := range allStayer {

		userName, err := UserService.GetUserNameByUserID(stayer.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomNameByRoomID(stayer.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		tagsGetResponse := make([]model.TagGetResponse, 0)

		tagsID, err := UserService.GetUserTagsID(stayer.UserID)
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
			tag := model.TagGetResponse{
				ID:   tagID,
				Name: tagName,
			}
			tagsGetResponse = append(tagsGetResponse, tag)
		}

		stayerGetResponse = append(stayerGetResponse, model.StayerGetResponse{
			ID:     stayer.UserID,
			Name:   userName,
			Room:   roomName,
			RoomID: int(stayer.RoomID),
			Tags:   tagsGetResponse,
		})
	}
	c.JSON(200, stayerGetResponse)
}

func Log(c *gin.Context) {
	RoomService := service.RoomService{}
	UserService := service.UserService{}

	//ページング処理
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	//ページごとにLogテーブルからデータを取得する
	pageLog, err := RoomService.GetLogsByPage(pageInt)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	logGetResponse := []model.LogGetResponse{}

	for _, log := range pageLog {

		userName, err := UserService.GetUserNameByUserID(log.UserID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		roomName, err := RoomService.GetRoomNameByRoomID(log.RoomID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		logGetResponse = append(logGetResponse, model.LogGetResponse{
			ID:      int64(log.ID),
			Name:    userName,
			Room:    roomName,
			StartAt: log.StartAt.Format("2006-01-02 15:04:05"),
			EndAt:   log.EndAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(200, logGetResponse)
}

// func SimultaneousList(c *gin.Context) {
// 	userID := c.Param("user_id")

// 	RoomService := service.RoomService{}

// 	//userIDをint64に変換
// 	userIDInt, err := strconv.ParseInt(userID, 10, 64)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Bad Request")
// 		return
// 	}

// 	SimultaneousList, err := RoomService.GetSimultaneousList(userIDInt)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}

// 	c.JSON(http.StatusOK, SimultaneousList)
// }

func LogGantt(c *gin.Context) {

	RoomService := service.RoomService{}
	GanttLogs, err := RoomService.GetGanttLog()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusOK, GanttLogs)
}
