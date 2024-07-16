package all_metrcis_handler

import (
	"github.com/sinobite/go-metrics/internal/storage"
	"io"
	"log"
	"net/http"
	"strings"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		metricsSlice, err := storage.FindAllMetrics()

		writer.Header().Set("Content-Type", "text/html")
		_, err = io.WriteString(writer, strings.Join(metricsSlice, ""))
		if err != nil {
			http.Error(writer, "Failed to find all metrics", http.StatusInternalServerError)
			log.Printf("Failed to find all metrics: %v", err)
		}
	}
}