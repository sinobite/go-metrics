package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentconfig"
	"github.com/sinobite/go-metrics/internal/services/metricsservice"
	"net/http"
)

func main() {
	cfg := agentconfig.New()
	cfg.Parse()

	ms := metricsservice.New(cfg)

	client := resty.New()

	ms.StartMonitoring(client)

	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		panic(err)
	}
}
