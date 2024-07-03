package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func main() {
	parseFlags()

	client := resty.New()

	go func() {
		for {
			monitoring()
			time.Sleep(time.Duration(pollInterval) * time.Second)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(reportInterval) * time.Second)
			sendMetric(m, client)
		}
	}()

	err := http.ListenAndServe(flagRunEndpoint, nil)
	if err != nil {
		panic(err)
	}
}

type Monitor struct {
	runtime.MemStats
	PollCount   int64
	RandomValue float64
}

var m Monitor

func monitoring() {
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

func sendMetric(m Monitor, client *resty.Client) {

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

	for _, m := range metricsTable {
		doRequest(m.metricType, m.metricName, m.metricValue, client)
	}
}

func doRequest(metricType string, metricName string, metricValue string, client *resty.Client) {
	_, err := client.R().Post(fmt.Sprintf("http://localhost:8080/update/%s/%s/%s", metricType, metricName, metricValue))
	if err != nil {
		panic(err)
	}
}
