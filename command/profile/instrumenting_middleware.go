package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/profile/service"
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

func (m instrumentingMiddleware) AddProfile(jsonProfile string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.AddProfile(jsonProfile)
	return
}

func (m instrumentingMiddleware) UpdateProfile(profileID string, jsonProfile string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfile(profileID, jsonProfile)
	return
}

func (m instrumentingMiddleware) UpdateProfileStatus(profileID string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileStatus"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileStatus(profileID, status)
	return
}

func (m instrumentingMiddleware) UpdateProfileBase(profileID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileBase"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileBase(profileID, mmap)
	return
}

func (m instrumentingMiddleware) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileAgencyMembers"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileAgencyMembers(profileID, agencyMembers)
	return
}
