package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/config/serverconfig"
	"github.com/sinobite/go-metrics/internal/handlers/allmetrcishandler"
	"github.com/sinobite/go-metrics/internal/handlers/metrichandler"
	"github.com/sinobite/go-metrics/internal/handlers/updatemetrichandler"
	"github.com/sinobite/go-metrics/internal/logger"
	"github.com/sinobite/go-metrics/internal/middleware/withlogger"
	"github.com/sinobite/go-metrics/internal/storage"
	"net/http"
)

func NewRouter(storage storage.Storage, log logger.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(withlogger.New(log))

	router.Get("/", allmetrcishandler.New(storage))
	router.Post("/update/{metricType}/{metricName}/{metricValue}", updatemetrichandler.New(storage))
	router.Get("/value/{metricType}/{metricName}", metrichandler.New(storage))

	return router
}

func main() {
	cfg := serverconfig.New()
	cfg.Parse()

	s := storage.New()

	log := logger.New("debug")

	router := NewRouter(s, log)

	err := http.ListenAndServe(cfg.FlagRunEndpoint, router)
	if err != nil {
		log.ZapLog.Fatalf("server stopped with error: %s", err)
	}
}
