package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agent-config"
	"github.com/sinobite/go-metrics/internal/services/metrics-service"
	"net/http"
)

func main() {
	cfg := agent_config.New()
	cfg.Parse()

	ms := metrics_service.New(cfg)

	client := resty.New()

	ms.StartMonitoring(client)

	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		panic(err)
	}
}
