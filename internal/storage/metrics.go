package storage

import (
	"errors"
	"fmt"
	"github.com/sinobite/go-metrics/internal/utils"
	"strconv"
)

type MemStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

func (impl Impl) UpdateGaugeMetric(metricName string, metricValue string) error {
	gaugeValue, err := utils.ConvertToGauge(metricValue)
	if err != nil {
		return err
	}

	impl.storage.Gauges[metricName] = gaugeValue
	return nil
}

func (impl Impl) UpdateCounterMetric(metricName string, metricValue string) error {
	counterValue, err := utils.ConvertToCounter(metricValue)
	if err != nil {
		return err
	}

	impl.storage.Counters[metricName] = impl.storage.Counters[metricName] + counterValue
	return nil
}

func (impl Impl) FindGaugeMetric(metricName string) (string, error) {
	value, ok := impl.storage.Gauges[metricName]
	if ok {
		return strconv.FormatFloat(value, 'f', -1, 64), nil
	} else {
		return "", errors.New("there is no metric")
	}
}

func (impl Impl) FindCounterMetric(metricName string) (string, error) {
	value, ok := impl.storage.Counters[metricName]
	if ok {
		return strconv.Itoa(int(value)), nil
	} else {
		return "", errors.New("there is no metric")
	}
}

func (impl Impl) FindAllMetrics() ([]string, error) {
	var metricsSlice []string
	for key, value := range impl.storage.Gauges {
		metricsSlice = append(metricsSlice, fmt.Sprintf("%s: %v \n", key, value))
	}
	for key, value := range impl.storage.Counters {
		metricsSlice = append(metricsSlice, fmt.Sprintf("%s: %v \n", key, value))
	}

	return metricsSlice, nil
}
