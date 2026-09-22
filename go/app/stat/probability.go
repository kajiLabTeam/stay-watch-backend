package stat

import "math"

// GetProbability estimates the probability of an action occurring by
// targetTime, given historical log times ("HH:MM") and the number of
// weeks of observation. When isForward is false, the result is inverted
// to represent the probability of the action occurring at or after
// targetTime instead of by targetTime.
func GetProbability(data []string, targetTime string, weeks int, isForward bool) (float64, error) {
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

	var p float64
	if len(dataMinutes) == 1 {
		if float64(targetMinutes) >= dataMinutes[0] {
			p = 1 / float64(weeks)
		}
	} else {
		clusters, err := Clustering(dataMinutes)
		if err != nil {
			return 0, err
		}
		for _, c := range clusters {
			switch {
			case len(c.Data) == 1:
				if float64(targetMinutes) >= c.Data[0] {
					p += 1 / float64(weeks)
				}
			default:
				scale := stddev(c.Data)
				if scale == 0 {
					if float64(targetMinutes) >= c.Center {
						p += float64(len(c.Data)) / float64(weeks)
					}
				} else {
					p += NormalCDF(float64(targetMinutes), c.Center, scale) * float64(len(c.Data)) / float64(weeks)
				}
			}
		}
	}

	if !isForward {
		p = float64(len(data))/float64(weeks) - p
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
