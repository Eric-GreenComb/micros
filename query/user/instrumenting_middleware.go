package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/user/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.UserService
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

func (mw instrumentingMiddleware) GetUserByEmail(email string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUserByEmail", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetUserByEmail(email)
	return
}

func (mw instrumentingMiddleware) GetUserByID(ID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUserByID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetUserByID(ID)
	return
}

func (mw instrumentingMiddleware) CountUser() (r int64) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CountUser", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CountUser()
	return
}
