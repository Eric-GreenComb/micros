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

func (m loggingMiddleware) SearchProfiles(json_search string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "SearchProfiles",
			"json_search", json_search,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.SearchProfiles(json_search, timestamp, pagesize)
	return
}
