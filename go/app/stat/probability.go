package stat

import (
	"fmt"
	"math"
)

const (
	// Beta(priorAlpha, priorBeta) prior on the per-weekday visit rate.
	priorAlpha = 1.0
	priorBeta  = 1.0
)

// visitRate returns the posterior-mean visit rate given n observed
// visits over weeks weeks, under a Beta(priorAlpha, priorBeta) prior.
func visitRate(n, weeks int) float64 {
	return (float64(n) + priorAlpha) / (float64(weeks) + priorAlpha + priorBeta)
}

// GetProbability estimates the probability of an action occurring by
// targetTime, given historical log times ("HH:MM", at most one per day)
// and the number of weeks of observation. The result is
//
//	P(by targetTime) = visitRate(n, weeks) * P(by targetTime | visits that day)
//
// where the conditional term is a CDF in [0, 1] estimated from the
// clustered log times. When isForward is false, the result is the
// probability of the action occurring at or after targetTime instead,
// i.e. visitRate * (1 - CDF).
func GetProbability(data []string, targetTime string, weeks int, isForward bool) (float64, error) {
	if weeks <= 0 {
		return 0, fmt.Errorf("weeks must be positive: %d", weeks)
	}
	dataMinutes := make([]float64, len(data))
	for i, d := range data {
		m, err := TimeToMinutes(d)
		if err != nil {
			return 0, err
		}
		dataMinutes[i] = float64(m)
	}
	targetMinutes, err := TimeToMinutes(targetTime)
	if err != nil {
		return 0, err
	}

	n := len(dataMinutes)
	var cdf float64
	if n == 1 {
		if float64(targetMinutes) >= dataMinutes[0] {
			cdf = 1
		}
	} else {
		clusters, err := Clustering(dataMinutes)
		if err != nil {
			return 0, err
		}
		for _, c := range clusters {
			share := float64(len(c.Data)) / float64(n)
			switch {
			case len(c.Data) == 1:
				if float64(targetMinutes) >= c.Data[0] {
					cdf += share
				}
			default:
				scale := stddev(c.Data)
				if scale == 0 {
					if float64(targetMinutes) >= c.Center {
						cdf += share
					}
				} else {
					cdf += NormalCDF(float64(targetMinutes), c.Center, scale) * share
				}
			}
		}
	}

	rate := visitRate(n, weeks)
	p := rate * cdf
	if !isForward {
		p = rate - p
	}
	return p, nil
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
