package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/user/service"
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

func (mw instrumentingMiddleware) CreateUser(mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateUser", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CreateUser(mmap)
	return
}

func (mw instrumentingMiddleware) ResetPwd(email string, newpwd string) (r bool) {
	defer func(begin time.Time) {
		lvs := []string{"method", "ResetPwd", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.ResetPwd(email, newpwd)
	return
}

func (mw instrumentingMiddleware) ActiveUser(email string) (r bool) {
	defer func(begin time.Time) {
		lvs := []string{"method", "ActiveUser", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.ActiveUser(email)
	return
}
