package collector

import (
	"math/rand"

	"github.com/kinduff/prometheus_exporter_template/internal/client"
	"github.com/kinduff/prometheus_exporter_template/internal/metrics"
	model "github.com/kinduff/prometheus_exporter_template/internal/models"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectThings() {
	thing := model.Thing{}
	if err := collector.client.DoAPIRequest(
		"demo",
		collector.config,
		&thing,
		&client.DoAPIRequestOptions{Id: "1"},
	); err != nil {
		log.Fatal(err)
	}

	metrics.Things.WithLabelValues(
		collector.config.BaseURL,
	).Set(rand.Float64())
}
