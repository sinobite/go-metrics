package metricsservice

import (
	"github.com/sinobite/go-metrics/internal/config/agentconfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMonitoring(t *testing.T) {
	cfg := agentconfig.New()
	ms := New(cfg)

	t.Run("MemStorage is changed", func(t *testing.T) {
		ms.Monitoring()

		assert.Equal(t, int64(1), ms.PollCount, "Poll count mismatch")
	})
}

// todo разобраться как тестировать горутины
//func TestStartMonitoring(t *testing.T) {
//	cfg := agentconfig.New()
//	ms := New(cfg)
//	client := resty.New()
//
//	t.Run("start Monitoring", func(t *testing.T) {
//		ctx := context.Background()
//
//		ms.StartMonitoring(ctx, client)
//		assert.Equal(t, int64(0), ms.PollCount, "Poll count not empty")
//		time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
//		assert.Equal(t, true, ms.PollCount > 0, "Poll count not empty")
//	})
//}
