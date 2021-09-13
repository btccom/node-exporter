package http

import "github.com/prometheus/client_golang/prometheus"

type rHttpSource struct {
	messageCount      prometheus.Counter
	messageEmptyCount prometheus.Counter
	duration          prometheus.Summary
	errors            *prometheus.CounterVec
}

type ExplorerCollector struct {
}

func (n *ExplorerCollector) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})
	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()
	n.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

func (n *ExplorerCollector) Collect(chan<- prometheus.Metric) {
	panic("implement me")
}

func NewPublicNodeCollector() *ExplorerCollector {
	return &ExplorerCollector{}
}
