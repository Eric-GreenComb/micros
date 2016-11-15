package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/profile/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.ProfileService
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

func (mw instrumentingMiddleware) AddProfile(jsonProfile string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddProfile", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.AddProfile(jsonProfile)
	return
}

func (mw instrumentingMiddleware) UpdateProfile(profileID string, jsonProfile string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateProfile", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateProfile(profileID, jsonProfile)
	return
}

func (mw instrumentingMiddleware) UpdateProfileStatus(profileID string, status bool) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateProfileStatus", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateProfileStatus(profileID, status)
	return
}

func (mw instrumentingMiddleware) UpdateProfileBase(profileID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateProfileBase", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateProfileBase(profileID, mmap)
	return
}

func (mw instrumentingMiddleware) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateProfileAgencyMembers", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateProfileAgencyMembers(profileID, agencyMembers)
	return
}
