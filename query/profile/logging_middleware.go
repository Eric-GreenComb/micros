package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/profile/service"
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

func (m loggingMiddleware) GetProfile(id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfile",
			"id", id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfile(id)
	return
}

func (m loggingMiddleware) GetProfilesByEmail(email string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfilesByEmail",
			"email", email,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfilesByEmail(email)
	return
}

func (m loggingMiddleware) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "SearchProfiles",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return
}
