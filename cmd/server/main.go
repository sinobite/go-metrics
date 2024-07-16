package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/config/server_config"
	"github.com/sinobite/go-metrics/internal/handlers/all_metrcis_handler"
	"github.com/sinobite/go-metrics/internal/handlers/metric_handler"
	"github.com/sinobite/go-metrics/internal/handlers/update_metric_handler"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
)

func NewRouter(storage storage.Storage) chi.Router {
	router := chi.NewRouter()

	router.Get("/", all_metrcis_handler.New(storage))
	router.Post("/update/{metricType}/{metricName}/{metricValue}", update_metric_handler.New(storage))
	router.Get("/value/{metricType}/{metricName}", metric_handler.New(storage))

	return router
}

func main() {
	cfg := server_config.New()

	s := storage.New()

	err := http.ListenAndServe(cfg.FlagRunEndpoint, NewRouter(s))
	if err != nil {
		panic(err)
	}
}
