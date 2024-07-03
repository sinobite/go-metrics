package main

import "flag"

var flagRunEndpoint string = "localhost:8080"

func parseFlags() {
	flag.StringVar(&flagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Parse()
}
