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
	// visitRate(1, 4) = 2/6, and 10:00 is after the only log.
	if math.Abs(p-2.0/6) > 1e-9 {
		t.Errorf("got %v, want %v", p, 2.0/6)
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
	// visitRate(3, 3) = 4/5, and all logs are before 10:00.
	if math.Abs(p-0.8) > 1e-9 {
		t.Errorf("got %v, want 0.8", p)
	}
}

func TestGetProbabilityMultiCluster(t *testing.T) {
	data := []string{"08:00", "08:00", "08:05", "18:00", "18:00", "18:05"}
	p, err := GetProbability(data, "12:00", 6, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Morning cluster (3 pts, ~08:00) should be fully counted by noon;
	// evening cluster (3 pts, ~18:00) should not: visitRate(6, 6) * 0.5.
	want := 7.0 / 8 * 0.5
	if math.Abs(p-want) > 0.05 {
		t.Errorf("got %v, want ~%v", p, want)
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
	want := visitRate(len(data), 4) - forward
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

func TestGetProbabilityNonPositiveWeeks(t *testing.T) {
	for _, weeks := range []int{0, -1} {
		if _, err := GetProbability([]string{"09:00"}, "10:00", weeks, true); err == nil {
			t.Errorf("expected error for weeks=%d", weeks)
		}
	}
}

func TestVisitRate(t *testing.T) {
	cases := []struct {
		n, weeks int
		want     float64
	}{
		{0, 1, 1.0 / 3},
		{1, 1, 2.0 / 3},
		{16, 20, 17.0 / 22},
	}
	for _, c := range cases {
		if got := visitRate(c.n, c.weeks); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("visitRate(%d, %d) = %v, want %v", c.n, c.weeks, got, c.want)
		}
	}
}

func TestGetProbabilityInUnitRange(t *testing.T) {
	logTimes := []string{"08:00", "09:30", "09:00", "13:00", "18:00", "08:30"}
	targets := []string{"00:00", "08:00", "09:00", "12:00", "18:00", "23:59"}
	for weeks := 1; weeks <= len(logTimes); weeks++ {
		for n := 1; n <= weeks; n++ {
			for _, target := range targets {
				for _, isForward := range []bool{true, false} {
					p, err := GetProbability(logTimes[:n], target, weeks, isForward)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if p < 0 || p > 1 {
						t.Errorf("n=%d weeks=%d target=%s forward=%v: p = %v out of [0, 1]", n, weeks, target, isForward, p)
					}
				}
			}
		}
	}
}
