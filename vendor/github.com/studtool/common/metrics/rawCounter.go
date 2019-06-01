package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type RawCounter struct {
	counter prometheus.Counter
}

type RawCounterParams struct {
	Name string
	Help string
}

func NewRawCounter(params RawCounterParams) *RawCounter {
	return &RawCounter{
		counter: promauto.NewGauge(prometheus.GaugeOpts{
			Name: params.Name,
			Help: params.Help,
		}),
	}
}

func (c *RawCounter) Inc() {
	c.counter.Inc()
}
