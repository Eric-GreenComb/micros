package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/profile/service"
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

func (m instrumentingMiddleware) GetProfileByCat(name string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetProfileByCat"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.GetProfileByCat(name)
	return
}

func (m instrumentingMiddleware) GetProfileBySubCat(name string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetProfileBySubCat"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.GetProfileBySubCat(name)
	return
}
