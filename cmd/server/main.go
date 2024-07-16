package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/config/serverconfig"
	"github.com/sinobite/go-metrics/internal/handlers/allmetrcishandler"
	"github.com/sinobite/go-metrics/internal/handlers/metrichandler"
	"github.com/sinobite/go-metrics/internal/handlers/updatemetrichandler"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
)

func NewRouter(storage storage.Storage) chi.Router {
	router := chi.NewRouter()

	router.Get("/", allmetrcishandler.New(storage))
	router.Post("/update/{metricType}/{metricName}/{metricValue}", updatemetrichandler.New(storage))
	router.Get("/value/{metricType}/{metricName}", metrichandler.New(storage))

	return router
}

func main() {
	cfg := serverconfig.New()

	s := storage.New()

	err := http.ListenAndServe(cfg.FlagRunEndpoint, NewRouter(s))
	if err != nil {
		panic(err)
	}
}
