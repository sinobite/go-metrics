package main

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentconfig"
	"github.com/sinobite/go-metrics/internal/services/metricsservice"
	"log"
	"net/http"
)

func main() {
	cfg := agentconfig.New()
	cfg.Parse()

	ms := metricsservice.New(cfg)

	client := resty.New()

	ctx := context.Background()

	ms.StartMonitoring(ctx, client)

	err := http.ListenAndServe(cfg.AgentRunEndpoint, nil)
	if err != nil {
		log.Fatalf("server stopped with error: %s", err)
	}
}
