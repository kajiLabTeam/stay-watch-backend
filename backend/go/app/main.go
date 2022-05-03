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
		AllowOrigins: []string{"http://localhost:3000", "https://stay-watch-front.vercel.app", "https://stay-watch-go.kajilab.tk"},
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
		}
	}
	engine.Run(":8080")
}

// package main

// import (
// 	"Stay_watch/service"
// 	"fmt"
// )

// func main() {
// 	RoomService := service.RoomService{}

// 	times, err := RoomService.GetTimesFromStartAtAndEntAt("2022-04-20 08:00:00", "2022-04-20 10:00:00")
// 	if err != nil {
// 		fmt.Println("err=", err)
// 	}
// 	fmt.Println("times=", times)
// }
