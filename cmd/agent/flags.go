package main

import (
	"flag"
	"time"
)

var flagRunEndpoint string = "localhost:8080"
var reportInterval time.Duration = 10
var pollInterval time.Duration = 2

func parseFlags() {
	flag.StringVar(&flagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.DurationVar(&reportInterval, "r", time.Duration(10*time.Second), "report Interval for metrics")
	flag.DurationVar(&pollInterval, "p", time.Duration(2*time.Second), "pool Interval for metrics")

	flag.Parse()
}
