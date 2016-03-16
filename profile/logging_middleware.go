package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/profile/service"
)

type loggingMiddleware struct {
	service.ProfileService
	log.Logger
}

func (m loggingMiddleware) GetProfile(profile_id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfile",
			"profile_id", profile_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfile(profile_id)
	return
}

func (m loggingMiddleware) GetProfileByCat(name string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfileByCat",
			"name", name,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfileByCat(name)
	return
}

func (m loggingMiddleware) GetProfileBySubCat(name string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfileBySubCat",
			"name", name,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfileBySubCat(name)
	return
}
