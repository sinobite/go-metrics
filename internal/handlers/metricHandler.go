package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
	"strconv"
)

func MetricHandler(writer http.ResponseWriter, request *http.Request) {
	metricType := chi.URLParam(request, "metricType")
	metricName := chi.URLParam(request, "metricName")
	switch metricType {
	case "gauge":
		value, ok := storage.Storage.Gauges[metricName]
		if ok {
			_, err := writer.Write([]byte(strconv.Itoa(int(value))))
			if err != nil {
				panic(err)
			}

		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	case "counter":
		value, ok := storage.Storage.Counters[metricName]
		if ok {
			_, err := writer.Write([]byte(strconv.Itoa(int(value))))
			if err != nil {
				panic(err)
			}

		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}
