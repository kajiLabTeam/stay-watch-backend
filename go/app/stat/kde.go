package stat

import "math"

const (
	// Lower bound on the kernel bandwidth, in minutes. Even perfectly
	// regular logs should not be treated as exact, and this also absorbs
	// BLE detection delay.
	minBandwidth = 15.0

	// Bisection settings for kdeMedian.
	medianTolerance = 0.01
	maxMedianIters  = 100
)

// bandwidth returns Silverman's rule-of-thumb bandwidth for xs,
// floored at minBandwidth.
func bandwidth(xs []float64) float64 {
	h := 1.06 * stddev(xs) * math.Pow(float64(len(xs)), -0.2)
	return math.Max(minBandwidth, h)
}

// kdeCDF returns the KDE-smoothed CDF of xs evaluated at t:
// (1/n) * Σ Φ((t - x_i) / h).
func kdeCDF(xs []float64, t, h float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += NormalCDF(t, x, h)
	}
	return sum / float64(len(xs))
}

// kdeMedian returns the t at which kdeCDF(xs, t, h) == 0.5,
// found by bisection.
func kdeMedian(xs []float64, h float64) float64 {
	lo, hi := xs[0], xs[0]
	for _, x := range xs {
		lo = math.Min(lo, x)
		hi = math.Max(hi, x)
	}
	lo -= 6 * h
	hi += 6 * h

	for i := 0; i < maxMedianIters && hi-lo >= medianTolerance; i++ {
		mid := (lo + hi) / 2
		if kdeCDF(xs, mid, h) < 0.5 {
			lo = mid
		} else {
			hi = mid
		}
	}
	return (lo + hi) / 2
}

// NormalCDF returns the cumulative distribution function of a normal
// distribution with the given mean and standard deviation, evaluated at x.
func NormalCDF(x, mean, std float64) float64 {
	return 0.5 * math.Erfc(-(x-mean)/(std*math.Sqrt2))
}

// stddev returns the population standard deviation (ddof=0) of data.
func stddev(data []float64) float64 {
	mean := 0.0
	for _, v := range data {
		mean += v
	}
	mean /= float64(len(data))

	variance := 0.0
	for _, v := range data {
		diff := v - mean
		variance += diff * diff
	}
	variance /= float64(len(data))
	return math.Sqrt(variance)
}
