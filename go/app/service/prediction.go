package service

import (
	"log"
	"sync"
	"time"

	"Stay_watch/model"
	"Stay_watch/stat"
)

type PredictionService struct{}

// goroutineを使って予測結果を取得する
func (PredictionService) GetPrediction(action string, userIDs []int64, weekday int) (model.PredictionResponse, error) {
	// 予測結果を格納するチャネル
	var wg sync.WaitGroup
	wg.Add(len(userIDs))
	ch := make(chan model.PredictionResult, len(userIDs))
	var rooms RoomService
	// ユーザーごとにgoroutineを生成して予測結果を取得
	for _, userID := range userIDs {
		go func(userID int64, wg *sync.WaitGroup) {
			weeks, err := rooms.GetWeeksSinceFirstLog(userID)
			defer wg.Done() // Goroutineが終了したらWaitGroupを減らす
			if err != nil {
				ch <- model.PredictionResult{UserID: userID, PredictionTime: ""}
				log.Println("Error getting weeks since first log:", err)
				return
			}
			// var logs []model.Log
			var times []time.Time
			switch action {
			case "visit":
				logs, err := rooms.GetEarliestEntryByUserAndWeekday(userID, weekday)
				if err != nil {
					ch <- model.PredictionResult{UserID: userID, PredictionTime: ""}
					log.Println("Error getting earliest entry logs:", err)
					return
				}
				for _, log := range logs {
					times = append(times, log.StartAt)
				}
			case "departure":
				logs, err := rooms.GetLatestExitByUserAndWeekday(userID, weekday)
				if err != nil {
					ch <- model.PredictionResult{UserID: userID, PredictionTime: ""}
					log.Println("Error getting latest exit logs:", err)
					return
				}
				for _, log := range logs {
					times = append(times, log.EndAt)
				}
			default:
				ch <- model.PredictionResult{UserID: userID, PredictionTime: ""}
				log.Println("Invalid action:", action)
				return
			}
			var p PredictionService
			result, err := p.PredictTime(times, weeks)
			if err != nil {
				ch <- model.PredictionResult{UserID: userID, PredictionTime: ""}
				log.Println("Error predicting time:", err)
				return
			}
			ch <- model.PredictionResult{UserID: userID, PredictionTime: result}
		}(userID, &wg)
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

// logから来訪する可能性の高い時刻を取得する
func (PredictionService) PredictTime(logs []time.Time, weeks int) (string, error) {
	// logからstart_atを”15:04”形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.Format("15:04"))
	}

	return stat.GetPredictionTime(startAt, weeks)
}