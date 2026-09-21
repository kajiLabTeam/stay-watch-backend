package service

import (
	"log"
	"sync"

	"Stay_watch/model"
	"Stay_watch/stat"
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
			defer wg.Done() // Goroutineが終了したらWaitGroupを減らす
			if err != nil {
				ch <- model.ProbabilityResult{UserID: userID, Probability: 0}
				log.Println("Error getting weeks since first log:", err)
				return
			}
			var logs []model.Log
			switch action {
			case "visit":
				logs, err = room.GetEarliestEntryByUserAndWeekday(userID, weekday)
			case "departure":
				logs, err = room.GetLatestExitByUserAndWeekday(userID, weekday)
			default:
				ch <- model.ProbabilityResult{UserID: userID, Probability: 0}
				log.Println("Invalid action:", action)
				return
			}
			if err != nil {
				ch <- model.ProbabilityResult{UserID: userID, Probability: 0}
				log.Println("Error getting logs:", err)
				return
			}
			var p ProbabilityService
			result, err := p.PredictProbability(logs, weeks, time, isForward)
			if err != nil {
				ch <- model.ProbabilityResult{UserID: userID, Probability: 0}
				log.Println("Error getting probability:", err)
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

// logsからtimeまでに(or以降に)来訪する確率の予測結果を取得する
func (ProbabilityService) PredictProbability(logs []model.Log, weeks int, time string, isForward bool) (float64, error) {
	// logsからstart_atを"15:04"形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.StartAt.Format("15:04"))
	}

	return stat.GetProbability(startAt, time, weeks, isForward)
}