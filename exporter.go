package main

import (
	"github.com/btccom/node-exporter/sources/peers"
	"github.com/btccom/node-exporter/sources/rpc"
	"github.com/btccom/node-exporter/sources/ss"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const (
	CollectorTypeSSName  = "ss"
	CollectorTypeRPCName = "rpc"
	CollectorTypePeerName = "peer"
)

type Exporter struct {
}

func (m *Exporter) Init() error {
	logrus.Info(Config.Name)
	return nil
}

// Handle 处理实际的
func (m *Exporter) Register() error {
	// 初始化各种源，根据各种源构造 Collector 以及注册
	for _, item := range Config.Sources {
		if item.Type == CollectorTypeSSName {
			c := ss.NewStratumServerCollector()
			prometheus.MustRegister(c)
		}
		if item.Type == CollectorTypeRPCName {
			c := rpc.NewStratumServerCollector()
			prometheus.MustRegister(c)
		}
		if item.Type == CollectorTypePeerName {
			c := peers.NewPublicNodeCollector()
			prometheus.MustRegister(c)
		}
	}
	return nil
}
