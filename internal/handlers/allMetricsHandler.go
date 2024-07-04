package handlers

import (
	"fmt"
	"github.com/sinobite/go-metrics/internal/storage"
	"io"
	"net/http"
	"strings"
)

func AllMetricsHandler(writer http.ResponseWriter, request *http.Request) {
	var metricsSlice []string
	for key, value := range storage.Storage.Gauges {
		metricsSlice = append(metricsSlice, fmt.Sprintf("%s: %v \n", key, value))
	}
	for key, value := range storage.Storage.Counters {
		metricsSlice = append(metricsSlice, fmt.Sprintf("%s: %v \n", key, value))
	}

	writer.Header().Set("Content-Type", "text/html")
	_, err := io.WriteString(writer, strings.Join(metricsSlice, ""))
	if err != nil {
		panic(err)
	}
}
