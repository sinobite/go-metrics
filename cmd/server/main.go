package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/handlers"
	"net/http"
	"os"
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

var flagRunEndpoint string = "localhost:8080"

func parseFlags() {
	flag.StringVar(&flagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		flagRunEndpoint = envRunAddr
	}
}
