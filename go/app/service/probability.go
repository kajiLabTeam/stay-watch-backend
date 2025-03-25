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
func (ProbabilityService) GetProbability(action string, userIDs []int64, weekday int, time string, isForward bool) (model.ProbabilityResponse, error) {
	room := RoomService{}
	// 予測結果を格納するチャネル
	var wg sync.WaitGroup
	wg.Add(len(userIDs))
	ch := make(chan model.ProbabilityResult, len(userIDs))
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64, wg *sync.WaitGroup) {
			weeks, err := room.GetWeeksSinceFirstLog(userID)
			if err != nil {
				ch <- model.ProbabilityResult{}
				return
			}
			var logs []model.Log
			switch action {
			case "visit":
				logs, err = room.GetEarliestEntryByUserAndWeekday(userID, weekday)
			case "departure":
				logs, err = room.GetLatestExitByUserAndWeekday(userID, weekday)
			default:
				ch <- model.ProbabilityResult{}
				return
			}
			if err != nil {
				ch <- model.ProbabilityResult{}
				return
			}
			defer wg.Done()
			var p ProbabilityService
			result, err := p.PredictProbability(logs, weeks, time, isForward)
			if err != nil {
				ch <- model.ProbabilityResult{}
				return
			}
			ch <- model.ProbabilityResult{UserID: userID, Probability: result}
		}(userID, &wg)
	}
	go func() {
		wg.Wait()
		close(ch) // すべての Goroutine が終了したらチャネルを閉じる
	}()
	// 予測結果を格納
	var results []model.ProbabilityResult
	for range userIDs {
		result := <-ch
		results = append(results, result)
	}
	// 予測結果を返す
	response := model.ProbabilityResponse{
		Weekday:   weekday,
		Time:      time,
		IsForward: isForward,
		Result:    results,
	}
	return response, nil
}

// pythonサーバにlogを送信してtimeまでに(or以降に)来訪する確率の予測結果を取得する
func (ProbabilityService) PredictProbability(logs []model.Log, weeks int, time string, isForward bool) (float64, error) {
	// logsからstart_atを"15:04"形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.StartAt.Format("15:04"))
	}

	// pythonサーバにstartAtとweeksを送信してtimeまでに来訪する確率の予測結果を取得
	baseUrl := "http://vol_prediction:8085/api/v1/prediction/probability"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return 0, err
	}
	q := u.Query()
	for _, t := range startAt {
		q.Add("logs", t)
	}
	q.Add("time", time)
	q.Add("weeks", fmt.Sprintf("%d", weeks))
	q.Add("is_forward", fmt.Sprintf("%t", isForward))
	u.RawQuery = q.Encode()
	// GETリクエストを送信して予測結果を取得
	res, err := http.Get(u.String())
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	// 予測結果を取得
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	var p model.Prediction
	if err = json.Unmarshal(b, &p); err != nil {
		return 0, err
	}
	// 予測結果を返す
	return p.Probability, nil
}
