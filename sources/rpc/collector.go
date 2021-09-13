package rpc

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type rRPCSource struct {
	messageCount      prometheus.Counter
	messageEmptyCount prometheus.Counter
	duration          prometheus.Summary
	errors            *prometheus.CounterVec
}

type NodeDaemonCollector struct {

}

func (n *NodeDaemonCollector) Describe(ch chan<- *prometheus.Desc) {
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

func (n *NodeDaemonCollector) Collect(ch chan<- prometheus.Metric) {
	log.Printf("StratumServerCollector Collect")
}

func NewNodeDaemonCollector() *NodeDaemonCollector {
	return &NodeDaemonCollector{}
}