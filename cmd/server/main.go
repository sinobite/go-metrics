package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/handlers"
	"net/http"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handlers.AllMetricsHandler)
	router.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateMetricHandler)
	router.Get("/value/{metricType}/{metricName}", handlers.MetricHandler)

	return router
}

func main() {
	parseFlags()

	err := http.ListenAndServe(flagRunEndpoint, NewRouter())
	if err != nil {
		panic(err)
	}
}
