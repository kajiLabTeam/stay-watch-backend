package service

import (
	"Stay_watch/model"
	"Stay_watch/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

type BotService struct{}

// BOTにメッセージを送信する
func (BotService) SendMessage(message string, channelID string) error {

	requestBody := &model.RequestBody{
		Text: message,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}

	endpoint := "https://hooks.slack.com/services/T04DMQ6PF/" + channelID
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}

	fmt.Printf("%#v", string(byteArray))

	return nil
}

// Botに電池切れのメッセージを送信する
func (BotService) NotifyOutOfBattery() error {
	RoomService := RoomService{}
	UserService := UserService{}
	UtilService := util.Util{}
	BotService := BotService{}

	//二週間分のログを取り出す
	logs, err := RoomService.GetLogWithinDate(14)
	if err != nil {
		fmt.Println(err)
	}

	usersName := make([]string, 0)
	for _, log := range logs {
		userID := log.UserID
		//ログのユーザIDからユーザ名を取得する
		userName, err := UserService.GetUserNameByUserID(userID)
		if err != nil {
			fmt.Println(err.Error())
		}
		usersName = append(usersName, userName)
	}
	uniqueUsersName := UtilService.SliceUniqueString(usersName)
	fmt.Println(uniqueUsersName)

	allUserName, err := UserService.GetAllUserName()
	if err != nil {
		fmt.Println(err.Error())
	}

	//電池切れがおそらくいない場合
	if len(uniqueUsersName) == len(allUserName) {
		message := "電池が切れている人がいない可能性が高いです"
		BotService.SendMessage(message, "B01QE009PPG/eZlyUIomDyqCcOfV8hcLV7AH")
		return nil
	}

	slackIdMap := map[string]string{"kaji": "U04DMQ6PX", "ogane": "U4YMSSLHY", "miyagawa-san": "UJ9J3SD25", "ken": "U036KC23XKJ", "suzaki": "U015235ME74", "toyama": "U021V7TDUBT", "kameda": "U021NPR1UDS", "akito": "U021DGMJWBF", "ohashi": "U021V7N9H33", "ukai": "U021GS24PEF", "maruyama": "U021DGPQU05", "rui": "U021V7NTR09", "fueta": "U0226E3CP96", "terada": "U021VQ4EEM7", "ayato": "U014VB0CKT8", "shamo": "U021NPWG2EQ", "ao": "U03EJ9WH73Q", "fuma": "U03EHR1M0MU", "isiguro": "U02D8MEFYHF", "iwaguti": "U03EMS4EEE8", "kazuo": "U03E7MAS8BH", "oiwa": "U03E7RCGKC7", "sakai": "U03ETK8GELU", "togawa": "U03ECGSKU78", "ueji": "U03EKCJR5CM", "yada": "U03F17PUCE5", "yokoyama": "U03EABEATST", "makino": "U03ENAGMXLL"}

	//全てのユーザと2週間以内に滞在したユーザの差分を求める
	for _, userName := range allUserName {
		//userNameがuniqueUsersNameに含まれていない場合、電池切れのメッセージを送信する
		if !UtilService.ArrayStringContains(uniqueUsersName, userName) {
			fmt.Println(userName)
			slackId := slackIdMap[userName]
			message := fmt.Sprintf("<@%s>", slackId)
			BotService.SendMessage(message, "B01QE009PPG/eZlyUIomDyqCcOfV8hcLV7AH")
		}
	}

	err = BotService.SendMessage("2週間以内の滞在履歴がありません 電池切れの可能性が高いです", "B01QE009PPG/eZlyUIomDyqCcOfV8hcLV7AH")
	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}

	return nil
}
