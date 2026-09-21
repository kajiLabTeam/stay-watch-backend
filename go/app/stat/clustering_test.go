package stat

import (
	"math"
	"testing"
)

func TestClusteringSinglePoint(t *testing.T) {
	clusters, err := Clustering([]float64{100})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(clusters) != 1 {
		t.Fatalf("expected 1 cluster, got %d", len(clusters))
	}
	if len(clusters[0].Data) != 1 || clusters[0].Data[0] != 100 {
		t.Errorf("unexpected cluster data: %+v", clusters[0])
	}
}

func TestClusteringTightGroup(t *testing.T) {
	data := []float64{480, 485, 490, 480, 480}
	clusters, err := Clustering(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	total := 0
	for _, c := range clusters {
		total += len(c.Data)
		if math.Abs(c.Center-480) > 15 {
			t.Errorf("cluster center %v far from tight group mean ~480", c.Center)
		}
	}
	if total != len(data) {
		t.Errorf("expected all %d points assigned, got %d", len(data), total)
	}
}

func TestClusteringBimodal(t *testing.T) {
	// Cluster around 08:00 (480 min) and 18:00 (1080 min).
	data := []float64{475, 480, 485, 480, 1075, 1080, 1085, 1080}
	clusters, err := Clustering(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(clusters) != 2 {
		t.Fatalf("expected 2 clusters for bimodal data, got %d: %+v", len(clusters), clusters)
	}

	morning, evening := clusters[0], clusters[1]
	if morning.Center > evening.Center {
		morning, evening = evening, morning
	}
	if math.Abs(morning.Center-480) > 15 {
		t.Errorf("morning cluster center = %v, want ~480", morning.Center)
	}
	if math.Abs(evening.Center-1080) > 15 {
		t.Errorf("evening cluster center = %v, want ~1080", evening.Center)
	}
	if len(morning.Data) != 4 || len(evening.Data) != 4 {
		t.Errorf("expected 4/4 split, got %d/%d", len(morning.Data), len(evening.Data))
	}
}

func TestClusteringCapsAtDataLength(t *testing.T) {
	// Only 2 points; component search must not exceed len(data).
	clusters, err := Clustering([]float64{100, 200})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	total := 0
	for _, c := range clusters {
		total += len(c.Data)
	}
	if total != 2 {
		t.Errorf("expected 2 points assigned total, got %d", total)
	}
}

func TestClusteringEmpty(t *testing.T) {
	if _, err := Clustering(nil); err == nil {
		t.Error("expected error for empty data")
	}
}

func TestNormalCDF(t *testing.T) {
	// CDF at the mean should be 0.5.
	got := NormalCDF(0, 0, 1)
	if math.Abs(got-0.5) > 1e-9 {
		t.Errorf("NormalCDF(0,0,1) = %v, want 0.5", got)
	}
}
