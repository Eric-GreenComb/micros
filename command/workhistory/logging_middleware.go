package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/workhistory/service"
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

func (m loggingMiddleware) UpdateWorkHistory(profile_id, json_workhistory string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateWorkHistory",
			"json_workhistory", json_workhistory,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.WorkHistoryService.UpdateWorkHistory(profile_id, json_workhistory)
	return
}
