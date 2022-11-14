package main

import (
	controller "Stay_watch/controller"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Start Server")
	SetUpServer().Run(":8080")

	// userEngine := engine.Group("/user")
	// {
	// 	v1 := userEngine.Group("/v1")
	// 	{
	// 		v1.GET("/check", controller.Check)
	// 		v1.GET("/list", controller.UserList)
	// 		v1.GET("/detail", controller.Detail)
	// 		// v1.GET("/list/simultaneous/:user_id", controller.SimultaneousStayUserList)
	// 		v1.POST("/registration", controller.Register)
	// 	}
	// }

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

func SetUpServer() *gin.Engine {

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

	versionEngine := engine.Group("/v1")
	{
		versionEngine.GET("/stayers", controller.Stayer)
		versionEngine.POST("/stayers", controller.Beacon)
		versionEngine.GET("/logs", controller.Log)
		versionEngine.GET("/logs/gantt", controller.LogGantt)
		versionEngine.GET("/users", controller.UserList)
		versionEngine.POST("/users", controller.CreateUser)
		versionEngine.GET("/check", controller.Check)
		versionEngine.POST("/attendance", controller.Attendance)
	}

	return engine
}
