package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/config/serverConfig"
	"github.com/sinobite/go-metrics/internal/handlers/allMetrcisHandler"
	"github.com/sinobite/go-metrics/internal/handlers/metricHandler"
	"github.com/sinobite/go-metrics/internal/handlers/updateMetricHandler"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
)

func NewRouter(storage storage.Storage) chi.Router {
	router := chi.NewRouter()

	router.Get("/", allMetrcisHandler.New(storage))
	router.Post("/update/{metricType}/{metricName}/{metricValue}", updateMetricHandler.New(storage))
	router.Get("/value/{metricType}/{metricName}", metricHandler.New(storage))

	return router
}

func main() {
	cfg := serverConfig.New()

	s := storage.New()

	err := http.ListenAndServe(cfg.FlagRunEndpoint, NewRouter(s))
	if err != nil {
		panic(err)
	}
}
