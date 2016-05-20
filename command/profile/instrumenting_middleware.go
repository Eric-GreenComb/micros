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

func (m instrumentingMiddleware) AddProfile(json_profile string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.AddProfile(json_profile)
	return
}

func (m instrumentingMiddleware) UpdateProfile(profile_id string, json_profile string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfile"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfile(profile_id, json_profile)
	return
}

func (m instrumentingMiddleware) UpdateProfileStatus(profile_id string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileStatus"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileStatus(profile_id, status)
	return
}

func (m instrumentingMiddleware) UpdateProfileBase(profile_id string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileBase"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileBase(profile_id, mmap)
	return
}

func (m instrumentingMiddleware) UpdateProfileAgencyMembers(profile_id string, agency_members string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateProfileAgencyMembers"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ProfileService.UpdateProfileAgencyMembers(profile_id, agency_members)
	return
}
