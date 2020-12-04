package rpc

import "github.com/prometheus/client_golang/prometheus"

type rHttpSource struct {
	messageCount      prometheus.Counter
	messageEmptyCount prometheus.Counter
	duration          prometheus.Summary
	errors            *prometheus.CounterVec
}
