package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/workhistory/service"
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

func (m instrumentingMiddleware) GetWorkHistory(profile_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetWorkHistory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.WorkHistoryService.GetWorkHistory(profile_id)
	return
}
