package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func main() {
	//select {
	//case <-time.After(2 * time.Second):
	//	go monitoring()
	//case <-time.After(10 * time.Second):
	//	go sendMetric(m)
	//}

	go func() {
		for {
			monitoring()
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			time.Sleep(10 * time.Second)
			sendMetric(m)
		}
	}()

	err := http.ListenAndServe("localhost:8088", nil)
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

	// Just encode to json and print
	//b, _ := json.Marshal(m)
	//fmt.Println(string(b))
}

func sendMetric(m Monitor) {
	client := &http.Client{}

	fmt.Println("sendMetric")

	doRequest("gauge", "Alloc", strconv.Itoa(int(m.Alloc)), *client)
	doRequest("gauge", "BuckHashSys", strconv.Itoa(int(m.BuckHashSys)), *client)
	doRequest("gauge", "Frees", strconv.Itoa(int(m.Frees)), *client)
	doRequest("gauge", "GCCPUFraction", strconv.Itoa(int(m.GCCPUFraction)), *client)
	doRequest("gauge", "GCSys", strconv.Itoa(int(m.GCSys)), *client)
	doRequest("gauge", "HeapAlloc", strconv.Itoa(int(m.HeapAlloc)), *client)
	doRequest("gauge", "HeapIdle", strconv.Itoa(int(m.HeapIdle)), *client)
	doRequest("gauge", "HeapInuse", strconv.Itoa(int(m.HeapInuse)), *client)
	doRequest("gauge", "HeapObjects", strconv.Itoa(int(m.HeapObjects)), *client)
	doRequest("gauge", "HeapReleased", strconv.Itoa(int(m.HeapReleased)), *client)
	doRequest("gauge", "HeapSys", strconv.Itoa(int(m.HeapSys)), *client)
	doRequest("gauge", "LastGC", strconv.Itoa(int(m.LastGC)), *client)
	doRequest("gauge", "Lookups", strconv.Itoa(int(m.Lookups)), *client)
	doRequest("gauge", "MCacheInuse", strconv.Itoa(int(m.MCacheInuse)), *client)
	doRequest("gauge", "MCacheSys", strconv.Itoa(int(m.MCacheSys)), *client)
	doRequest("gauge", "MSpanInuse", strconv.Itoa(int(m.MSpanInuse)), *client)
	doRequest("gauge", "MSpanSys", strconv.Itoa(int(m.MSpanSys)), *client)
	doRequest("gauge", "Mallocs", strconv.Itoa(int(m.Mallocs)), *client)
	doRequest("gauge", "NextGC", strconv.Itoa(int(m.NextGC)), *client)
	doRequest("gauge", "NumForcedGC", strconv.Itoa(int(m.NumForcedGC)), *client)
	doRequest("gauge", "NumGC", strconv.Itoa(int(m.NumGC)), *client)
	doRequest("gauge", "OtherSys", strconv.Itoa(int(m.OtherSys)), *client)
	doRequest("gauge", "PauseTotalNs", strconv.Itoa(int(m.PauseTotalNs)), *client)
	doRequest("gauge", "StackInuse", strconv.Itoa(int(m.StackInuse)), *client)
	doRequest("gauge", "StackSys", strconv.Itoa(int(m.StackSys)), *client)
	doRequest("gauge", "Sys", strconv.Itoa(int(m.Sys)), *client)
	doRequest("gauge", "TotalAlloc", strconv.Itoa(int(m.TotalAlloc)), *client)
	doRequest("counter", "PollCount", strconv.Itoa(int(m.PollCount)), *client)
	doRequest("gauge", "RandomValue", strconv.Itoa(int(m.RandomValue)), *client)

}

func doRequest(metricType string, metricName string, metricValue string, client http.Client) {
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/update/counter/someMetric/527", nil)
	if err != nil {
		panic(err)
	}

	request.SetPathValue("metricType", metricType)
	request.SetPathValue("metricName", metricName)
	request.SetPathValue("metricValue", metricValue)

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}
}
