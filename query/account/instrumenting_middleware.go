package main

import (
	"github.com/banerwai/micros/query/account/service"
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

func (mw instrumentingMiddleware) GetAccountByUserID(userID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAccountByUserID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetAccountByUserID(userID)
	return
}

func (mw instrumentingMiddleware) GetBillingByID(ID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetBillingByID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetBillingByID(ID)
	return
}

func (mw instrumentingMiddleware) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetDealBillingByUserID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetDealBillingByUserID(userID, timestamp, pagesize)
	return
}

func (mw instrumentingMiddleware) GetBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetBillingByUserID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetBillingByUserID(userID, timestamp, pagesize)
	return
}
