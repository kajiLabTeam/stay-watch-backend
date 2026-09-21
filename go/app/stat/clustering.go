package stat

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	maxClusters   = 4
	maxEMIters    = 100
	emTolerance   = 1e-6
	varianceFloor = 1e-6
)

// ClusterResult holds the data points assigned to a Gaussian mixture
// component along with the component's fitted mean (center).
type ClusterResult struct {
	Data   []float64
	Center float64
}

// gaussianComponent is a single 1-D Gaussian mixture component.
type gaussianComponent struct {
	weight   float64
	mean     float64
	variance float64
}

// Clustering fits a 1-D Gaussian Mixture Model to data, selecting the
// number of components (1..min(maxClusters, len(data))) that minimizes
// the Bayesian Information Criterion (BIC), then hard-assigns each data
// point to its most likely component.
func Clustering(data []float64) ([]ClusterResult, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("clustering: data must not be empty")
	}

	upper := maxClusters
	if len(data) < upper {
		upper = len(data)
	}

	var bestComponents []gaussianComponent
	bestBIC := math.Inf(1)
	for k := 1; k <= upper; k++ {
		components, logLikelihood := fitGMM(data, k)
		b := bic(logLikelihood, k, len(data))
		if b < bestBIC {
			bestBIC = b
			bestComponents = components
		}
	}

	return assignClusters(data, bestComponents), nil
}

// fitGMM fits a k-component 1-D Gaussian mixture to data via EM,
// starting from a deterministic quantile-based initialization, and
// returns the fitted components and the final log-likelihood.
func fitGMM(data []float64, k int) ([]gaussianComponent, float64) {
	components := initComponents(data, k)

	n := len(data)
	responsibilities := make([][]float64, n)
	for i := range responsibilities {
		responsibilities[i] = make([]float64, k)
	}

	prevLogLikelihood := math.Inf(-1)
	for iter := 0; iter < maxEMIters; iter++ {
		// E-step
		logLikelihood := 0.0
		for i, x := range data {
			total := 0.0
			densities := make([]float64, k)
			for c, comp := range components {
				d := comp.weight * normalPDF(x, comp.mean, math.Sqrt(comp.variance))
				densities[c] = d
				total += d
			}
			if total == 0 {
				// Numerically degenerate; assign uniformly to avoid NaN.
				for c := range components {
					responsibilities[i][c] = 1.0 / float64(k)
				}
				continue
			}
			for c := range components {
				responsibilities[i][c] = densities[c] / total
			}
			logLikelihood += math.Log(total)
		}

		// M-step
		for c := range components {
			sumR := 0.0
			for i := range data {
				sumR += responsibilities[i][c]
			}
			if sumR == 0 {
				continue
			}
			mean := 0.0
			for i, x := range data {
				mean += responsibilities[i][c] * x
			}
			mean /= sumR

			variance := 0.0
			for i, x := range data {
				diff := x - mean
				variance += responsibilities[i][c] * diff * diff
			}
			variance /= sumR
			if variance < varianceFloor {
				variance = varianceFloor
			}

			components[c].mean = mean
			components[c].variance = variance
			components[c].weight = sumR / float64(n)
		}

		if math.Abs(logLikelihood-prevLogLikelihood) < emTolerance {
			prevLogLikelihood = logLikelihood
			break
		}
		prevLogLikelihood = logLikelihood
	}

	return components, prevLogLikelihood
}

// initComponents deterministically initializes k components by sorting
// the data and splitting it into k contiguous buckets.
func initComponents(data []float64, k int) []gaussianComponent {
	sorted := append([]float64(nil), data...)
	sort.Float64s(sorted)

	components := make([]gaussianComponent, k)
	n := len(sorted)
	for c := 0; c < k; c++ {
		start := c * n / k
		end := (c + 1) * n / k
		if end <= start {
			end = start + 1
		}
		if end > n {
			end = n
		}
		bucket := sorted[start:end]

		mean := 0.0
		for _, v := range bucket {
			mean += v
		}
		mean /= float64(len(bucket))

		variance := 0.0
		for _, v := range bucket {
			diff := v - mean
			variance += diff * diff
		}
		variance /= float64(len(bucket))
		if variance < varianceFloor {
			variance = varianceFloor
		}

		components[c] = gaussianComponent{
			weight:   float64(len(bucket)) / float64(n),
			mean:     mean,
			variance: variance,
		}
	}
	return components
}

// assignClusters hard-assigns each data point to the component with the
// highest weighted density, dropping any component with no assigned points.
func assignClusters(data []float64, components []gaussianComponent) []ClusterResult {
	buckets := make([][]float64, len(components))
	for _, x := range data {
		best := 0
		bestDensity := math.Inf(-1)
		for c, comp := range components {
			d := comp.weight * normalPDF(x, comp.mean, math.Sqrt(comp.variance))
			if d > bestDensity {
				bestDensity = d
				best = c
			}
		}
		buckets[best] = append(buckets[best], x)
	}

	var results []ClusterResult
	for c, bucket := range buckets {
		if len(bucket) == 0 {
			continue
		}
		results = append(results, ClusterResult{
			Data:   bucket,
			Center: components[c].mean,
		})
	}
	return results
}

// bic computes the Bayesian Information Criterion for a k-component
// 1-D Gaussian mixture: BIC = -2*logLikelihood + numParams*log(n).
// numParams = k means + k variances + (k-1) independent weights.
func bic(logLikelihood float64, k, n int) float64 {
	numParams := 2*k - 1
	return -2*logLikelihood + float64(numParams)*math.Log(float64(n))
}

func normalPDF(x, mean, std float64) float64 {
	if std <= 0 {
		std = math.Sqrt(varianceFloor)
	}
	return distuv.Normal{Mu: mean, Sigma: std}.Prob(x)
}

// NormalCDF returns the cumulative distribution function of a normal
// distribution with the given mean and standard deviation, evaluated at x.
func NormalCDF(x, mean, std float64) float64 {
	return distuv.Normal{Mu: mean, Sigma: std}.CDF(x)
}
