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

func (m loggingMiddleware) GetProfile(profileID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfile",
			"profileID", profileID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfile(profileID)
	return
}

func (m loggingMiddleware) GetProfilesByUserID(userID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetProfilesByUserID",
			"userID", userID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.GetProfilesByUserID(userID)
	return
}

func (m loggingMiddleware) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "SearchProfiles",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	return
}
