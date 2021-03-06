package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/workhistory/service"
)

type loggingMiddleware struct {
	service.WorkHistoryService
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
	r = m.WorkHistoryService.Ping()
	return
}

func (m loggingMiddleware) GetWorkHistory(profileID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetWorkHistory",
			"profile_id", profileID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.WorkHistoryService.GetWorkHistory(profileID)
	return
}
