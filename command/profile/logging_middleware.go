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

func (m loggingMiddleware) AddProfile(jsonProfile string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "AddProfile",
			"jsonProfile", jsonProfile,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.AddProfile(jsonProfile)
	return
}

func (m loggingMiddleware) UpdateProfile(profileID string, jsonProfile string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfile",
			"jsonProfile", jsonProfile,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfile(profileID, jsonProfile)
	return
}

func (m loggingMiddleware) UpdateProfileStatus(profileID string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileStatus",
			"profileID", profileID,
			"status", status,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileStatus(profileID, status)
	return
}

func (m loggingMiddleware) UpdateProfileBase(profileID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileBase",
			"profileID", profileID,
			"mmap", mmap,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileBase(profileID, mmap)
	return
}

func (m loggingMiddleware) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfileAgencyMembers",
			"profileID", profileID,
			"agencyMembers", agencyMembers,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfileAgencyMembers(profileID, agencyMembers)
	return
}
