package metricsservice

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sinobite/go-metrics/internal/config/agentconfig"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Monitor struct {
	runtime.MemStats
	PollCount   int64
	RandomValue float64
	cfg         agentconfig.EnvConfig
}

func New(cfg agentconfig.EnvConfig) Monitor {
	return Monitor{
		MemStats:    runtime.MemStats{},
		PollCount:   0,
		RandomValue: 0,
		cfg:         cfg,
	}
}

func (m *Monitor) Monitoring() {
	var rtm runtime.MemStats

	// Read full mem stats
	runtime.ReadMemStats(&rtm)

	// Misc memory stats
	m.Alloc = rtm.Alloc
	m.BuckHashSys = rtm.BuckHashSys
	m.Frees = rtm.Frees
	m.GCCPUFraction = rtm.GCCPUFraction
	m.GCSys = rtm.GCSys
	m.HeapAlloc = rtm.HeapAlloc
	m.HeapIdle = rtm.HeapIdle
	m.HeapInuse = rtm.HeapInuse
	m.HeapObjects = rtm.HeapObjects
	m.HeapReleased = rtm.HeapReleased
	m.HeapSys = rtm.HeapSys
	m.LastGC = rtm.LastGC
	m.Lookups = rtm.Lookups
	m.MCacheInuse = rtm.MCacheInuse
	m.MCacheSys = rtm.MCacheSys
	m.MSpanInuse = rtm.MSpanInuse
	m.MSpanSys = rtm.MSpanSys
	m.Mallocs = rtm.Mallocs
	m.NextGC = rtm.NextGC
	m.NumForcedGC = rtm.NumForcedGC
	m.NumGC = rtm.NumGC
	m.OtherSys = rtm.OtherSys
	m.PauseTotalNs = rtm.PauseTotalNs
	m.StackInuse = rtm.StackInuse
	m.StackSys = rtm.StackSys
	m.Sys = rtm.Sys
	m.TotalAlloc = rtm.TotalAlloc
	m.PollCount = m.PollCount + 1
	m.RandomValue = m.RandomValue + 1
}

func (m *Monitor) SendMetric(client *resty.Client) {
	var metricsTable = []struct {
		metricType  string
		metricName  string
		metricValue string
	}{
		{"gauge", "Alloc", strconv.Itoa(int(m.Alloc))},
		{"gauge", "BuckHashSys", strconv.Itoa(int(m.BuckHashSys))},
		{"gauge", "Frees", strconv.Itoa(int(m.Frees))},
		{"gauge", "GCCPUFraction", strconv.Itoa(int(m.GCCPUFraction))},
		{"gauge", "GCSys", strconv.Itoa(int(m.GCSys))},
		{"gauge", "HeapAlloc", strconv.Itoa(int(m.HeapAlloc))},
		{"gauge", "HeapIdle", strconv.Itoa(int(m.HeapIdle))},
		{"gauge", "HeapInuse", strconv.Itoa(int(m.HeapInuse))},
		{"gauge", "HeapObjects", strconv.Itoa(int(m.HeapObjects))},
		{"gauge", "HeapReleased", strconv.Itoa(int(m.HeapReleased))},
		{"gauge", "HeapSys", strconv.Itoa(int(m.HeapSys))},
		{"gauge", "LastGC", strconv.Itoa(int(m.LastGC))},
		{"gauge", "Lookups", strconv.Itoa(int(m.Lookups))},
		{"gauge", "MCacheInuse", strconv.Itoa(int(m.MCacheInuse))},
		{"gauge", "MCacheSys", strconv.Itoa(int(m.MCacheSys))},
		{"gauge", "MSpanInuse", strconv.Itoa(int(m.MSpanInuse))},
		{"gauge", "MSpanSys", strconv.Itoa(int(m.MSpanSys))},
		{"gauge", "Mallocs", strconv.Itoa(int(m.Mallocs))},
		{"gauge", "NextGC", strconv.Itoa(int(m.NextGC))},
		{"gauge", "NumForcedGC", strconv.Itoa(int(m.NumForcedGC))},
		{"gauge", "NumGC", strconv.Itoa(int(m.NumGC))},
		{"gauge", "OtherSys", strconv.Itoa(int(m.OtherSys))},
		{"gauge", "PauseTotalNs", strconv.Itoa(int(m.PauseTotalNs))},
		{"gauge", "StackInuse", strconv.Itoa(int(m.StackInuse))},
		{"gauge", "StackSys", strconv.Itoa(int(m.StackSys))},
		{"gauge", "Sys", strconv.Itoa(int(m.Sys))},
		{"gauge", "TotalAlloc", strconv.Itoa(int(m.TotalAlloc))},
		{"counter", "PollCount", strconv.Itoa(int(m.PollCount))},
		{"gauge", "RandomValue", strconv.Itoa(int(m.RandomValue))},
	}

	for _, mt := range metricsTable {
		m.doRequest(mt.metricType, mt.metricName, mt.metricValue, client)
	}
}

func (m *Monitor) doRequest(metricType string, metricName string, metricValue string, client *resty.Client) {
	_, err := client.R().SetHeader("Content-Type", "text/plain").Post(fmt.Sprintf("http://%s/update/%s/%s/%s", m.cfg.FlagRunEndpoint, metricType, metricName, metricValue))
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Monitor) StartMonitoring(ctx context.Context, client *resty.Client) {
	var wg sync.WaitGroup

	pollTimer := time.NewTimer(time.Duration(m.cfg.PollInterval))
	reportTimer := time.NewTimer(time.Duration(m.cfg.ReportInterval))
	defer func() {
		pollTimer.Stop()
		reportTimer.Stop()
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		case <-pollTimer.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				m.Monitoring()
				pollTimer.Reset(time.Duration(m.cfg.PollInterval))
			}()
		case <-reportTimer.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				m.SendMetric(client)
				reportTimer.Reset(time.Duration(m.cfg.ReportInterval))
			}()
		}
	}
}
