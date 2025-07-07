// Package middleware provides middleware handlers for authentication and other HTTP middleware features.
package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
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
	ctx := context.Background()

	// IDトークンの検証(JWTのベリファイ)
	_, err := firebaseAuth.VerifyIDToken(ctx, tokenID)
	if err != nil {
		log.Printf("Token verify failed: %v\n", err)
		return err
	}

	return nil
}

func checkAPIKey(apiKey string) error {
	if apiKey != os.Getenv("STAYWATCH_API_KEY") {
		err := errors.New("invalid API Key")
		log.Printf("%v", err)
		return err
	}
	return nil
}
