package ss

import "github.com/prometheus/client_golang/prometheus"

type StratumServerCollector struct {

}

func (StratumServerCollector) Describe(chan<- *prometheus.Desc) {
	panic("implement me")
}

func (StratumServerCollector) Collect(chan<- prometheus.Metric) {
	panic("implement me")
}

func NewStratumServerCollector() StratumServerCollector {
	return StratumServerCollector{}
}