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
		AllowOrigins: []string{"*", "http://localhost:3000", "https://stay-watch-front.vercel.app", "https://stay-watch-go.kajilab.tk"},
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
			v1.GET("/log/gantt", controller.LogGantt)
			v1.POST("/beacon", controller.Beacon)
			v1.GET("/list/simultaneous/:user_id", controller.SimultaneousList)
		}
	}
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.GET("/check", controller.Check)
			v1.GET("/list", controller.UserList)
			v1.GET("/detail", controller.Detail)
			v1.GET("/list/simultaneous/:user_id", controller.SimultaneousStayUserList)
			v1.POST("/registration", controller.Register)
			v1.POST("/attendance", controller.Attendance)
		}
	}
	engine.Run(":8080")

	// BotService := service.BotService{}
	// //2週間に一度定期的実行
	// ticker := time.NewTicker(time.Hour * 24 * 14)
	// defer ticker.Stop()
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		BotService.NotifyOutOfBattery()
	// 	}
	// }

}
