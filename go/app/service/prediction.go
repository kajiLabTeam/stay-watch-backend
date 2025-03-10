package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type PredictionService struct{}

// pythonサーバにlogを送信してtimeまでに(or以降に)来訪する確率の予測結果を取得する
func (PredictionService) PredictVisitProbability(userID int64, weekday int, time string, isForward bool) (float64, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetEarliestEntryByUserAndWeekday(userID, weekday)
	if err != nil {
		return 0, err
	}
	// logsからstart_atを"15:04:05"形式に変換してスライスに格納
	var startAt []string
	for _, log := range logs {
		startAt = append(startAt, log.StartAt.Format("15:04:05"))
	}

	// user_IDを持つuserが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return 0, err
	}

	// pythonサーバにstartAtとweeksを送信してtimeまでに来訪する確率の予測結果を取得
	baseUrl := "http://localhost:5000/predict"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return 0, err
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
		return 0, err
	}
	defer res.Body.Close()
	// 予測結果を取得
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	probability, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0, err
	}
	// 予測結果を返す
	return probability, nil
}

// pythonサーバにlogを送信してtimeまでに(or以降に)帰宅する確率の予測結果を取得する
func (PredictionService) PredictDepartureProbability(userID int64, weekday int, time string, isForward bool) (float64, error) {
	var room RoomService
	// user_IDとweekdayを元にlogを取得
	logs, err := room.GetLatestExitByUserAndWeekday(userID, weekday)
	if err != nil {
		return 0, err
	}
	// logsからstart_atを"15:04:05"形式に変換してスライスに格納
	// start_atの日付とend_atの日付が異なる場合はスルーする
	var endAt []string
	for _, log := range logs {
		if log.StartAt.Day() != log.EndAt.Day() {
			continue
		}
		endAt = append(endAt, log.EndAt.Format("15:04:05"))
	}

	// user_IDを持つuserが所属を始めてからの週数を取得
	weeks, err := room.GetWeeksSinceFirstLog(userID)
	if err != nil {
		return 0, err
	}

	// pythonサーバにstartAtとweeksを送信してtimeまでに来訪する確率の予測結果を取得
	baseUrl := "http://localhost:5000/predict"
	u, err := url.Parse(baseUrl)
	if err != nil {
		return 0, err
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
		return 0, err
	}
	defer res.Body.Close()
	// 予測結果を取得
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	probability, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0, err
	}
	// 予測結果を返す
	return probability, nil
}
