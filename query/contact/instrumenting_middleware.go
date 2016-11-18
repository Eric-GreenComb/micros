package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/contact/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.ContactService
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

func (mw instrumentingMiddleware) GetContactTpl(tplname string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetContactTpl", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetContactTpl(tplname)
	return
}

func (mw instrumentingMiddleware) GetContact(contactid string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetContact(contactid)
	return
}

func (mw instrumentingMiddleware) GetContactSignStatus(contactid string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetContactSignStatus", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetContactSignStatus(contactid)
	return
}

func (mw instrumentingMiddleware) GetClientContact(clientemail string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetClientContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetClientContact(clientemail)
	return
}

func (mw instrumentingMiddleware) GetFreelancerContact(freelanceremail string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetFreelancerContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetFreelancerContact(freelanceremail)
	return
}
