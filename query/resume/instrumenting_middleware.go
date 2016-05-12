package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/resume/service"
)

type instrumentingMiddleware struct {
	service.ResumeService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.Ping()
	return
}

func (m instrumentingMiddleware) GetResume(id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetResume"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.GetResume(id)
	return
}
