package utils

import (
	"strconv"
	"strings"
)

func ConvertToGauge(value string) (float64, error) {
	gauge, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0, err
	}

	return gauge, nil
}
func ConvertToCounter(value string) (int64, error) {
	counter, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return counter, nil
}
