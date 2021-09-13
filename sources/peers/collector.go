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

func (n *PublicNodeCollector) Describe(ch chan<- *prometheus.Desc) {
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

func (n *PublicNodeCollector) Collect(chan<- prometheus.Metric) {
	panic("implement me")
}

func NewPublicNodeCollector() *PublicNodeCollector {
	return &PublicNodeCollector{}
}
