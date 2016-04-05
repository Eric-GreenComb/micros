package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/profile/service"
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
)

type instrumentingMiddleware struct {
	service.ProfileService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) GetProfile(profile_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.GetProfile(profile_id)
	return
}

func (m instrumentingMiddleware) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "SearchProfiles"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.SearchProfiles(profile_search_condition, timestamp, pagesize)
	return
}
