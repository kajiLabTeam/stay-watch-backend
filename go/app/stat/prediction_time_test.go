package stat

import "testing"

func TestGetPredictionTimeSingleLog(t *testing.T) {
	got, err := GetPredictionTime([]string{"09:00"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "09:00" {
		t.Errorf("got %q, want %q", got, "09:00")
	}
}

func TestGetPredictionTimeSingleGroup(t *testing.T) {
	got, err := GetPredictionTime([]string{"09:00", "09:00", "09:05"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got < "09:00" || got > "09:05" {
		t.Errorf("got %q, want between 09:00 and 09:05", got)
	}
}

func TestGetPredictionTimeLeansToLargerPeak(t *testing.T) {
	// Four logs around 08:00 and two around 18:00: the median should sit on
	// the morning side of the midpoint (13:00) between the two peaks.
	data := []string{"08:00", "08:00", "08:00", "08:00", "18:00", "18:00"}
	got, err := GetPredictionTime(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got < "08:00" || got >= "13:00" {
		t.Errorf("got %q, want between 08:00 and 13:00", got)
	}
}

func TestGetPredictionTimeInvalid(t *testing.T) {
	if _, err := GetPredictionTime([]string{"bad"}); err == nil {
		t.Error("expected error for invalid log time")
	}
}

func TestGetPredictionTimeNoLogs(t *testing.T) {
	if _, err := GetPredictionTime(nil); err == nil {
		t.Error("expected error for no logs")
	}
}
