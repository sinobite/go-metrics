package handlers

import (
	"github.com/sinobite/go-metrics/internal/storage"
	"github.com/sinobite/go-metrics/internal/utils"
	"net/http"
)

func UpdateMetricHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		metricType := request.PathValue("metricType")
		metricName := request.PathValue("metricName")
		metricValue := request.PathValue("metricValue")

		if metricType == "gauge" {
			gaugeValue, err := utils.ConvertToGauge(metricValue)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			}

			storage.Storage.Gauges[metricName] = gaugeValue
			writer.WriteHeader(http.StatusOK)
		} else if metricType == "counter" {
			counterValue, err := utils.ConvertToCounter(metricValue)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			}

			storage.Storage.Counters[metricName] = storage.Storage.Counters[metricName] + counterValue
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusBadRequest)

		}

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}

}
