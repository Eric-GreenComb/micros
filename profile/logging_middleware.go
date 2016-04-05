package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/profile/service"
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
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

func (m loggingMiddleware) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "SearchProfiles",
			"profile_search_condition", profile_search_condition,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ProfileService.SearchProfiles(profile_search_condition, timestamp, pagesize)
	return
}
