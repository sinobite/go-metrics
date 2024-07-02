package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinobite/go-metrics/internal/handlers"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateMetricHandler)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic(err)
	}
}
