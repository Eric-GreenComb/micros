package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/resume/service"
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

func (m loggingMiddleware) AddResume(json_resume string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "AddResume",
			"json_resume", json_resume,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.AddResume(json_resume)
	return
}

func (m loggingMiddleware) UpdateResume(json_resume string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResume",
			"json_resume", json_resume,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResume(json_resume)
	return
}
