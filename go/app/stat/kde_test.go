package stat

import (
	"math"
	"testing"
)

func TestNormalCDF(t *testing.T) {
	// CDF at the mean should be 0.5.
	got := NormalCDF(0, 0, 1)
	if math.Abs(got-0.5) > 1e-9 {
		t.Errorf("NormalCDF(0,0,1) = %v, want 0.5", got)
	}
	// Φ(1) ≈ 0.841345.
	got = NormalCDF(1, 0, 1)
	if math.Abs(got-0.841345) > 1e-6 {
		t.Errorf("NormalCDF(1,0,1) = %v, want ~0.841345", got)
	}
}

func TestBandwidthSingleLog(t *testing.T) {
	if got := bandwidth([]float64{540}); got != minBandwidth {
		t.Errorf("bandwidth = %v, want %v", got, minBandwidth)
	}
}

func TestBandwidthSilverman(t *testing.T) {
	xs := []float64{480, 600, 720, 840, 960}
	want := 1.06 * stddev(xs) * math.Pow(float64(len(xs)), -0.2)
	if want <= minBandwidth {
		t.Fatalf("test data should exceed minBandwidth, got %v", want)
	}
	if got := bandwidth(xs); math.Abs(got-want) > 1e-9 {
		t.Errorf("bandwidth = %v, want %v", got, want)
	}
}

func TestKdeCDFMonotonic(t *testing.T) {
	xs := []float64{480, 500, 540, 1080}
	h := bandwidth(xs)
	prev := 0.0
	for tm := 0.0; tm <= 1440; tm += 5 {
		got := kdeCDF(xs, tm, h)
		if got < prev {
			t.Fatalf("kdeCDF decreased at t=%v: %v < %v", tm, got, prev)
		}
		prev = got
	}
}

func TestKdeCDFTails(t *testing.T) {
	xs := []float64{480, 500, 540}
	h := bandwidth(xs)
	if got := kdeCDF(xs, 0, h); got > 1e-6 {
		t.Errorf("kdeCDF far below data = %v, want ~0", got)
	}
	if got := kdeCDF(xs, 1439, h); got < 1-1e-6 {
		t.Errorf("kdeCDF far above data = %v, want ~1", got)
	}
}

func TestKdeCDFAtMedian(t *testing.T) {
	xs := []float64{480, 500, 540}
	if got := kdeCDF(xs, 500, bandwidth(xs)); math.Abs(got-0.5) > 0.05 {
		t.Errorf("kdeCDF at median = %v, want ~0.5", got)
	}
}

func TestKdeCDFExample(t *testing.T) {
	// (Φ(1)·2 + Φ(−1)) / 3
	got := kdeCDF([]float64{540, 540, 570}, 555, 15)
	want := (2*NormalCDF(1, 0, 1) + NormalCDF(-1, 0, 1)) / 3
	if math.Abs(got-want) > 1e-9 || math.Abs(got-0.614) > 1e-3 {
		t.Errorf("kdeCDF = %v, want %v (~0.614)", got, want)
	}
}

func TestKdeMedianSingle(t *testing.T) {
	if got := kdeMedian([]float64{540}, minBandwidth); math.Abs(got-540) > 0.01 {
		t.Errorf("kdeMedian = %v, want 540", got)
	}
}

func TestKdeMedianSymmetric(t *testing.T) {
	xs := []float64{480, 510, 540, 570, 600}
	if got := kdeMedian(xs, bandwidth(xs)); math.Abs(got-540) > 0.01 {
		t.Errorf("kdeMedian = %v, want 540", got)
	}
}
