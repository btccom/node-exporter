package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

var Version = "1.0.0"

func init() {
	InitConfig("./configs/config.yaml")
}

func main() {
	// Init Exporter / 初始化 Exporter
	exporter := newExporter()
	if err := exporter.Init(); err != nil {
		panic(err)
	}
	if err := exporter.Register(); err != nil {
		panic(err)
	}
	listenAddr := ":8080"
	http.Handle("/metrics", promhttp.Handler())
	logrus.Infof("Now listening on %s", listenAddr)
	logrus.Fatal(http.ListenAndServe(listenAddr, nil))
}
