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
	SetUpServer().Run(":8082")
	// v1.GET("/list/simultaneous/:user_id", controller.SimultaneousStayUserList
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

	versionEngine := engine.Group("api/v1")
	{
		versionEngine.GET("/stayers", controller.Stayer)
		versionEngine.POST("/stayers", controller.Beacon)
		versionEngine.GET("/logs", controller.Log)
		versionEngine.GET("/logs/gantt", controller.LogGantt)
		// versionEngine.GET("/users", controller.UserList)
		versionEngine.GET("/users/:communityId", controller.UserList)
		versionEngine.GET("/users/extended", controller.ExtendedUserList)
		// versionEngine.POST("/users", controller.PastCreateUser)
		versionEngine.POST("/users", controller.CreateUser)
		versionEngine.GET("/check", controller.Check)
		versionEngine.POST("/attendance", controller.Attendance)
		versionEngine.GET("/rooms/:communityID", controller.GetRoomsByCommunityID)
		versionEngine.PUT("/rooms", controller.UpdateRoom)
		// versionEngine.POST("/updateroom", controller.UpdateRoom)
		versionEngine.GET("/buildings/editor", controller.GetBuildingsEditor)
		versionEngine.GET("/signup", controller.SignUp)
	}

	return engine
}
