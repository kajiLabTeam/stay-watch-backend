package stat

import "testing"

func TestGetPredictionTimeSingleLog(t *testing.T) {
	got, err := GetPredictionTime([]string{"09:00"}, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "09:00" {
		t.Errorf("got %q, want %q", got, "09:00")
	}
}

func TestGetPredictionTimeSingleCluster(t *testing.T) {
	got, err := GetPredictionTime([]string{"09:00", "09:00", "09:05"}, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "09:00" {
		t.Errorf("got %q, want %q", got, "09:00")
	}
}

func TestGetPredictionTimeMultiCluster(t *testing.T) {
	// Equal-sized clusters at 08:00 and 18:00 -> weighted average is 13:00.
	data := []string{"08:00", "08:00", "08:00", "18:00", "18:00", "18:00"}
	got, err := GetPredictionTime(data, 6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "13:00" {
		t.Errorf("got %q, want %q", got, "13:00")
	}
}

func TestGetPredictionTimeInvalid(t *testing.T) {
	if _, err := GetPredictionTime([]string{"bad"}, 4); err == nil {
		t.Error("expected error for invalid log time")
	}
}
