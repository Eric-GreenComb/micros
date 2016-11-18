package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/profile/service"
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

func (mw instrumentingMiddleware) GetProfile(profileID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetProfile", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetProfile(profileID)
	return
}

func (mw instrumentingMiddleware) GetProfilesByUserID(userID string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetProfilesByUserID", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetProfilesByUserID(userID)
	return
}

func (mw instrumentingMiddleware) GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetProfilesByCategory", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetProfilesByCategory(categoryID, timestamp, pagesize)
	return
}

func (mw instrumentingMiddleware) GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetProfilesBySubCategory", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.GetProfilesBySubCategory(subcategoryID, timestamp, pagesize)
	return
}

func (mw instrumentingMiddleware) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SearchProfiles", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	return
}
