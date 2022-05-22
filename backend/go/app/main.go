package main

import (
	controller "Stay_watch/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	// ミドルウェア
	// engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://stay-watch-front.vercel.app", "https://stay-watch-go.kajilab.tk", "https://stay-watch-front.*happy663.vercel.app"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	roomEngine := engine.Group("/room")
	{
		v1 := roomEngine.Group("/v1")
		{
			v1.GET("/stayer", controller.Stayer)
			v1.GET("/log", controller.Log)
			v1.POST("/beacon", controller.Beacon)
			v1.GET("/list/simultaneous/:user_id", controller.SimultaneousList)
		}
	}
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.GET("/list", controller.List)
			v1.GET("/detail", controller.Detail)
			v1.POST("/register", controller.Register)
			v1.POST("/attendance", controller.Attendance)
		}
	}

	engine.Run(":8080")
}

// package main

// func main() {

// 	UserService := service.UserService{}
// 	UserService.TemporarilySavedAttendance(1, 1)

// }

// func main() {

// 	doPeriodically()

// 	// log.SetFlags(log.Lmicroseconds)
// 	// // ticker := time.NewTicker(time.Millisecond * 1000)
// 	// ticker := time.NewTicker(time.Hour * 12)
// 	// defer ticker.Stop()
// 	// count := 0
// 	// for {
// 	// 	select {
// 	// 	case <-ticker.C:
// 	// 		curentTime := time.Now()

// 	// 		//曜日が火曜日の場合
// 	// 		if curentTime.Weekday() == time.Tuesday {
// 	// 			log.Println("火曜日です")
// 	// 			//現在時刻が午後だったら
// 	// 			if curentTime.Hour() >= 12 {
// 	// 				count++
// 	// 				log.Println("count:", count)
// 	// 			}
// 	// 		}

// 	// 	}
// 	// }
// }

// func doPeriodically() {
// 	/* do something */
// 	RoomService := service.RoomService{}

// 	currentTime := time.Now().Format("2006-01-02 15:04:05")
// 	currenDay := currentTime[:10]
// 	fmt.Println("currentDay:", currenDay)
// 	startAt := fmt.Sprintf("%s 07:00:00", currenDay)
// 	endAt := fmt.Sprintf("%s 12:00:00", currenDay)
// 	fmt.Println("startAt:", startAt)
// 	fmt.Println("endAt:", endAt)
// 	logs, err := RoomService.GetLogsFromStartAtAndEntAt(startAt, endAt)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(logs)

// 	//火曜9時〜12時のログのUserIDsを取得
// 	logUserIDs := make([]string, 0)
// 	for _, log := range logs {
// 		logUserIDs = append(logUserIDs, log.UserID)
// 	}
// 	util := util.Util{}
// 	logUniqueUserIDs := util.SliceUniqueString(logUserIDs)

// 	UserService := service.UserService{}

// 	//全てのユーザーリストの構造体を取得
// 	allUser, err := UserService.GetAllUser()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//全てのuserIDのみ取得する
// 	allUserID := make([]string, 0)
// 	for _, user := range allUser {
// 		allUserID = append(allUserID, user.ID)
// 	}

// 	for _, userID := range allUserID {
// 		if util.ArrayStringContains(logUniqueUserIDs, userID) {
// 			fmt.Println("true=userID:", userID)
// 			UserService.RegisterAttendance(userID, currenDay, true)
// 		} else {
// 			fmt.Println("flase=userID:", userID)
// 			UserService.RegisterAttendance(userID, currenDay, false)
// 		}
// 	}

// }
