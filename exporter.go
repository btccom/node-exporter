package main

import (
	"github.com/sirupsen/logrus"
)

type Exporter struct {
}

func (m *Exporter) Init() error {
	logrus.Info(Config.Name)
	return nil
}

// Handle 处理实际的
func (m *Exporter) Handle(title string) error {
	// 初始化各种源
	// c := newCollector(logger)
	// 根据各种源构造 Collector
	// 注册
	// prometheus.MustRegister(c)
	return nil
}
