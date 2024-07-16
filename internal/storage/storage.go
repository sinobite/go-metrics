package storage

type Impl struct {
	storage MemStorage
}

type Storage interface {
	UpdateGaugeMetric(metricName string, metricValue string) error
	UpdateCounterMetric(metricName string, metricValue string) error
	FindGaugeMetric(metricName string) (string, error)
	FindCounterMetric(metricName string) (string, error)
	FindAllMetrics() ([]string, error)
}

func New() Impl {
	return Impl{
		storage: MemStorage{
			Gauges:   make(map[string]float64),
			Counters: make(map[string]int64),
		}}
}
