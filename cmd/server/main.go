package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/handlers"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", handlers.AllMetricsHandler)
	router.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateMetricHandler)
	router.Get("/value/{metricType}/{metricName}", handlers.MetricHandler)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic(err)
	}
}
