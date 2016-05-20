package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/profile/service"
)

type loggingMiddleware struct {
	service.ProfileService
	log.Logger
}

func (m loggingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.Ping()
	return
}

func (m loggingMiddleware) AddProfile(json_profile string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "AddProfile",
			"json_profile", json_profile,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.AddProfile(json_profile)
	return
}

func (m loggingMiddleware) UpdateProfile(profile_id string, json_profile string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfile",
			"json_profile", json_profile,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfile(profile_id, json_profile)
	return
}

func (m loggingMiddleware) UpdateProfileStatus(profile_id string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileStatus",
			"profile_id", profile_id,
			"status", status,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileStatus(profile_id, status)
	return
}

func (m loggingMiddleware) UpdateProfileBase(profile_id string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileBase",
			"profile_id", profile_id,
			"mmap", mmap,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileBase(profile_id, mmap)
	return
}

func (m loggingMiddleware) UpdateProfileAgencyMembers(profile_id string, agency_members string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileAgencyMembers",
			"profile_id", profile_id,
			"agency_members", agency_members,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileAgencyMembers(profile_id, agency_members)
	return
}
