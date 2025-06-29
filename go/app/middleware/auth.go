package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// RequestTimer は各リクエストの処理時間をログに出力するミドルウェアだ
func RequestTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 次のハンドラーへ処理を移す
		duration := time.Since(startTime)
		log.Printf("リクエスト %s %s の処理時間: %v", c.Request.Method, c.Request.RequestURI, duration)
	}
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Firebase App 初期化
		ctx := context.Background()
		opt := option.WithCredentialsFile("./credentials/firebase.json")
		if os.Getenv("ENVIRONMENT") == "production" {
			opt = option.WithCredentialsFile("/app/credentials/firebase.json")
		}

		conf := &firebase.Config{ProjectID: "stay-watch-a616f"}
		// OAuth2.0更新トークン対応用
		app, err := firebase.NewApp(ctx, conf, opt)
		// OAuth2.0を用いない場合はconfをnilにする
		if err != nil {
			fmt.Printf("Cannot initialize firebase app: %v\n", err)
		}
		authClient, err := app.Auth(ctx)
		if err != nil {
			log.Printf("Cannot initialize firebase auth: %v\n", err)
			log.Println("認証失敗3")
			c.Status(http.StatusUnauthorized)
		}

		// Authorization ヘッダー取得
		authHeader := c.GetHeader("Authorization") // クライアントからJWTを取得する
		tokenID := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenID == "" {
			c.Status(http.StatusUnauthorized)
		}

		// IDトークンの検証(JWTのベリファイ)
		token, err := authClient.VerifyIDToken(ctx, tokenID)
		if err != nil {
			log.Printf("Token verify failed: %v\n", err)
			log.Println("認証失敗2")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		uid := token.UID // 認証に成功した場合はuidを取得する
		user, err := authClient.GetUser(ctx, uid)
		if err != nil {
			log.Printf("Cannot get user: %v\n", err)
			log.Println("認証失敗1")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User fetch failed"})
		}

		// コンテキストに埋め込む
		c.Set("user", map[string]string{
			"Name":            user.DisplayName,
			"ProfileImageURL": user.PhotoURL,
			"FirebaseUID":     uid,
			"Email":           user.Email,
		})

		log.Println("認証成功")
		// log.Println(user.DisplayName, user.PhotoURL, user.Email)

		c.Next()
	}
}
