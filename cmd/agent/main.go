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
			sendMetric(m)
			time.Sleep(10 * time.Second)
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

	err := doRequest("gauge", "Alloc", strconv.Itoa(int(m.Alloc)), *client)
	err = doRequest("gauge", "BuckHashSys", strconv.Itoa(int(m.BuckHashSys)), *client)
	err = doRequest("gauge", "Frees", strconv.Itoa(int(m.Frees)), *client)
	err = doRequest("gauge", "GCCPUFraction", strconv.Itoa(int(m.GCCPUFraction)), *client)
	err = doRequest("gauge", "GCSys", strconv.Itoa(int(m.GCSys)), *client)
	err = doRequest("gauge", "HeapAlloc", strconv.Itoa(int(m.HeapAlloc)), *client)
	err = doRequest("gauge", "HeapIdle", strconv.Itoa(int(m.HeapIdle)), *client)
	err = doRequest("gauge", "HeapInuse", strconv.Itoa(int(m.HeapInuse)), *client)
	err = doRequest("gauge", "HeapObjects", strconv.Itoa(int(m.HeapObjects)), *client)
	err = doRequest("gauge", "HeapReleased", strconv.Itoa(int(m.HeapReleased)), *client)
	err = doRequest("gauge", "HeapSys", strconv.Itoa(int(m.HeapSys)), *client)
	err = doRequest("gauge", "LastGC", strconv.Itoa(int(m.LastGC)), *client)
	err = doRequest("gauge", "Lookups", strconv.Itoa(int(m.Lookups)), *client)
	err = doRequest("gauge", "MCacheInuse", strconv.Itoa(int(m.MCacheInuse)), *client)
	err = doRequest("gauge", "MCacheSys", strconv.Itoa(int(m.MCacheSys)), *client)
	err = doRequest("gauge", "MSpanInuse", strconv.Itoa(int(m.MSpanInuse)), *client)
	err = doRequest("gauge", "MSpanSys", strconv.Itoa(int(m.MSpanSys)), *client)
	err = doRequest("gauge", "Mallocs", strconv.Itoa(int(m.Mallocs)), *client)
	err = doRequest("gauge", "NextGC", strconv.Itoa(int(m.NextGC)), *client)
	err = doRequest("gauge", "NumForcedGC", strconv.Itoa(int(m.NumForcedGC)), *client)
	err = doRequest("gauge", "NumGC", strconv.Itoa(int(m.NumGC)), *client)
	err = doRequest("gauge", "OtherSys", strconv.Itoa(int(m.OtherSys)), *client)
	err = doRequest("gauge", "PauseTotalNs", strconv.Itoa(int(m.PauseTotalNs)), *client)
	err = doRequest("gauge", "StackInuse", strconv.Itoa(int(m.StackInuse)), *client)
	err = doRequest("gauge", "StackSys", strconv.Itoa(int(m.StackSys)), *client)
	err = doRequest("gauge", "Sys", strconv.Itoa(int(m.Sys)), *client)
	err = doRequest("gauge", "TotalAlloc", strconv.Itoa(int(m.TotalAlloc)), *client)
	err = doRequest("counter", "PollCount", strconv.Itoa(int(m.PollCount)), *client)
	err = doRequest("gauge", "RandomValue", strconv.Itoa(int(m.RandomValue)), *client)

	if err != nil {
		panic(err)
	}

}

func doRequest(metricType string, metricName string, metricValue string, client http.Client) error {
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8087/update/counter/someMetric/527", nil)
	if err != nil {
		return err
	}

	request.SetPathValue("metricType", metricType)
	request.SetPathValue("metricName", metricName)
	request.SetPathValue("metricValue", metricValue)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	//b, _ := json.Marshal(m)
	//fmt.Println(string(b))
	fmt.Println(resp.Status)
	return nil
}
