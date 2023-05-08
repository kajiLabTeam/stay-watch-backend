// main_test.go
package main

import (
	controller "Stay_watch/controller"
	"Stay_watch/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostStayer(t *testing.T) {
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)

	beaconsRoom := model.BeaconRoom{
		Beacons: []*model.BeaconSignal{
			{
				Uuid: "e7d61ea3f8dd49c88f2ff2484c07ac00",
				Rssi: -60,
			},
		},
		RoomID: 1,
	}
	//jsonに変換
	jsonBeaconsRoom, err := json.Marshal(beaconsRoom)
	if err != nil {
		t.Fatal(err)
	}

	// リクエスト情報をコンテキストに入れる
	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/stayers", nil)
	ginContext.Request.Header.Set("Content-Type", "application/json")
	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBeaconsRoom))
	controller.Beacon(ginContext)
	asserts := assert.New(t)
	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusCreated, response.Code)

	fmt.Println("TestPostStayer通過")
}

func TestGetStayer(t *testing.T) {
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)

	// リクエストの生成
	// 今回はmiddlewareのテストのためpathはなんでも可
	req, _ := http.NewRequest(http.MethodGet, "/stayers", nil)

	// リクエスト情報をコンテキストに入れる
	ginContext.Request = req
	controller.Stayer(ginContext)

	asserts := assert.New(t)

	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusOK, response.Code)
	// レスポンスのボディを構造体に変換
	var responseStayer []model.StayerGetResponse
	json.Unmarshal(response.Body.Bytes(), &responseStayer)
	// レスポンスのボディの確認
	asserts.Equal("kaji", responseStayer[0].Name)
	asserts.Equal("B1", responseStayer[0].Tags[0].Name)
	asserts.Equal(2, int(responseStayer[0].Tags[0].ID))
	asserts.Equal(2, int(responseStayer[0].ID))
	asserts.Equal(1, int(responseStayer[0].RoomID))

	fmt.Println("TestGetStayer通過")
}

func TestGetLog(t *testing.T) {
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	// リクエストの生成
	// 今回はmiddlewareのテストのためpathはなんでも可
	req, _ := http.NewRequest(http.MethodGet, "/logs", nil)
	// リクエスト情報をコンテキストに入れる
	ginContext.Request = req
	controller.Log(ginContext)
	asserts := assert.New(t)
	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusOK, response.Code)
	// レスポンスのボディを構造体に変換
	var responseLog []model.LogGetResponse
	json.Unmarshal(response.Body.Bytes(), &responseLog)
	// レスポンスのボディの確認
	asserts.Equal("kaji", responseLog[0].Name)
	asserts.Equal("梶研-学生部屋", responseLog[0].Room)
	// asserts.Equal(1, int(responseLog[0].ID))
	// asserts.Equal("2021-05-01 00:00:00", responseLog[0].StartAt)
	// asserts.Equal("2021-05-01 00:00:00", responseLog[0].EndAt)

	fmt.Println("TestGetLog通過")
}

func TestGetUser(t *testing.T) {
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	// リクエストの生成
	// 今回はmiddlewareのテストのためpathはなんでも可
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	// リクエスト情報をコンテキストに入れる
	ginContext.Request = req
	controller.PastUserList(ginContext)
	asserts := assert.New(t)
	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusOK, response.Code)
	// レスポンスのボディを構造体に変換
	var responseUser []model.UserInformationGetResponse
	json.Unmarshal(response.Body.Bytes(), &responseUser)
	// レスポンスのボディの確認
	//fmt.Println(responseUser)
	asserts.Equal("test", responseUser[0].Name)
	asserts.Equal("テストタグ", responseUser[0].Tags[0].Name)
	asserts.Equal(1, int(responseUser[0].ID))

	fmt.Println("TestGetUser通過")
}

// 管理者画面でのユーザ取得API
func TestGetEditorUser(t *testing.T) {

	router := gin.Default()
	router.GET("/api/v1/admin/users/:communityId", controller.AdminUserList)

	asserts := assert.New(t)

	lastSumUsers := 0
	isAllSumEqual := true

	for i := 0; i < 10; i++ {
		// HTTPリクエストの生成
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/users/"+strconv.Itoa(i), nil)

		// レスポンスのレコーダーを作成
		res := httptest.NewRecorder()

		// リクエストをハンドル
		router.ServeHTTP(res, req)
		// レスポンスのステータスコードの確認
		asserts.Equal(http.StatusOK, res.Code)

		// レスポンスのボディを構造体に変換
		var responseUser []model.UserEditorResponse
		json.Unmarshal(res.Body.Bytes(), &responseUser)

		// community_id=1 のテストユーザの情報が正しく取れているか確認
		if i == 1 {
			asserts.Equal("test", responseUser[0].Name)
			asserts.Equal("テストタグ", responseUser[0].Tags[0].Name)
			asserts.Equal("a7d61ea3f8dd49c88f2ff2484c07ac00", responseUser[0].Uuid)
			asserts.Equal("", responseUser[0].Email)
			asserts.Equal("FCS1301", responseUser[0].BeaconName)
		}

		if lastSumUsers != len(responseUser) {
			isAllSumEqual = false
		}
	}
	if isAllSumEqual {
		// 全てユーザ数が同じ場合は正常なら存在しないため
		// community_idによる絞り込みができていないか、データがそもそも取れていないかなど
		t.Fatalf("All community users have the same count")
	}
	fmt.Println("TestGetEditorUser通過")
}

func TestGetTagNames(t *testing.T) {

	router := gin.Default()
	router.GET("/api/v1/tags/:communityId/names", controller.GetTagNamesByCommunityId)

	asserts := assert.New(t)

	lastSumUsers := 0
	isAllSumEqual := true

	for i := 0; i < 10; i++ {
		// HTTPリクエストの生成
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/tags/"+strconv.Itoa(i)+"/names", nil)

		// レスポンスのレコーダーを作成
		res := httptest.NewRecorder()

		// リクエストをハンドル
		router.ServeHTTP(res, req)
		// レスポンスのステータスコードの確認
		asserts.Equal(http.StatusOK, res.Code)

		// レスポンスのボディを構造体に変換
		var responseTagNames []model.TagsNamesGetResponse
		json.Unmarshal(res.Body.Bytes(), &responseTagNames)

		// community_id=1 のテストタグの情報が正しく取れているか確認
		if i == 1 {
			asserts.Equal("B1", responseTagNames[0].Name)
		}

		if lastSumUsers != len(responseTagNames) {
			isAllSumEqual = false
		}
	}
	if isAllSumEqual {
		// 全てユーザ数が同じ場合は正常なら存在しないため
		// community_idによる絞り込みができていないか、データがそもそも取れていないかなど
		t.Fatalf("All community tags have the same count")
	}

	fmt.Println("TestGetTagNames通過")
}

func TestGetBeacons(t *testing.T) {

	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)

	// リクエストの生成
	// 今回はmiddlewareのテストのためpathはなんでも可
	req, _ := http.NewRequest(http.MethodGet, "/beacons", nil)

	// リクエスト情報をコンテキストに入れる
	ginContext.Request = req
	controller.GetBeacon(ginContext)

	asserts := assert.New(t)

	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusOK, response.Code)

	// レスポンスのボディを構造体に変換
	var responseBeacon []model.BeaconGetResponse

	json.Unmarshal(response.Body.Bytes(), &responseBeacon)
	// レスポンスのボディの確認
	asserts.Equal("FCS1301", responseBeacon[0].BeaconName)

	fmt.Println("TestGetBeacon通過")
}

func TestGetCommunityByUserId(t *testing.T) {

	router := gin.Default()
	router.GET("/api/v1/communities/:userId", controller.GetCommunityByUserIdHandler)

	asserts := assert.New(t)

	// HTTPリクエストの生成
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/communities/"+strconv.Itoa(1), nil)

	// レスポンスのレコーダーを作成
	res := httptest.NewRecorder()

	// リクエストをハンドル
	router.ServeHTTP(res, req)
	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusOK, res.Code)

	// レスポンスのボディを構造体に変換
	var responseCommunity model.CommunityGetResponse
	json.Unmarshal(res.Body.Bytes(), &responseCommunity)

	asserts.Equal("テスト研究室", responseCommunity.Name)

	fmt.Println("TestGetCommunityByUserId通過")
}

// ユーザの新規登録API
// 未登録の場合
func TestCreateUser(t *testing.T) {

	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)

	// 未登録ユーザ
	unregisteredUser := model.UserCreateRequest{
		Name:        "tarou",
		Uuid:        "",
		Email:       "toge7113+unregistaserd@gmail.com",
		Role:        1,
		CommunityId: 1,
		BeaconName:  "FCS1301",
		TagIds:      []int64{1, 2, 5},
	}
	jsonUnegisteredUser, err := json.Marshal(unregisteredUser)
	if err != nil {
		t.Fatal(err)
	}
	// リクエスト情報をコンテキストに入れる
	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/users", nil)
	ginContext.Request.Header.Set("Content-Type", "application/json")
	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonUnegisteredUser))
	controller.CreateUser(ginContext)
	asserts := assert.New(t)
	fmt.Println(response)
	// レスポンスのステータスコードの確認(未登録ユーザ)
	asserts.Equal(http.StatusCreated, response.Code)

	fmt.Println("TestCreateUser通過")
}

// ユーザの更新API
func TestPutUser(t *testing.T) {

	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)

	user := model.UserUpdateRequest{
		ID:          10,
		Name:        "test",
		Uuid:        "e7d61ea3f8dd49c88f2ffaefae484c07ab00",
		Email:       "toge7113+test-put-user@gmail.com",
		Role:        1,
		CommunityId: 1,
		BeaconName:  "FCS1301",
		TagIds:      []int64{1, 2, 5},
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	// リクエスト情報をコンテキストに入れる
	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/users", nil)
	ginContext.Request.Header.Set("Content-Type", "application/json")
	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonUser))
	controller.UpdateUser(ginContext)
	asserts := assert.New(t)
	// レスポンスのステータスコードの確認
	asserts.Equal(http.StatusCreated, response.Code)
	fmt.Println("TestPutUser通過")
}

func TestDeleteUser(t *testing.T) {

	router := gin.Default()
	router.DELETE("/api/v1/users/:userId", controller.DeleteUser)

	asserts := assert.New(t)

	for i := 1; i < 5; i++ {
		// HTTPリクエストの生成
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/"+strconv.Itoa(i), nil)

		// レスポンスのレコーダーを作成
		res := httptest.NewRecorder()

		// リクエストをハンドル
		router.ServeHTTP(res, req)
		// レスポンスのステータスコードの確認
		asserts.Equal(http.StatusCreated, res.Code)

		// レスポンスのボディを構造体に変換
		var responseUser []model.UserEditorResponse
		json.Unmarshal(res.Body.Bytes(), &responseUser)

	}

	fmt.Println("TestDeleteUser通過")
}

// func TestCreateUserUnRegistered(t *testing.T) {

// 	response := httptest.NewRecorder()
// 	ginContext, _ := gin.CreateTestContext(response)

// 	// 未登録ユーザ
// 	unregisteredUser := model.UserCreateRequest{
// 		Name:        "tarou",
// 		Uuid:        "",
// 		Email:       "toge711312121212@gmail.com",
// 		Role:        1,
// 		CommunityId: 1,
// 		BeaconName:  "FCS1301",
// 		TagIds:      []int64{1, 2, 5},
// 	}
// 	jsonUnegisteredUser, err := json.Marshal(unregisteredUser)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// リクエスト情報をコンテキストに入れる
// 	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/users", nil)
// 	ginContext.Request.Header.Set("Content-Type", "application/json")
// 	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonUnegisteredUser))
// 	controller.CreateUser(ginContext)
// 	asserts := assert.New(t)
// 	fmt.Println(response)
// 	// レスポンスのステータスコードの確認(未登録ユーザ)
// 	asserts.Equal(http.StatusCreated, response.Code)
// }
