package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentConfig"
	"github.com/sinobite/go-metrics/internal/services/metricsService"
	"net/http"
)

func main() {
	cfg := agentConfig.New()
	cfg.Parse()

	ms := metricsService.New(cfg)

	client := resty.New()

	ms.StartMonitoring(client)

	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		panic(err)
	}
}
