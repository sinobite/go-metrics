package main

import (
	"net/http"
	"strconv"
	"strings"
)

func convertToGauge(value string) (float64, error) {
	gauge, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0, err
	}

	return gauge, nil
}
func convertToCounter(value string) (int64, error) {
	counter, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

type MemStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

var memStorage MemStorage = MemStorage{
	Gauges:   make(map[string]float64),
	Counters: make(map[string]int64),
}

func updateMetricHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		metricType := request.PathValue("metricType")
		metricName := request.PathValue("metricName")
		metricValue := request.PathValue("metricValue")

		if metricType == "gauge" {
			gaugeValue, err := convertToGauge(metricValue)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			}

			memStorage.Gauges[metricName] = gaugeValue
			writer.WriteHeader(http.StatusOK)
		} else if metricType == "counter" {
			counterValue, err := convertToCounter(metricValue)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			}

			memStorage.Counters[metricName] = memStorage.Counters[metricName] + counterValue
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusBadRequest)

		}

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/{metricType}/{metricName}/{metricValue}", updateMetricHandler)

	err := http.ListenAndServe("localhost:8087", mux)
	if err != nil {
		panic(err)
	}
}
