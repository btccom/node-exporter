package peers

import "github.com/prometheus/client_golang/prometheus"

type rPeersSource struct {
	messageCount      prometheus.Gauge
	messageEmptyCount prometheus.Counter
	duration          prometheus.Summary
	errors            *prometheus.CounterVec
}

func (r *rPeersSource) Handle() {
}

type PublicNodeCollector struct {
}

func (n PublicNodeCollector) Describe(chan<- *prometheus.Desc) {
	panic("implement me")
}

func (n PublicNodeCollector) Collect(chan<- prometheus.Metric) {
	panic("implement me")
}

func NewPublicNodeCollector() PublicNodeCollector {
	return PublicNodeCollector{}
}
