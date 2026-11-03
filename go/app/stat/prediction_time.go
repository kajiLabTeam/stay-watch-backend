package stat

import "fmt"

// GetPredictionTime predicts the most likely "HH:MM" time for an action,
// given historical log times ("HH:MM"). The prediction is the median of
// the kernel density estimate of the log times.
func GetPredictionTime(data []string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("no logs")
	}
	dataMinutes := make([]float64, len(data))
	for i, d := range data {
		m, err := TimeToMinutes(d)
		if err != nil {
			return "", err
		}
		dataMinutes[i] = float64(m)
	}

	return MinutesToTime(kdeMedian(dataMinutes, bandwidth(dataMinutes))), nil
}
