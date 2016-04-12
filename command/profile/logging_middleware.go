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

func (m loggingMiddleware) UpdateProfile(json_profile string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateProfile",
			"json_profile", json_profile,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.UpdateProfile(json_profile)
	return
}

func (m loggingMiddleware) DeleteProfile(id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DeleteProfile",
			"id", id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.DeleteProfile(id)
	return
}
