package metricsservice

import (
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentconfig"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMonitoring(t *testing.T) {
	cfg := agentconfig.New()
	ms := New(cfg)

	t.Run("MemStorage is changed", func(t *testing.T) {
		ms.Monitoring()

		assert.Equal(t, int64(1), ms.PollCount, "Poll count mismatch")
	})
}

func TestStartMonitoring(t *testing.T) {
	cfg := agentconfig.New()
	ms := New(cfg)
	client := resty.New()

	t.Run("start Monitoring", func(t *testing.T) {
		ms.StartMonitoring(client)
		assert.Equal(t, int64(0), ms.PollCount, "Poll count not empty")
		time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
		assert.Equal(t, int64(1), ms.PollCount, "Poll count not empty")
	})
}
