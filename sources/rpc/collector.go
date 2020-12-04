package rpc

import "github.com/prometheus/client_golang/prometheus"

type rRPCSource struct {
	messageCount      prometheus.Counter
	messageEmptyCount prometheus.Counter
	duration          prometheus.Summary
	errors            *prometheus.CounterVec
}

type NodeDaemonCollector struct {

}

func (n NodeDaemonCollector) Describe(chan<- *prometheus.Desc) {
	panic("implement me")
}

func (n NodeDaemonCollector) Collect(chan<- prometheus.Metric) {
	panic("implement me")
}

func NewStratumServerCollector() NodeDaemonCollector {
	return NodeDaemonCollector{}
}