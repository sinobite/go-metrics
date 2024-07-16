package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/handlers/allMetrcisHandler"
	"github.com/sinobite/go-metrics/internal/handlers/metricHandler"
	"github.com/sinobite/go-metrics/internal/handlers/updateMetricHandler"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
	"os"
)

func NewRouter(storage storage.Storage) chi.Router {
	router := chi.NewRouter()

	router.Get("/", allMetrcisHandler.New(storage))
	router.Post("/update/{metricType}/{metricName}/{metricValue}", updateMetricHandler.New(storage))
	router.Get("/value/{metricType}/{metricName}", metricHandler.New(storage))

	return router
}

func main() {
	parseFlags()

	s := storage.New()

	err := http.ListenAndServe(flagRunEndpoint, NewRouter(s))
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
