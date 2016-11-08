package main

import (
	"github.com/banerwai/micros/command/account/service"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.AccountService
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

func (mw instrumentingMiddleware) CreateAccount(jsonAccount string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateAccount", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CreateAccount(jsonAccount)
	return
}

func (mw instrumentingMiddleware) CreateBilling(jsonBilling string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateBilling", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CreateBilling(jsonBilling)
	return
}

func (mw instrumentingMiddleware) DealBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DealBilling", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.DealBilling(billingID)
	return
}

func (mw instrumentingMiddleware) CancelBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CancelBilling", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CancelBilling(billingID)
	return
}

func (mw instrumentingMiddleware) GenAccount(billingID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GenAccount", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GenAccount(billingID)
	return
}
