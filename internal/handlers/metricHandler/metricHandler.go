package metricHandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		metricType := chi.URLParam(request, "metricType")
		metricName := chi.URLParam(request, "metricName")
		switch metricType {
		case "gauge":
			gaugeValue, err := storage.FindGaugeMetric(metricName)
			if err != nil {
				http.Error(writer, "Failed to find gauge metrics", http.StatusNotFound)
			}

			_, err = writer.Write([]byte(gaugeValue))
			if err != nil {
				http.Error(writer, "Failed to find gauge metrics", http.StatusNotFound)
			}
		case "counter":
			counterValue, err := storage.FindCounterMetric(metricName)
			if err != nil {
				http.Error(writer, "Failed to find counter metrics", http.StatusNotFound)
			}

			_, err = writer.Write([]byte(counterValue))
			if err != nil {
				http.Error(writer, "Failed to find counter metrics", http.StatusNotFound)
			}
		default:
			http.Error(writer, "Metric type is incorrect", http.StatusBadRequest)
		}
	}
}
