package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/contact/service"
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

func (mw instrumentingMiddleware) CreateContact(jsonContact string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.CreateContact(jsonContact)
	return
}

func (mw instrumentingMiddleware) ClientSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "ClientSignContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.ClientSignContact(contactID, status)
	return
}

func (mw instrumentingMiddleware) FreelancerSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "FreelancerSignContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.FreelancerSignContact(contactID, status)
	return
}

func (mw instrumentingMiddleware) DealContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DealContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.DealContact(contactID, status)
	return
}

func (mw instrumentingMiddleware) UpdateContact(contactID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateContact", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateContact(contactID, mmap)
	return
}
