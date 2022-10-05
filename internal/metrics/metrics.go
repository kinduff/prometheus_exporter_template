// Package metrics sets and initializes Prometheus metrics.
package metrics

import (
	"github.com/kinduff/prometheus_exporter_template/config"

	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

var (
	namespace = "prometheus_exporter"

	Things = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "things_metrics",
			Help:      "Total number of things",
			Namespace: namespace,
		},
		[]string{"base_url"},
	)
)

// Init initializes all Prometheus metrics
func Init(config *config.Config) {
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	initMetric("things", Things)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}
