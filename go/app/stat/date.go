package stat

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// TimeToMinutes converts a "HH:MM" time string to minutes since midnight,
// rounded to the nearest 30-minute mark.
func TimeToMinutes(t string) (int, error) {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid time format: %s", t)
	}
	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid time format: %s", t)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid time format: %s", t)
	}
	total := h*60 + m
	rounded := math.Round(float64(total)/30) * 30
	return int(rounded), nil
}

// MinutesToTime converts minutes since midnight to a "HH:MM" time string.
// Values outside a single day are wrapped, so 1440 (e.g. "23:59" rounded
// up by TimeToMinutes) becomes "00:00".
func MinutesToTime(m float64) string {
	total := ((int(m) % 1440) + 1440) % 1440
	hours, minutes := total/60, total%60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
