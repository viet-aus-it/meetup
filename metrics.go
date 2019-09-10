package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type metrics struct {
	holes_available prometheus.Gauge
}

func (m *metrics) register() error {
	if err := prometheus.Register(m.holes_available); nil != err {
		logrus.WithError(err).Errorln("can not register prometheus collector")
		return err
	}

	return nil
}

func newMetrics() *metrics {
	return &metrics {
		holes_available: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "gopher",
			Name: "holes_available",
		}),
	}
}