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

func (m instrumentingMiddleware) AddWorkHistory(json_workhistory string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddWorkHistory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.WorkHistoryService.AddWorkHistory(json_workhistory)
	return
}

func (m instrumentingMiddleware) UpdateWorkHistory(json_workhistory string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateWorkHistory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.WorkHistoryService.UpdateWorkHistory(json_workhistory)
	return
}
