package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/workhistory/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.WorkHistoryService
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

func (mw instrumentingMiddleware) GetWorkHistory(profileID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetWorkHistory", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetWorkHistory(profileID)
	return
}
