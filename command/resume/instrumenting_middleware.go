package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/resume/service"
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

func (m instrumentingMiddleware) AddResume(json_resume string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddResume"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.AddResume(json_resume)
	return
}

func (m instrumentingMiddleware) UpdateResume(json_resume string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResume"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResume(json_resume)
	return
}
