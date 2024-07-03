package main

import "flag"

var flagRunEndpoint string

func parseFlags() {
	flag.StringVar(&flagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Parse()
}
