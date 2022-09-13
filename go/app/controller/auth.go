package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func verifyCheck(r *http.Request) (map[string]string, error) {

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	opt := option.WithCredentialsFile("/app/credentials/firebase.json")
	conf := &firebase.Config{ProjectID: "stay-watch"}
	//OAuth2.0更新トークン対応用
	app, err := firebase.NewApp(ctx, conf, opt)
	//OAuth2.0を用いない場合はconfをnilにする
	if err != nil {
		fmt.Printf("Cannot initialize firebase app: %v\n", err)
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		fmt.Printf("Cannot initialize firebase auth: %v\n", err)
	}

	header := r.Header.Get("Authorization") //クライアントからJWTを取得する
	tokenID := strings.Replace(header, "Bearer ", "", 1)
	fmt.Println(tokenID)
	//fmt.Println(token_id)
	//JWTのベリファイ
	gotToken, err := auth.VerifyIDToken(ctx, tokenID)
	if err != nil { //認証に失敗した場合(JWTが不正な場合)は、エラーを返す
		fmt.Printf("Cannot verify token_id: %v\n", err)
		return nil, err
	}

	log.Printf("Verified ID token: %v\n", gotToken)

	uid := gotToken.UID //認証に成功した場合はuidを取得する
	log.Printf("Verified user_id: %v\n", uid)

	user, err := auth.GetUser(ctx, uid)
	if err != nil {
		log.Printf("Cannot get user: %v\n", err)
		return nil, err
	} //UIDからユーザー情報を取得する(ユーザ画像，ユーザ名)
	log.Println(user.DisplayName, user.PhotoURL, user.Email)

	userData := map[string]string{
		"Name":            user.DisplayName,
		"ProfileImageURL": user.PhotoURL,
		"FirebaseUID":     uid,
		"Email":           user.Email,
	} //取得したデータを連想配列で格納し，返す
	//fmt.Println(userData)
	return userData, nil
}
