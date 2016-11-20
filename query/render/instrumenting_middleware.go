package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/render/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.RenderService
}

func (mw instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Ping", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.Ping()
	return
}

func (mw instrumentingMiddleware) RenderTpl(tplname string, keyMap map[string]string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RenderTpl", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.RenderTpl(tplname, keyMap)
	return
}
