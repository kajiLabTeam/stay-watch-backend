package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ヘッダー取得
		authHeader := c.GetHeader("Authorization") // クライアントからAuthorizationを取得
		tokenID := strings.Replace(authHeader, "Bearer ", "", 1)
		apiKey := c.GetHeader("X-API-Key") // クライアントからAPIキーを取得

		// トークンIDとAPIキーが両方空の場合は拒否する
		if tokenID == "" && apiKey == "" {
			log.Printf("Authorization and X-API-Key are empty")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// トークンIDかAPIキーどちらかが正しければリクエストを受理する
		if tokenID != "" {
			// トークンIDの検証
			err := checkFirebaseAuth(tokenID)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else if apiKey != "" {
			// APIキーの検証
			err := checkAPIKey(apiKey)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		c.Next()
	}
}

func checkFirebaseAuth(tokenID string) error {
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
		log.Printf("Cannot initialize firebase app: %v\n", err)
		return err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Cannot initialize firebase auth: %v\n", err)
		return err
	}

	// IDトークンの検証(JWTのベリファイ)
	_, err = authClient.VerifyIDToken(ctx, tokenID)
	if err != nil {
		log.Printf("Token verify failed: %v\n", err)
		return err
	}

	return nil
}

func checkAPIKey(apiKey string) error {
	if apiKey != "abcde1" {
		err := errors.New("invalid API Key")
		log.Printf("%v", err)
		return err
	}
	return nil
}
