package stat

import (
	"math"
	"testing"
)

func TestGetProbabilitySingleLog(t *testing.T) {
	p, err := GetProbability([]string{"09:00"}, "10:00", 4, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(p-0.25) > 1e-9 {
		t.Errorf("got %v, want 0.25", p)
	}

	p, err = GetProbability([]string{"09:00"}, "08:00", 4, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != 0 {
		t.Errorf("got %v, want 0", p)
	}
}

func TestGetProbabilityZeroVarianceCluster(t *testing.T) {
	// All logs at the same rounded time -> single cluster, zero variance.
	p, err := GetProbability([]string{"09:00", "09:00", "09:05"}, "10:00", 3, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(p-1.0) > 1e-9 {
		t.Errorf("got %v, want 1.0", p)
	}
}

func TestGetProbabilityMultiCluster(t *testing.T) {
	data := []string{"08:00", "08:00", "08:05", "18:00", "18:00", "18:05"}
	p, err := GetProbability(data, "12:00", 6, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Morning cluster (3 pts, ~08:00) should be fully counted by noon;
	// evening cluster (3 pts, ~18:00) should not.
	if math.Abs(p-0.5) > 0.05 {
		t.Errorf("got %v, want ~0.5", p)
	}
}

func TestGetProbabilityIsForwardInversion(t *testing.T) {
	data := []string{"09:00"}
	forward, err := GetProbability(data, "10:00", 4, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	backward, err := GetProbability(data, "10:00", 4, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := float64(len(data))/4 - forward
	if math.Abs(backward-want) > 1e-9 {
		t.Errorf("backward = %v, want %v", backward, want)
	}
}

func TestGetProbabilityInvalidTime(t *testing.T) {
	if _, err := GetProbability([]string{"09:00"}, "bad", 4, true); err == nil {
		t.Error("expected error for invalid target time")
	}
	if _, err := GetProbability([]string{"bad"}, "09:00", 4, true); err == nil {
		t.Error("expected error for invalid log time")
	}
}
