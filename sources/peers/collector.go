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
