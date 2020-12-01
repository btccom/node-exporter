package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

func setGauge(name string, help string, callback func() float64) {
	gaugeFunc := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "bitcoind",
		Subsystem: "blockchain",
		Name:      name,
		Help:      help,
	}, callback)
	prometheus.MustRegister(gaugeFunc)
}

func init() {
	// init configs
}

func main() {
	listenAddr := ":8080"
	http.Handle("/metrics", promhttp.Handler())
	logrus.Info("Now listening on 8080")
	logrus.Fatal(http.ListenAndServe(listenAddr, nil))
}
