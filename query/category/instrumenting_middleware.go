package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/category/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.CategoryService
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

func (mw instrumentingMiddleware) LoadCategory(path string) (r bool) {
	defer func(begin time.Time) {
		lvs := []string{"method", "LoadCategory", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.LoadCategory(path)
	return
}

func (mw instrumentingMiddleware) GetCategories() (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetCategories", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetCategories()
	return
}

func (mw instrumentingMiddleware) GetSubCategories(serialnumber int32) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSubCategories", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetSubCategories(serialnumber)
	return
}
