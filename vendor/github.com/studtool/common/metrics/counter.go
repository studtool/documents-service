package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Counter struct {
	gauge  prometheus.Gauge
	ticker *time.Ticker
}

type CounterParams struct {
	Name          string
	Help          string
	ClearInterval time.Duration
}

func NewCounter(params CounterParams) *Counter {
	return &Counter{
		gauge: promauto.NewGauge(prometheus.GaugeOpts{
			Name: params.Name,
			Help: params.Help,
		}),
		ticker: time.NewTicker(params.ClearInterval),
	}
}

func (c *Counter) Run() {
	go func() {
		for range c.ticker.C {
			c.clear()
		}
	}()
}

func (c *Counter) Inc() {
	c.gauge.Inc()
}

func (c *Counter) clear() {
	c.gauge.Set(0)
}
