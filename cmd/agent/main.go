package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentConfig"
	"github.com/sinobite/go-metrics/internal/services/metricsService"
	"net/http"
	"time"
)

func main() {
	cfg := agentConfig.New()
	ms := metricsService.New(cfg)

	client := resty.New()

	go func() {
		for {
			ms.Monitoring()
			time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(cfg.ReportInterval) * time.Second)
			ms.SendMetric(client)
		}
	}()

	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		panic(err)
	}
}
