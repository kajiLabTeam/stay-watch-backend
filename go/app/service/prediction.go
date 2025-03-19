package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"Stay_watch/model"
)

type PredictionService struct{}

// goroutineを使って予測結果を取得する
func (PredictionService) GetVisitPrediction(userIDs []int64, weekday int, time string, isForward bool) (model.PredictionResponse, error) {
	// 予測結果を格納するチャネル
	var ch = make(chan model.PredictionResult, len(userIDs))
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64) {
			var p PredictionService
			result, err := p.PredictVisitTime(userID, weekday, time, isForward)
			if err != nil {
				ch <- model.PredictionResult{}
				return
			}
			ch <- result
		}(userID)
	}
	// 予測結果を格納
	var results []model.PredictionResult
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	responce := model.PredictionResponse{
		Weekday:   weekday,
		Time:      time,
		IsForward: isForward,
		Result:    results,
	}
	return responce, nil
}

// pythonサーバにlogを送信して来訪する可能性の高い時刻を取得する
func (PredictionService) PredictVisitTime(userID int64, weekday int, time string, isForward bool) (model.PredictionResult, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetEarliestEntryByUserAndWeekday(userID, weekday)
	if err != nil {
		return model.PredictionResult{}, err
	}
	// logからstart_atを”15:04”形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.StartAt.Format("15:04"))
	}
	// userが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return model.PredictionResult{}, err
	}
	// pythonサーバに送信して予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/time/visit"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return model.PredictionResult{}, err
	}
	q := u.Query()
	for _, t := range startAt {
		q.Add("start_at", t)
	}
	q.Add("time", time)
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	q.Add("is_forward", fmt.Sprintf("%t", isForward))
	u.RawQuery = q.Encode()
	// 予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return model.PredictionResult{}, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return model.PredictionResult{}, err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return model.PredictionResult{}, err
	}
	// 予測結果を返す
	result := model.PredictionResult{
		UserID:         userID,
		PredictionTime: p.Time,
	}
	return result, nil
}

// goroutineを使って予測結果を取得する
func (PredictionService) GetDeparturePrediction(userIDs []int64, weekday int, time string, isForward bool) (model.PredictionResponse, error) {
	// 予測結果を格納するチャネル
	var ch = make(chan model.PredictionResult, len(userIDs))
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64) {
			var p PredictionService
			result, err := p.PredictDepartureTime(userID, weekday, time, isForward)
			if err != nil {
				ch <- model.PredictionResult{}
				return
			}
			ch <- result
		}(userID)
	}
	// 予測結果を格納
	var results []model.PredictionResult
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	responce := model.PredictionResponse{
		Weekday:   weekday,
		Time:      time,
		IsForward: isForward,
		Result:    results,
	}
	return responce, nil
}

// pythonサーバにlogを送信して帰宅する可能性の高い時刻を取得する
func (PredictionService) PredictDepartureTime(userID int64, weekday int, time string, isForward bool) (model.PredictionResult, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetLatestExitByUserAndWeekday(userID, weekday)
	if err != nil {
		return model.PredictionResult{}, err
	}
	// logからend_atを”15:04”形式に変換してスライスに格納
	var endAt []string
	for _, log := range logs {
		endAt = append(endAt, log.EndAt.Format("15:04"))
	}
	// userが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return model.PredictionResult{}, err
	}
	// pythonサーバに送信して予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/time/departure"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return model.PredictionResult{}, err
	}
	q := u.Query()
	for _, t := range endAt {
		q.Add("end_at", t)
	}
	q.Add("time", time)
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	q.Add("is_forward", fmt.Sprintf("%t", isForward))
	u.RawQuery = q.Encode()
	// 予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return model.PredictionResult{}, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return model.PredictionResult{}, err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return model.PredictionResult{}, err
	}
	// 予測結果を返す
	result := model.PredictionResult{
		UserID:         userID,
		PredictionTime: p.Time,
	}
	return result, nil
}