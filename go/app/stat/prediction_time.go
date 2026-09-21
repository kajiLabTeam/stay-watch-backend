package stat

// GetPredictionTime predicts the most likely "HH:MM" time for an action,
// given historical log times ("HH:MM") and the number of weeks of
// observation.
func GetPredictionTime(data []string, weeks int) (string, error) {
	dataMinutes := make([]float64, len(data))
	for i, d := range data {
		m, err := TimeToMinutes(d)
		if err != nil {
			return "", err
		}
		dataMinutes[i] = float64(m)
	}

	if len(dataMinutes) == 1 {
		return MinutesToTime(dataMinutes[0]), nil
	}

	clusters, err := Clustering(dataMinutes)
	if err != nil {
		return "", err
	}

	var sum float64
	for _, c := range clusters {
		sum += c.Center * (float64(len(c.Data)) / float64(len(dataMinutes)))
	}
	return MinutesToTime(sum), nil
}
