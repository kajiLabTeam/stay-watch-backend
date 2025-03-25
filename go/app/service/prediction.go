package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"Stay_watch/model"
)

type PredictionService struct{}

// goroutineを使って予測結果を取得する
func (PredictionService) GetPrediction(action string, userIDs []int64, weekday int) (model.PredictionResponse, error) {
	// 予測結果を格納するチャネル
	ch := make(chan model.PredictionResult, len(userIDs))
	var rooms RoomService
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64) {
			weeks, err := rooms.GetWeeksSinceFirstLog(userID)
			if err != nil {
				return
			}
			// var logs []model.Log
			var times []time.Time
			switch action {
			case "visit":
				logs, err := rooms.GetEarliestEntryByUserAndWeekday(userID, weekday)
				if err != nil {
					ch <- model.PredictionResult{}
					return
				}
				for _, log := range logs {
					times = append(times, log.StartAt)
				}
			case "departure":
				logs, err := rooms.GetLatestExitByUserAndWeekday(userID, weekday)
				if err != nil {
					ch <- model.PredictionResult{}
					return
				}
				for _, log := range logs {
					times = append(times, log.EndAt)
				}
			default:
				ch <- model.PredictionResult{}
				return
			}
			var p PredictionService
			result, err := p.PredictTime(times, weeks)
			if err != nil {
				ch <- model.PredictionResult{}
				return
			}
			ch <- model.PredictionResult{UserID: userID, PredictionTime: result}
		}(userID)
	}
	// 予測結果を格納
	var results []model.PredictionResult
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	response := model.PredictionResponse{
		Weekday: weekday,
		Result:  results,
	}
	return response, nil
}

// pythonサーバにlogを送信して来訪する可能性の高い時刻を取得する
func (PredictionService) PredictTime(logs []time.Time, weeks int) (string, error) {
	// logからstart_atを”15:04”形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.Format("15:04"))
	}
	// pythonサーバに送信して予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/time"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for _, t := range startAt {
		q.Add("logs", t)
	}
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	u.RawQuery = q.Encode()
	// 予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return "", err
	}
	// 予測結果を返す
	return p.Time, nil
}
