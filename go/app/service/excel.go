package service

import (
	"Stay_watch/model"
	"context"
	"fmt"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var spreadsheetID = "1wzKljj2j8EmRHQ5wK6A6RDtfyJbKzUXUfbKyIHURliU"

type ExcelService struct{}

func (ExcelService) WriteExcel(attendancesTmp []model.AttendanceTmp, meetingID int64) error {
	credential := option.WithCredentialsFile("/app/credentials/secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		return fmt.Errorf(" failed: %w", err)
	}

	//3年セミナー
	if meetingID == 1 {
		//3年セミナー 2022
		// readRange := "シート1!A1:A10"
		readRange := "3年セミナー 2022!A2:A55"
		resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
		if err != nil {
			return fmt.Errorf(" failed: %w", err)
		}
		if len(resp.Values) == 0 {
			return fmt.Errorf(" failed: %w", err)
		}
		nowDate := time.Now().Format("2006/01/02 15:04:05")
		dateIndex := 0
		for index, row := range resp.Values {
			if row[0] == nowDate[5:10] {
				dateIndex = index
			}
		}

		mozis := make([]string, 0)

		for _, attendance_tmp := range attendancesTmp[16:] {
			if attendance_tmp.Flag == 0 {
				// mozis[attendance_tmp.UserID] = "欠"
				mozis = append(mozis, "欠")
			} else {
				// mozis[attendance_tmp.UserID] = "出"
				mozis = append(mozis, "出")
			}

			UserService := UserService{}
			//レコードを初期状態に戻す
			UserService.TemporarilySavedAttendance(attendance_tmp.UserID, 0)

		}

		AbstractSlice(mozis)

		if dateIndex != 0 {
			//書き込み
			var vr sheets.ValueRange
			vr.Values = append(vr.Values, AbstractSlice(mozis))
			updateRange := fmt.Sprintf("3年セミナー 2022!C%d:M%d", dateIndex+2, dateIndex+2)
			if _, err = srv.Spreadsheets.Values.Update(spreadsheetID, updateRange, &vr).ValueInputOption("RAW").Do(); err != nil {
				return fmt.Errorf(" failed: %w", err)
			}
			srv.Spreadsheets.Values.Update(spreadsheetID, updateRange, &vr).ValueInputOption("RAW").Do()
		}

		return nil
	} else {
		// readRange := "シート1!A1:A10"
		readRange := "全体ミーティング(toyama)!A2:A55"
		resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
		if err != nil {
			return fmt.Errorf(" failed: %w", err)
		}
		if len(resp.Values) == 0 {
			return fmt.Errorf(" failed: %w", err)
		}
		nowDate := time.Now().Format("2006/01/02 15:04:05")

		dateIndex := 0

		for index, row := range resp.Values {
			if row[0] == nowDate[5:10] {
				dateIndex = index
			}
		}

		mozis := make([]string, 0)

		// for i := 0; i < 16; i++ {
		// 	mozis = append(mozis, "出")
		// }

		for _, attendance_tmp := range attendancesTmp[:16] {
			if attendance_tmp.Flag == 0 {
				// mozis[attendance_tmp.UserID] = "欠"
				mozis = append(mozis, "欠")
			} else {
				// mozis[attendance_tmp.UserID] = "出"
				mozis = append(mozis, "出")
			}

			UserService := UserService{}
			//レコードを初期状態に戻す
			UserService.TemporarilySavedAttendance(attendance_tmp.UserID, 0)
		}

		AbstractSlice(mozis)

		if dateIndex != 0 {
			//書き込み
			var vr sheets.ValueRange
			vr.Values = append(vr.Values, AbstractSlice(mozis))
			updateRange := fmt.Sprintf("全体ミーティング(toyama)!C%d:R%d", dateIndex+2, dateIndex+2)

			if _, err = srv.Spreadsheets.Values.Update(spreadsheetID, updateRange, &vr).ValueInputOption("RAW").Do(); err != nil {
				return fmt.Errorf(" failed: %w", err)
			}
			srv.Spreadsheets.Values.Update(spreadsheetID, updateRange, &vr).ValueInputOption("RAW").Do()
		}

		return nil
	}

}

func AbstractSlice(arr interface{}) []interface{} {
	dest := []interface{}{}
	switch sl := arr.(type) {
	case []interface{}:
		// こういうのは無理なんですよ
	case string:
		// rangeかけるためにcase内じゃないとダメ
		for _, b := range sl {
			dest = append(dest, b)
		}
	case []string:
		for _, str := range sl {
			dest = append(dest, str)
		}
	case []int:
		for _, i := range sl {
			dest = append(dest, i)
		}
	}
	return dest
}
