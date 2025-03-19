package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"Stay_watch/model"
)

type ProbabilityService struct{}

// goroutineを使って予測結果を取得する
func (ProbabilityService) GetVisitProbability(userIDs []int64, weekday int, time string, isForward bool) ([]model.ProbabilityResponse, error) {
	// 予測結果を格納するチャネル
	var wg sync.WaitGroup
	wg.Add(len(userIDs))
	ch := make(chan model.ProbabilityResponse, len(userIDs))
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64, wg *sync.WaitGroup) {
			defer wg.Done()
			var p ProbabilityService
			result, err := p.PredictVisitProbability(userID, weekday, time, isForward)
			if err != nil {
				ch <- model.ProbabilityResponse{}
				return
			}
			ch <- result
		}(userID, &wg)
	}
	go func() {
		wg.Wait()
		close(ch) // すべての Goroutine が終了したらチャネルを閉じる
	}()
	// 予測結果を格納
	var results []model.ProbabilityResponse
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	return results, nil
}

// pythonサーバにlogを送信してtimeまでに(or以降に)来訪する確率の予測結果を取得する
func (ProbabilityService) PredictVisitProbability(userID int64, weekday int, time string, isForward bool) (model.ProbabilityResponse, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetEarliestEntryByUserAndWeekday(userID, weekday)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	// logsからstart_atを"15:04"形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.StartAt.Format("15:04"))
	}

	// user_IDを持つuserが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}

	// pythonサーバにstartAtとweeksを送信してtimeまでに来訪する確率の予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/probability/visit"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	q := u.Query()
	for _, t := range startAt {
		q.Add("start_at", t)
	}
	q.Add("time", time)
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	q.Add("is_forward", fmt.Sprintf("%t", isForward))
	u.RawQuery = q.Encode()
	// GETリクエストを送信して予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	defer res.Body.Close()
	// 予測結果を取得
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return model.ProbabilityResponse{}, err
	}
	// 予測結果を返す
	result := model.ProbabilityResponse{
		UserID:      userID,
		Weekday:     weekday,
		Time:        time,
		IsForward:   isForward,
		Probability: p.Probability,
	}
	return result, nil
}

// goroutineを使って予測結果を取得する
func (ProbabilityService) GetDepartureProbability(userIDs []int64, weekday int, time string, isForward bool) ([]model.ProbabilityResponse, error) {
	// 予測結果を格納するチャネル
	var wg sync.WaitGroup
	wg.Add(len(userIDs))
	ch := make(chan model.ProbabilityResponse, len(userIDs))
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64, wg *sync.WaitGroup) {
			defer wg.Done()
			var p ProbabilityService
			result, err := p.PredictDepartureProbability(userID, weekday, time, isForward)
			if err != nil {
				ch <- model.ProbabilityResponse{}
				return
			}
			ch <- result
		}(userID, &wg)
	}
	go func() {
		wg.Wait()
		close(ch) // すべての Goroutine が終了したらチャネルを閉じる
	}()
	// 予測結果を格納
	var results []model.ProbabilityResponse
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	return results, nil
}

// pythonサーバにlogを送信してtimeまでに(or以降に)帰宅する確率の予測結果を取得する
func (ProbabilityService) PredictDepartureProbability(userID int64, weekday int, time string, isForward bool) (model.ProbabilityResponse, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetLatestExitByUserAndWeekday(userID, weekday)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	// logsからstart_atを"15:04"形式に変換してスライスに格納
	// start_atの日付とend_atの日付が異なる場合はスルーする
	var endAt []string
	for _, log := range logs {
		if log.StartAt.Day() != log.EndAt.Day() {
			continue
		}
		endAt = append(endAt, log.EndAt.Format("15:04"))
	}

	// user_IDを持つuserが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}

	// pythonサーバにstartAtとweeksを送信してtimeまでに来訪する確率の予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/probability/departure"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	q := u.Query()
	for _, t := range endAt {
		q.Add("end_at", t)
	}
	q.Add("time", time)
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	q.Add("is_forward", fmt.Sprintf("%t", isForward))
	u.RawQuery = q.Encode()
	// GETリクエストを送信して予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	defer res.Body.Close()
	// 予測結果を取得
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return model.ProbabilityResponse{}, err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return model.ProbabilityResponse{}, err
	}
	// 予測結果を返す
	result := model.ProbabilityResponse{
		UserID:      userID,
		Weekday:     weekday,
		Time:        time,
		IsForward:   isForward,
		Probability: p.Probability,
	}
	return result, nil
}
