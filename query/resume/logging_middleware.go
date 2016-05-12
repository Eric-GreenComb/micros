package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/resume/service"
)

type loggingMiddleware struct {
	service.ResumeService
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
	r = m.ResumeService.Ping()
	return
}

func (m loggingMiddleware) GetResume(id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetResume",
			"id", id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.GetResume(id)
	return
}
