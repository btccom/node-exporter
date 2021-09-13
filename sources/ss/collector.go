package ss

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type StratumServerCollector struct {
}

func (c *StratumServerCollector) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})
	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()
	c.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

func (c *StratumServerCollector) Collect(ch chan<- prometheus.Metric) {
	log.Printf("StratumServerCollector Collect")
	// test
	// ⽣成采集的指标名
	name := prometheus.BuildFQName("sys", "", "mem_usage")
	// ⽣成 NewDesc 类型的数据格式，该指标⽆维度，[] string {} 为空
	desc := prometheus.NewDesc(name, "Gauge metric with mem_usage", []string{}, nil)
	// ⽣成具体的采集信息并写⼊ ch 通道
	metric := prometheus.MustNewConstMetric(desc,
		prometheus.GaugeValue, 22)
	ch <- metric
}

func NewStratumServerCollector() *StratumServerCollector {
	return &StratumServerCollector{}
}
