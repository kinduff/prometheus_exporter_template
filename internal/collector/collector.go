package collector

import (
	"time"

	"github.com/kinduff/prometheus_exporter_template/config"
	"github.com/kinduff/prometheus_exporter_template/internal/client"

	log "github.com/sirupsen/logrus"
)

type collector struct {
	config *config.Config
	client *client.Client
}

// NewCollector provides an interface to collector player statistics.
func NewCollector(config *config.Config) *collector {
	return &collector{
		config: config,
		client: client.NewClient(),
	}
}

func (collector *collector) Scrape() {
	collector.collectAll()

	for range time.Tick(collector.config.ScrapeInterval) {
		log.Info("Heartbeat, thump-thump...")
		collector.collectAll()
	}
}

func (collector *collector) collectAll() {
	go collector.collectThings()
}
