package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/workhistory/service"
)

type instrumentingMiddleware struct {
	service.WorkHistoryService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.WorkHistoryService.Ping()
	return
}

func (m instrumentingMiddleware) UpdateWorkHistory(profile_id, json_workhistory string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateWorkHistory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.WorkHistoryService.UpdateWorkHistory(profile_id, json_workhistory)
	return
}
