package metrics

import (
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Counter struct {
	value  uint64
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
		value:  0,
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
	v := atomic.LoadUint64(&c.value)
	v++
	atomic.StoreUint64(&c.value, v)
}

func (c *Counter) clear() {
	c.gauge.Set(float64(atomic.LoadUint64(&c.value)))
}
