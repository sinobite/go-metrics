package storage

type MemStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

var Storage MemStorage = MemStorage{
	Gauges:   make(map[string]float64),
	Counters: make(map[string]int64),
}
