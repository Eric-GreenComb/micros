package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/profile/service"
)

type instrumentingMiddleware struct {
	service.ProfileService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.Ping()
	return
}

func (m instrumentingMiddleware) GetProfile(id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.GetProfile(id)
	return
}

func (m instrumentingMiddleware) GetProfilesByEmail(email string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetProfilesByEmail"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.GetProfilesByEmail(email)
	return
}

func (m instrumentingMiddleware) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "SearchProfiles"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return
}
