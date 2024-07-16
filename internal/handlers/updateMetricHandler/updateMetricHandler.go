package updateMetricHandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/storage"
	"log"
	"net/http"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		metricType := chi.URLParam(request, "metricType")
		metricName := chi.URLParam(request, "metricName")
		metricValue := chi.URLParam(request, "metricValue")
		switch metricType {
		case "gauge":
			err := storage.UpdateGaugeMetric(metricName, metricValue)
			if err != nil {
				http.Error(writer, "Failed to update gauge metric", http.StatusBadRequest)
				log.Printf("Failed to update gauge metric: %v", err)
			}
			writer.WriteHeader(http.StatusOK)
		case "counter":
			err := storage.UpdateCounterMetric(metricName, metricValue)
			if err != nil {
				http.Error(writer, "Failed to update counter metric", http.StatusBadRequest)
				log.Printf("Failed to update counter metric: %v", err)
			}
			writer.WriteHeader(http.StatusOK)
		default:
			http.Error(writer, "Metric type is incorrect", http.StatusBadRequest)
		}
	}
}
