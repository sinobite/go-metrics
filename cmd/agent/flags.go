package main

import "flag"

var flagRunEndpoint string
var reportInterval int64
var pollInterval int64

func parseFlags() {
	flag.StringVar(&flagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Int64Var(&reportInterval, "r", 10, "report Interval for metrics")
	flag.Int64Var(&pollInterval, "p", 2, "pool Interval for metrics")

	flag.Parse()
}
